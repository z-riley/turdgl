package main

import (
	"image/color"
	"time"

	"github.com/z-riley/turdgl"
)

func main() {
	win, err := turdgl.NewWindow(turdgl.WindowCfg{
		Title:  "Moving Snake Example",
		Width:  1024,
		Height: 768,
	})
	if err != nil {
		panic(err)
	}
	defer win.Destroy()

	snake := NewSnake(turdgl.Vec{X: 400, Y: 100})
	instruction := turdgl.NewText("Use WASD to move", turdgl.Vec{X: 10}, "../../fonts/arial.ttf")

	win.RegisterKeybind(turdgl.KeyEscape, turdgl.KeyPress, func() { win.Quit() })

	prevTime := time.Now()

	for win.IsRunning() {
		dt := time.Since(prevTime)
		prevTime = time.Now()

		// React to pressed keys
		const speed = 1000
		if win.KeyIsPressed(turdgl.KeyW) {
			snake.velocity = &turdgl.Vec{Y: -speed}
			snake.Update(dt, win.Framebuffer)
		}
		if win.KeyIsPressed(turdgl.KeyA) {
			snake.velocity = &turdgl.Vec{X: -speed}
			snake.Update(dt, win.Framebuffer)
		}
		if win.KeyIsPressed(turdgl.KeyS) {
			snake.velocity = &turdgl.Vec{Y: speed}
			snake.Update(dt, win.Framebuffer)
		}
		if win.KeyIsPressed(turdgl.KeyD) {
			snake.velocity = &turdgl.Vec{X: speed}
			snake.Update(dt, win.Framebuffer)
		}

		win.SetBackground(color.RGBA{39, 45, 53, 0})

		win.Draw(snake)
		win.Draw(instruction)

		win.Update()
	}
}
