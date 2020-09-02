package pigl

import "testing"

func Test_pigx(t *testing.T) {

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
	}
	for _, tt := range tests {
		t.Run(tt.word, func(t *testing.T) {
			got := string(Pigx([]byte(tt.word)))

			if got != tt.want {
				t.Errorf("Pigl() = %v, want %v", got, tt.want)
			}
		})
	}
}
