package main

/*
#cgo linux LDFLAGS: -lX11 -Wl,--allow-multiple-definition
#cgo darwin LDFLAGS: -lX11 -L/opt/homebrew/lib
#cgo darwin CFLAGS: -I/opt/homebrew/include -fcommon
#include <stdio.h>
#include <X11/Xlib.h>

static int* getScreenSize(){
	Display* d = XOpenDisplay(NULL);
	Screen*  s = DefaultScreenOfDisplay(d);
	static int r[2];
	r[0] = s->width;
	r[1] = s->height;
	return r;
}
*/
import "C"
import "unsafe"

func getDisplaySize() (int, int) {
	var srceenSize *C.int = C.getScreenSize()
	length := 2
	slice := (*[1 << 28]C.int)(unsafe.Pointer(srceenSize))[:length:length]
	return (int)(slice[0]), (int)(slice[1])
}
