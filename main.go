package main

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"sync/atomic"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/event"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// MainWindow - holds the app window struct and all of the clickable button actions
type MainWindow struct {
	*app.Window

	ClickMeBtnAction widget.Clickable
}

var windowSize int32

func main() {

	// create the main window
	buildWindow()
	app.Main()

}

func buildWindow() {
	atomic.AddInt32(&windowSize, +1)
	go func() {
		w := &MainWindow{}
		w.Window = app.NewWindow()
		if err := w.loop(w.Events()); err != nil {
			log.Fatal(err)
		}
		if c := atomic.AddInt32(&windowSize, -1); c == 0 {
			os.Exit(0)
		}
	}()
}

func (w *MainWindow) loop(evts <-chan event.Event) error {
	th := material.NewTheme(gofont.Collection())
	var ops op.Ops
	for {
		e := <-evts
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:

			for w.ClickMeBtnAction.Clicked() {
				fmt.Println("Opening a new window...")
				buildWindow()
			}

			gtx := layout.NewContext(&ops, e)
			globalColor := color.RGBA{255, 255, 255, 255}

			label := material.H1(th, "Testing")
			label.Color = color.RGBA{0, 0, 255, 255}
			label.Alignment = text.Middle

			btn := material.Button(th, &w.ClickMeBtnAction, "Click Me!")
			btn.Color = globalColor

			layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{
					Alignment: layout.Middle,
					Axis:      layout.Vertical,
				}.Layout(gtx,
					RigidInset(label.Layout),
					RigidInset(btn.Layout),
				)
			})

			// create the frame
			e.Frame(gtx.Ops)
		}
	}
}

// RigidInset - returns a flex child based on the passed widget
func RigidInset(w layout.Widget) layout.FlexChild {
	return layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		return layout.UniformInset(unit.Sp(5)).Layout(gtx, w)
	})
}
