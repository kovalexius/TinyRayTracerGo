package geometry

import (
	"math"
)

type Sphere struct {
	Center Vec3f
	Radius float64
	Material Material
}

func NewSphere(center Vec3f, radius float64, material Material) Sphere {
	sphere := Sphere{Center: center, Radius: radius, Material: material}
	return sphere
}

func (sphere Sphere) RayIntersect (orig *Vec3f, dir *Vec3f) (result bool, t0 float64) {
	t0 = 0.0
	result = false
	L := sphere.Center.Sub(orig)
	var tca float64 = L.ScalarMul(dir)
	var d2 float64 = L.ScalarMul(&L) - tca*tca
	if d2 > sphere.Radius * sphere.Radius {
		return
	}
	var thc float64 = math.Sqrt(sphere.Radius*sphere.Radius - d2)
	t0 = tca - thc
	var t1 float64 = tca + thc
	if t0 < 0 {
		t0 = t1
	}
	if t0 < 0 {
		return
	}
	return true, t0
}




