package character

import (
	"fmt"
	"strings"
)


type Menu interface {
	Display() string

	Add(Item)
	Remove()
}

type Item interface {
	GetDisplay() string

	Name() string
	Desc() string
}

type ItemImpl struct {
	name string
	desc string
}

func NewItem(name string, desc string) Item {
	return &ItemImpl{
		name: name,
		desc: desc,
	}
}

func (i ItemImpl) GetDisplay() string {
	return fmt.Sprint("Name: ", i.name, "Description: ", i.desc)
}

func (i ItemImpl) Name() string {
	return strings.ToUpper(i.name)
}

func (i ItemImpl) Desc() string {
	return i.desc
}
