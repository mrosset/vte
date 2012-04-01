package vte

import (
	//"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"os"
	"testing"
)

func TestVte(t *testing.T) {
	gtk.Init(&os.Args)
	window := gtk.Window(gtk.GTK_WINDOW_TOPLEVEL)
	terminal := NewTerminal()
	swin := gtk.ScrolledWindow(nil, nil)
	swin.SetPolicy(gtk.GTK_POLICY_NEVER, gtk.GTK_POLICY_NEVER)
	terminal.Fork([]string{"bash", "--login"})
	terminal.Connect("child-exited", gtk.MainQuit)
	swin.Add(terminal)
	window.Add(swin)
	window.SetSizeRequest(309, 99)
	window.ShowAll()
	terminal.SetColors(MikePal)
	gtk.Main()
}
