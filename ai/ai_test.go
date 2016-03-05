package ai

import "testing"

func TestNew(t *testing.T) {
	ai := New()

	if ai.Board == nil {
		t.Errorf("expected exists")
	}
}

func TestCloneMem(t *testing.T) {
	ai := New()
	ai.Board.Mem = []uint8{1, 1, 1, 1, 1, 1, 1, 1, 1}
	mem := ai.cloneMem()
	mem[0] = 2

	if ai.Board.Mem[0] != 1 {
		t.Errorf("expected 1, actual %v", ai.Board.Mem[0])
	}

	if mem[0] != 2 {
		t.Errorf("expected 2, actual %v", mem[0])
	}
}

func TestLine(t *testing.T) {
	mem := []uint8{1, 2, 0, 1, 2, 0, 1, 2, 0}

	l := line(mem, 0, 1, 2)
	if l != "012" {
		t.Errorf("expected 012, actual %v", l)
	}

	l = line(mem, 1, 2, 0)
	if l != "012" {
		t.Errorf("expected 012, actual %v", l)
	}
}

func TestLines(t *testing.T) {
	ai := New()
	ai.Board.Mem = []uint8{1, 2, 0, 1, 2, 0, 1, 2, 0}
	lines := ai.lines(2)

	if lines[0] != "122" {
		t.Errorf("expected 122, actual %v", lines[0])
	}
}

func TestCalc(t *testing.T) {
	ai := New()
	ai.Board.Mem = []uint8{1, 0, 1, 0, 2, 0, 0, 2, 0}

	score := ai.calc(1)

	if score != bingo+plain+plain+block {
		t.Errorf("expected %v, actual %v", bingo+plain+plain+block, score)
	}
}

func TestMax(t *testing.T) {
	m := map[int]int{
		0: 10,
		1: 20,
		2: 5,
		3: 12,
	}
	r := max(m)

	if r != 1 {
		t.Errorf("expected 1, actual %v", r)
	}
}

func TestThink(t *testing.T) {
	ai := New()
	ai.Board.Mem = []uint8{1, 0, 1, 0, 2, 0, 0, 2, 0}

	r := ai.Think()

	if r != 1 {
		t.Errorf("expected 1, actual %v", r)
	}
}

func TestTick(t *testing.T) {
	ai := New()
	isContinue, _ := ai.Tick(1)

	if !isContinue {
		t.Errorf("expected continue")
	}

	isContinue, _ = ai.Tick(1)
	if isContinue {
		t.Errorf("expected not continue")
	}

}

func TestWait(t *testing.T) {
	ai := New()

	ai.Wait()
}
