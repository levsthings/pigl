package main

import (
	"fmt"

	"gitlab.com/levsthings/pigl"
)

func main() {
	word := "I"

	ret := pigl.Pigx([]byte(word))
	fmt.Println(string(ret))
}
