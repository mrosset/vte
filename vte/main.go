package main

import (
	"github.com/mattn/go-gtk/gtk"
	"os"
	"vte"
)

func main() {
	gtk.Init(&os.Args)
	window := gtk.Window(gtk.GTK_WINDOW_TOPLEVEL)
	swin := gtk.ScrolledWindow(nil, nil)
	swin.SetPolicy(gtk.GTK_POLICY_AUTOMATIC, gtk.GTK_POLICY_AUTOMATIC)
	swin.SetShadowType(gtk.GTK_SHADOW_IN)
	terminal := vte.NewTerminal()
	terminal.Fork("bash")
	swin.Add(terminal)
	window.Add(swin)
	window.ShowAll()
	gtk.Main()
}
