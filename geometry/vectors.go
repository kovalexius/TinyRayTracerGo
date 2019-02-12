package geometry

import (
	"math"
)

type Vec2f struct {
	Data_ [2]float64
}

func (vec Vec2f) GetComponent(index int) float64 {
	if index < 2 {
		return vec.Data_[index]
	}
	panic("Vec2f size exceeded")
}

func NewVec2f(x float64, y float64) Vec2f {
	result := Vec2f{[2]float64{x, y}}
	return result
}





type Vec3f struct {
	Data_ [3]float64
}

func NewVec3f(x float64, y float64, z float64) Vec3f {
	result := Vec3f{[3]float64{x, y, z}}
	return result
}

func (vec Vec3f) GetComponent(index int) float64 {
	if index < 3 {
		return vec.Data_[index]
	}
	panic("Vec3f size exceeded")
}

func (vec Vec3f) Sub(other Vec3f) Vec3f {
	result := vec
	for i := 0; i < 3; i++ {
		result.Data_[i] -= other.Data_[i]
	}
	return result
}

func (vec Vec3f) Add(other Vec3f) Vec3f {
	result := vec
	for i := 0; i < 3; i++ {
		result.Data_[i] += other.Data_[i]
	}
	return result
}

func (vec Vec3f) ScalarMul(other Vec3f) float64 {
	result := 0.0
	for i := 0; i < 3; i++ {
		result += vec.Data_[i] * other.Data_[i]
	}
	return result
}

func (vec Vec3f) Scaling(rhs float64) Vec3f {
	var ret Vec3f
	for i := 0; i < 3; i++ {
		ret.Data_[i] = vec.Data_[i]*rhs
	}
	return ret
}

func (vec Vec3f) Norm() float64 {
	var sqrSum float64 = 0.0
	for i := 0; i < 3; i++ {
		sqrSum += vec.Data_[i] * vec.Data_[i]
	}
	return math.Sqrt(sqrSum)
}

func (vec Vec3f) Normalize(l ...float64) Vec3f {
	L := 1.0
	if len(l) > 0 {
		L = l[0]
	}
	var result Vec3f = vec.Scaling(L / vec.Norm())
	return result
}

func (vec Vec3f) negative() Vec3f {
	return vec.Scaling(-1.0)
}




type Vec4f struct {
	Data_ [4]float64
}

func (vec Vec4f) GetComponent(index int) float64 {
	if index < 4 {
		return vec.Data_[index]
	}
	panic("Vec4f size exceeded")
}

func NewVec4f(x float64, y float64, z float64, w float64) Vec4f {
	result := Vec4f{[4]float64{x, y, z, w}}
	return result
}
