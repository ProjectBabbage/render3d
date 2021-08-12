package datatypes

type Material struct {
	A float64 // alpha
	// the material refraction index, supposing N1 is the air refraction index
	N2 float64
	// Ambient, Diffuse, Specular, Reflected, Refracted
	Ka, Kd, Ks, Kr, Kra Col
}
