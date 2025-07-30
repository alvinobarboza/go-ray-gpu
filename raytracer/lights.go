package raytracer

import rl "github.com/gen2brain/raylib-go/raylib"

type lightType int

const (
	AMBIENT lightType = iota
	POINT
	DIRECTIONAL
)

type Lights struct {
	TypeL     lightType
	Intensity float32
	Direction Vec3
	Position  Vec3
}

func SetupLights(shader rl.Shader) []Lights {
	return []Lights{
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
}
