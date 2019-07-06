package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/gogo/protobuf/proto"
	"github.com/mikedanese/pwstore/pwdb"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	root := &cobra.Command{
		Use: "pwstore",
	}
	addSub(root, &copyCmd{})

	raw := &cobra.Command{
		Use: "raw",
	}
	root.AddCommand(raw)

	addSub(raw, &getCmd{})
	addSub(raw, &listCmd{})
	addSub(raw, &putCmd{})

	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type cmd interface {
	cmd() *cobra.Command
	bindFlags(fs *pflag.FlagSet)
}

func addSub(root *cobra.Command, sub cmd) *cobra.Command {
	cmd := sub.cmd()
	sub.bindFlags(cmd.Flags())
	root.AddCommand(cmd)
	return cmd
}

type getCmd struct {
	name string
}

func (c *getCmd) cmd() *cobra.Command {
	return &cobra.Command{
		Use: "get",
		Run: c.run,
	}
}

func (c *getCmd) bindFlags(fs *pflag.FlagSet) {
	fs.StringVar(&c.name, "name", "", "")
	cobra.MarkFlagRequired(fs, "name")
}

func (c *getCmd) run(cmd *cobra.Command, args []string) {
	db, err := pwdb.Open()
	if err != nil {
		cmd.PrintErrf("failed to open pwdb: %v", err)
		return
	}
	r, err := db.Get(c.name)
	if err != nil {
		cmd.PrintErrf("failed to get %q: %v", c.name, err)
		return
	}
	fmt.Print(proto.MarshalTextString(r))
}

type listCmd struct {
}

func (c *listCmd) cmd() *cobra.Command {
	return &cobra.Command{
		Use: "list",
		Run: c.run,
	}
}

func (c *listCmd) bindFlags(fs *pflag.FlagSet) {
}

func (c *listCmd) run(cmd *cobra.Command, args []string) {
	db, err := pwdb.Open()
	if err != nil {
		cmd.PrintErrf("failed to open pwdb: %v", err)
		return
	}
	for _, name := range db.List() {
		cmd.Println(name)
	}
}

type putCmd struct {
	name string
	file string
}

func (c *putCmd) cmd() *cobra.Command {
	return &cobra.Command{
		Use: "put",
		Run: c.run,
	}
}

func (c *putCmd) bindFlags(fs *pflag.FlagSet) {
	fs.StringVar(&c.name, "name", "", "")
	cobra.MarkFlagRequired(fs, "name")
	fs.StringVar(&c.file, "file", "", "")
	cobra.MarkFlagRequired(fs, "file")
}

func (c *putCmd) run(cmd *cobra.Command, args []string) {
	db, err := pwdb.Open()
	if err != nil {
		cmd.PrintErrf("failed to open pwdb: %v", err)
		return
	}
	b, err := ioutil.ReadFile(c.file)
	if err != nil {
		cmd.PrintErrf("failed to read from stdin: %v", err)
		return
	}
	var r pwdb.Record
	if err := proto.UnmarshalText(string(b), &r); err != nil {
		cmd.PrintErrf("failed to read from stdin: %v", err)
		return
	}
	if err := db.Put(c.name, &r); err != nil {
		cmd.PrintErrf("failed to put password: %v", err)
		return
	}
	cmd.Println("ok")
}

type copyCmd struct {
	name     string
	username bool
}

func (c *copyCmd) cmd() *cobra.Command {
	return &cobra.Command{
		Use: "copy",
		Run: c.run,
	}
}

func (c *copyCmd) bindFlags(fs *pflag.FlagSet) {
	fs.StringVar(&c.name, "name", "", "")
	cobra.MarkFlagRequired(fs, "name")
	fs.BoolVarP(&c.username, "username", "u", false, "")
}

func (c *copyCmd) run(cmd *cobra.Command, args []string) {
	db, err := pwdb.Open()
	if err != nil {
		cmd.PrintErrf("failed to open pwdb: %v", err)
		return
	}
	r, err := db.Get(c.name)
	if err != nil {
		cmd.PrintErrf("failed to copy %q: %v", c.name, err)
		return
	}
	out := r.Password
	if c.username {
		out = r.Username
	}
	fmt.Printf("\x1b]52;c;" + base64.StdEncoding.EncodeToString([]byte(out)) + "\x07")
	cmd.Println("ok")
}
