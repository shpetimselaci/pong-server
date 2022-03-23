package pong

import (
	"net"

	"github.com/hajimehoshi/ebiten"
)

type Game struct {
	State    GameState
	Ball     *Ball
	Player1  *Paddle
	Player2  *Paddle
	Rally    int
	Level    int
	MaxScore int
	Ws       *net.Conn
}

type KeyAction int

type UserAction struct {
	Key    ebiten.Key `json:"key"`
	Action KeyAction  `json:"action"`
}

const (
	initBallVelocity = 5.0
	initPaddleSpeed  = 10.0
	speedUpdateCount = 6
	speedIncrement   = 0.5
)

const (
	windowWidth  = 800
	windowHeight = 600
)

// func respondToPlayerAction(action UserAction, g Game) {

// 	g.player1.
// }
// NewGame creates an initializes a new game
func NewGame() *Game {
	g := &Game{}
	g.init("player1", "player2")
	return g
}

func (g *Game) init(player1Name string, player2Name string) {
	g.State = StartState
	g.MaxScore = 11

	g.Player1 = &Paddle{
		Position: Position{
			X: InitPaddleShift,
			Y: float32(windowHeight / 2)},
		Score:      0,
		Speed:      initPaddleSpeed,
		Width:      InitPaddleWidth,
		Height:     InitPaddleHeight,
		Up:         ebiten.KeyW,
		Down:       ebiten.KeyS,
		PlayerName: player1Name,
	}

	g.Player2 = &Paddle{
		Position: Position{
			X: windowWidth - InitPaddleShift - InitPaddleWidth,
			Y: float32(windowHeight / 2)},
		Score:      0,
		Speed:      initPaddleSpeed,
		Width:      InitPaddleWidth,
		Height:     InitPaddleHeight,
		Up:         ebiten.KeyUp,
		Down:       ebiten.KeyDown,
		PlayerName: player2Name,
	}
	g.Ball = &Ball{
		Position: Position{
			X: float32(windowWidth / 2),
			Y: float32(windowHeight / 2)},
		Radius:    InitBallRadius,
		XVelocity: initBallVelocity,
		YVelocity: initBallVelocity,
	}
	g.Level = 0
}

func (g *Game) reset(state GameState) {
	w := windowWidth
	g.State = state
	g.Rally = 0
	g.Level = 0
	if state == StartState {
		g.Player1.Score = 0
		g.Player2.Score = 0
	}
	g.Player1.Position = Position{
		X: InitPaddleShift, Y: GetCenter().Y}
	g.Player2.Position = Position{
		X: float32(w - InitPaddleShift - InitPaddleWidth), Y: GetCenter().Y}
	g.Ball.Position = GetCenter()
	g.Ball.XVelocity = initBallVelocity
	g.Ball.YVelocity = initBallVelocity
}

// Update updates the game state

func (g *Game) UpdateGameState(action UserAction) error {
	switch g.State {
	case StartState:
		if IsKeyJustPressed(action, ebiten.KeySpace) {
			g.State = PlayState
		}

	case PlayState:
		w := windowWidth

		g.Player1.Update(action)
		g.Player2.Update(action)

		xV := g.Ball.XVelocity
		g.Ball.Update(g.Player1, g.Player2)
		// rally count
		if xV*g.Ball.XVelocity < 0 {
			// score up when ball touches human player's paddle
			if g.Ball.X < float32(w/2) {
				g.Player1.Score++
			}

			g.Rally++

			// spice things up
			if (g.Rally)%speedUpdateCount == 0 {
				g.Level++
				g.Ball.XVelocity += speedIncrement
				g.Ball.YVelocity += speedIncrement
				g.Player1.Speed += speedIncrement
				g.Player2.Speed += speedIncrement
			}
		}

		if g.Ball.X < 0 {
			g.Player2.Score++
			g.reset(StartState)
		} else if g.Ball.X > float32(w) {
			g.Player1.Score++
			g.reset(StartState)
		}

		if g.Player1.Score == g.MaxScore || g.Player2.Score == g.MaxScore {
			g.State = GameOverState
		}

	case GameOverState:
		if IsKeyJustPressed(action, ebiten.KeySpace) {
			g.reset(StartState)
		}
	}
	return nil
}

func (g *Game) Update(screen *ebiten.Image) error {

	g.UpdateGameState(UserAction{})
	g.Draw(screen)
	// screen.Fill(g.BgColor)

	// y, x := screen.Size()
	// fmt.Println("screen", y, x)
	// fmt.Println("Rendering", ebiten.CurrentTPS())
	return nil
}

// Draw updates the game screen elements drawn
func (g *Game) Draw(screen *ebiten.Image) error {
	screen.Fill(BgColor)
	g.Player1.Draw(screen)
	g.Player2.Draw(screen)
	g.Ball.Draw(screen)
	if g.Ws != nil {
		sendGameState(g)
	}
	// ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
	// fmt.Println("Rendering", ebiten.CurrentTPS())
	return nil
}

// // Layout sets the screen layout
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}
