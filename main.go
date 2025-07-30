package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.SetConfigFlags(rl.FlagWindowResizable | rl.FlagMsaa4xHint)
	rl.InitWindow(screenWidth, screenHeight, "go gpu raytracer - raylib screen texture")

	camera := rl.Camera{
		Position:   rl.Vector3{X: 0, Y: 0, Z: 0},
		Target:     rl.Vector3{X: 0, Y: 0, Z: 0},
		Up:         rl.Vector3{X: 0, Y: 1, Z: 0},
		Fovy:       45,
		Projection: rl.CameraPerspective,
	}

	shader := rl.LoadShader("", "raytrace.fs")
	defer rl.UnloadShader(shader)

	res := rl.GetShaderLocation(shader, "res")

	rl.SetShaderValue(
		shader, res,
		[]float32{float32(screenWidth), float32(screenHeight)},
		rl.ShaderUniformVec2,
	)

	// target := rl.LoadRenderTexture(screenWidth, screenHeight)
	// defer rl.UnloadRenderTexture(target)

	rl.SetTargetFPS(60)

	sChange := false
	for !rl.WindowShouldClose() {
		sW := int32(rl.GetScreenWidth())
		sH := int32(rl.GetScreenHeight())

		if sW != screenWidth {
			screenWidth = sW
			sChange = true
		}

		if sH != screenHeight {
			screenHeight = sH
			sChange = true
		}

		if sChange {
			rl.SetShaderValue(
				shader, res,
				[]float32{float32(screenWidth), float32(screenHeight)},
				rl.ShaderUniformVec2,
			)
			sChange = false
		}

		rl.UpdateCamera(&camera, rl.CameraFirstPerson)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.BeginShaderMode(shader)
		rl.DrawRectangle(0, 0, screenWidth, screenHeight, rl.White)
		rl.EndShaderMode()

		rl.DrawText("Test shader", screenWidth-140, screenHeight-20, 20, rl.Gray)
		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}

}
