package main

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	screenWidth  = int32(1200)
	screenHeight = int32(720)
	cubeSize     = float32(10)
	camera       rl.Camera3D
	scale        = float32(0.01)
	stepTime     = time.Millisecond
	lastUpdate   time.Time
)

func updateCameraMovement() {
	// camera movement
	move := rl.NewVector3(0.0, 0.0, 0.0)
	rotate := rl.NewVector3(0.0, 0.0, 0.0)
	speed := float32(2)
	mouseSens := float32(0.3)

	// movements
	// left
	if rl.IsKeyDown(rl.KeyA) {
		move.Y = -speed
	}
	// right
	if rl.IsKeyDown(rl.KeyD) {
		move.Y = speed
	}
	// forward
	if rl.IsKeyDown(rl.KeyW) {
		move.X = speed
	}
	// backward
	if rl.IsKeyDown(rl.KeyS) {
		move.X = -speed
	}
	// up
	if rl.IsKeyDown(rl.KeyE) {
		move.Z = speed
	}
	// down
	if rl.IsKeyDown(rl.KeyQ) {
		move.Z = -speed
	}

	// move movement
	mouseDelta := rl.GetMouseDelta()
	rotate.X = mouseDelta.X * mouseSens
	rotate.Y = mouseDelta.Y * mouseSens

	rl.UpdateCameraPro(&camera, move, rotate, 0.0)
}

func update() {
	// update camera
	updateCameraMovement()

	if time.Since(lastUpdate) > stepTime {
		connect()
		lastUpdate = time.Now()
	}
}

func renderSimulation() {
	// render unconnected cubes
	for _, box := range boxes {
		rl.DrawCube(box.Vector3, cubeSize, cubeSize, cubeSize, box.color)
	}

	// render liens
	for _, line := range lines {
		rl.DrawLine3D(line[0], line[1], rl.Black)
	}
}

func renderData() {
	currentConntxt := fmt.Sprintf("Current Conn: %d", ConnCount)
	maxConntxt := fmt.Sprintf("Max Conn: %d", MaxConn)

	var setCountTxt string
	var status string

	if !finished {
		status = "running"
	} else {
		setCountTxt = fmt.Sprintf("final answer: %v", (myset.GetSetCount()))
		status = "completed"
	}

	xpos := int32(10)
	ypos := int32(10)
	fontSize := int32(10)

	rl.DrawText(currentConntxt, xpos, ypos, fontSize, rl.Black)
	ypos += fontSize

	rl.DrawText(maxConntxt, xpos, ypos, fontSize, rl.Black)
	ypos += fontSize

	rl.DrawText(status, xpos, ypos, fontSize, rl.Black)
	ypos += fontSize

	rl.DrawText(setCountTxt, xpos, ypos, fontSize, rl.Black)
	ypos += fontSize
}

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.White)

	rl.BeginMode3D(camera)
	rl.SetClipPlanes(0.1, 10000.0)

	renderSimulation()

	rl.DrawGrid(100, 10)
	rl.EndMode3D()

	renderData()

	rl.EndDrawing()
}

func main() {
	// problem logic
	getJunctionBoxes()
	computeDistances()

	// rendering logic
	rl.InitWindow(screenWidth, screenHeight, "trying 3D for the 1st time")
	rl.DisableCursor()
	rl.SetTargetFPS(60)

	pos := rl.NewVector3(0, 0, 5)
	target := rl.NewVector3(0, 0, 0)
	up := rl.NewVector3(0, 1, 0)
	fovy := float32(60)
	ct := rl.CameraPerspective
	camera = rl.NewCamera3D(pos, target, up, fovy, ct)

	for !rl.WindowShouldClose() {
		update()
		render()
	}
}
