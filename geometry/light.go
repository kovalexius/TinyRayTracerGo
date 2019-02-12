package geometry

import (
)


type Light struct {
	Position Vec3f
	Intensity float64
}
func NewLight(position Vec3f, intensity float64) Light {
	light := Light{Position: position, Intensity: intensity}
	return light
}
