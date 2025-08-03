package main

import (
	"fmt"

	"github.com/alvinobarboza/go-ray-gpu/raytracer"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Same as in shader
const MaxBounces = int32(6)

func main() {
	maxBounces := int32(3)

	screenWidth := int32(1000)
	screenHeight := int32(1000)

	moveSpeed := float32(4)
	turnSpeed := float32(25.0)

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

	bouncesLoc := rl.GetShaderLocation(shader, "maxBounces")
	raytracer.SetIntShader(shader, bouncesLoc, maxBounces)

	rl.SetTargetFPS(200)

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

		if rl.IsKeyReleased(rl.KeyQ) {
			maxBounces--
			if maxBounces < 2 {
				maxBounces = 2
			}
			raytracer.SetIntShader(shader, bouncesLoc, maxBounces)
		}
		if rl.IsKeyReleased(rl.KeyE) {
			maxBounces++
			if maxBounces > 6 {
				maxBounces = 6
			}
			raytracer.SetIntShader(shader, bouncesLoc, maxBounces)
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.BeginShaderMode(shader)
		rl.DrawRectangle(0, 0, screenWidth, screenHeight, rl.White)
		rl.EndShaderMode()

		rl.DrawText("Test shader", screenWidth-140, screenHeight-20, 20, rl.Gray)

		rl.DrawRectangle(2, 2, 305, 230, rl.Fade(rl.DarkGray, 0.6))
		rl.DrawRectangleLines(2, 2, 305, 230, rl.Gray)

		rl.DrawFPS(10, 10)
		rl.DrawText(
			fmt.Sprintf("Light bounce count: %01d E+/Q-", maxBounces),
			10, 30, 20, rl.White)
		rl.DrawText(
			fmt.Sprintf("Cam: \nX:%01f \nY:%01f \nZ:%01f", camera.Position.X, camera.Position.Y, camera.Position.Z),
			10, 50, 20, rl.White)
		rl.DrawText("Move: A/W/S/D",
			10, 140, 20, rl.White)
		rl.DrawText("Mouse view moviment: \nL Click: Lock view\nR Click: Unlock",
			10, 162, 20, rl.White)

		rl.EndDrawing()
	}

}
