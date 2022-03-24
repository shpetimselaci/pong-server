package pong

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/hajimehoshi/ebiten"
)

type StateResponse struct {
	State    GameState
	Ball     *Ball
	Player1  *Paddle
	Player2  *Paddle
	Rally    int
	Level    int
	MaxScore int
	KeyPads  []ebiten.Key
}

type CloseResponse struct {
	Placeholder string
	TryAgain    bool
}

func sendGameState(g *Game) error {

	state := StateResponse{
		Player1:  g.Player1,
		Player2:  g.Player2,
		State:    g.State,
		Ball:     g.Ball,
		MaxScore: g.MaxScore,
		Level:    g.Level,
		Rally:    g.Rally,
	}

	for _, playerConnection := range g.Ws {
		state.KeyPads = playerConnection.padKeys

		out, err := json.Marshal(state)
		if err != nil {
			return err
		}
		msgErr := wsutil.WriteServerText(*playerConnection.Connection, out)
		if msgErr != nil {
			return err
		}

	}
	return nil
}

func contains(elems []ebiten.Key, v ebiten.Key) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func ListenAndServe(g *Game) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		r.Header.Set("Access-Control-Allow-Origin", "*")
		conn, _, _, err := ws.DefaultHTTPUpgrader.Upgrade(r, w)
		if len(g.Ws) == 2 {
			out, _ := json.Marshal(CloseResponse{Placeholder: "Two players already connected!", TryAgain: true})

			wsutil.WriteServerText(conn, out)
			conn.Close()
			return
		}

		pads := g.freePads[0:2]
		fmt.Println("prev", g.freePads)

		g.freePads = g.freePads[2:]
		g.Ws[conn] = PlayerConnection{
			Connection: &conn,
			padKeys:    pads,
		}
		if err != nil {
			// handle error
			fmt.Println(err, "err")
		}

		go func() {
			defer func() {
				if g.Ws[conn].padKeys != nil {
					g.freePads = append(g.freePads, g.Ws[conn].padKeys...)
					fmt.Println(g.freePads)
				}

				conn.Close()
				delete(g.Ws, conn)
			}()

			for {
				msg, _, clientConnError := wsutil.ReadClientData(conn)
				if clientConnError != nil {
					fmt.Println(clientConnError, "something happened")
					// handle error
					return
				}

				userAction := &UserAction{}

				if err := json.Unmarshal(msg, &userAction); err == nil {

					if contains([]ebiten.Key{ebiten.KeySpace}, userAction.Key) || contains(g.Ws[conn].padKeys, userAction.Key) {
						g.UpdateGameState(*userAction)
					}

				}

			}
		}()
	})
	err := http.ListenAndServe("0.0.0.0:8000", nil)

	if err != nil {
		fmt.Println("Could not start server", err)
		panic(err)
	}
}
