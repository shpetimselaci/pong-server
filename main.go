package main

import (
	"fmt"
	"runtime"

	"github.com/hajimehoshi/ebiten"
	"github.com/shpetimselaci/pong-server/pong"
)

func main() {
	runtime.GOMAXPROCS(2) // Set the maximum number of threads/processes

	g := pong.NewGame()
	fmt.Println("screen size")
	go pong.ListenAndServe(g)
	ebiten.SetRunnableOnUnfocused(true)
	if err := ebiten.RunGame(g); err != nil {
		fmt.Println("Err", err)
		panic(err)
	}

}
