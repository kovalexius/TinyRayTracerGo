package geometry



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

func (vec Vec3f) GetComponent(index int) float64 {
	if index < 3 {
		return vec.Data_[index]
	}
	panic("Vec3f size exceeded")
}

func NewVec3f(x float64, y float64, z float64) Vec3f {
	result := Vec3f{[3]float64{x, y, z}}
	return result
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
