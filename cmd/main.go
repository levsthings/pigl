package main

import (
	"fmt"

	"gitlab.com/levsthings/pigl"
)

func main() {
	word := "I"

	ret := pigl.Translate(word)
	fmt.Println(ret)
}
