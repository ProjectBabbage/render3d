package render

import (
	"broengine/config"
	"broengine/datatypes"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

func RenderScreen(screen Screen) {
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

	if config.RenderBackend != "gpu" {
		fmt.Println("Using the CPU")
		// SET THE SCREEN ON THE SURFACE AND UPDATE
		for I := 0; I <= config.PixelsX; I++ {
			for J := 0; J <= config.PixelsY; J++ {
				var pixelColor = screen.Pixels[I][J]
				surface.Set(I, J, pixelColor)
			}
		}
		window.UpdateSurface()

	} else {
		fmt.Println("Using the GPU")
		rend, err := window.GetRenderer()
		if err != nil {
			fmt.Println("Could not find a renderer ", err)
		}

		// SET THE SCREEN ON THE SURFACE AND UPDATE
		for I := 0; I <= config.PixelsX; I++ {
			for J := 0; J <= config.PixelsY; J++ {
				r, g, b, a := screen.Pixels[I][J].RGBA()
				rend.SetDrawColor(uint8(r), uint8(g), uint8(b), uint8(a))
				rend.DrawPoint(int32(I), int32(J))
			}
		}
		rend.Present()
	}

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

func Render(scene datatypes.Scene) {
	screen := CastAll(scene)
	RenderScreen(screen)
}
