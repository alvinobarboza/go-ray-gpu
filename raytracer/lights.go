package raytracer

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type lightType int

const (
	AMBIENT lightType = iota
	POINT
	DIRECTIONAL
)

type Lights struct {
	TypeL        lightType
	Intensity    float32
	Direction    Vec3
	Position     Vec3
	typeLoc      int32
	intensityLoc int32
	directionLoc int32
	positionLoc  int32
}

func (l Lights) UpdateShaderValues(shader rl.Shader) {
	SetIntShader(shader, l.typeLoc, int32(l.TypeL))
	SetFloatShader(shader, l.intensityLoc, l.Intensity)
	SetVec3Shader(shader, l.directionLoc, l.Direction)
	SetVec3Shader(shader, l.positionLoc, l.Position)
}

func SetupLights(shader rl.Shader) []Lights {
	lights := []Lights{
		{
			TypeL:     AMBIENT,
			Intensity: 0.2,
		},
		{
			TypeL:     POINT,
			Intensity: 0.6,
			Position:  Vec3{X: 2, Y: 1, Z: 0},
		},
		{
			TypeL:     DIRECTIONAL,
			Intensity: 0.2,
			Direction: Vec3{X: 1, Y: 4, Z: 4},
		},
	}

	for i := range lights {
		lights[i].typeLoc = rl.GetShaderLocation(shader, fmt.Sprintf("lights[%d].type", i))
		lights[i].intensityLoc = rl.GetShaderLocation(shader, fmt.Sprintf("lights[%d].intensity", i))
		lights[i].directionLoc = rl.GetShaderLocation(shader, fmt.Sprintf("lights[%d].direction", i))
		lights[i].positionLoc = rl.GetShaderLocation(shader, fmt.Sprintf("lights[%d].position", i))
	}

	return lights
}

func UpdateAllLightsShader(lights []Lights, shader rl.Shader) {
	for _, l := range lights {
		l.UpdateShaderValues(shader)
	}
}
