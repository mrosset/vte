package vte

/*
#include <stdlib.h>
#include <vte/vte.h>
static VteTerminal* to_VteTerminal(void* w) { return VTE_TERMINAL(w); }
*/
// #cgo pkg-config: vte
import "C"

import (
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

func NewTerminal() *Terminal {
	return &Terminal{*gtk.WidgetFromNative(unsafe.Pointer(C.vte_terminal_new()))}
}
