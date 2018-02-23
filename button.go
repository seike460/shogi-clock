package main

import (
	"context"

	termbox "github.com/nsf/termbox-go"
)

const (
	STOP = iota
	LOSE
	PLAY
	RESET
	QUIT
)

type Notified interface {
	Notify(kind int)
}

type Reserve struct {
	key      termbox.Key
	kind     int
	notified Notified
}

type Button struct {
	reserves []Reserve
}

func NewButton(reserves []Reserve) *Button {
	var button Button
	button.reserves = reserves
	return &button
}

func (button *Button) Start(ctx context.Context) {
	go func(button *Button, ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
			default:
				switch ev := termbox.PollEvent(); ev.Type {
				case termbox.EventKey:
					for _, reserve := range button.reserves {
						if reserve.key == ev.Key {
							reserve.notified.Notify(reserve.kind)
						}
					}
				}
			}
		}
	}(button, ctx)
}
