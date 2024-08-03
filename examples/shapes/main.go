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
		Width:  1024,
		Height: 768,
	})
	if err != nil {
		panic(err)
	}
	defer win.Destroy()

	// For measuring FPS
	frames := 0
	second := time.Tick(time.Second)

	// Shapes
	rect := tgl.NewRect(
		120, 90,
		tgl.Vec{X: 200, Y: 200},
		tgl.WithStyle(tgl.Style{Colour: color.RGBA{0, 0, 255, 1}, Thickness: 0}),
	)
	rect2 := tgl.NewRect(
		120, 90,
		tgl.Vec{X: 200, Y: 200},
		tgl.WithStyle(tgl.Style{Colour: color.RGBA{255, 0, 0, 1}, Thickness: 4}),
	)
	circle := tgl.NewCircle(
		100,
		tgl.Vec{X: 500, Y: 200},
		tgl.WithStyle(tgl.Style{Colour: color.RGBA{255, 0, 0, 1}, Thickness: 10}),
	)
	txt := tgl.NewText("Hello there", tgl.Vec{X: 100, Y: 600}).
		SetColour(color.RGBA{255, 255, 255, 255})

	// Keybinds
	win.RegisterKeybind(tgl.KeyEscape, func() { win.Quit() })
	win.RegisterKeybind(tgl.KeyLCtrl, func() { win.Quit() })
	win.RegisterKeybind(tgl.KeyE, func() { rect.Move(tgl.Vec{X: 2, Y: 2}) })

	for win.IsRunning() {
		win.Framebuffer.SetBackground(color.RGBA{39, 45, 53, 255})

		// Draw shapes
		win.Draw(rect)
		win.Draw(rect2)
		win.Draw(circle)
		win.Draw(txt)

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