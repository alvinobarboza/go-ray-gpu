package raytracer

import (
	"math"
	"unsafe"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	TAU        = 2 * math.Pi
	DEG_TO_RAD = TAU / 360
)

func VecDot(v1, v2 Vec3) float32 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

func VecAdd(v1, v2 Vec3) Vec3 {
	return Vec3{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
		Z: v1.Z + v2.Z,
	}
}

func CrossProdutc(v, w Vec3) Vec3 {
	return Vec3{
		X: v.Y*w.Z - v.Z*w.Y,
		Y: v.Z*w.X - v.X*w.Z,
		Z: v.X*w.Y - v.Y*w.X,
	}
}

func SetVec3Shader(shader rl.Shader, pos int32, vec Vec3) {
	rl.SetShaderValue(
		shader, pos,
		[]float32{vec.X, vec.Y, vec.Z},
		rl.ShaderUniformVec3,
	)
}

func SetFloatShader(shader rl.Shader, pos int32, value float32) {
	rl.SetShaderValue(
		shader, pos,
		[]float32{value},
		rl.ShaderUniformFloat,
	)
}

func SetIntShader(shader rl.Shader, pos int32, value int32) {
	rl.SetShaderValue(
		shader, pos,
		unsafe.Slice((*float32)(unsafe.Pointer(&value)), 4),
		rl.ShaderUniformInt,
	)
}
