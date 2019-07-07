package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"unicode/utf8"

	"github.com/gogo/protobuf/proto"
	"github.com/google/tink/go/subtle/random"
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
	addSub(root, &genCmd{})

	raw := &cobra.Command{
		Use:   "raw",
		Short: "Raw database access.",
	}
	root.AddCommand(raw)

	completion := &cobra.Command{
		Use:   "completion",
		Short: "Generates bash completion scripts",
		Run: func(cmd *cobra.Command, args []string) {
			root.GenBashCompletion(os.Stdout)
		},
	}
	root.AddCommand(completion)

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
	ansiCopy(cmd.OutOrStdout(), out)
	cmd.Println("ok")
}

func ansiCopy(w io.Writer, s string) {
	fmt.Fprintf(os.Stdout, "\x1b]52;c;%s\x07", base64.StdEncoding.EncodeToString([]byte(s)))
}

type genCmd struct {
	length int
}

func (c *genCmd) cmd() *cobra.Command {
	return &cobra.Command{
		Use: "gen",
		Run: c.run,
	}
}

func (c *genCmd) bindFlags(fs *pflag.FlagSet) {
	fs.IntVarP(&c.length, "length", "l", 20, "")
}

func (c *genCmd) run(cmd *cobra.Command, args []string) {
	var buf bytes.Buffer
	for i := 0; i < c.length; {
		bs := random.GetRandomBytes(4)
		if utf8.Valid(bs) {
			buf.Write(bs)
			i++
		}
	}
	cmd.Println(buf.String())
}
