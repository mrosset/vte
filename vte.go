package vte

/*
#include <stdlib.h>
#include <vte/vte.h>
#include <gdk/gdk.h>


void get_colors(GdkColor*);

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
	return (*C.VteTerminal)(unsafe.Pointer(v.Widget))
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

type Palette struct {
}

var pal = map[int]string{}

func init() {
	pal[Black] = "#000000"
	pal[BlackLight] = "#252525"

	pal[Red] = "#803232"
	pal[RedLight] = "#982B2B"

	pal[Green] = "#85A136"
	pal[GreenLight] = "#85A136"

	pal[Yellow] = "#AA9943"
	pal[YellowLight] = "#EFEF60"

	pal[Blue] = "#324C80"
	pal[BlueLight] = "#4186BE"

	pal[Magenta] = "#706C9A"
	pal[MagentaLight] = "#826AB1"

	pal[Cyan] = "#92B19E"
	pal[CyanLight] = "#A1CDCD"

	pal[White] = "#E7E7E7"
	pal[WhiteLight] = "#E7E7E&"
}

const (
	Black = iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
	BlackLight
	RedLight
	GreenLight
	YellowLight
	BlueLight
	MagentaLight
	CyanLight
	WhiteLight
)

func (v *Terminal) SetBgColor(s string) {
	C.vte_terminal_set_color_background(v.getTerminal(), getColor(s))
}

func (v *Terminal) SetFgColor(s string) {
	C.vte_terminal_set_color_foreground(v.getTerminal(), getColor(s))
}

func (v *Terminal) SetColors() {
	colors := new([16]C.GdkColor)
	for i := 0; i < len(colors); i++ {
		C.gdk_color_parse((*C.gchar)(C.CString(pal[i])), &colors[i])
	}
	C.vte_terminal_set_colors(
		v.getTerminal(),
		nil, nil,
		(*C.GdkColor)(unsafe.Pointer(colors)),
		16)
}

func NewTerminal() *Terminal {
	return &Terminal{*gtk.WidgetFromNative(unsafe.Pointer(C.vte_terminal_new()))}
}

func getColor(s string) *C.GdkColor {
	c := gdk.Color(s).Color
	return (*C.GdkColor)(unsafe.Pointer(&c))
}
