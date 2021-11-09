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
	Ka: IsoCol(50), Kd: IsoCol(50),
}

var RED_DIFFUSE_MEDIUM = Material{
	A:  20,
	Ka: RedCol(20), Kd: RedCol(80), Ks: IsoCol(50),
}

var BLUE_DIFFUSE_MEDIUM = Material{
	A:  20,
	Ka: BlueCol(20), Kd: BlueCol(80), Ks: IsoCol(50),
}

// OBJECTS
var s = NewSphere(Vector{}, 1, SPECIAL_RED_DIFFUSE)
var sphere = Object{[]Surface{&s}}
var s1 = NewSphere(Vector{}, 1, SPECIAL_BLUE_DIFFUSE)
var sphere_blue = Object{[]Surface{&s1}}
var plane = stl.Parse(FilesPath+"plane.stl", WHITE_DIFFUSE_MEDIUM)
var cube = stl.Parse(FilesPath+"cube.stl", GLASS)
var cube_front = stl.Parse(FilesPath+"cube.stl", GLASS)
var cube_rotated = stl.Parse(FilesPath+"cube_rotated.stl", GLASS)

var all_objects = []Object{
	plane,
	sphere,
	sphere_blue,
	cube,
	cube_front,
	cube_rotated,
}

// CONFIG
var multiply_resolution = 1
var conf = NewConfig(Config{
	D:          2, // screen Z distance to the Eye at (0,0,0)
	Msaa:       3,
	MaxBounces: 10,
	ScreenX:    1920. / 1080.,
	PixelsX:    1920 * multiply_resolution,
	PixelsY:    1080 * multiply_resolution,
	SaveAsPNG:  true,
})

// The scene moves all the OBJECTS at the right position
func SFinal() (Scene, Config) {
	plane.Dilate(10)
	plane.Translate(Vector{0, 0.5, 0})
	sphere.Translate(Vector{0, -0.5, 8})
	sphere_blue.Dilate(0.2)
	sphere_blue.Translate(Vector{-0.4, 0.3, 4.5})
	cube.Dilate(0.2)
	cube.Translate(Vector{-0.6, -0.6, 4.5})
	cube_front.Dilate(0.2)
	cube_front.Translate(Vector{0, 0.3, 4})
	cube_rotated.Dilate(0.2)
	cube_rotated.Translate(Vector{0.3, -0.5, 6})

	scene := NewEmptyScene()

	scene.AddObjects(all_objects...)
	scene.AddLights(LightLeftBehind)

	return scene, conf
}
