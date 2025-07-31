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

func RGBToShaderVec3Normalized(color rl.Color) Vec3 {
	return Vec3{X: float32(color.R) / 255, Y: float32(color.G) / 255, Z: float32(color.B) / 255}
}

func RotationMatrixXYZ(angle Vec3) []float32 {
	cosa := float32(math.Cos(float64(angle.X * -DEG_TO_RAD)))
	sina := float32(math.Sin(float64(angle.X * -DEG_TO_RAD)))

	cosb := float32(math.Cos(float64(angle.Y * -DEG_TO_RAD)))
	sinb := float32(math.Sin(float64(angle.Y * -DEG_TO_RAD)))

	cosga := float32(math.Cos(float64(angle.Z * -DEG_TO_RAD)))
	singa := float32(math.Sin(float64(angle.Z * -DEG_TO_RAD)))

	// Formula for general 3D roation using matrix
	matrix := []float32{
		cosb * cosga, sina*sinb*cosga - cosa*singa, cosa*sinb*cosga + sina*singa,
		cosb * singa, sina*sinb*singa + cosa*cosga, cosa*sinb*singa - sina*cosga,
		-sinb, sina * cosb, cosa * cosb,
	}
	return matrix
}

func MatrixToGlslMatrix(matrix []float32) rl.Matrix {

	glMatrix := rl.Matrix{}

	glMatrix.M0 = matrix[0]
	glMatrix.M1 = matrix[3]
	glMatrix.M2 = matrix[6]

	glMatrix.M4 = matrix[1]
	glMatrix.M5 = matrix[4]
	glMatrix.M6 = matrix[7]

	glMatrix.M8 = matrix[2]
	glMatrix.M9 = matrix[5]
	glMatrix.M10 = matrix[8]

	return glMatrix
}
