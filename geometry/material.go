package geometry

type Material struct {
    Refractive_index float64
    Albedo *Vec4f
    Diffuse_color *Vec3f
    Specular_exponent float64
}

func NewMaterial(args ...interface{}) *Material {
	number := len(args)
	if number == 0 {
		return &Material {
			Refractive_index:1.0, 
			Albedo: NewVec4f(1.0, 0.0, 0.0, 0.0), 
			Diffuse_color: &Vec3f{},
		}
	}
	
	if  number != 4 {
		panic("Wrong number of arguments")
	}
	
	refractive, ok := args[0].(float64)
	if !ok {
		panic("First argument must be float64")
	}
	albedo, ok := args[1].(*Vec4f)
	if !ok {
		panic("Second argument must be Vec4f")
	}
	diffuse, ok := args[2].(*Vec3f)
	if !ok {
		panic("Third argument must be Vec3f")
	}
	specular, ok := args[3].(float64)
	if !ok {
		panic("4-th argument must be float64")
	}
	
	return &Material {
		Refractive_index: refractive,
		Albedo: albedo,
		Diffuse_color: diffuse,
		Specular_exponent: specular,
	}
}
