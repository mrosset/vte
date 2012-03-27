package vte

import (
	"github.com/mattn/go-gtk/gtk"
	"os"
	"testing"
)

func TestVte(t *testing.T) {
	gtk.Init(&os.Args)
	window := gtk.Window(gtk.GTK_WINDOW_TOPLEVEL)
	terminal := NewTerminal()
	terminal.Fork("bash")
	terminal.Connect("child-exited", gtk.MainQuit)
	window.Add(terminal)
	window.ShowAll()
	terminal.SetColors()
	gtk.Main()
}
