package cui

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/jroimartin/gocui"
	"github.com/mikedanese/pwstore/bazel-pwstore/external/com_github_golang_protobuf/proto"
	"github.com/nsf/termbox-go"
)

type ui struct {
	db        *db.DB
	g         *gocui.Gui
	viewStack []string
}

func (u *ui) Run() error {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return err
	}
	defer g.Close()

	g.SetManagerFunc(u.layout)
	g.Highlight = true
	g.Cursor = true
	g.SelFgColor = gocui.Attribute(termbox.AttrBold | termbox.ColorMagenta)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, u.quit); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyEsc, gocui.ModNone, u.popView); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowLeft, gocui.ModNone, u.popView); err != nil {
		return err
	}
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		return err
	}
	return nil
}

func (u *ui) layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	v, err := g.SetView("main", 0, 0, maxX-1, maxY-1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		if err := u.pushView(g, v); err != nil {
			return err
		}
	}
	if err := u.layoutMain(v); err != nil {
		return err
	}
	g.DeleteKeybindings("main")
	if err := g.SetKeybinding("main", gocui.KeyArrowUp, gocui.ModNone, up); err != nil {
		return err
	}
	if err := g.SetKeybinding("main", gocui.KeyArrowDown, gocui.ModNone, down); err != nil {
		return err
	}
	if err := g.SetKeybinding("main", gocui.KeyArrowRight, gocui.ModNone, u.focus); err != nil {
		return err
	}
	return nil
}

func (u *ui) layoutMain(v *gocui.View) error {
	v.Clear()
	v.Title = "Passwords"
	v.Highlight = true
	v.SelFgColor = gocui.Attribute(termbox.AttrBold | termbox.AttrUnderline | termbox.ColorDefault)
	for _, name := range u.db.List() {
		fmt.Fprintln(v, name)
	}
	return nil
}

func (u *ui) focus(g *gocui.Gui, v *gocui.View) error {
	maxX, maxY := v.Size()
	_, y := v.Cursor()
	name, err := v.Line(y)
	if err != nil {
		return err
	}
	v, err = g.SetView("focus", 25, 1, maxX-1, maxY)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		if err := u.pushView(g, v); err != nil {
			return err
		}
	}
	if err := u.layoutFocus(v, name); err != nil {
		return err
	}
	g.DeleteKeybindings("focus")
	if err := g.SetKeybinding("focus", 'e', gocui.ModNone, u.editFocus); err != nil {
		return err
	}
	return nil
}

func (u *ui) layoutFocus(v *gocui.View, name string) error {
	v.Clear()
	v.Title = name
	r, err := u.db.Get(name)
	if err != nil {
		fmt.Fprintf(v, "Error: %v", err)
		return nil
	}
	fmt.Fprint(v, proto.MarshalTextString(r))
	return nil
}

func up(g *gocui.Gui, v *gocui.View) error {
	v.MoveCursor(0, -1, false)
	return nil
}

func down(g *gocui.Gui, v *gocui.View) error {
	v.MoveCursor(0, 1, false)
	return nil
}

func (u *ui) quit(_ *gocui.Gui, _ *gocui.View) error {
	return gocui.ErrQuit
}

func (u *ui) pushView(g *gocui.Gui, v *gocui.View) error {
	u.viewStack = append(u.viewStack, v.Name())
	if _, err := g.SetCurrentView(v.Name()); err != nil {
		return err
	}
	return nil
}

func (u *ui) popView(g *gocui.Gui, v *gocui.View) error {
	if u.viewStack[len(u.viewStack)-1] != v.Name() {
		panic("uh oh")
	}
	if len(u.viewStack) == 1 {
		// Already on main
		return u.quit(g, v)
	}
	u.viewStack = u.viewStack[:len(u.viewStack)-1]
	if err := g.DeleteView(v.Name()); err != nil {
		return err
	}
	if _, err := g.SetCurrentView("main"); err != nil {
		return err
	}
	return nil
}

func (u *ui) editFocus(g *gocui.Gui, v *gocui.View) error {
	b, err := systemEdit(g, v)
	if err != nil {
		return err
	}
	v.Clear()

	var r db.Record
	if err := proto.UnmarshalText(string(b), &r); err != nil {
		return err
	}
	if err := u.db.Put(v.Title, &r); err != nil {
		return err
	}

	if _, err := v.Write(b); err != nil {
		return err
	}
	return nil
}

func systemEdit(g *gocui.Gui, v *gocui.View) ([]byte, error) {
	const editTemp = "/dev/shm/pwstore.tmp/tmp"
	e := os.Getenv("EDITOR")
	if e == "" {
		return nil, errors.New("EDITOR not set")
	}
	if err := os.MkdirAll(filepath.Dir(editTemp), 0700); err != nil {
		return nil, err
	}

	defer func() {
		if err := os.Remove(editTemp); err != nil {
			panic(err)
		}
	}()
	if err := ioutil.WriteFile(editTemp, []byte(v.Buffer()), 0600); err != nil {
		return nil, err
	}

	termbox.Close()

	cmd := exec.Command(e, editTemp)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	termbox.Init()
	inputMode := termbox.InputAlt
	if g.InputEsc {
		inputMode = termbox.InputEsc
	}
	if g.Mouse {
		inputMode |= termbox.InputMouse
	}
	termbox.SetInputMode(inputMode)

	return ioutil.ReadFile(editTemp)
}
