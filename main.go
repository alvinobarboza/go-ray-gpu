package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "go gpu raytracer - raylib screen texture")

	rl.SetConfigFlags(rl.FlagMsaa4xHint)

	camera := rl.Camera{
		Position:   rl.Vector3{X: 0, Y: 0, Z: 0},
		Target:     rl.Vector3{X: 0, Y: 0, Z: 0},
		Up:         rl.Vector3{X: 0, Y: 1, Z: 0},
		Fovy:       45,
		Projection: rl.CameraPerspective,
	}

	shader := rl.LoadShader("", "raytrace.fs")
	defer rl.UnloadShader(shader)

	resolution := rl.GetShaderLocation(shader, "resolution")

	rl.SetShaderValue(
		shader, resolution,
		[]float32{float32(screenWidth), float32(screenHeight)},
		rl.ShaderUniformIvec2,
	)

	// target := rl.LoadRenderTexture(screenWidth, screenHeight)
	// defer rl.UnloadRenderTexture(target)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&camera, rl.CameraFirstPerson)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.BeginShaderMode(shader)
		rl.DrawRectangle(0, 0, screenWidth, screenHeight, rl.White)
		rl.EndShaderMode()

		rl.DrawText("Test shader", screenWidth-100, screenHeight-20, 10, rl.Gray)
		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}

}
