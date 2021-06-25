package main

import (
	. "broengine/assets"
	"broengine/render"
)

func main() {
	// render.Render(SSphere())
	// render.Render(SSpherePlane())
	// render.Render(SSpherePlaneShadow())
	// render.Render(SCubeRotated())
	// render.Render(SCubeManuallyRotated())
	// render.Render(STwoTrianglesPlane())
	// render.Render(STwoTrianglesPlane2())
	// render.Render(STrueSphere())
	// render.Render(STrueSpherePlane())
	render.Render(STrueSphereInside())

	// render.Render(STrueSphereInside2())

	// scene := SFaces("front", "bottom", "left", "right", "front", "back")
	// scene.Print()
	// scene.TranslateObjects(Vector{3, -3, 0})
	// render.Render(scene)

	// scene := SSimpleTriangle()
	// ray := NewRay(Vector{2, 2, 0}, Vector{0, 0, 1})
	// intensity := render.Cast(ray, scene)
	// fmt.Println(intensity)

}
