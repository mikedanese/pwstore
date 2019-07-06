package pwdb

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"syscall"

	"github.com/golang/protobuf/proto"
	"github.com/google/tink/go/aead"
	"github.com/google/tink/go/keyset"
	"github.com/google/tink/go/subtle/random"
	"github.com/google/tink/go/tink"
	"github.com/mikedanese/pwstore/passwd"
	"golang.org/x/sys/unix"
)

func Open() (*DB, error) {
	// We want the permissions we specify to be respected.
	syscall.Umask(0)

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("unable to find user home dir: %v", err)
	}
	pwDir := filepath.Join(homeDir, ".pwstore")
	if err := os.MkdirAll(pwDir, 0700); err != nil {
		return nil, err
	}

	fd, err := unix.Open(filepath.Join(pwDir, "lock"), unix.O_CREAT|unix.O_WRONLY, 0600)
	if err != nil {
		return nil, err
	}
	// Hold lock until process exits.
	if err := unix.Flock(fd, unix.LOCK_EX|unix.LOCK_NB); err != nil {
		return nil, fmt.Errorf("failed to acquire DB lock: %v", err)
	}

	key, err := loadMasterAEAD(pwDir)
	if err != nil {
		return nil, err
	}

	db := &DB{
		dir:     pwDir,
		records: make(map[string][]byte),
		master:  key,
	}
	if err := db.load(); err != nil {
		return nil, err
	}
	return db, nil
}

type DB struct {
	dir     string
	master  tink.AEAD
	records map[string][]byte
}

func (db *DB) List() []string {
	names := []string{}
	for name, _ := range db.records {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func (db *DB) Get(name string) (*Record, error) {
	c, ok := db.records[name]
	if !ok {
		return nil, fmt.Errorf("password %q not found", name)
	}
	b, err := db.master.Decrypt(c, []byte(name))
	if err != nil {
		return nil, err
	}
	var out Record
	if err := proto.Unmarshal(b, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (db *DB) Put(name string, r *Record) error {
	b, err := proto.Marshal(r)
	if err != nil {
		return err
	}
	c, err := db.master.Encrypt(b, []byte(name))
	if err != nil {
		return err
	}
	db.records[name] = c
	return db.commit()
}

func (db *DB) load() error {
	pwPath := filepath.Join(db.dir, "pw.db")
	var rs RecordSet
	b, err := ioutil.ReadFile(pwPath)
	if err != nil {
		if os.IsNotExist(err) {
			return db.commit()
		}
		return err
	}
	if err := proto.Unmarshal(b, &rs); err != nil {
		return err
	}
	records := make(map[string][]byte)
	for _, env := range rs.Records {
		records[env.Name] = env.Data
	}
	db.records = records
	return nil
}

func (db *DB) commit() error {
	pwPath := filepath.Join(db.dir, "pw.db")
	var rs RecordSet
	for name, val := range db.records {
		rs.Records = append(rs.Records, &Envelope{
			Name: name,
			Data: val,
		})
	}
	b, err := proto.Marshal(&rs)
	if err != nil {
		return err
	}
	return writeFile(pwPath, b)
}

func loadMasterAEAD(pwDir string) (tink.AEAD, error) {
	// load salt
	saltPath := filepath.Join(pwDir, "salt")
	salt, err := ioutil.ReadFile(saltPath)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, fmt.Errorf("failed to read salt from %q: %v", saltPath, err)
		}
		salt = random.GetRandomBytes(16)
		if err := writeFile(saltPath, salt); err != nil {
			return nil, fmt.Errorf("failed to write initial salt to %q: %v", saltPath, err)
		}
	}

	pwKey, err := passwd.Read(salt)
	if err != nil {
		return nil, fmt.Errorf("failed to read password: %v", err)
	}

	// load master secret
	masterPath := filepath.Join(pwDir, "master")
	masterb, err := ioutil.ReadFile(masterPath)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, fmt.Errorf("failed to read master from %q: %v", masterPath, err)
		}

		h, err := keyset.NewHandle(aead.XChaCha20Poly1305KeyTemplate())
		if err != nil {
			return nil, err
		}

		var buf bytes.Buffer
		if err := h.Write(keyset.NewBinaryWriter(&buf), pwKey); err != nil {
			return nil, fmt.Errorf("failed to write initial master keyset: %v", err)
		}

		if err := writeFile(masterPath, buf.Bytes()); err != nil {
			return nil, fmt.Errorf("failed to write initial master keyset to %q: %v", masterPath, err)
		}
		masterb = buf.Bytes()
	}
	ks, err := keyset.Read(keyset.NewBinaryReader(bytes.NewReader(masterb)), pwKey)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt master keyset: %v", err)
	}
	key, err := aead.New(ks)
	if err != nil {
		return nil, err
	}
	return key, nil
}
