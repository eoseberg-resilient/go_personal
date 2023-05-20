package character

import (
	//"PERSONAL/txt_game/locations"
)

type Character interface {
	Display()

	SetInventory(map[string]string)
	Inventory() Menu

	//TryGetCmd() Cmd
}


type Player struct {
	health 	int
	currloc	int
	inv		Menu

	cmd_handler CmdHandler
}

type Position struct {
	room int
	pos  int
}

func NewPlayer() Character {
	return &Player {
		health: 10,
		inv: NewInventory(),
		currloc: -1,
		cmd_handler: NewPlayerCmdHandler(),
	}
}

func (p Player) Display() {

}

func (p Player) Inventory() Menu {
	return p.inv
}

func (p *Player) SetInventory(items map[string]string) {
	for k, v := range items {
		p.inv.Add(NewItem(k, v))
	}
}

func (p *Player) SetCmd(s string) {
	cmd, err := p.cmd_handler.TryGetCmd(s)
	if err != nil {
		panic("Not implemented")
	}

	p.cmd_handler.SetCmd(cmd)
}
