package vte

/*
#include <stdlib.h>
#include <vte/vte.h>
#include <gdk/gdk.h>

static VteTerminal* to_VteTerminal(void* w) { return VTE_TERMINAL(w); }
*/
// #cgo pkg-config: vte
import "C"

import (
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/gtk"
	"unsafe"
)

type Terminal struct {
	gtk.GtkWidget
}

func (v *Terminal) getTerminal() *C.VteTerminal {
	return C.to_VteTerminal(unsafe.Pointer(v.Widget))
}

func (v *Terminal) Feed(m string) {
	c := C.CString(m)
	defer C.free(unsafe.Pointer(c))
	C.vte_terminal_feed(v.getTerminal(), C.CString(m), -1)
}

func (v *Terminal) Fork(a string) {
	arg := C.CString(a)
	defer C.free(unsafe.Pointer(arg))
	C.vte_terminal_fork_command_full(v.getTerminal(),
		C.VTE_PTY_DEFAULT,
		nil,
		&arg,
		nil,
		C.G_SPAWN_SEARCH_PATH,
		nil,
		nil,
		nil, nil)
}

func (v *Terminal) BgColor(s string) {
	c := gdk.Color(s)
	cc := (*C.GdkColor)(unsafe.Pointer(&c.Color))
	C.vte_terminal_set_color_background(v.getTerminal(), cc)
}

func (v *Terminal) FgColor(s string) {
	c := gdk.Color(s)
	cc := (*C.GdkColor)(unsafe.Pointer(&c.Color))
	C.vte_terminal_set_color_foreground(v.getTerminal(), cc)
}

func NewTerminal() *Terminal {
	return &Terminal{*gtk.WidgetFromNative(unsafe.Pointer(C.vte_terminal_new()))}
}
