package main

import (
	"fmt"
	"image/color"
	"time"

	tgl "github.com/zac460/turdgl"
)

func main() {
	win, err := tgl.NewWindow(tgl.WindowCfg{
		Title:  "Basic Shapes Example",
		Width:  1200,
		Height: 768,
	})
	if err != nil {
		panic(err)
	}
	defer win.Destroy()

	// For measuring FPS
	frames := 0
	second := time.Tick(time.Second)

	// Shapes can be filled or have an outline of specified thickness
	rectSolid := tgl.NewRect(
		120, 90,
		tgl.Vec{X: 50, Y: 50},
		tgl.WithStyle(tgl.Style{Colour: color.RGBA{0, 0, 255, 255}, Thickness: 0}),
	)
	rectOutline := tgl.NewRect(
		120, 90,
		tgl.Vec{X: 50, Y: 50},
		tgl.WithStyle(tgl.Style{Colour: color.RGBA{255, 0, 0, 255}, Thickness: 2}),
	)
	curvedRect := tgl.NewCurvedRect(
		120, 90, 12,
		tgl.Vec{X: 50, Y: 200},
		tgl.WithStyle(tgl.Style{Colour: color.RGBA{240, 170, 90, 255}, Thickness: 8}),
	)
	bgRect := tgl.NewRect(
		120, 90,
		tgl.Vec{X: 250, Y: 200},
		tgl.WithStyle(tgl.Style{Colour: color.RGBA{90, 65, 48, 255}, Thickness: 8}),
	)
	fgRect := tgl.NewRect(
		120, 90,
		tgl.Vec{X: 280, Y: 230},
		tgl.WithStyle(tgl.Style{Colour: color.RGBA{175, 136, 90, 255}, Thickness: 8}),
	)

	// Buttons can be constructed from shapes
	styleUnpressed := tgl.Style{Colour: color.RGBA{80, 0, 0, 255}, Thickness: 30}
	stylePressed := tgl.Style{Colour: color.RGBA{255, 0, 0, 255}, Thickness: 0}
	c := tgl.NewCircle(100, tgl.Vec{X: 300, Y: 100}, tgl.WithStyle(styleUnpressed))
	circleButton := tgl.NewButton(c).SetText("Press me")
	circleButton.Label.SetSize(16)
	circleButton.Label.SetColour(color.White)
	circleButton.SetCallback(func(m tgl.MouseState) {
		if m == tgl.LeftClick {
			c.SetStyle(stylePressed)
			circleButton.SetText("Pressed!")
		} else {
			c.SetStyle(styleUnpressed)
			circleButton.SetText("Press me")
		}
	})
	circleButton.Label.SetAlignment(tgl.AlignBottomCentre)

	// More complex shapes with limited feature sets can also be created
	triangle := tgl.NewTriangle(
		tgl.Vec{X: 400, Y: 50},
		tgl.Vec{X: 580, Y: 120},
		tgl.Vec{X: 450, Y: 130},
	).SetStyle(tgl.Style{Colour: color.RGBA{100, 10, 100, 255}})

	txt := tgl.NewText("Hello there", tgl.Vec{X: 800, Y: 80}).
		SetColour(color.RGBA{255, 255, 255, 255})

	// Register window-level keybinds
	win.RegisterKeybind(tgl.KeyEscape, func() { win.Quit() })
	win.RegisterKeybind(tgl.KeyLCtrl, func() { win.Quit() })
	win.RegisterKeybind(tgl.KeyE, func() { rectSolid.Move(tgl.Vec{X: 2, Y: 2}) })

	for win.IsRunning() {
		win.SetBackground(color.RGBA{39, 45, 53, 255})

		// Draw foreground shapes
		for _, shape := range []tgl.Drawable{
			rectSolid,
			rectOutline,
			curvedRect,
			triangle,
			txt,
			circleButton,
			fgRect,
		} {
			win.Draw(shape)
		}

		// Shapes drawn to the background appear behind foreground shapes
		win.DrawBackground(bgRect)

		win.SetTitle(fmt.Sprint(win.MouseLocation()))

		circleButton.Update(win)
		win.Update()

		// Count FPS
		frames++
		select {
		case <-second:
			txt.SetText(fmt.Sprintf("FPS: %d", frames))
			frames = 0
		default:
		}
	}
}
