package main

import (
	."./geometry"
	"testing"
	"math"
)

func BenchmarkMain(b *testing.B) {
	main()
}

func BenchmarkRenderRow(b *testing.B) {
	const width int = 1024
	const height int = 768
	const fov float64 = math.Pi/3.0;
	framebuffer := make([]Vec3f, width*height)
	
	ivory := NewMaterial(1.0, NewVec4f(0.6,  0.3, 0.1, 0.0), NewVec3f(0.4, 0.4, 0.3), 50.0)
    glass := NewMaterial(1.5, NewVec4f(0.0,  0.5, 0.1, 0.8), NewVec3f(0.6, 0.7, 0.8),  125.0)
    red_rubber := NewMaterial(1.0, NewVec4f(0.9,  0.1, 0.0, 0.0), NewVec3f(0.3, 0.1, 0.1), 10.0)
    mirror := NewMaterial(1.0, NewVec4f(0.0, 10.0, 0.8, 0.0), NewVec3f(1.0, 1.0, 1.0), 1425.0)

    var spheres = []Sphere {
			NewSphere(*NewVec3f(-3,    0,   -16), 2, ivory),
			NewSphere(*NewVec3f(-1.0, -1.5, -12), 2, glass),
			NewSphere(*NewVec3f( 1.5, -0.5, -18), 3, red_rubber),
			NewSphere(*NewVec3f( 7,    5,   -18), 4, mirror),
	}

    var lights = []Light {
			NewLight(*NewVec3f(-20, 20,  20), 1.5),
			NewLight(*NewVec3f( 30, 50, -25), 1.8),
			NewLight(*NewVec3f( 30, 20,  30), 1.7),
	}
	
	//renderRow(0, height, width, fov, spheres, lights, framebuffer[:])
	for j := 0; j < height; j++ {
		renderRow(j, height, width, fov, &spheres, &lights, &framebuffer)
	}
}

func BenchmarkRenderRowParallel(b *testing.B) {
	const width int = 1024
	const height int = 768
	const fov float64 = math.Pi/3.0;
	framebuffer := make([]Vec3f, width*height)
	
	ivory := NewMaterial(1.0, NewVec4f(0.6,  0.3, 0.1, 0.0), NewVec3f(0.4, 0.4, 0.3), 50.0)
    glass := NewMaterial(1.5, NewVec4f(0.0,  0.5, 0.1, 0.8), NewVec3f(0.6, 0.7, 0.8),  125.0)
    red_rubber := NewMaterial(1.0, NewVec4f(0.9,  0.1, 0.0, 0.0), NewVec3f(0.3, 0.1, 0.1), 10.0)
    mirror := NewMaterial(1.0, NewVec4f(0.0, 10.0, 0.8, 0.0), NewVec3f(1.0, 1.0, 1.0), 1425.0)

    var spheres = []Sphere {
			NewSphere(*NewVec3f(-3,    0,   -16), 2, ivory),
			NewSphere(*NewVec3f(-1.0, -1.5, -12), 2, glass),
			NewSphere(*NewVec3f( 1.5, -0.5, -18), 3, red_rubber),
			NewSphere(*NewVec3f( 7,    5,   -18), 4, mirror),
	}

    var lights = []Light {
			NewLight(*NewVec3f(-20, 20,  20), 1.5),
			NewLight(*NewVec3f( 30, 50, -25), 1.8),
			NewLight(*NewVec3f( 30, 20,  30), 1.7),
	}
	
	//renderRow(0, height, width, fov, spheres, lights, framebuffer[:])
	for j := 0; j < height; j++ {
		go renderRow(j, height, width, fov, &spheres, &lights, &framebuffer)
	}
}

func BenchmarkCastRay(b *testing.B) {
	const width int = 1024
	const height int = 768
	const fov float64 = math.Pi/3.0
	const total int = width * height
	var row int = height/2
	var col int = width/2
	framebuffer := make([]Vec3f, total)
	
	ivory := NewMaterial(1.0, NewVec4f(0.6,  0.3, 0.1, 0.0), NewVec3f(0.4, 0.4, 0.3), 50.0)
    glass := NewMaterial(1.5, NewVec4f(0.0,  0.5, 0.1, 0.8), NewVec3f(0.6, 0.7, 0.8),  125.0)
    red_rubber := NewMaterial(1.0, NewVec4f(0.9,  0.1, 0.0, 0.0), NewVec3f(0.3, 0.1, 0.1), 10.0)
    mirror := NewMaterial(1.0, NewVec4f(0.0, 10.0, 0.8, 0.0), NewVec3f(1.0, 1.0, 1.0), 1425.0)
	
	var spheres = []Sphere {
			NewSphere(*NewVec3f(-3,    0,   -16), 2, ivory),
			NewSphere(*NewVec3f(-1.0, -1.5, -12), 2, glass),
			NewSphere(*NewVec3f( 1.5, -0.5, -18), 3, red_rubber),
			NewSphere(*NewVec3f( 7,    5,   -18), 4, mirror),
	}

    var lights = []Light {
			NewLight(*NewVec3f(-20, 20,  20), 1.5),
			NewLight(*NewVec3f( 30, 50, -25), 1.8),
			NewLight(*NewVec3f( 30, 20,  30), 1.7),
	}
	
	_ = spheres
	_ = lights
	
	
	for row = 0; row < height; row++ {
		for col = 0; col < width; col++ {
			var dir_x float64 =  (float64(col) + 0.5) - float64(width)/2.0;
			var dir_y float64 = -(float64(row) + 0.5) + float64(height)/2.0;    // this flips the image at the same time
			var dir_z float64 = -float64(height)/(2.0*math.Tan(fov/2.0));
			_ = dir_x
			_ = dir_y
			_ = dir_z
			framebuffer[row] = *CastRay(NewVec3f(0.0, 0.0 ,0.0 ), NewVec3f(dir_x, dir_y, dir_z).Normalize(), &spheres, &lights);
			//_ = vec
		}
	}

	/*
	for row = 0; row < total; row++ {
		var dir_x float64 =  (float64(col) + 0.5) - float64(width)/2.0;
		var dir_y float64 = -(float64(row) + 0.5) + float64(height)/2.0;    // this flips the image at the same time
		var dir_z float64 = -float64(height)/(2.0*math.Tan(fov/2.0));
		framebuffer[row] = CastRay(NewVec3f(0.0, 0.0 ,0.0 ), NewVec3f(dir_x, dir_y, dir_z).Normalize(), spheres, lights);
		//vec := CastRay(NewVec3f(0.0, 0.0 ,0.0 ), NewVec3f(dir_x, dir_y, dir_z), spheres, lights);
		//_ = vec
	}
	*/
	_ = framebuffer
}


