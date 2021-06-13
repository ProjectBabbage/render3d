package render

import (
	"broengine/config"
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// Render the screen
func Rendering(screen *Screen) {
	// INIT SDL
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow(
		"test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		config.PixelsWidth, config.PixelsHeight,
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	// GET SURFACE
	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	// SET THE SCREEN ON THE SURFACE AND UPDATE
	var i int
	var j int
	for i = 0; i < config.PixelsWidth; i++ {
		for j = 0; j < config.PixelsHeight; j++ {
			var pixel = screen.Pixels[i][j]
			fmt.Println(pixel)
			surface.Set(i, j, pixel)
		}
	}
	window.UpdateSurface()

	// WAIT FOR AN EXIT SIGNAL
	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}
	}
}
