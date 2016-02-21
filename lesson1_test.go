package lesson1

import (
	"sync"
	"testing"
)

type AssertWriter struct {
	str string
}

func (w *AssertWriter) Write(p []byte) (n int, err error) {
	w.str = string(p)
	return 0, nil
}

func TestSolve(t *testing.T) {
	m := new(sync.Mutex)
	m.Lock()
	defer m.Unlock()

	tests := map[string]string{
		"79538246":     "x won.",
		"35497162193":  "x won.",
		"61978543":     "x won.",
		"254961323121": "x won.",
		"6134278187":   "x won.",
		"4319581":      "Foul : x won.",
		"9625663381":   "Foul : x won.",
		"7975662":      "Foul : x won.",
		"2368799597":   "Foul : x won.",
		"18652368566":  "Foul : x won.",
		"965715":       "o won.",
		"38745796":     "o won.",
		"371929":       "o won.",
		"758698769":    "o won.",
		"42683953":     "o won.",
		"618843927":    "Foul : o won.",
		"36535224":     "Foul : o won.",
		"882973":       "Foul : o won.",
		"653675681":    "Foul : o won.",
		"9729934662":   "Foul : o won.",
		"972651483927": "Draw game.",
		"142583697":    "Draw game.",
		"5439126787":   "Draw game.",
		"42198637563":  "Draw game.",
		"657391482":    "Draw game.",
	}

	i := 0
	for k, v := range tests {
		i = i + 1
		var b *board = newBoard()
		var w *AssertWriter = &AssertWriter{}
		b.solve(read(k), w)

		if w.str != v {
			t.Errorf("%d: expected \"%s\", actual \"%s\"", i, v, w.str)
		}

	}
}
