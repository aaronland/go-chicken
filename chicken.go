package chicken

import (
	"errors"
	"github.com/thisisaaronland/go-chicken/strings"
)

func GetChickenForLanguageTag(tag string) (string, error) {

	// see also: https://github.com/thisisaaronland/go-chicken/issues/1
	
	c, ok := chicken.CHICKENS[tag]

	if !ok {
		return "", errors.New("unknown (or untranslated) chicken")
	}

	return c, nil
}
