package pigl

import (
	"testing"
)

func Test_TranslateWord(t *testing.T) {
	translate := new(Translator)
	tests := []struct {
		word string
		want string
	}{
		{"smile", "ilesmay"},
		{"string", "ingstray"},
		{"stupid", "upidstay"},
		{"glove", "oveglay"},
		{"trash", "ashtray"},
		{"floor", "oorflay"},
		{"store", "orestay"},
		{"pig", "igpay"},
		{"latin", "atinlay"},
		{"banana", "ananabay"},
		{"will", "illway"},
		{"butler", "utlerbay"},
		{"happy", "appyhay"},
		{"duck", "uckday"},
		{"me", "emay"},
		{"eat", "eatay"},
		{"omelet", "omeletay"},
		{"are", "areay"},
		{"egg", "eggay"},
		{"explain", "explainay"},
		{"always", "alwaysay"},
		{"ends", "endsay"},
		{"honest", "onesthay"},
		{"I", "Iay"},
		{"scheme", "emeschay"},
		{"is", "isay"},
	}
	for _, tt := range tests {
		t.Run(tt.word, func(t *testing.T) {
			got := translate.Word(tt.word)

			if got != tt.want {
				t.Errorf("TranslateWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_TranslateSentence(t *testing.T) {
	translate := new(Translator)
	tests := []struct {
		sentence string
		want     string
	}{
		{"smile string stupid glove trash floor store pig latin banana will butler happy duck me eat omelet are egg explain always ends honest I scheme is",
			"ilesmay ingstray upidstay oveglay ashtray oorflay orestay igpay atinlay ananabay illway utlerbay appyhay uckday emay eatay omeletay areay eggay explainay alwaysay endsay onesthay Iay emeschay isay"},
	}
	for _, tt := range tests {
		t.Run(tt.sentence, func(t *testing.T) {
			got := translate.Sentence(tt.sentence)
			if got != tt.want {
				t.Errorf("TranslateSentence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_TranslateWord(b *testing.B) {
	translate := new(Translator)
	for i := 0; i < b.N; i++ {
		translate.Word("scheme")
	}
}

func Benchmark_TranslateSentence(b *testing.B) {
	translate := new(Translator)

	for i := 0; i < b.N; i++ {
		translate.Sentence("smile string stupid glove trash floor store pig latin banana will butler happy duck me eat omelet are egg explain always ends honest I scheme")
	}
}
