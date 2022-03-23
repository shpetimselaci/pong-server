package pong

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

type StateResponse struct {
	State    GameState
	Ball     *Ball
	Player1  *Paddle
	Player2  *Paddle
	Rally    int
	Level    int
	MaxScore int
}

func sendGameState(g *Game) error {

	out, err := json.Marshal(StateResponse{
		Player1:  g.Player1,
		Player2:  g.Player2,
		State:    g.State,
		Ball:     g.Ball,
		MaxScore: g.MaxScore,
		Level:    g.Level,
		Rally:    g.Rally,
	})
	if err != nil {
		return err
	}
	for _, conn := range g.Ws {
		msgErr := wsutil.WriteServerText(*conn, out)
		if msgErr != nil {
			return err
		}

	}
	return nil
}

func ListenAndServe(g *Game) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		r.Header.Set("Access-Control-Allow-Origin", "*")

		fmt.Println("client", r.Host, r.Method, r.URL.Host)
		conn, _, _, err := ws.DefaultHTTPUpgrader.Upgrade(r, w)

		g.Ws[conn] = &conn
		if err != nil {
			// handle error
			fmt.Println(err, "err")
		}

		go func() {
			defer func() {
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

				if err := json.Unmarshal(msg, &userAction); err != nil {
					//
				}
				if err == nil {
					fmt.Println(userAction.Action, userAction.Key)
					g.UpdateGameState(*userAction)
				}
			}
		}()
	})
	err := http.ListenAndServe("0.0.0.0:8080", nil)

	if err != nil {
		fmt.Println("Could not start server", err)
		panic(err)
	}
}
