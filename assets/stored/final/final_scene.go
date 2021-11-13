// put the scene we want to print on a poster here :)
package stored

import (
	"broengine/assets/stl"
	. "broengine/assets/stored"
	. "broengine/config"
	. "broengine/datatypes"
)

// LIGHTS

// COLORS

// MATERIAL
var NORMAL = Material{
	A:  1,
	Ka: IsoCol(10), Kd: IsoCol(180),
}

var RED_DIFFUSE_MEDIUM = Material{
	A:  20,
	Ka: RedCol(20), Kd: RedCol(80), Ks: IsoCol(50),
}

var BLUE_DIFFUSE_MEDIUM = Material{
	A:  20,
	Ka: BlueCol(20), Kd: BlueCol(80), Ks: IsoCol(50),
}

var ULTRA_WHITE = Material{
	A:  20,
	Ka: IsoCol(10),
	Kd: IsoCol(210),
	Ks: IsoCol(20),
}

var MIROIR = Material{
	Ka: IsoCol(10),
	Kr: IsoCol(0.8),
	Kd: IsoCol(20),
}

// OBJECTS
var s = NewSphere(Vector{}, 1, SPECIAL_RED_DIFFUSE)
var sphere_red = Object{[]Surface{&s}}
var s1 = NewSphere(Vector{}, 1, SPECIAL_BLUE_DIFFUSE)
var sphere_blue = Object{[]Surface{&s1}}
var s2 = NewSphere(Vector{}, 1, ULTRA_WHITE)
var sphere_white = Object{[]Surface{&s2}}
var plane = stl.Parse(FilesPath+"plane.stl", NORMAL)
var miroir = stl.Parse(FilesPath+"plane.stl", MIROIR)
var miroir_edge = stl.Parse(FilesPath+"plane.stl", WHITE_DIFFUSE_MEDIUM)
var cube_top_left = stl.Parse(FilesPath+"cube.stl", GLASS)
var cube_front = stl.Parse(FilesPath+"cube.stl", GLASS)
var cube_rotated = stl.Parse(FilesPath+"cube.stl", GLASS)

var all_objects = []Object{
	plane,
	miroir,
	// miroir_edge,
	sphere_red,
	sphere_blue,
	sphere_white,
	// cube_top_left,
	cube_front,
	// cube_rotated,
}

// CONFIG
var multiply_resolution = 1
var conf = NewConfig(Config{
	D:          2, // screen Z distance to the Eye at (0,0,0)
	Msaa:       3,
	MaxBounces: 5,
	ScreenX:    1920. / 1080.,
	PixelsX:    1920 * multiply_resolution,
	PixelsY:    1080 * multiply_resolution,
	SaveAsPNG:  true,
})

// The scene moves all the OBJECTS at the right position
func SFinal() (Scene, Config) {
	plane.Dilate(10)
	plane.Translate(Vector{0, 0.5, 0})

	var miroir_dilate float64 = 0.3
	var miroir_x_rotate float64 = 90.
	var miroir_y_rotate float64 = 23.
	var miroir_translate Vector = Vector{3.5, 0, 13}
	miroir.Dilate(miroir_dilate)
	miroir.Rotate(XAxis, miroir_x_rotate)
	miroir.Rotate(YAxis, miroir_y_rotate)
	miroir.Translate(miroir_translate)
	miroir_edge.Dilate(miroir_dilate * 1.02)
	miroir_edge.Rotate(XAxis, miroir_x_rotate)
	miroir_edge.Rotate(YAxis, miroir_y_rotate)
	miroir_edge.Translate(miroir_translate.Add(Vector{X: 0.0001, Y: 0, Z: 0.0001}))

	sphere_red.Translate(Vector{0, -0.5, 8})

	sphere_blue.Dilate(0.2)
	sphere_blue.Translate(Vector{-0.4, 0.3, 4.5})

	sphere_white.Dilate(0.2)
	sphere_white.Translate(Vector{-0.6, -0.6, 4.5})

	cube_top_left.Dilate(0.2)
	cube_top_left.Translate(Vector{-0.6, -0.6, 4})

	cube_front.Dilate(0.2)
	cube_front.Translate(Vector{0, 0.3, 3})

	cube_rotated.Rotate(XAxis, 30)
	cube_rotated.Rotate(YAxis, 50)
	cube_rotated.Dilate(0.2)
	cube_rotated.Translate(Vector{0.45, -0.65, 6})

	scene := NewEmptyScene()

	scene.AddObjects(all_objects...)
	scene.AddLights(LightLeftBehind)

	return scene, conf
}
