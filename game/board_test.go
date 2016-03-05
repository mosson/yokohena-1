package game

import "testing"

func TestNew(t *testing.T) {
	board := New()
	if board == nil {
		t.Errorf("expected not eq nil")
	}
}

func TestIsAvailable(t *testing.T) {
	board := New()
	board.Mem = []uint8{1, 1, 1, 1, 1, 1, 1, 1, 1}

	if board.IsAvailable() {
		t.Errorf("expected false")
	}

	board.Mem = []uint8{2, 2, 2, 2, 2, 2, 2, 2, 2}
	if board.IsAvailable() {
		t.Errorf("exptected false")
	}

	board.Mem = []uint8{1, 2, 1, 2, 1, 2, 1, 2, 1}
	if board.IsAvailable() {
		t.Errorf("exptected false")
	}

	board.Mem = []uint8{1, 2, 0, 1, 2, 0, 1, 2, 0}
	if !board.IsAvailable() {
		t.Errorf("expected true")
	}
}

func TestTurn(t *testing.T) {
	board := New()
	player, err := board.turn(0, 1)

	if player != 1 {
		t.Errorf("expected 1, actual %v", player)
	}

	if err != nil {
		t.Errorf("expected nil")
	}

	player, err = board.turn(0, 1)

	if player != 1 {
		t.Errorf("expected 1, actual %v", player)
	}

	if err == nil {
		t.Errorf("expected not eq nil")
	}
}

func TestCheck(t *testing.T) {
	board := New()
	player := board.Check()

	if player != 0 {
		t.Errorf("expected 0, actual %v", player)
	}

	board.Mem = []uint8{1, 1, 1, 0, 0, 0, 0, 0, 0}

	player = board.Check()

	if player != 1 {
		t.Errorf("expected 1, actual %v", player)
	}
}

func TestTick(t *testing.T) {
	board := New()
	b, _ := board.Tick(1)

	b, message := board.Tick(-12)

	if message == "" {
		t.Errorf("expected exists")
	}

	if !b {
		t.Errorf("expected true")
	}

	b, _ = board.Tick(2)

	if !b {
		t.Errorf("expected true")
	}

	b, _ = board.Tick(4)

	if !b {
		t.Errorf("expected true")
	}

	b, _ = board.Tick(5)

	if !b {
		t.Errorf("expected true")
	}

	b, _ = board.Tick(7)

	if b {
		t.Errorf("expected false")
	}

	b, message = board.Tick(7)

	if message == "" {
		t.Errorf("expected message exists")
	}

}

func TestDescription(t *testing.T) {
	b := New()
	b.Description()
}
