package raytracer

import "math"

type Vec3 struct {
	X, Y, Z float32
}

func (v Vec3) VecDot() float32 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vec3) VecLen() float32 {
	return float32(math.Sqrt(float64(v.VecDot())))
}

func (v Vec3) VecNormal() Vec3 {
	n := v.VecLen()
	return Vec3{
		X: v.X / n,
		Y: v.Y / n,
		Z: v.Z / n,
	}
}

func (v Vec3) VecMultiply(n float32) Vec3 {
	return Vec3{
		X: v.X * n,
		Y: v.Y * n,
		Z: v.Z * n,
	}
}

func (v Vec3) MatrixMultiplication(m []float32) Vec3 {
	result := []float32{0, 0, 0}
	vec := []float32{v.X, v.Y, v.Z}

	length := 3

	for h := range length {
		for w := range length {
			result[h] += vec[w] * m[length*h+w]
		}
	}
	return Vec3{X: result[0], Y: result[1], Z: result[2]}
}

func (v Vec3) RotateXYZ(angle Vec3) Vec3 {
	matrix := RotationMatrixXYZ(angle)
	value := v.MatrixMultiplication(matrix)
	return value
}
