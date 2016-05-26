package chicken

import (
	"errors"
	"github.com/thisisaaronland/go-chicken/strings"
)

func GetChickenForLanguageTag(tag string) (string, error) {

	c, ok := chicken.CHICKENS[tag]

	if !ok {
		return "", errors.New("unknown (or untranslated) chicken")
	}

	return c, nil
}
