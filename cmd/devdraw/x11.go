package main

import (
	"log"

	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/keybind"
	"github.com/BurntSushi/xgbutil/mousebind"
	"github.com/BurntSushi/xgbutil/xevent"
	"github.com/BurntSushi/xgbutil/xwindow"
)

const (
	btnPressMask     = xproto.EventMaskButtonPress
	btnReleaseMask   = xproto.EventMaskButtonRelease
	ptrMotionMask    = xproto.EventMaskPointerMotion
	btn1MotionMask   = xproto.EventMaskButton1Motion
	btn2MotionMask   = xproto.EventMaskButton2Motion
	btn3MotionMask   = xproto.EventMaskButton3Motion
	exposureMask     = xproto.EventMaskExposure
	structNotifyMask = xproto.EventMaskStructureNotify
	keyPressMask     = xproto.EventMaskKeyPress
	keyReleaseMask   = xproto.EventMaskKeyRelease
	enterWinMask     = xproto.EventMaskEnterWindow
	leaveWinMask     = xproto.EventMaskLeaveWindow
	focusChangeMask  = xproto.EventMaskFocusChange
)

func SrvInit() {
	display, err := xgbutil.NewConn()
	if err != nil {
		log.Fatal(err)
	}
	keybind.Initialize(display)
	mousebind.Initialize(display)

	win, err := xwindow.Generate(display)
	if err != nil {
		log.Fatalf("Could not generate a new window X id: %s", err)
	}
	win.Create(display.RootWin(), 0, 0, 500, 500, xproto.CwBackPixel, 0xffffffff)

	win.Listen(ptrMotionMask, exposureMask, structNotifyMask, focusChangeMask,
		btnPressMask, btnReleaseMask, btn1MotionMask, btn2MotionMask, btn3MotionMask,
		keyPressMask, keyReleaseMask, enterWinMask, leaveWinMask)
	win.Map()
	xevent.KeyPressFun(
		func(display *xgbutil.XUtil, e xevent.KeyPressEvent) {
			// keybind.LookupString does the magic of implementing parts of
			// the X Keyboard Encoding to determine an english representation
			// of the modifiers/keycode tuple.
			// N.B. It's working for me, but probably isn't 100% correct in
			// all environments yet.
			modStr := keybind.ModifierString(e.State)
			keyStr := keybind.LookupString(display, e.State, e.Detail)
			if len(modStr) > 0 {
				log.Printf("Key: %s-%s\n", modStr, keyStr)
			} else {
				log.Println("Key:", keyStr)
			}

			if keybind.KeyMatch(display, "q", e.State, e.Detail) {
				if e.State&xproto.ModMaskControl > 0 {
					log.Println("Control-q detected. Quitting...")
					xevent.Quit(display)
				}
			}
		}).Connect(display, win.Id)

	xevent.ButtonPressFun(func(display *xgbutil.XUtil, e xevent.ButtonPressEvent) {
		mousebind.DeduceButtonInfo(e.State, e.Detail)
	}).Connect(display, win.Id)

	log.Println("Program initialized. Start pressing keys!")
	xevent.Main(display)
}
