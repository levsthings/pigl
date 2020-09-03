package pigl

import (
	"strings"
)

type Translate struct {
	buff []byte
}

func (t *Translate) Word(input string) string {
	t.pigify([]byte(input))
	return string(t.buff)
}

func (t *Translate) Sentence(input string) string {
	words := strings.Split(input, " ")

	var sentence []string
	for _, word := range words {
		t.pigify([]byte(word))
		sentence = append(sentence, string(t.buff))
		t.buff = nil
	}

	res := strings.Join(sentence, " ")
	return res
}

func (t *Translate) pigify(word []byte) {
	switch word[0] {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		t.buff = append(word, 'a', 'y')
	default:
		t.buff = append(word[1:], word[0])
		t.pigify(t.buff)
	}
}
