package render

import (
	"image/color"

	"github.com/veandco/go-sdl2/sdl"
)

// TODO : prendre un screen
func Rendering() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	surface.FillRect(nil, 0)
	rect := sdl.Rect{X: 0, Y: 0, W: 800, H: 600}
	surface.FillRect(&rect, 0xffffffff)

	// TODO : utiliser ce screen ici
	var i int = 0
	var j int = 0
	for i = 0; i < 100; i++ {
		for j = 0; j < 100; j++ {
			// rend.SetDrawColor(uint8(2*i), uint8(2*j), 0, 0)
			surface.Set(300+i, 300+j, color.RGBA{uint8(2 * i), uint8(2 * j), 0, 0})
			window.UpdateSurface()
		}
	}

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
