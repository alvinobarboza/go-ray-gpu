package raytracer

import rl "github.com/gen2brain/raylib-go/raylib"

type Camera struct {
	Position    Vec3
	Rotation    Vec3
	Direction   Vec3
	Fov         Vec3
	MoveSpeed   float32
	TurnSpeed   float32
	rotationLoc int32
	positionLoc int32
	fovLoc      int32
}

func (c *Camera) MoveForward(unit float32) {
	direction := c.Direction.RotateXYZ(c.Rotation)
	lenD := direction.VecLen()
	normalDirection := Vec3{
		X: direction.X / lenD,
		Y: direction.Y / lenD,
		Z: direction.Z / lenD,
	}
	c.Position.X += normalDirection.X * unit
	c.Position.Y += normalDirection.Y * unit
	c.Position.Z += normalDirection.Z * unit
}

func (c *Camera) MoveBackward(unit float32) {
	c.MoveForward(-unit)
}

func (c *Camera) MoveLeft(unit float32) {
	direction := c.Direction.RotateXYZ(c.Rotation)
	sideDirection := CrossProdutc(direction, Vec3{Y: 1})

	lenD := sideDirection.VecLen()
	normalDirection := Vec3{
		X: sideDirection.X / lenD,
		Y: sideDirection.Y / lenD,
		Z: sideDirection.Z / lenD,
	}
	c.Position.X += normalDirection.X * unit
	c.Position.Y += normalDirection.Y * unit
	c.Position.Z += normalDirection.Z * unit
}

func (c *Camera) MoveRight(unit float32) {
	c.MoveLeft(-unit)
}

func (c *Camera) UpdateCamera() bool {
	updated := false
	if rl.IsKeyDown(rl.KeyW) {
		c.MoveForward(c.MoveSpeed * rl.GetFrameTime())
		updated = true
	}
	if rl.IsKeyDown(rl.KeyS) {
		c.MoveBackward(c.MoveSpeed * rl.GetFrameTime())
		updated = true
	}
	if rl.IsKeyDown(rl.KeyA) {
		c.MoveLeft(c.MoveSpeed * rl.GetFrameTime())
		updated = true
	}
	if rl.IsKeyDown(rl.KeyD) {
		c.MoveRight(c.MoveSpeed * rl.GetFrameTime())
		updated = true
	}
	if rl.IsKeyDown(rl.KeySpace) {
		c.Position.Y += c.MoveSpeed * rl.GetFrameTime()
		updated = true
	}
	if rl.IsKeyDown(rl.KeyLeftControl) {
		c.Position.Y -= c.MoveSpeed * rl.GetFrameTime()
		updated = true
	}

	if rl.IsKeyDown(rl.KeyRight) {
		c.Rotation.Y -= c.TurnSpeed * rl.GetFrameTime()
		updated = true
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		c.Rotation.Y += c.TurnSpeed * rl.GetFrameTime()
		updated = true
	}
	if rl.IsKeyDown(rl.KeyUp) {
		c.Rotation.X += c.TurnSpeed * rl.GetFrameTime()
		if c.Rotation.X >= 90 {
			c.Rotation.X = 89
		}
		updated = true
	}
	if rl.IsKeyDown(rl.KeyDown) && c.Rotation.X < 90 {
		c.Rotation.X -= c.TurnSpeed * rl.GetFrameTime()
		if c.Rotation.X <= -90 {
			c.Rotation.X = -89
		}
		updated = true
	}
	return updated
}

func (c *Camera) UpdateShaderValues(shader rl.Shader) {
	SetVec3Shader(shader, c.rotationLoc, c.Rotation)
	SetVec3Shader(shader, c.positionLoc, c.Position)
	SetVec3Shader(shader, c.fovLoc, c.Fov)
}

func SetupCamera(moveSpeed, turnSpeed float32, shader rl.Shader) Camera {
	camera := Camera{
		Position:  Vec3{X: 0, Y: 0, Z: 0},
		Rotation:  Vec3{X: 0, Y: 0, Z: 0},
		Direction: Vec3{X: 0, Y: 0, Z: 1},
		MoveSpeed: moveSpeed,
		TurnSpeed: turnSpeed,
	}

	camera.rotationLoc = rl.GetShaderLocation(shader, "camera.rotation")
	camera.positionLoc = rl.GetShaderLocation(shader, "camera.position")
	camera.fovLoc = rl.GetShaderLocation(shader, "camera.fov")

	camera.UpdateShaderValues(shader)

	return camera
}
