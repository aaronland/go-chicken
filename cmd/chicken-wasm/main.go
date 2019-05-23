package main

import (
	"github.com/aaronland/go-chicken"
	"log"
	"syscall/js"
)

var chicken_func js.Func

func main() {

	chicken_func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		if len(args) != 2 {
			log.Println("Invalid arguments")
			return nil
		}

		lang := args[0].String()
		text := args[1].String()

		ch, err := chicken.GetChickenForLanguageTag(lang, false)

		if err != nil {
			log.Println(err)
			return nil
		}

		return ch.TextToChicken(text)
	})

	defer chicken_func.Release()

	js.Global().Set("chicken", chicken_func)

	c := make(chan struct{}, 0)

	log.Println("WASM chicken initialized")
	<-c
}
