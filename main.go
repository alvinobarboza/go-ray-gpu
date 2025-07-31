package main

import (
	"github.com/alvinobarboza/go-ray-gpu/raytracer"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	moveSpeed := float32(4)
	turnSpeed := float32(70.0)

	rl.SetConfigFlags(rl.FlagWindowResizable | rl.FlagMsaa4xHint)
	rl.InitWindow(screenWidth, screenHeight, "go gpu raytracer - raylib screen texture")

	shader := rl.LoadShader("", "raytrace.fs")
	defer rl.UnloadShader(shader)

	camera := raytracer.SetupCamera(
		moveSpeed,
		turnSpeed,
		shader,
		float32(screenWidth), float32(screenHeight))

	spheres := raytracer.SetupSpheres(shader)
	raytracer.UpdateAllSpheresShaders(spheres, shader)

	lights := raytracer.SetupLights(shader)
	raytracer.UpdateAllLightsShader(lights, shader)

	res := rl.GetShaderLocation(shader, "res")

	rl.SetShaderValue(
		shader, res,
		[]float32{float32(screenWidth), float32(screenHeight)},
		rl.ShaderUniformVec2,
	)

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
			w, h := float32(screenWidth), float32(screenHeight)

			camera.UpdateFov(w, h)
			camera.UpdateShaderValues(shader)

			rl.SetShaderValue(
				shader, res,
				[]float32{w, h},
				rl.ShaderUniformVec2,
			)
			sChange = false
		}

		if camera.UpdateCamera() {
			camera.UpdateShaderValues(shader)
		}

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
