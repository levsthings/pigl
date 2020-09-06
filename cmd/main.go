package main

import (
	"fmt"

	"gitlab.com/levsthings/pigl"
)

func main() {
	word := "scheme"
	sent := "scheme is really cool"

	translate := new(pigl.Translator)
	pw := translate.Word(word)
	ps := translate.Sentence(sent)
	fmt.Println(pw)
	fmt.Println(ps)
}
