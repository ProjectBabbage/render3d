package render

import "github.com/veandco/go-sdl2/sdl"

// TODO : prendre un screen
func render() {
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

	rect := sdl.Rect{0, 0, 800, 600}
	surface.FillRect(&rect, 0x00000000)

	window.UpdateSurface()

	rend, err := window.GetRenderer()
	rend.SetDrawColor(255, 255, 0, 0)
	// TODO : utiliser ce screen ici
	// var i int32 = 0
	// var j int32 = 0
	// for i = 0; i < 100; i++ {
	// 	for j = 0; j < 100; j++ {
	// 		rend.SetDrawColor(uint8(2*i), uint8(2*j), 0, 0)
	// 		rend.DrawPoint(300+i, 300+j)

	// 	}
	// }
	rend.Present()

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
