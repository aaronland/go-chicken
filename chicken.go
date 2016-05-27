package chicken

import (
	"errors"
	"github.com/thisisaaronland/go-chicken/strings"
	"math/rand"
	"strings"
	"unicode"
)

func GetChickenForLanguageTag(tag string, clucking bool) (*Chicken, error) {

	// see also: https://github.com/thisisaaronland/go-chicken/issues/1

	c, ok := chicken.CHICKENS[tag]

	if !ok {
		return nil, errors.New("unknown (or untranslated) chicken")
	}

	ch := Chicken{tag, c, clucking}
	return &ch, nil
}

type Chicken struct {
	tag      string // sudo make me a https://godoc.org/golang.org/x/text/language thingy
	chicken  string
	clucking bool
}

func (ch *Chicken) String() string {

	if ch.clucking {
		return ch.Cluck()
	}

	return ch.Chicken()
}

func (ch *Chicken) Chicken() string {

	return ch.chicken
}

func (ch *Chicken) Cluck() string {

	tag := ch.tag
	clucks, ok := chicken.CLUCKING[tag]

	if !ok {
		return ch.Chicken()
	}

	count := len(clucks)
	idx := rand.Intn(count)

	return clucks[idx]
}

func (ch *Chicken) TextToChicken(txt string) string {

	// https://golang.org/pkg/unicode/#IsLetter

	buf := make([]string, 0)
	word := false

	for _, char := range txt {

		// fmt.Println(idx, char, unicode.IsLetter(char), unicode.IsSpace(char))

		if unicode.IsSpace(char) {

			buf = append(buf, string(char))
			word = false

		} else if unicode.IsLetter(char) {

			if !word {
				buf = append(buf, ch.String())
				word = true
			}

		} else {
			buf = append(buf, string(char))
			word = false
		}
	}

	return strings.Join(buf, "")
}
