package main

import (
	"context"
	"time"

	"github.com/nsf/termbox-go"
)

const (
	LEFT = iota
	RIGHT
	SIDES
)

type Time struct {
	min   int
	sec   int
	blink bool
	disp  bool
	color int
}
type Display struct {
	times [SIDES]Time
}

const Y_PER_DISIT = 7
const X_PER_DISIT = 5
const X_PADDING = 1
const X_MARGIN = 7
const Y_MARGIN = 7
const DIV = 10
const CHARCTOR = 'X'

const BLINK_PERIOD_MS = 200

const COLOR_DEFAULT = 15
const COLOR_WIN = 2
const COLOR_LOSE = 1
const COLOR_CREATED_BY = 6

var COLON_MAP = [][]bool{
	{false, false, false, false, false},
	{false, false, true, false, false},
	{false, false, false, false, false},
	{false, false, false, false, false},
	{false, false, false, false, false},
	{false, false, true, false, false},
	{false, false, false, false, false}}

var NUM_MAP = [][][]bool{
	{{true, true, true, true, true},
		{true, false, false, false, true},
		{true, false, false, false, true},
		{true, false, false, false, true},
		{true, false, false, false, true},
		{true, false, false, false, true},
		{true, true, true, true, true}},
	{{false, true, true, false, false},
		{false, false, true, false, false},
		{false, false, true, false, false},
		{false, false, true, false, false},
		{false, false, true, false, false},
		{false, false, true, false, false},
		{false, true, true, true, false}},
	{{true, true, true, true, true},
		{false, false, false, false, true},
		{false, false, false, false, true},
		{true, true, true, true, true},
		{true, false, false, false, false},
		{true, false, false, false, false},
		{true, true, true, true, true}},
	{{true, true, true, true, true},
		{false, false, false, false, true},
		{false, false, false, false, true},
		{true, true, true, true, true},
		{false, false, false, false, true},
		{false, false, false, false, true},
		{true, true, true, true, true}},
	{{true, false, false, true, false},
		{true, false, false, true, false},
		{true, false, false, true, false},
		{true, true, true, true, true},
		{false, false, false, true, false},
		{false, false, false, true, false},
		{false, false, false, true, false}},
	{{true, true, true, true, true},
		{true, false, false, false, false},
		{true, false, false, false, false},
		{true, true, true, true, true},
		{false, false, false, false, true},
		{false, false, false, false, true},
		{true, true, true, true, true}},
	{{true, true, true, true, true},
		{true, false, false, false, false},
		{true, false, false, false, false},
		{true, true, true, true, true},
		{true, false, false, false, true},
		{true, false, false, false, true},
		{true, true, true, true, true}},
	{{true, true, true, true, true},
		{false, false, false, false, true},
		{false, false, false, false, true},
		{false, false, false, false, true},
		{false, false, false, false, true},
		{false, false, false, false, true},
		{false, false, false, false, true}},
	{{true, true, true, true, true},
		{true, false, false, false, true},
		{true, false, false, false, true},
		{true, true, true, true, true},
		{true, false, false, false, true},
		{true, false, false, false, true},
		{true, true, true, true, true}},
	{{true, true, true, true, true},
		{true, false, false, false, true},
		{true, false, false, false, true},
		{true, true, true, true, true},
		{false, false, false, false, true},
		{false, false, false, false, true},
		{true, true, true, true, true}}}

func NewDisplay(ctx context.Context) *Display {
	var display Display
	for side := 0; side < SIDES; side++ {
		display.times[side].min = 0
		display.times[side].sec = 0
		display.times[side].blink = false
		display.times[side].disp = false
		display.times[side].color = COLOR_DEFAULT
	}
	go func(display *Display, ctx context.Context) {
		t := time.NewTicker(BLINK_PERIOD_MS * time.Millisecond)
		for {
			select {
			case <-ctx.Done():
				return
			case <-t.C:
				termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
				printTitle()
				offset := 0
				printBothTime(&display.times, offset)
				termbox.Flush()
			}
		}
	}(&display, ctx)
	return &display
}

func (display *Display) Print(side int, sec int) {
	display.times[side].min = sec / 60
	display.times[side].sec = sec % 60
	display.times[side].disp = true
}

func (display *Display) BlinkOn(side int) {
	display.times[side].blink = true
}

func (display *Display) BlinkOff(side int) {
	display.times[side].blink = false
}

func (display *Display) SetColor(side int, color int) {
	display.times[side].color = color
}

func printTitle() {
	x := 3
	termbox.SetCell(x, 1, 'F', termbox.Attribute(COLOR_DEFAULT+1), termbox.ColorDefault)
	x++
	termbox.SetCell(x, 1, 'u', termbox.Attribute(COLOR_DEFAULT+1), termbox.ColorDefault)
	x++
	termbox.SetCell(x, 1, 's', termbox.Attribute(COLOR_DEFAULT+1), termbox.ColorDefault)
	x++
	termbox.SetCell(x, 1, 'i', termbox.Attribute(COLOR_DEFAULT+1), termbox.ColorDefault)
	x++
	termbox.SetCell(x, 1, 'c', termbox.Attribute(COLOR_DEFAULT+1), termbox.ColorDefault)
	x++
	termbox.SetCell(x, 1, ' ', termbox.Attribute(COLOR_DEFAULT+1), termbox.ColorDefault)
	x++
	termbox.SetCell(x, 1, 'S', termbox.Attribute(COLOR_DEFAULT+1), termbox.ColorDefault)
	x++
	termbox.SetCell(x, 1, 'h', termbox.Attribute(COLOR_DEFAULT+1), termbox.ColorDefault)
	x++
	termbox.SetCell(x, 1, 'o', termbox.Attribute(COLOR_DEFAULT+1), termbox.ColorDefault)
	x++
	termbox.SetCell(x, 1, 'g', termbox.Attribute(COLOR_DEFAULT+1), termbox.ColorDefault)
	x++
	termbox.SetCell(x, 1, 'i', termbox.Attribute(COLOR_DEFAULT+1), termbox.ColorDefault)
	x++
	termbox.SetCell(x, 1, '-', termbox.Attribute(COLOR_DEFAULT+1), termbox.ColorDefault)
	x++
	termbox.SetCell(x, 1, 'C', termbox.Attribute(COLOR_DEFAULT+1), termbox.ColorDefault)
	x++
	termbox.SetCell(x, 1, 'l', termbox.Attribute(COLOR_DEFAULT+1), termbox.ColorDefault)
	x++
	termbox.SetCell(x, 1, 'o', termbox.Attribute(COLOR_DEFAULT+1), termbox.ColorDefault)
	x++
	termbox.SetCell(x, 1, 'c', termbox.Attribute(COLOR_DEFAULT+1), termbox.ColorDefault)
	x++
	termbox.SetCell(x, 1, 'k', termbox.Attribute(COLOR_DEFAULT+1), termbox.ColorDefault)
	x = 54
	termbox.SetCell(x, 19, 'C', termbox.Attribute(COLOR_CREATED_BY+1), termbox.ColorDefault)
	x++
	termbox.SetCell(x, 19, 'r', termbox.Attribute(COLOR_CREATED_BY+1), termbox.ColorDefault)
	x++
	termbox.SetCell(x, 19, 'e', termbox.Attribute(COLOR_CREATED_BY+1), termbox.ColorDefault)
	x++
	termbox.SetCell(x, 19, 'a', termbox.Attribute(COLOR_CREATED_BY+1), termbox.ColorDefault)
	x++
	termbox.SetCell(x, 19, 't', termbox.Attribute(COLOR_CREATED_BY+1), termbox.ColorDefault)
	x++
	termbox.SetCell(x, 19, 'e', termbox.Attribute(COLOR_CREATED_BY+1), termbox.ColorDefault)
	x++
	termbox.SetCell(x, 19, 'd', termbox.Attribute(COLOR_CREATED_BY+1), termbox.ColorDefault)
	x++
	termbox.SetCell(x, 19, ' ', termbox.Attribute(COLOR_CREATED_BY+1), termbox.ColorDefault)
	x++
	termbox.SetCell(x, 19, 'b', termbox.Attribute(COLOR_CREATED_BY+1), termbox.ColorDefault)
	x++
	termbox.SetCell(x, 19, 'y', termbox.Attribute(COLOR_CREATED_BY+1), termbox.ColorDefault)
	x++
	termbox.SetCell(x, 19, ' ', termbox.Attribute(COLOR_CREATED_BY+1), termbox.ColorDefault)
	x++
	termbox.SetCell(x, 19, 'Y', termbox.Attribute(COLOR_CREATED_BY+1), termbox.ColorDefault)
	x++
	termbox.SetCell(x, 19, '.', termbox.Attribute(COLOR_CREATED_BY+1), termbox.ColorDefault)
	x++
	termbox.SetCell(x, 19, 'O', termbox.Attribute(COLOR_CREATED_BY+1), termbox.ColorDefault)
	x++
	termbox.SetCell(x, 19, 'k', termbox.Attribute(COLOR_CREATED_BY+1), termbox.ColorDefault)
	x++
	termbox.SetCell(x, 19, 'a', termbox.Attribute(COLOR_CREATED_BY+1), termbox.ColorDefault)
	x++
	termbox.SetCell(x, 19, 'z', termbox.Attribute(COLOR_CREATED_BY+1), termbox.ColorDefault)
	x++
	termbox.SetCell(x, 19, 'a', termbox.Attribute(COLOR_CREATED_BY+1), termbox.ColorDefault)
	x++
	termbox.SetCell(x, 19, 'k', termbox.Attribute(COLOR_CREATED_BY+1), termbox.ColorDefault)
	x++
	termbox.SetCell(x, 19, 'i', termbox.Attribute(COLOR_CREATED_BY+1), termbox.ColorDefault)
	x++
}

func printBothTime(times *[SIDES]Time, offset int) int {
	for side := 0; side < SIDES; side++ {
		offset = printTime(&times[side], offset)
		offset++
	}
	return offset
}
func printTime(time *Time, offset int) int {
	if time.blink == true {
		time.disp = !(time.disp)
	} else {
		time.disp = true
	}
	if time.disp == true {
		offset = printTowDisit(time.min, time.color, offset)
		offset = printColon(time.color, offset)
		offset = printTowDisit(time.sec, time.color, offset)
	} else {
		offset += 2 + 1 + 2
	}
	return offset
}

func printTowDisit(num int, color int, offset int) int {
	for disit := 0; disit < 2; disit++ {
		output := num / DIV % DIV
		offset = printDisit(output, color, offset)
		num *= DIV
	}
	return offset
}

func printDisit(num int, color int, offset int) int {
	offsetX := offset*(X_PADDING+X_PER_DISIT) + X_MARGIN
	offsetY := Y_MARGIN
	for x := 0; x < X_PER_DISIT; x++ {
		for y := 0; y < Y_PER_DISIT; y++ {
			if NUM_MAP[num][y][x] {
				termbox.SetCell(x+offsetX, y+offsetY, CHARCTOR, termbox.Attribute(color+1), termbox.ColorDefault)
			}
		}
	}
	return offset + 1
}

func printColon(color int, offset int) int {
	offsetX := offset*(X_PADDING+X_PER_DISIT) + X_MARGIN
	offsetY := Y_MARGIN
	for x := 0; x < X_PER_DISIT; x++ {
		for y := 0; y < Y_PER_DISIT; y++ {
			if COLON_MAP[y][x] {
				termbox.SetCell(x+offsetX, y+offsetY, CHARCTOR, termbox.Attribute(color+1), termbox.ColorDefault)
			}
		}
	}
	return offset + 1
}
