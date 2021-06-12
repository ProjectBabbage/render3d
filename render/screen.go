package render

import "github.com/veandco/go-sdl2/sdl"

type Color struct {
	R, G, B, A uint8
}

type Screen struct {
	window   sdl.Window
	renderer sdl.Renderer
	pixels   [][]Color
}

func (s Screen) Init() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	s.window = *window
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

}

// Draw each pixels
func (s Screen) Render() {

	s.renderer.Present()
}
