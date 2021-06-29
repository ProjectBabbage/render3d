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
}
