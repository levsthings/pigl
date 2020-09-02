package pigl

var pig []byte

func Pigx(word []byte) []byte {
	switch word[0] {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		pig = append(word, 'a', 'y')
	default:
		pig = append(word[1:], word[0])
		Pigx(pig)
	}
	return pig
}
