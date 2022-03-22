package pong

import (
	"image/color"
)

const (
	InitBallVelocity = 5.0
	InitPaddleSpeed  = 10.0
	SpeedUpdateCount = 6
	SpeedIncrement   = 0.5
)

const (
	WindowWidth  = 800
	WindowHeight = 600
)

const (
	KeyPressed  KeyAction = iota
	KeyReleased           = 1
)

type GameState byte

const (
	StartState GameState = iota
	PlayState
	GameOverState
)

var (
	BgColor  = color.Black
	ObjColor = color.RGBA{120, 226, 160, 255}
)
