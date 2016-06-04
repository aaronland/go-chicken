package chicken

import (
	"errors"
	"fmt"
	"github.com/cooperhewitt/go-ucd"
	emoji "github.com/thisisaaronland/go-chicken/emoji"
	"github.com/thisisaaronland/go-chicken/strings"
	"math/rand"
	"regexp"
	"strconv"
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

func (ch *Chicken) ExpandAlphaCodes(txt string) string {

	cb := func(in []byte) []byte {

		str_in := string(in)

		code, ok := emoji.EmojiAlphaCodes[str_in]

		if !ok {
			return in
		}

		num, _ := strconv.ParseInt(code, 0, 32)
		rune := fmt.Sprintf("%c", num)

		return []byte(rune)
	}

	re := regexp.MustCompile(`:(\w+):`)

	bytes_in := []byte(txt)
	bytes_out := re.ReplaceAllFunc(bytes_in, cb)

	return string(bytes_out)
}

func (ch *Chicken) TextToChicken(txt string) string {

	re := regexp.MustCompile(`:(\w+):`)

	if re.MatchString(txt) {
		// fmt.Printf("BEFORE %s\n", txt)
		txt = ch.ExpandAlphaCodes(txt)
		// fmt.Printf("AFTER %s\n", txt)
	}

	buf := make([]string, 0)
	word := false

	for _, char := range txt {

		// https://golang.org/pkg/unicode/#IsLetter
		// fmt.Println(idx, char, unicode.IsLetter(char), unicode.IsSpace(char))

		if unicode.IsSpace(char) {

			buf = append(buf, string(char))
			word = false

		} else if unicode.IsSymbol(char) {

			// fmt.Printf("BEFORE %s\n", char)
			n := ucd.Name(string(char))
			// fmt.Printf("AFTER %s\n", n.Name)
			b := ch.TextToChicken(n.Name)
			buf = append(buf, b)

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
