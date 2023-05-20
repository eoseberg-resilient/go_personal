package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

var MissingFileNameError = errors.New("FileName could not be found")
var FileReadingError = errors.New("Error while reading the file")
var CriticalError = errors.New("Fuck if I know dude")


type VerbLoader interface {
	Load(string) (Lookup, error)
}

type VerbLoaderImpl struct {}

func NewVerbLoader() *VerbLoaderImpl {
	p := new(VerbLoaderImpl)

	fmt.Println("Created loader!")

	return p
}

func (loader *VerbLoaderImpl) Load(verbFile string, phraseFile string) (LookupImpl, error) {
	vf, verr := os.Open(verbFile)
	pf, perr := os.Open(phraseFile)

	if verr != nil || perr != nil {
		return LookupImpl{}, MissingFileNameError
	}
	defer func() {
		vclose := vf.Close()
		pclose := pf.Close()
		if vclose != nil {
			panic(vclose)
		}
		if pclose != nil {
			panic(pclose)
		}
	}()

	reader_verbs := bufio.NewReader(vf)
	reader_phrases := bufio.NewReader(pf)

	look := LookupImpl{
		words: NewMap[string, []string](),
	}

	for {
		vbytes, _, err := reader_verbs.ReadLine()
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return LookupImpl{}, FileReadingError
		}
		if len(vbytes) == 0 {
			break
		}

		word := string(vbytes)
		if word[0] != 45 {
			indx := 0
			phrases_for_word := make(map[int]string)
			for {
				pbytes, _, err := reader_phrases.ReadLine()
				if err != nil && err != io.EOF {
					fmt.Println(err)
					return LookupImpl{}, FileReadingError
				}
				if len(pbytes) == 0 {
					break
				}
				if word != string(pbytes) {
					phrases_for_word[indx] = string(pbytes)
				}
				indx += 1
			}

			look.words.Push(word, *ToArray(phrases_for_word))
		}
	}
	return look, nil
}


var WordNotFoundError = errors.New("No such word loaded")
var NoInputError = errors.New("No input")

type Lookup interface {
	GetPhrasesFromWord(string) (LookupResult, error)
	TryGetWordFromString(string) ([]string, error)
}

type LookupImpl struct {
	words Map[string, []string]
}

type LookupResult struct {
	phrases []string
}

func (l LookupImpl) TryGetWordFromString(s string) ([]string, error) {
	if len(s) == 0 {
		return nil, NoInputError
	}

	trywords := make([]string, 0)
	for _, ss := range l.words.Keys() {
		if strings.HasPrefix(ss,s) {
			trywords = append(trywords, ss)
		}
	}

	return trywords, nil
}

func (l LookupImpl) GetPhrasesFromWord(wordInp string) (LookupResult, error) {
	res := LookupResult{}
	var err = WordNotFoundError
	if l.words.Has(wordInp) {
		phrases := l.words.Get(wordInp)
		res.phrases = phrases
		err = nil
	}
	return res, err
}
