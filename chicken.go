package chicken

import (
	"errors"
	"github.com/thisisaaronland/go-chicken/strings"
	"strings"
)

func GetChickenForLanguageTag(tag string) (*Chicken, error) {

	// see also: https://github.com/thisisaaronland/go-chicken/issues/1

	c, ok := chicken.CHICKENS[tag]

	if !ok {
		return nil, errors.New("unknown (or untranslated) chicken")
	}

	ch := Chicken{tag, c}
	return &ch, nil
}

type Chicken struct {
	tag     string
	chicken string
}

func (ch *Chicken) String() string {
	return ch.chicken
}

func (ch *Chicken) TextToChicken(txt string) string {

	chickens := make([]string, 0)

	/*

		things the following doesn't do well (or at all) yet
		- account for not-english languages
		- distinguish between "words" and punctuation, etc.
	*/

	matches := strings.Fields(txt)
	count := len(matches)

	for i := 0; i < count; i++ {
		chickens = append(chickens, ch.String())
	}

	return strings.Join(chickens, " ")
}
