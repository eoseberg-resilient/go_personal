package system

import (
	"fmt"
	"strings"
)

type OutputHandler interface {
	WriteOutput() bool

	SetOutputText(s string) error
}

type TextOutputHandlerImpl struct {
	column_width int
	tick_length int

	text string
	slice TextSlice
}

type TextSlice struct {
	start_index int
	stop_index int
}

type MenuOutputHandler struct {
	column_width int
	tick_length int

	menu []string
	indx int
}

func NewTextOutputHandler(columnWidth int, tickLength int) OutputHandler {
	return &TextOutputHandlerImpl{
		column_width: columnWidth,
		tick_length: tickLength,
		slice: TextSlice{
			start_index: 0,
			stop_index: 0,
		},
		text: "",
	}
}

func (handl *TextOutputHandlerImpl) WriteOutput() bool {
	isDone := false
	handl.slice.start_index = handl.slice.stop_index
	handl.slice.stop_index += handl.tick_length
	if handl.slice.stop_index > len(handl.text) {
		handl.slice.stop_index = len(handl.text)
		isDone = true
	}

	fmt.Print(handl.text[handl.slice.start_index:handl.slice.stop_index])

	return isDone
}

func (handl *TextOutputHandlerImpl) SetOutputText(s string) error {
	handl.slice = TextSlice{}
	handl.text = string(s)
	return nil
}

func NewMenuOutputHandler(columnWidth int, tickLength int) OutputHandler {
	return &MenuOutputHandler{
		column_width: columnWidth,
		tick_length: tickLength,
		menu: make([]string, 0),
		indx: 0,
	}
}

func (handl *MenuOutputHandler) WriteOutput() bool {
	isDone := false
	indx_end := handl.indx+handl.tick_length
	if indx_end > len(handl.menu) {
		indx_end = len(handl.menu)
		isDone = true
	}
	for i := handl.indx; i < indx_end; i++ {
		fmt.Println("  " + handl.menu[i])
	}
	handl.indx = indx_end
	return isDone
}

func (handl *MenuOutputHandler) SetOutputText(s string) error {
	handl.indx = 0
	handl.menu = strings.Split(s, "\n")
	return nil
}
