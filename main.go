package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/shpetimselaci/pong-server/pong"
)

func main() {
	g := pong.NewGame()
	fmt.Println("screen size")
	go pong.ListenAndServe(g)

	ebiten.SetRunnableOnUnfocused(true)
	if err := ebiten.RunGame(g); err != nil {
		fmt.Println("Err", err)
		panic(err)
	}

}
