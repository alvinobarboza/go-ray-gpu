package raytracer

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Sphere struct {
	Center          Vec3
	Radius          float32
	Color           rl.Color
	Specular        int32
	Reflective      float32
	Opacity         float32
	RefractionIndex float32
}

func SetupSpheres(shader rl.Shader) []Sphere {
	return []Sphere{
		{
			Center:     Vec3{X: 0, Y: -1, Z: 3},
			Radius:     1,
			Color:      rl.Red,
			Specular:   500,
			Reflective: 0.2,
		},
		{
			Center:     Vec3{X: 2, Y: 0, Z: 4},
			Radius:     1,
			Color:      rl.Blue,
			Specular:   500,
			Reflective: 0.001,
		},
		{
			Center:     Vec3{X: -2, Y: 0, Z: 4},
			Radius:     1,
			Color:      rl.Green,
			Specular:   10,
			Reflective: 0.1,
		},
		{
			Center:          Vec3{X: -.5, Y: 0, Z: 2},
			Radius:          .4,
			Color:           rl.Blue,
			Specular:        200,
			Reflective:      0.2,
			Opacity:         0.9,
			RefractionIndex: 1.33,
		},
		{
			Center:     Vec3{X: 0, Y: -501, Z: 0},
			Radius:     500,
			Color:      rl.DarkGreen,
			Specular:   200,
			Reflective: 0.1,
		},
	}
}
