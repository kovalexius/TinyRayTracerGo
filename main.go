package main

import (
	."./geometry"
	"fmt"
	"os"
	"math"
)

/*
func cast_ray(const Vec3f &orig, const Vec3f &dir, const std::vector<Sphere> &spheres, const std::vector<Light> &lights, size_t depth=0) Vec3f {
	point, N Vec3f
	material Material

    if (depth>4 || !scene_intersect(orig, dir, spheres, point, N, material)) {
        return Vec3f(0.2, 0.7, 0.8); // background color
    }

    Vec3f reflect_dir = reflect(dir, N).normalize();
    Vec3f refract_dir = refract(dir, N, material.refractive_index).normalize();
    Vec3f reflect_orig = reflect_dir*N < 0 ? point - N*1e-3 : point + N*1e-3; // offset the original point to avoid occlusion by the object itself
    Vec3f refract_orig = refract_dir*N < 0 ? point - N*1e-3 : point + N*1e-3;
    Vec3f reflect_color = cast_ray(reflect_orig, reflect_dir, spheres, lights, depth + 1);
    Vec3f refract_color = cast_ray(refract_orig, refract_dir, spheres, lights, depth + 1);

    float diffuse_light_intensity = 0, specular_light_intensity = 0;
    for (size_t i=0; i<lights.size(); i++) {
        Vec3f light_dir      = (lights[i].position - point).normalize();
        float light_distance = (lights[i].position - point).norm();

        Vec3f shadow_orig = light_dir*N < 0 ? point - N*1e-3 : point + N*1e-3; // checking if the point lies in the shadow of the lights[i]
        Vec3f shadow_pt, shadow_N;
        Material tmpmaterial;
        if (scene_intersect(shadow_orig, light_dir, spheres, shadow_pt, shadow_N, tmpmaterial) && (shadow_pt-shadow_orig).norm() < light_distance)
            continue;

        diffuse_light_intensity  += lights[i].intensity * std::max(0.f, light_dir*N);
        specular_light_intensity += powf(std::max(0.f, -reflect(-light_dir, N)*dir), material.specular_exponent)*lights[i].intensity;
    }
    return material.diffuse_color * diffuse_light_intensity * material.albedo[0] + Vec3f(1., 1., 1.)*specular_light_intensity * material.albedo[1] + reflect_color*material.albedo[2] + refract_color*material.albedo[3];
}
*/


func render() {
	const width int = 1024
	const height int = 768
	
	var framebuffer [width*height] Vec3f

	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			framebuffer[i+j*width] = Vec3f{[3]float64{float64(j)/float64(height), float64(i)/float64(width), 0.0}}
		}
	}
	
	file , err := os.Create("./out.ppm")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	
	fmt.Fprintf(file, "P6\n%d %d \n255\n", width, height)
	buf := []byte{}
	for i := 0; i < height*width; i++ {
		for j := 0; j < 3; j++ {
			
			var fComponent float64 = 255 * math.Max(0.0, math.Min(1.0, framebuffer[i].GetComponent(j)))
			var component byte = byte(fComponent)
			buf = append(buf, component)
		}
	}
	file.Write(buf)
}

func main() {
	render()
	material := NewMaterial(1.0, Vec4f{}, Vec3f{}, 1.0)
	_ = material
}
