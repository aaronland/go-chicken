package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/thisisaaronland/go-chicken"
	"os"
	"strings"
)

func main() {

	var lang = flag.String("language", "zxx", "A valid ISO-639-3 language code.")
	flag.Parse()

	ch, err := chicken.GetChickenForLanguageTag(*lang)

	if err != nil {
		panic(err)
	}

	for _, path := range flag.Args() {

		var buf *bufio.Scanner

		if path == "-" {
			buf = bufio.NewScanner(os.Stdin)
		} else {

			fh, err := os.Open(path)

			if err != nil {
				panic(err)
			}

			buf = bufio.NewScanner(fh)
		}

		for buf.Scan() {

			txt := buf.Text()

			chickens := make([]string, 0)

			/*

				things the following doesn't do well (or at all) yet
				- account for not-english languages
				- distinguish between "words" and punctuation, etc.
			*/

			matches := strings.Fields(txt)
			count := len(matches)

			for i := 0; i < count; i++ {
				chickens = append(chickens, ch)
			}

			fmt.Println(strings.Join(chickens, " "))
		}
	}
}
