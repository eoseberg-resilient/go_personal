package locations

import (
	//"fmt"

	"PERSONAL/txt_game/player"
)

type Location interface {
	GetDisplay() string
}

type LocationImpl struct {
	location_indx int
	name string
	desc string

	chars []character.Character
}

func NewLocation(indx int, name string, desc string) Location {
	return &LocationImpl{
		location_indx: indx,
		name: name,
		desc: desc,
	}
}

func (loc LocationImpl) GetDisplay() string {
	return loc.name + "\n" + loc.desc
}
