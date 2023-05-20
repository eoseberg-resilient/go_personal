package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mattn/go-tty"
)

func TryGetWordSearch(s string) {
	outp(s)
}

func GetKeypress(ch chan KeyEvent) {
	tty, err := tty.Open()
	defer tty.Close()
	if err != nil {
		panic(err)
	}
	for {
		r, err := tty.ReadRune()
		if err != nil || r == 13 { //break on newline (enter key) or error
			close(ch)
			return
		}
		if r == 127 { // backspace char
			ch <- KeyEvent{
				isBackspace: true,
				isTab: false,
				letter: "",
			}
		} else if r == 9 { // tab char
			ch <- KeyEvent{
				isBackspace: false,
				isTab: true,
				letter: "",
			}
		} else {
			ch <- DoKey(r)
		}
	}
}

type KeyEvent struct {
	letter string
	isBackspace bool
	isTab bool
}

func DoKey(r rune) KeyEvent {
	return KeyEvent{
		letter: string(r),
		isBackspace: false,
		isTab: false,
	}
}

func main() {
	inp_ch := make(chan KeyEvent)
	go GetKeypress(inp_ch)

	noutp("Welcome to Verb Phrase Checker!")

	loader := NewVerbLoader()
	lookup, err := loader.Load("verbs.txt", "phrases.txt")
	if err != nil {
		panic(err)
	}
	noutp(fmt.Sprint("Loaded ", lookup.words.Len()," words!"))

	noutp("Please enter your word:")
	noutp("")
	inp_buffer := make([]string, 0)
	var tryget_words []string
	tictoc := false
	selected_indx := -1
	render := 0
	stdinloop:for {
		select {
		case keypress, ok := <- inp_ch:
			if !ok {
				break stdinloop
			} else {
				if keypress.isBackspace {
					selected_indx = -1
					if len(inp_buffer) > 0 {
						inp_buffer = inp_buffer[:len(inp_buffer)-1]
					}
					render = 0
				} else if keypress.isTab {
					if selected_indx != len(tryget_words)-1 {
						selected_indx += 1
					} else {
						selected_indx = 0
					}
					render = 1
				} else {
					selected_indx = -1
					inp_buffer = append(inp_buffer, keypress.letter)
					render = 0
				}
				s := strings.Join(inp_buffer, "")
				tryget_words, err = lookup.TryGetWordFromString(s)
			}
		case <-time.After(1 * time.Second):
			render = 1
		}

		switch render {
		case 0: {
			if tictoc {
				clearline()
				fmt.Print("\033[F") //return to start of prev line
				tictoc = false
			}
			s := strings.Join(inp_buffer, "")
			outp(s)
			}
		case 1: {
			if len(inp_buffer) == 0	{
				continue
			}
			if !tictoc {
				fmt.Println()
				tictoc = true
			}

			wordlist := "["
			for i, s := range tryget_words {
				prefix := "   "
				if selected_indx == i {
					prefix = " ->"
				}
				wordlist += prefix + s
			}
			wordlist += "   ]"
			clearline()
			fmt.Print("\r"+wordlist)
			}
		}
	}
	if err == nil {
		text := strings.Join(inp_buffer, "")
		if selected_indx != -1 {
			text = tryget_words[selected_indx]
		}
		text = TrimInput(text)
		if tictoc {
			fmt.Print("\033[F")
			tictoc = false
		}
		clearline()
		fmt.Print("\033[F") //return to start of prev line
		noutp("You entered: " + text)

		res, err := lookup.GetPhrasesFromWord(text)
		if err == WordNotFoundError {
			noutp("Sorry! We don't seem to have '" + text + "' in the system.")
			noutp("Consider adding it!")
			fmt.Println()
			os.Exit(0)
		}
		fmt.Println()
		soutp("There are", len(res.phrases), "phrases using the word", text, "!")
		noutp("These are: ")
		fmt.Println(res.phrases)
	}
}

func TrimInput(inp string) string {
	if strings.HasSuffix(inp, "\n") {
		inp = inp[:len(inp)-len("\n")]
	}
	return inp
}

func clearline() {
	fmt.Print("\r                                                 ")
}

func outp(s string) {
	clearline()
    fmt.Print("\r-> " + s)
}
func noutp(s string) {
	fmt.Println()
	clearline()
    fmt.Print("\r-> " + s)
}
func soutp(ss ...any) {
	outp(fmt.Sprint(ss))
}
