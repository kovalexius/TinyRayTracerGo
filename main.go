package main

import (
	."./geometry"
	"fmt"
	"os"
	"math"
)


/*
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
*/

func renderRow(row int, 
			   height int, 
			   width int,
			   fov float64,
			   spheres []Sphere,
			   lights []Light, 
			   framebuffer []Vec3f) {
	for i := 0; i < width; i++ {
		var dir_x float64 =  (float64(i) + 0.5) - float64(width)/2.0;
        var dir_y float64 = -(float64(row) + 0.5) + float64(height)/2.0;    // this flips the image at the same time
        var dir_z float64 = -float64(height)/(2.0*math.Tan(fov/2.0));
		framebuffer[i+row*width] = CastRay(NewVec3f(0.0, 0.0 ,0.0 ), NewVec3f(dir_x, dir_y, dir_z).Normalize(), spheres, lights);
	}
}


func render(spheres []Sphere, lights []Light) {
	const width int = 1024
	const height int = 768
	const fov float64 = math.Pi/3.0;
	var framebuffer [width*height] Vec3f
	//framebuffer

	for j := 0; j < height; j++ {				// actual rendering loop
		go renderRow(j, height, width, fov, spheres, lights, framebuffer[:])
	}
	
	file , err := os.Create("./outGoLang.ppm")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	
	fmt.Fprintf(file, "P6\n%d %d \n255\n", width, height)
	buf := []byte{}
	for i := 0; i < height*width; i++ {
        var c Vec3f = framebuffer[i]
        var max float64 = math.Max(c.Data_[0], math.Max(c.Data_[1], c.Data_[2]))
        if max>1 {
			c = c.Scaling(1.0/max);
		}
        for j := 0; j<3; j++ {
            var fComponent float64 = 255 * math.Max(0.0, math.Min(1.0, framebuffer[i].GetComponent(j)))
			var component byte = byte(fComponent)
			buf = append(buf, component)
        }
    }
	
	file.Write(buf)
}


func main() {
	ivory := NewMaterial(1.0, NewVec4f(0.6,  0.3, 0.1, 0.0), NewVec3f(0.4, 0.4, 0.3), 50.0)
    glass := NewMaterial(1.5, NewVec4f(0.0,  0.5, 0.1, 0.8), NewVec3f(0.6, 0.7, 0.8),  125.0)
    red_rubber := NewMaterial(1.0, NewVec4f(0.9,  0.1, 0.0, 0.0), NewVec3f(0.3, 0.1, 0.1), 10.0)
    mirror := NewMaterial(1.0, NewVec4f(0.0, 10.0, 0.8, 0.0), NewVec3f(1.0, 1.0, 1.0), 1425.0)

    var spheres = []Sphere {
			NewSphere(NewVec3f(-3,    0,   -16), 2, ivory),
			NewSphere(NewVec3f(-1.0, -1.5, -12), 2, glass),
			NewSphere(NewVec3f( 1.5, -0.5, -18), 3, red_rubber),
			NewSphere(NewVec3f( 7,    5,   -18), 4, mirror),
	}

    var lights = []Light {
			NewLight(NewVec3f(-20, 20,  20), 1.5),
			NewLight(NewVec3f( 30, 50, -25), 1.8),
			NewLight(NewVec3f( 30, 20,  30), 1.7),
	}
	
	render(spheres, lights)
}
