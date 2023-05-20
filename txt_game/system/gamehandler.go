package system

import (
	"fmt"
	"time"
)

type GameState int;

const (
	StateUserInput GameState = iota
	StatePrintText
	StateViewMenu
	StateDoCommand
	StateBetween
)

type GameStateInt interface {
	Tick() bool
	Save() bool
	Wait(int)
	SetGameState(GameState) error

	State() GameState
	IsWaiting() bool
}

type GameStateHolder struct {
	curr_time time.Time
	tick_rate int64
	pause int

	state GameState
}

func (g *GameStateHolder) Tick() bool {
	now := time.Now()
	if now.Sub(g.curr_time).Milliseconds() > g.tick_rate {
		g.curr_time = now
		if g.pause == 0 {
			return true
		} else {
			g.pause -= 1
			return false
		}
	}
	return false
}

func (g GameStateHolder) Save() bool {
	panic("Save game not implemented!")
}

func (g *GameStateHolder) Wait(ticks int) {
	g.pause = ticks
}

func (g *GameStateHolder) SetGameState(s GameState) error {
	g.pause = 3
	g.state = s
	if g.state != StateBetween {
		fmt.Println()
	}
	return nil
}

func (g GameStateHolder) State() GameState {
	return g.state
}

func (g GameStateHolder) IsWaiting() bool {
	return g.pause > 0
}

func NewGame(tickrate int64, startstate GameState) GameStateInt {
	return &GameStateHolder{
		curr_time: time.Now(),
		tick_rate: tickrate,

		state: startstate,
	}
}

func (g *GameStateHolder) Pause(dur int) {
	g.pause = dur
}
