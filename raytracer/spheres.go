package raytracer

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Sphere struct {
	Center             Vec3
	Radius             float32
	Color              rl.Color
	Specular           int32
	Reflective         float32
	Opacity            float32
	RefractionIndex    float32
	centerLoc          int32
	radiusLoc          int32
	colorLoc           int32
	specularLoc        int32
	reflectiveLoc      int32
	opacityLoc         int32
	refractionIndexLoc int32
}

func (s Sphere) UpdateShaderValues(shader rl.Shader) {
	SetVec3Shader(shader, s.centerLoc, s.Center)
	SetFloatShader(shader, s.radiusLoc, s.Radius)
	SetVec3Shader(shader, s.colorLoc, RGBToShaderVec3Normalized(s.Color))
	SetIntShader(shader, s.specularLoc, s.Specular)
	SetFloatShader(shader, s.reflectiveLoc, s.Reflective)
	SetFloatShader(shader, s.opacityLoc, s.Opacity)
	SetFloatShader(shader, s.refractionIndexLoc, s.RefractionIndex)
}

func SetupSpheres(shader rl.Shader) []Sphere {
	spheres := []Sphere{
		{
			Center:     Vec3{X: 0, Y: -1, Z: 3},
			Radius:     1,
			Color:      rl.Red,
			Specular:   500,
			Reflective: 0.9,
		},
		{
			Center:     Vec3{X: 2, Y: 0, Z: 4},
			Radius:     1,
			Color:      rl.Blue,
			Specular:   500,
			Reflective: 0,
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
			Reflective:      0.9,
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

	for i := range spheres {
		spheres[i].centerLoc = rl.GetShaderLocation(shader, fmt.Sprintf("spheres[%d].center", i))
		spheres[i].radiusLoc = rl.GetShaderLocation(shader, fmt.Sprintf("spheres[%d].radius", i))
		spheres[i].colorLoc = rl.GetShaderLocation(shader, fmt.Sprintf("spheres[%d].color", i))
		spheres[i].specularLoc = rl.GetShaderLocation(shader, fmt.Sprintf("spheres[%d].specular", i))
		spheres[i].reflectiveLoc = rl.GetShaderLocation(shader, fmt.Sprintf("spheres[%d].reflective", i))
		spheres[i].opacityLoc = rl.GetShaderLocation(shader, fmt.Sprintf("spheres[%d].opacity", i))
		spheres[i].refractionIndexLoc = rl.GetShaderLocation(shader, fmt.Sprintf("spheres[%d].refractionIndex", i))
	}

	return spheres
}

func UpdateAllSpheresShaders(spheres []Sphere, shader rl.Shader) {
	for _, s := range spheres {
		s.UpdateShaderValues(shader)
	}
}
