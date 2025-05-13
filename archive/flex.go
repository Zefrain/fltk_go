package fltk_go

/*
#include "flex.h"
*/
import "C"
import "unsafe"

type Flex struct {
	Group
}

func NewFlex(x, y, w, h int, text ...string) *Flex {
	p := &Flex{}
	initWidget(p, unsafe.Pointer(C.go_fltk_new_Flex(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return p
}

type FlexType uint8

var (
	COLUMN = FlexType(C.go_FL_FLEX_COLUMN)
	ROW    = FlexType(C.go_FL_FLEX_ROW)
)

func (f *Flex) SetType(flexType FlexType) {
	f.widget.SetType(uint8(flexType))
}
func (f *Flex) SetGap(spacing int) {
	C.go_fltk_Flex_set_gap((*C.Fl_Flex)(f.ptr()), C.int(spacing))
}
func (f *Flex) Fixed(w Widget, size int) {
	C.go_fltk_Flex_fixed((*C.Fl_Flex)(f.ptr()), w.getWidget().ptr(), C.int(size))
}
func (f *Flex) End() {
	C.go_fltk_Flex_end((*C.Fl_Flex)(f.ptr()))
}
func (f *Flex) Spacing() int {
	return int(C.go_fltk_Flex_spacing((*C.Fl_Flex)(f.ptr())))
}
func (f *Flex) SetSpacing(spacing int) {
	C.go_fltk_Flex_set_spacing((*C.Fl_Flex)(f.ptr()), C.int(spacing))
}
func (f *Flex) Margin() int {
	return int(C.go_fltk_Flex_margin((*C.Fl_Flex)(f.ptr())))
}
func (f *Flex) SetMargin(margin int, gap ...int) {
	g := -1
	if len(gap) > 0 {
		g = gap[0]
	}
	C.go_fltk_Flex_set_margin((*C.Fl_Flex)(f.ptr()), C.int(margin), C.int(g))
}

// Layout calculates the layout of the widget and redraws it.
func (f *Flex) Layout() {
	C.go_fltk_Flex_layout((*C.Fl_Flex)(f.ptr()))
}
