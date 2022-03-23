# Pong in Go

This project is a simpler version of https://github.com/dstoiko/go-pong-wasm,
used as a demonstration of a native app running pong in golang. If working 
on [the pong exercise](https://docs.google.com/document/d/1Q8iGyyG-pv1GJTFA7hctVbwlHdvh0YM2JzmrVLl4dOE/edit),
you can use the logic in this project as an inspiration.

However, there are crucial differences between this project and the pong exercise:

 - The pong exercise is a online multiplayer pong written using Golang and React. This project is a native app letting two players play locally.
 - The pong exercise expect the client side displaying the game to be written in React, and the server side only maintaining state. This project actually draws the game state using Golang and the Ebiten framework.

## Features

- [x] Works on desktop (Linux, MacOS, Windows)
- [x] 2-player "VS" mode with same keyboard

## Build instructions

First, `git clone` and `cd` into this repo.

1. Run `make build` to build for native desktop (Linux, MacOS, Windows)
2. Run `make run` to start running the game.

## Caution

This project was written having no prior experience with GO language, thus I would recommend not taking this an example to write something off from.


## ToDo List

* ~~Learn how to write and understand Golang to some extent~~
* ~~Remove unnecessary drawing~~
* ~~Add web sockets~~
* ~~Make sure state is updated when user sends action~~
* ~~Send socket updates~~
* Create CI/CD to deploy at GCP
* ~~Avoid having more than 2 web socket connections to prevent more than 2 players~~