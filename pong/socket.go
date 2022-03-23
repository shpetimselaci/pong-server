package pong

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

// type StateResponse struct {
// 	GameState pong.GameState `json: "gameState"`
// 	Player1   pong.Paddle    `json: "player1"`
// 	Player2   pong.Paddle    `json: "player2"`
// 	Ball      pong.Ball      `json: "ball"`
// }

func sendGameState(g *Game) error {

	out, err := json.Marshal(g)

	if err != nil {
		return err
	}

	return wsutil.WriteServerText(*g.Ws, out)
}

func ListenAndServe(g *Game) {
	fmt.Println("Listening!")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		r.Header.Set("Access-Control-Allow-Origin", "*")

		fmt.Println("client", r.Host, r.Method, r.URL.Host)
		conn, _, _, err := ws.DefaultHTTPUpgrader.Upgrade(r, w)
		g.Ws = &conn
		if err != nil {
			// handle error
			fmt.Println(err, "err")
		}

		go func() {
			defer conn.Close()

			for {
				msg, _, _ := wsutil.ReadClientData(conn)

				if err != nil {
					fmt.Println(err, "something happened")
					// handle error

					return
				}

				userAction := &UserAction{}
				if err := json.Unmarshal(msg, &userAction); err != nil {
					fmt.Println("err", err, msg)
				}

				g.UpdateGameState(*userAction)
			}
		}()
	})
	err := http.ListenAndServe("0.0.0.0:8080", nil)

	if err != nil {
		fmt.Println("Could not start server", err)
		panic(err)
	}
}
