package main

import (
	"context"

	termbox "github.com/nsf/termbox-go"
)

type Game struct {
	c            chan int
	firstPlayer  *Player
	secondPlayer *Player
	firstButton  *Button
	secondButton *Button
	display      *Display
}

func NewGame(firstPlayer *Player, secondPlayer *Player, firstButton *Button, secondButton *Button, display *Display) *Game {
	var game Game
	game.firstPlayer = firstPlayer
	game.secondPlayer = secondPlayer
	game.firstButton = firstButton
	game.secondButton = secondButton
	game.display = display
	return &game
}

func (game *Game) Start() {
	for {
		game.firstPlayer.Reset(10, game.display)
		game.secondPlayer.Reset(10, game.display)
		{
			game.c = make(chan int)
			ctx, cancel := context.WithCancel(context.Background())
			NewButton(
				[]Reserve{
					Reserve{termbox.KeyCtrlP, PLAY, game},
					Reserve{termbox.KeyCtrlQ, QUIT, game},
				},
			).Start(ctx)
			kind := <-game.c
			cancel()
			if kind == QUIT {
				break
			}
		}
		for {
			if game.firstPlayer.Turn(game.display, game.firstButton) {
				game.firstPlayer.Lose(game.display)
				game.secondPlayer.Win(game.display)
				break
			}
			if game.secondPlayer.Turn(game.display, game.secondButton) {
				game.firstPlayer.Win(game.display)
				game.secondPlayer.Lose(game.display)
				break
			}
		}
		{
			game.c = make(chan int)
			ctx, cancel := context.WithCancel(context.Background())
			NewButton(
				[]Reserve{
					Reserve{termbox.KeyCtrlR, PLAY, game},
					Reserve{termbox.KeyCtrlQ, QUIT, game},
				},
			).Start(ctx)
			kind := <-game.c
			cancel()
			if kind == QUIT {
				break
			}
		}
	}
}

func (game *Game) Notify(kind int) {
	game.c <- kind
}
