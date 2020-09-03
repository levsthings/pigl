package pigl

import "strings"

var pig []byte

// TranslateWord translates an ASCII word into pig latin
func TranslateWord(word string) string {
	piglr([]byte(word))

	return string(pig)
}

// TranslateSentence translates an ASCII sentence into pig latin. It does not
// handle trailing whitespaces or punctuations.
func TranslateSentence(input string) string {
	words := strings.Split(input, " ")

	var sentence []string
	for _, word := range words {
		piglr([]byte(word))
		sentence = append(sentence, string(pig))
		pig = nil
	}

	res := strings.Join(sentence, " ")
	return res
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
