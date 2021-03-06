package stored

import (
	. "render3d/datatypes"
	"fmt"
	"os"
	"path"
)

// FILEPATH RESOLUTION
func get_assets_files_path() string {
	fp, _ := os.Getwd()
	base := path.Base(fp)
	// in case the working directory is not the root of the project
	// but from a direct subdirectory
	if base != "render3d" {
		fp = path.Dir(fp) // path/render3d/subdirectory becomes path/render3d
	}
	return fmt.Sprintf("%s/assets/stl/files/", fp)
}

var FilesPath = get_assets_files_path()

// LIGHTS
var LightLeftBehind = Light{
	Vector{-10, -15, -10},
	IsoCol(1), IsoCol(1), IsoCol(1),
}

// COLORS
var special_red = Col{R: 252, G: 47, B: 47}
var special_blue = Col{R: 94, G: 94, B: 249}

var SPECIAL_RED_DIFFUSE = Material{
	A:  20,
	Ka: special_red.DilateColor(0.10),
	Kd: special_red.DilateColor(0.6),
	Ks: special_red.DilateColor(0.3),
}

var SPECIAL_BLUE_DIFFUSE = Material{
	A:  20,
	Ka: special_blue.DilateColor(0.15),
	Kd: special_blue.DilateColor(0.6),
	Ks: special_blue.DilateColor(0.3),
}

// MATERIAL
var WHITE_DIFFUSE_MEDIUM = Material{
	A:  1,
	Ka: IsoCol(10), Kd: IsoCol(180), Ks: IsoCol(10),
}

var BLACK_DIFFUSE = Material{}

var GLASS = Material{
	N2:  1.524, // indice de réfraction du verre
	Kr:  IsoCol(0.2),
	Kra: IsoCol(0.6),
	Ka:  IsoCol(10),
}
