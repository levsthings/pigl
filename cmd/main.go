package main

import (
	"fmt"

	"gitlab.com/levsthings/pigl"
)

func main() {
	word := "scheme"
	sent := "scheme is really cool"

	pw := pigl.TranslateWord(word)
	ps := pigl.TranslateSentence(sent)
	fmt.Println(pw)
	fmt.Println(ps)

	translate := new(pigl.Translator)
	pw2 := translate.Word("scheme")
	ps2 := translate.Sentence("scheme is really cool")
	fmt.Println(pw2)
	fmt.Println(ps2)
}
