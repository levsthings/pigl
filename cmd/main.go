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
}
