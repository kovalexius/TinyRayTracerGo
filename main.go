package main

import (
	."./geometry"
	"fmt"
	"os"
	"math"
	"strconv"
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
			   raybuffer [][]Vec3f,
			   spheres []Sphere,
			   lights []Light, 
			   framebuffer [][]Vec3f) {
	orig := NewVec3f(0.0, 0.0 ,0.0 )
	for col := 0; col < width; col++ {
		framebuffer[row][col] = CastRay(&orig, &raybuffer[row][col], &spheres, &lights);
	}
}


func render(spheres []Sphere, lights []Light, width int, height int, fov float64, framebuffer [][]Vec3f, raybuffer [][]Vec3f, filename string) {
	
	for j := 0; j < height; j++ {				// actual rendering loop
		go renderRow(j, height, width, fov, raybuffer, spheres, lights, framebuffer)
	}
	
	file , err := os.Create(filename)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	
	fmt.Fprintf(file, "P6\n%d %d \n255\n", width, height)
	buf := []byte{}
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			var c Vec3f = framebuffer[row][col]
			var max float64 = math.Max(c.Data_[0], math.Max(c.Data_[1], c.Data_[2]))
			if max>1 {
				c = c.Scaling(1.0/max);
			}
			for j := 0; j<3; j++ {
				var fComponent float64 = 255 * math.Max(0.0, math.Min(1.0, c.GetComponent(j)))
				var component byte = byte(fComponent)
				buf = append(buf, component)
			}
		}
    }
	
	file.Write(buf)
}


func main() {
	const width int = 1024
	const height int = 768
	const fov float64 = math.Pi/3.0;
	framebuffer := make([][]Vec3f, width)
	for i := 0; i < height; i++ {
		framebuffer[i] = make([]Vec3f, width)
	}
	raybuffer := make([][]Vec3f, width)
	for i:= 0; i < height; i++ {
		raybuffer[i] = make([]Vec3f, width)
	}
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			var dir_x float64 =  (float64(col) + 0.5) - float64(width)/2.0;
			var dir_y float64 = -(float64(row) + 0.5) + float64(height)/2.0;    // this flips the image at the same time
			var dir_z float64 = -float64(height)/(2.0*math.Tan(fov/2.0));
			raybuffer[row][col] = NewVec3f(dir_x, dir_y, dir_z).Normalize()
		}
	}
	
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
	
	name := "outGoLang"
	for i := 0; i < 10; i++ {
		filename := name + strconv.Itoa(i) + ".ppm"
		spheres[3].Center.Data_[0] -= 0.2
		render(spheres, lights, width, height, fov, framebuffer, raybuffer, filename)
	}
}
