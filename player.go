package main

import (
	"context"
	"time"
)

const COUNT_DOWN_SEC = 1

type Player struct {
	sec  int
	c    chan bool
	side int
}

func (player *Player) Notify(kind int) {
	switch kind {
	case STOP:
		player.c <- false
	case LOSE:
		player.c <- true
	}
}

func NewPlayer(side int) *Player {
	var player Player
	player.side = side
	return &player
}

func (player *Player) Turn(display *Display, button *Button) bool {
	t := time.NewTicker(COUNT_DOWN_SEC * time.Second)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	display.BlinkOn(player.side)
	defer display.BlinkOff(player.side)

	player.c = make(chan bool)
	button.Start(ctx)
	for {
		select {
		case lose := <-player.c:
			return lose
		case <-t.C:
			player.sec--
			display.Print(player.side, player.sec)
			if player.sec == 0 {
				return true
			}
		}
	}
}

func (player *Player) Reset(sec int, display *Display) {
	player.sec = sec
	display.Print(player.side, player.sec)
}

func (player *Player) Win(display *Display) {
	display.SetColor(player.side, COLOR_WIN)
}

func (player *Player) Lose(display *Display) {
	display.SetColor(player.side, COLOR_LOSE)
}
