package pigl

import (
	"strings"
)

// Translator translates English words to Pig Latin.
// It only works with ASCII characters, and does not
// support punctuation characters.
type Translator struct {
	buff []byte
}

// Word method is used for translating a single word.
func (t *Translator) Word(input string) string {
	t.pigify([]byte(input))
	return string(t.buff)
}

// Sentence method is used for translating sentences.
func (t *Translator) Sentence(input string) string {
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

// conjure pigify for pigification.
func (t *Translator) pigify(word []byte) {
	switch word[0] {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		t.buff = append(word, 'a', 'y')
	default:
		t.buff = append(word[1:], word[0])
		t.pigify(t.buff)
	}
}
