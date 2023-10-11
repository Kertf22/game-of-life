package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Life struct {
	status bool
	// posX   int32
	// posY   int32
	width  int32
	height int32
}

type Environment struct {
	lifes  []Life
	board  [160][160]Life
	count  int
	status bool
}

func (e *Environment) rules() {
	var local_board [160][160]Life
	copy(local_board[:], e.board[:])

	for i := range e.board {
		for j := range e.board[i] {
			count := 0

			for k := i - 1; k <= i+1; k++ {
				for l := j - 1; l <= j+1; l++ {
					if k < 0 || l < 0 || k > 159 || l > 159 || (k == i && l == j) {
						continue
					}
					if e.board[k][l].status {
						count++
					}
				}
			}

			if count < 2 || count > 3 {
				local_board[i][j].status = false
			}

			if count == 3 {
				local_board[i][j].status = true
			}
		}
	}

	e.board = local_board
}

// Any live cell with fewer than two live neighbours dies, as if by underpopulation.
// Any live cell with more than three live neighbours dies, as if by overpopulation.
// Any dead cell with three live neighbours becomes a live cell.

func (e *Environment) handleMouse() {
	mouse := rl.GetMousePosition()
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		posX := int32(mouse.X) - int32(mouse.X)%4
		posY := int32(mouse.Y) - int32(mouse.Y)%4

		life := Life{true, 4, 4}
		e.board[posX/4][posY/4] = life
	}

	if rl.IsMouseButtonPressed(rl.MouseRightButton) {
		e.status = !e.status
	}
}
func (e *Environment) draw() {
	for i := range e.board {
		for j := range e.board[i] {
			if e.board[i][j].status {
				rl.DrawRectangle(int32(i*4), int32(j*4), 4, 4, rl.Black)
			}
		}
	}
}

func main() {
	rl.InitWindow(640, 640, "raylib [core] example - basic window")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	board := [160][160]Life{}

	for i := range board {
		for j := range board[i] {
			board[i][j] = Life{false, 4, 4}
		}
	}

	e := Environment{
		lifes:  make([]Life, 0),
		count:  0,
		status: false,
		board:  board,
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		text := ""
		if e.status {
			text = "Running"
		} else {
			text = "Paused"
		}
		rl.DrawText(text, 10, 10, 20, rl.Black)

		if e.count%10 == 0 && e.status {
			e.rules()
		}

		e.draw()
		e.handleMouse()
		e.count++
		rl.EndDrawing()
	}
}
