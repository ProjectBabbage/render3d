// put the scene we want to print on a poster here :)
package dota

import (
	"broengine/assets/stl"
	. "broengine/assets/stored"
	. "broengine/config"
	. "broengine/datatypes"
)

// LIGHT
var OneMoreLight = Light{
	Vector{-15, -15, 90},
	IsoCol(1), IsoCol(1), IsoCol(1),
}

// COLORS
var special_red2 = Col{R: 120, G: 47, B: 47}

// MATERIAL
var GLASS2 = Material{
	N2:  1.524, // indice de réfraction du verre
	Kra: IsoCol(0.6),
	Kr:  IsoCol(0.2),
	Kd:  IsoCol(0),
	Ka:  IsoCol(10),
}

var SPECIAL_RED2_DIFFUSE = Material{
	A:  20,
	Kd: special_red2.DilateColor(0.6),
	Ks: special_red2.DilateColor(0.3),
}

// OBJECTS
var ls1 = NewSphere(Vector{}, 1, GLASS)
var lsphere1 = Object{[]Surface{&ls1}}

// var ls2 = NewSphere(Vector{}, 1, SPECIAL_RED2_DIFFUSE)
// var lsphere2 = Object{[]Surface{&ls2}}
// var ls3 = NewSphere(Vector{}, 1, SPECIAL_RED2_DIFFUSE)
// var lsphere3 = Object{[]Surface{&ls3}}
// var ls4 = NewSphere(Vector{}, 1, SPECIAL_RED_DIFFUSE)
// var lsphere4 = Object{[]Surface{&ls4}}

var plane = stl.Parse(FilesPath+"plane.stl", GLASS)

var logo = stl.Parse(FilesPath+"logo.stl", SPECIAL_RED_DIFFUSE)

// var base1 = stl.Parse(FilesPath+"base1.stl", SPECIAL_RED_DIFFUSE)
// var base2 = stl.Parse(FilesPath+"base2.stl", SPECIAL_RED_DIFFUSE)

// var tower = stl.Parse(FilesPath+"tower.stl", SPECIAL_RED2_DIFFUSE)
// var river = stl.Parse(FilesPath+"river.stl", SPECIAL_RED_DIFFUSE)

var lall_objects = []Object{
	plane,
	lsphere1,
	// lsphere2,
	// lsphere3,
	// lsphere4,
	// cube,
	logo,
	// base1,
	// base2,
	// tower,
	// river,
}

// CONFIG
var multiply_resolution = .5
var lconf = NewConfig(Config{
	D:          2, // screen Z distance to the Eye at (0,0,0)
	Msaa:       3,
	MaxBounces: 3,
	ScreenX:    1920. / 1080.,
	PixelsX:    int(1920 * multiply_resolution),
	PixelsY:    int(1080 * multiply_resolution),
	SaveAsPNG:  true,
})

// The scene moves all the OBJECTS at the right position
func SDota() (Scene, Config) {
	// tower.Rotate(XAxis, 180)
	// tower.Translate(Vector{0, 0.5, 5})

	// river.Rotate(XAxis, 180)
	// river.Rotate(ZAxis, 45)
	// river.Dilate(1.2)
	// river.Translate(Vector{0, 0, 10})

	// base1.Rotate(XAxis, 180)
	// base1.Rotate(ZAxis, 40)
	// base1.Translate(Vector{1.8, -1.6, 10})

	// base2.Rotate(XAxis, 180)
	// base2.Rotate(ZAxis, 40)
	// base2.Translate(Vector{-1.6, 1.5, 10})

	// paved.Dilate(3.5)
	// paved.Rotate(XAxis, 180)
	// paved.Translate(Vector{0, 0, 15})

	lsphere1.Dilate(1.5)
	lsphere1.Translate(Vector{6, 6, 90})

	plane.Rotate(XAxis, 90)
	plane.Dilate(3.5)
	plane.Translate(Vector{0, 0, 99.5})

	// lsphere2.Dilate(5)
	// lsphere2.Translate(Vector{6, -6, 30})

	// lsphere3.Dilate(5)
	// lsphere3.Translate(Vector{-6, 6, 30})

	// lsphere4.Dilate(30)
	// lsphere4.Translate(Vector{4, 4, 130})

	logo.Rotate(XAxis, 180)
	logo.Dilate(20)
	logo.Translate(Vector{0, 0, 100})

	scene := NewEmptyScene()

	scene.AddObjects(lall_objects...)
	scene.AddLights(OneMoreLight)

	return scene, lconf
}
