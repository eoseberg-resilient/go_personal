package character

import (
	"strings"
	"errors"
)

//"fmt"

type CmdIntention int;
type CmdSynonym int;

const (
	Describe = iota + 1
	Do
	Ask
)



type Cmd interface {
	Intention() string
	Synonyms() []string
}

type CmdImpl struct {
	intention string
	synonyms []string
}

func NewCmd(intention string, synonyms []string) Cmd {
	return &CmdImpl {
		intention: intention,
		synonyms: synonyms,
	}
}

func (cmd CmdImpl) Intention() string {
	return cmd.intention
}

func (cmd CmdImpl) Synonyms() []string {
	return cmd.synonyms
}

type CmdHandler interface {
	TryGetCmd(string) (Cmd, error)
	SetCmd(Cmd) error
}

type PlayerCmds struct {
	cmds []Cmd
}

// TODO more logic here
func (pcmds PlayerCmds) TryGetCmd(inp string)(Cmd, error) {
	for _, cmd := range pcmds.cmds {
		for _, s := range cmd.Synonyms() {
			if strings.Contains(s, inp) {
				return cmd, nil
			}
		}
	}

	return CmdImpl{}, errors.New("Not found")
}

func (pcmds PlayerCmds) SetCmd(cmd Cmd) error {
	return nil
}

func NewPlayerCmdHandler() CmdHandler {
	// TODO get file and extract cmds
	// for now we hardcode

	cmds := []Cmd{
		NewCmd("describe", []string{
			"tell me about",
			"describe",
			"what is",
			"what are",
		}),
	}

	return &PlayerCmds{
		cmds: cmds,
	}
}
