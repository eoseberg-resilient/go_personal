package system

import (
	"fmt"

	"github.com/mattn/go-tty"
)

type KeyInput int;

const (
	Enter 		KeyInput = -1
	Backspace	KeyInput = -2
	Tab			KeyInput = -3
)

type InputHandler interface {
	InitInput() error
	TearDown()

	CheckForKeypress() error
	OnKeyPress(k KeyInput) error

	InputDone() bool
	GetInput() string

	Display()
}

type InputHandlerImpl struct {
	key_event *tty.TTY
	key_map map[rune]KeyInput
	input_channel chan KeyInput
	curr_input string

	IsDone bool
}

func NewInputHandler() InputHandler {
	// INIT KEYPRESS CODES
	keymap := map[rune]KeyInput {
		127: Backspace,
		9: Tab,
		13: Enter,
	}
	tty, err := tty.Open()
	if err != nil {
		panic("TTY error")
	}
	return &InputHandlerImpl {
			input_channel: make(chan KeyInput),
			key_event: tty,
			key_map: keymap,
			curr_input: "",

			IsDone: false,
		}
}

func (handl *InputHandlerImpl) InitInput() error {
	handl.curr_input = ""
	handl.IsDone = false
	go func() {
		for {
			r, err := handl.key_event.ReadRune()
			if err == nil {
				val, has := handl.key_map[r]
				if !has {
					handl.input_channel <- KeyInput(r)
				} else {
					handl.input_channel <- val
				}
			}
			if handl.IsDone {
				break
			}
		}
	}()
	return nil
}

func (handl *InputHandlerImpl) TearDown() {
	handl.key_event.Close()
	close(handl.input_channel)
}

func (handl *InputHandlerImpl) CheckForKeypress() error {
	select {
	case keypress, ok := <- handl.input_channel: {
		if !ok {
			//fmt.Println("CLOSED FOR BUSINESS")
		} else {
			handl.OnKeyPress(keypress)
		}
		}
	}
	return nil
}

func (handl *InputHandlerImpl) OnKeyPress(k KeyInput) error {
	switch k {
	case Backspace: {
		//fmt.Println("Backspace")
		if len(handl.curr_input) > 0 {
			handl.curr_input = handl.curr_input[:len(handl.curr_input)-1]
		}
	}
	case Enter: {
		//fmt.Println("Enter")
		handl.IsDone = true
	}
	case Tab: {
		panic("ERR: Backspace not implemented")
	}
	default: {
		handl.curr_input += string(k)
	}
	}

	return nil
}

func (handl InputHandlerImpl) InputDone() bool {
	return handl.IsDone
}

func (handl InputHandlerImpl) GetInput() string {
	return handl.curr_input
}

func (handl InputHandlerImpl) Display() {
	fmt.Print("\r                                                                                                       ")
	fmt.Print("\r>>\t"+handl.curr_input)
}

func (handl InputHandlerImpl) Input() string {
	return handl.curr_input
}

func GiveCommand(s string) error {
	fmt.Print("   This is a test CMD")
	fmt.Println()
	return nil
}
