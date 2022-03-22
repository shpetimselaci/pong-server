package pong

import (
	"github.com/hajimehoshi/ebiten"
)

// Position is a set of coordinates in 2-D plan
type Position struct {
	X, Y float32
}

// GetCenter returns the center position on screen
func GetCenter() Position {
	w, h := WindowWidth, WindowHeight
	return Position{
		X: float32(w / 2),
		Y: float32(h / 2),
	}
}

func IsKeyJustPressed(action UserAction, key ebiten.Key) bool {
	return action.Action == 1 && key == action.Key
}

func IsKeyJustReleased(action UserAction, key ebiten.Key) bool {
	return action.Action == 0 && key == action.Key
}
