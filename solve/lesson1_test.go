package solve

import (
	"os"
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

	for k, v := range tests {
		b := newBoard(os.Stdout)
		w := &AssertWriter{}
		b.solve(read(k), w)

		if w.str != v {
			t.Errorf("%s: expected \"%s\", actual \"%s\"", k, v, w.str)
		}

	}
}

func BenchmarkAllocate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = newBoard(nil)
	}
}

func BenchmarkRead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = read("79538246")
	}
}

func BenchmarkLines(b *testing.B) {
	board := newBoard(nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		board.lines()
	}
}

func BenchmarkDrawCheck(b *testing.B) {
	board := newBoard(nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		board.drawCheck()
	}
}

func BenchmarkCheck(b *testing.B) {
	board := newBoard(nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		board.check()
	}
}

func BenchmarkSolveXWon(b *testing.B) {
	w := &AssertWriter{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b := newBoard(nil)

		b.solve(read("79538246"), w)
	}
}

func BenchmarkSolveFoulXWon(b *testing.B) {
	w := &AssertWriter{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b := newBoard(nil)
		b.out = nil

		b.solve(read("4319581"), w)
	}
}

func BenchmarkSolveOWon(b *testing.B) {
	w := &AssertWriter{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b := newBoard(nil)

		b.solve(read("965715"), w)
	}
}

func BenchmarkSolveDraw(b *testing.B) {
	w := &AssertWriter{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b := newBoard(nil)

		b.solve(read("972651483927"), w)
	}
}

func BenchmarkSolveFoulOWon(b *testing.B) {
	w := &AssertWriter{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b := newBoard(nil)

		b.solve(read("618843927"), w)
	}
}
