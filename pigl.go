package pigl

var pig []byte

func Translate(word string) string {
	piglr([]byte(word))

	return string(pig)
}

func piglr(word []byte) {
	switch word[0] {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		pig = append(word, 'a', 'y')
	default:
		pig = append(word[1:], word[0])
		piglr(pig)
	}
}
