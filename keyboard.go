package main

/*
#cgo linux LDFLAGS: -lX11 -Wl,--allow-multiple-definition
#cgo darwin LDFLAGS: -lX11 -L/opt/homebrew/lib
#cgo darwin CFLAGS: -I/opt/homebrew/include -fcommon
#include <stdio.h>
#include <ctype.h>
#include <X11/X.h>
#include <X11/Xlib.h>
#include <X11/Xutil.h>
#include <X11/XKBlib.h>

extern void keyboardEvent(char*,int);

Bool running;

static void startKeyboardListener() {
	Display *display = XOpenDisplay(NULL);
	XGrabKeyboard(display, DefaultRootWindow(display), True, GrabModeAsync, GrabModeAsync, CurrentTime);
	XEvent event;
	running = True;
	while(running) {
			XNextEvent(display, &event);
			switch (event.type){
					case KeyPress: {
							int a;
							char *key = XKeysymToString(*XGetKeyboardMapping(display,event.xkey.keycode,1,&a));
							keyboardEvent(key,1);
							continue;
					}
					case KeyRelease: {
							int a;
							char *key = XKeysymToString(*XGetKeyboardMapping(display,event.xkey.keycode,1,&a));
							keyboardEvent(key,0);
							continue;
					}
			}
	}
	XUngrabKeyboard(display, CurrentTime);
	XFlush(display);
}
static void releaseKeyboard() {
	running = False;
}
*/
import "C"

var keyboardCallback func(string, bool)

func startKeyboardListen(calback func(string, bool)) {
	keyboardCallback = calback
	go C.startKeyboardListener()
}

//export keyboardEvent
func keyboardEvent(key *C.char, pressed C.int) {
	sKey := C.GoString(key)
	press := false
	if pressed == 1 {
		press = true
	}
	keyboardCallback(sKey, press)
}

func releaseKeyboard() {
	C.releaseKeyboard()
}
