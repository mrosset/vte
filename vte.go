package vte

/*
#include <stdlib.h>
#include <vte/vte.h>

static inline char** make_strings(int count) {
	return (char**)malloc(sizeof(char*) * count);
}

static inline void set_string(char** strings, int n, char* str) {
	strings[n] = str;
}

*/
// #cgo pkg-config: vte
import "C"

import (
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/gtk"
	"unsafe"
)

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

var MikePal = map[int]string{
	Black:        "#000000",
	BlackLight:   "#252525",
	Red:          "#803232",
	RedLight:     "#982B2B",
	Green:        "#85A136",
	GreenLight:   "#85A136",
	Yellow:       "#AA9943",
	YellowLight:  "#EFEF60",
	Blue:         "#324C80",
	BlueLight:    "#4186BE",
	Magenta:      "#706C9A",
	MagentaLight: "#826AB1",
	Cyan:         "#92B19E",
	CyanLight:    "#A1CDCD",
	White:        "#E7E7E7",
	WhiteLight:   "#E7E7E&",
}

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

func (v *Terminal) Fork(args []string) {
	cargs := C.make_strings(C.int(len(args)))
	for i, j := range args {
		ptr := C.CString(j)
		defer C.free(unsafe.Pointer(ptr))
		C.set_string(cargs, C.int(i), ptr)
	}
	C.vte_terminal_fork_command_full(v.getTerminal(),
		C.VTE_PTY_DEFAULT,
		nil,
		cargs,
		nil,
		C.G_SPAWN_SEARCH_PATH,
		nil,
		nil,
		nil, nil)
}

type Palette struct {
}

func (v *Terminal) SetBgColor(s string) {
	C.vte_terminal_set_color_background(v.getTerminal(), getColor(s))
}

func (v *Terminal) SetFgColor(s string) {
	C.vte_terminal_set_color_foreground(v.getTerminal(), getColor(s))
}

func (v *Terminal) SetColors(pal map[int]string) {
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
