package render

import (
	"broengine/config"

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
		config.PixelsX+1, config.PixelsY+1,
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
	for I := 0; I <= config.PixelsX; I++ {
		for J := 0; J <= config.PixelsY; J++ {
			var pixelColor = screen.Pixels[I][J]
			surface.Set(I, J, pixelColor)
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
