package main

import (
	"context"

	termbox "github.com/nsf/termbox-go"
)

func main() {
	termbox.Init()
	defer termbox.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	disp := NewDisplay(ctx)
	players := []*Player{
		NewPlayer(LEFT),
		NewPlayer(RIGHT),
	}
	NewGame(
		players[LEFT],
		players[RIGHT],
		NewButton([]Reserve{Reserve{termbox.KeySpace, STOP, players[LEFT]}}),
		NewButton([]Reserve{Reserve{termbox.KeyEnter, STOP, players[RIGHT]}}),
		disp,
	).Start()
}
