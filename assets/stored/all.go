package stored

import (
	"broengine/assets/stl"
	. "broengine/config"
	. "broengine/datatypes"
)

var mat0 = Material{
	A:  1,
	Ka: IsoCol(2), Kd: IsoCol(10), Ks: IsoCol(1),
}

var glass = Material{
	A:   1,
	N2:  1.5, // indice de réfraction du verre
	Kra: IsoCol(0.7),
}

var paperMat = Material{
	A:  1,
	Ka: IsoCol(5), Kd: IsoCol(5), Ks: IsoCol(1),
}

func STrueSpherePlane() (Scene, Config) {
	conf := NewConfig(Config{
		Msaa:       4,
		MaxBounces: 10,
	})

	s := NewSphere(Vector{}, 1, glass)
	sphere := Object{[]Surface{&s}}
	plane := stl.Parse(FilesPath+"plane.stl", mat0)
	plane.Rotate(XAxis, 10)
	plane.Dilate(10)

	sphere.Translate(Vector{0, 0, 10})
	plane.Translate(Vector{0, 0, 15})

	scene := NewEmptyScene()
	scene.AddObjects(sphere, plane)
	scene.AddLights(LightLeftBehind)

	return scene, conf
}

func SRefractionExperiment() (Scene, Config) {
	conf := NewConfig(Config{
		PixelsX:    700,
		PixelsY:    700,
		Msaa:       4,
		MaxBounces: 10,
	})

	t1 := NewTriangle(Vector{0, 0, 0}, Vector{1, 0, 0}, Vector{0, 1, 0}, Vector{0, 0, -1}, glass)
	triangle := Object{[]Surface{&t1}}
	triangle.Rotate(XAxis, 45)
	triangle.Translate(Vector{Z: 5})

	t2 := NewTriangle(Vector{0, 0, 0}, Vector{1, 0, 0}, Vector{0, 1, 0}, Vector{0, 0, -1}, paperMat)
	paper := Object{[]Surface{&t2}}
	paper.Translate(Vector{X: -0.1, Y: 0.1, Z: 6})
	paper.Rotate(ZAxis, 25)

	plane := stl.Parse(FilesPath+"plane.stl", mat0)
	plane.Rotate(XAxis, 10)
	plane.Dilate(10)

	plane.Translate(Vector{0, 0, 15})

	scene := NewEmptyScene()
	scene.AddObjects(triangle, plane, paper)
	scene.AddLights(LightLeftBehind)

	return scene, conf
}

// var matSphere = Material{
// 	A:  30,
// 	Ka: c, Kd: c, Ks: c, Kr: IsoCol(150),
// }

var c0 = Col{}
var ca1 = Col{30, 50, 40}

// var ca2 = Col{50, 30, 40}
// var ca3 = Col{30, 40, 50}
// var ca1b = ca1.DilateColor(.5)
// var ca2b = ca2.DilateColor(.5)
// var ca3b = ca3.DilateColor(.5)

// var cd1 = Col{100, 250, 100}
var cd2 = Col{200, 100, 100}

// var cd3 = Col{100, 200, 100}
// var cd1b = cd1.DilateColor(.5)
// var cd2b = cd2.DilateColor(.5)
// var cd3b = cd3.DilateColor(.5)

// var cs1 = Col{70, 20, 110}
// var cs2 = Col{70, 60, 20}
// var cs3 = Col{20, 90, 90}

// var L0 = Light{Vector{0, -50, -50}, ca1, cd1, c0}
var L1 = Light{Vector{0, -50, -50}, ca1, cd2, c0}

// var L2 = Light{Vector{0, 0, -50}, ca1b, cd1, c0}
// var L3 = Light{Vector{0, 0, 10}, ca2, cd2, c0}
// var L4 = Light{Vector{-2, -2, 5}, ca2, cd1, cs1}
// var L5 = Light{Vector{0, 0, 0}, ca3, cd3, c0}
// var L6 = Light{Vector{-8, 0, 0}, ca3, cd3, cs2}
// var L7 = Light{Vector{0, -50, -50}, ca3, cd2, cs3}
// var L8 = Light{Vector{-5, -5, 0}, ca2b, cd2b, cs1}
// var L9 = Light{Vector{3, 3, 0}, ca3b, cd3b, cs2}

// var Lstandard = Light{Vector{0, -50, -50}, IsoCol(30), IsoCol(150), IsoCol(70)}
// var Lstandard2 = Light{Vector{-2, -2, 5}, IsoCol(80), IsoCol(250), c0}

func SSphere() (Scene, Config) {
	conf := NewConfig(Config{
		PixelsX: 500,
		PixelsY: 500,
	})

	sphere := stl.Parse(FilesPath+"sphere.stl", mat0)

	sphere.Translate(Vector{0, 0, 12})

	scene := NewEmptyScene()
	scene.AddObjects(sphere)
	scene.AddLights(L1)

	return scene, conf
}

// func STwoSpheresOnePlane() (Scene, Config) {
// 	conf := NewConfig(Config{
// 		MaxBounces: 2,
// 	})

// 	mat1 := Material{
// 		A:  1,
// 		Ka: c0, Kd: c, Ks: c0, Kr: IsoCol(150),
// 	}
// 	mat2 := Material{
// 		A:  1,
// 		Ka: c, Kd: c, Ks: c0, Kr: IsoCol(200),
// 	}

// 	sphere := stl.Parse(FilesPath+"sphere_high_definition.stl", mat1)
// 	sphere2 := stl.Parse(FilesPath+"sphere.stl", mat1)
// 	plane := stl.Parse(FilesPath+"plane.stl", mat2)

// 	sphere.Translate(Vector{1, 0, 8})
// 	sphere2.Translate(Vector{-0.5, -1, 9})
// 	plane.Translate(Vector{0, 1, 7})

// 	scene := NewEmptyScene()
// 	scene.AddObjects(sphere, sphere2, plane)
// 	scene.AddLights(L4)

// 	return scene, conf
// }

// func SSpherePlaneShadow() (Scene, Config) {
// 	conf := NewConfig(Config{})

// 	sphere := stl.Parse(FilesPath+"sphere.stl", mat0)
// 	plane := stl.Parse(FilesPath+"plane.stl", mat0)

// 	sphere.Translate(Vector{0, 0, 12})
// 	plane.Translate(Vector{0, 1, 7})

// 	scene := NewEmptyScene()
// 	scene.AddObjects(sphere, plane)
// 	scene.AddLights(L3)

// 	return scene, conf
// }

// func SCubeRotated() (Scene, Config) {
// 	conf := NewConfig(Config{})

// 	cube_rotated := stl.Parse(FilesPath+"cube_rotated.stl", mat0)

// 	cube_rotated.Translate(Vector{2, 0, 15})
// 	scene := NewEmptyScene()
// 	scene.AddObjects(cube_rotated)
// 	scene.AddLights(L1)

// 	return scene, conf
// }

// func SCubeManuallyRotated() (Scene, Config) {
// 	conf := NewConfig(Config{})

// 	cube := stl.Parse(FilesPath+"cube.stl", mat0)
// 	cube.Rotate(XAxis, 20)
// 	cube.Rotate(YAxis, 20)
// 	cube.Translate(Vector{2, 0, 15})
// 	scene := NewEmptyScene()
// 	scene.AddObjects(cube)
// 	scene.AddLights(L1)

// 	return scene, conf
// }

// func SFaces(listIndex ...string) (Scene, Config) {
// 	conf := NewConfig(Config{})

// 	var objects = []Object{}

// 	for _, face := range listIndex {
// 		filename := fmt.Sprintf(FilesPath+"faces/%s.stl", face)
// 		o := stl.Parse(filename, mat0)
// 		objects = append(objects, o)
// 	}
// 	scene := NewEmptyScene()
// 	scene.AddObjects(objects...)
// 	scene.AddLights(L1)
// 	scene.TranslateObjects(Vector{0, 0, 20})

// 	return scene, conf
// }

// func SSimpleTriangle() (Scene, Config) {
// 	conf := NewConfig(Config{})
// 	var distance float64 = 100

// 	p1 := Vector{0, 0, distance}
// 	p2 := Vector{0, 25, distance}
// 	p3 := Vector{25, 0, distance}

// 	newTriangle := NewTriangle(p1, p2, p3, Vector{0, 0, 0}, mat0)
// 	newTriangle.RecomputeNormal()
// 	o := Object{[]Surface{&newTriangle}}
// 	scene := NewEmptyScene()
// 	scene.AddObjects(o)
// 	scene.AddLights(L2)

// 	// scene.TranslateObjects(Vector{4, -4, 40})

// 	return scene, conf
// }

// func STwoTrianglesPlane() (Scene, Config) {
// 	conf := NewConfig(Config{})

// 	triangles := stl.Parse(FilesPath+"two_triangles.stl", mat0)
// 	plane := stl.Parse(FilesPath+"plane.stl", mat0)
// 	// plane.Rotate(XAxis, 90)
// 	triangles.Rotate(YAxis, -45)

// 	triangles.Translate(Vector{0, 0, 4})
// 	plane.Translate(Vector{0, 2, 10})

// 	scene := NewEmptyScene()
// 	scene.AddObjects(triangles, plane)
// 	scene.AddLights(L5)
// 	scene.Print()

// 	return scene, conf
// }

// func STwoTrianglesPlane2() (Scene, Config) {
// 	conf := NewConfig(Config{})

// 	q1 := Vector{0.22975452523155737, 0, 3.026750959038924}
// 	q2 := Vector{-0.3162252123460546, 0, 3.051315849298438}
// 	x := -0.11487165708492686
// 	y := 0.49999698996543884
// 	z := 3.141625397815853
// 	triangle1 := NewTriangle(q1, q2, Vector{x, y, z}, Vector{}, mat0)
// 	triangle1.RecomputeNormal()
// 	triangle2 := NewTriangle(q1, Vector{x, -y, z}, q2, Vector{}, mat0)
// 	triangle2.RecomputeNormal()
// 	triangles := Object{[]Surface{&triangle1, &triangle2}}
// 	plane := stl.Parse(FilesPath+"plane.stl", mat0)
// 	// plane.Rotate(XAxis, 90)
// 	// triangles.Rotate(YAxis, -45)

// 	triangles.Translate(Vector{0, 0, 0})
// 	plane.Translate(Vector{0, 2, 10})

// 	scene := NewEmptyScene()
// 	scene.AddObjects(triangles, plane)
// 	scene.AddLights(L5)
// 	scene.Print()

// 	return scene, conf
// }

// func STrueSphere() (Scene, Config) {
// 	conf := NewConfig(Config{Msaa: 1})

// 	s := NewSphere(Vector{}, 1, matSphere)
// 	sphere := Object{[]Surface{&s}}

// 	sphere.Translate(Vector{0, 0, 10})

// 	scene := NewEmptyScene()
// 	scene.AddObjects(sphere)
// 	scene.AddLights(L7)

// 	return scene, conf
// }

// func STrueSphereInside() (Scene, Config) {
// 	conf := NewConfig(Config{Msaa: 3})

// 	mat := Material{
// 		A:  30,
// 		Ka: Col{0, 1, 1}, Kd: c, Ks: c,
// 	}

// 	s1 := NewSphere(Vector{}, 25, mat)
// 	s2 := NewSphere(Vector{}, 1, matSphere)

// 	sphere1 := Object{[]Surface{&s1}}
// 	sphere2 := Object{[]Surface{&s2}}
// 	sphere2.Translate(Vector{0, 0, 10})

// 	scene := NewEmptyScene()
// 	scene.AddObjects(sphere1, sphere2)
// 	var L1 = Light{Vector{-5, -5, 0}, RedCol(20), RedCol(150), c0}
// 	var L2 = Light{Vector{3, 3, 0}, Col{50, 100, 50}, BlueCol(100), c0}
// 	scene.AddLights(L1, L2)

// 	return scene, conf
// }

// func STrueSphereInsideNonIsoChannels() (Scene, Config) {
// 	conf := NewConfig(Config{Msaa: 3})

// 	mat := Material{
// 		A:  30,
// 		Ka: Col{0, 1, 1}, Kd: c, Ks: c,
// 	}

// 	s1 := NewSphere(Vector{}, 25, mat)
// 	s2 := NewSphere(Vector{}, 1, matSphere)

// 	sphere1 := Object{[]Surface{&s1}}
// 	sphere2 := Object{[]Surface{&s2}}
// 	sphere2.Translate(Vector{0, 0, 10})

// 	scene := NewEmptyScene()
// 	scene.AddObjects(sphere1, sphere2)
// 	var L1 = Light{Vector{-5, -5, 0}, RedCol(20), RedCol(150), c0}
// 	var L2 = Light{Vector{3, 3, 0}, Col{50, 100, 50}, BlueCol(100), c0}
// 	scene.AddLights(L1, L2)

// 	return scene, conf
// }
