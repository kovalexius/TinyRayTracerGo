package geometry

import (
	"math"
	//"fmt"
)



func Reflect(I *Vec3f, N *Vec3f) *Vec3f {
	return I.Sub(N.Scaling(2.0 * I.ScalarMul(N)))
}


func Refract(I *Vec3f, N *Vec3f, eta_t float64, eta_i ...float64) *Vec3f {			 // Snell's law
	Eta_i := 1.0
	if len(eta_i) > 0 {
		Eta_i = eta_i[0]
	}
	var cosi float64 = -math.Max(-1.0, math.Min(1.0, I.ScalarMul(N)))
	if cosi < 0 {
		return Refract(I, N.negative(), Eta_i, eta_t)			// if the ray comes from the inside the object, swap the air and the media
	}
	var eta float64 = Eta_i / eta_t
	var k float64 = 1.0 - eta*eta*(1.0 - cosi*cosi)
	
	// k<0 = total reflection, no ray to refract. I refract it anyways, this has no physical meaning
	if k<0 {
		return NewVec3f(1.0, 0.0, 0.0)
	}
	
	return I.Scaling(eta).Add(N.Scaling(eta*cosi - math.Sqrt(k)))
}


func SceneIntersect(orig *Vec3f, 
					dir *Vec3f, 
					spheres *[]*Sphere,
					hit *Vec3f, 
					N *Vec3f, 
					material *Material) (result bool) {
	var spheres_dist float64 = math.MaxFloat64
	for i := 0; i < len(*spheres); i++ {
		if ok, dist_i := (*spheres)[i].RayIntersect(orig, dir); ok && (dist_i < spheres_dist) {
			spheres_dist = dist_i
			*hit =  *dir.Scaling(dist_i).Add(orig)
			*N = *hit.Sub((*spheres)[i].Center_).Normalize();
			*material = *(*spheres)[i].Material_
		}
	}
	
	var checkerboard_dist float64 = math.MaxFloat64
	if math.Abs(dir.Data_[1]) > 1e-3 {
		var d float64 = -(orig.Data_[1]+4.0)/dir.Data_[1]   // the checkerboard plane has equation y = -4
		vecDir := dir.Scaling(d)
		pt := orig.Add(vecDir)
		if (d > 0) && (math.Abs(pt.Data_[0])<10.0) && (pt.Data_[2]<(-10.0)) && (pt.Data_[2]>(-30.0)) && (d<spheres_dist) {
			checkerboard_dist = d
			*hit = *pt
			*N = *NewVec3f(0.0, 1.0, 0.0)
			if ((int(0.5*hit.Data_[0]+1000) + int(0.5*hit.Data_[2])) & 1) != 0 {
				material.Diffuse_color = NewVec3f(0.3, 0.3, 0.3)
			} else {
				material.Diffuse_color = NewVec3f(0.1, 0.7, 0.1)
			}
		}
	}
	result = math.Min(spheres_dist, checkerboard_dist) < 1000.0
	return 
}


func CastRay (orig *Vec3f, 
			  dir *Vec3f, 
			  spheres *[]*Sphere,
			  lights *[]*Light,
			  depth  ...int) *Vec3f {
	var Depth int = 0
	if len(depth) > 0 {
		Depth = depth[0]
	}
	
	var point, N Vec3f
    material := NewMaterial()
	ok := SceneIntersect(orig, dir, spheres, &point, &N, material)
	if (Depth>4) || !ok {
		return NewVec3f(0.2, 0.7, 0.8) 			// background color
	}
	
	reflect_dir := Reflect(dir, &N).Normalize()
	refract_dir := Refract(dir, &N, material.Refractive_index).Normalize()
	
	var reflect_orig *Vec3f						// offset the original point to avoid occlusion by the object itself
	vecN := N.Scaling(1e-3)
	if reflect_dir.ScalarMul(&N) < 0 {
		reflect_orig = point.Sub(vecN)
	} else {
		reflect_orig = point.Add(vecN)
	}
	
	var refract_orig *Vec3f
	if refract_dir.ScalarMul(&N) < 0 {
		refract_orig = point.Sub(vecN)
	} else {
		refract_orig = point.Add(vecN)
	}
	
	reflect_color := CastRay(reflect_orig, reflect_dir, spheres, lights, Depth + 1)
	refract_color := CastRay(refract_orig, refract_dir, spheres, lights, Depth + 1)
	
	var diffuse_light_intensity, specular_light_intensity float64 = 0.0, 0.0
	
	for i:=0; i<len(*lights); i++ {
		light_dir := (*lights)[i].Position.Sub(&point).Normalize();
		var light_distance float64 = (*lights)[i].Position.Sub(&point).Norm();
		
		var shadow_orig *Vec3f					// checking if the point lies in the shadow of the lights[i]
		if light_dir.ScalarMul(&N) < 0 {
			shadow_orig = point.Sub(vecN)
		} else {
			shadow_orig = point.Add(vecN)
		}
		
		var shadow_pt Vec3f
		ok := SceneIntersect(shadow_orig, light_dir, spheres, &shadow_pt, &Vec3f{}, &Material{})
		if ok && shadow_pt.Sub(shadow_orig).Norm() < light_distance {
			continue
		}
		
		diffuse_light_intensity  += (*lights)[i].Intensity * math.Max(0.0, light_dir.ScalarMul(&N))
		var max float64 = math.Max(0.0, Reflect(light_dir.negative(), &N).negative().ScalarMul(dir))
		var pow float64 = math.Pow(max, material.Specular_exponent)
		specular_light_intensity += pow * (*lights)[i].Intensity;
	}
	
	firstComponent := material.Diffuse_color.Scaling(diffuse_light_intensity * material.Albedo.Data_[0])
	secondComponent := NewVec3f(1.0, 1.0, 1.0).Scaling(specular_light_intensity * material.Albedo.Data_[1])
	thirdComponent := reflect_color.Scaling(material.Albedo.Data_[2])
	forthComponent := refract_color.Scaling(material.Albedo.Data_[3])
	
	return firstComponent.Add(secondComponent).Add(thirdComponent).Add(forthComponent)
}








