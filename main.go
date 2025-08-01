package main

import (
	"fmt"

	"github.com/alvinobarboza/go-ray-gpu/raytracer"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	screenWidth := int32(1000)
	screenHeight := int32(1000)

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

	bgColorLoc := rl.GetShaderLocation(shader, "backgroundColor")
	raytracer.SetVec3Shader(shader, bgColorLoc, raytracer.RGBToShaderVec3Normalized(rl.Gray))

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

		rl.DrawRectangle(2, 2, 445, 195, rl.Fade(rl.DarkGray, 0.6))
		rl.DrawRectangleLines(2, 2, 445, 195, rl.Gray)

		rl.DrawFPS(10, 10)
		rl.DrawText(
			fmt.Sprintf("Cam-> \nX:%01f \nY:%01f \nZ:%01f", camera.Position.X, camera.Position.Y, camera.Position.Z),
			10, 30, 20, rl.White)
		rl.DrawText("Move: A/W/S/D\nControl Camera: UP/DOWN/LEFT/RIGHT",
			10, 120, 20, rl.White)
		rl.DrawText("I'm a bit too lazy to make \nit work with the mouse...",
			10, 160, 10, rl.White)

		rl.EndDrawing()
	}

}
