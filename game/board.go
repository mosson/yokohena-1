package game

import (
	"errors"
	"fmt"
)

// Board ゲーム・フィールド
type Board struct {
	Mem    []uint8
	Player uint8
}

// New returns new board
func New() *Board {
	return &Board{Mem: make([]uint8, 9), Player: 2}
}

// Check 勝利者を返す
func (b *Board) Check() uint8 {
	checks := []uint8{
		(b.Mem[0] & b.Mem[1] & b.Mem[2]),
		(b.Mem[3] & b.Mem[4] & b.Mem[5]),
		(b.Mem[6] & b.Mem[7] & b.Mem[8]),
		(b.Mem[0] & b.Mem[3] & b.Mem[6]),
		(b.Mem[1] & b.Mem[4] & b.Mem[7]),
		(b.Mem[2] & b.Mem[5] & b.Mem[8]),
		(b.Mem[0] & b.Mem[4] & b.Mem[8]),
		(b.Mem[2] & b.Mem[4] & b.Mem[6]),
	}

	for _, n := range checks {
		if n > 0 {
			return n
		}
	}

	return 0
}

func (b *Board) turn(index int, player uint8) (uint8, error) {
	if b.Mem[index] > 0 {
		return player, errors.New("fault")
	}

	b.Mem[index] = player
	return player, nil
}

// IsAvailable まだ打てるかどうかを返す
func (b *Board) IsAvailable() bool {
	for _, i := range b.Mem {
		if i < 1 {
			return true
		}
	}
	return false
}

// Tick 次の一手を指定する
func (b *Board) Tick(index int) (bool, string) {
	if index > 8 || index < 0 {
		return true, "please type 1-9"
	}

	if b.Player == 2 {
		b.Player = 1
	} else {
		b.Player = 2
	}

	player, err := b.turn(index, b.Player)
	if err != nil {
		return false, fmt.Sprintf("player %d fault.", player)
	}

	winner := b.Check()

	if winner > 0 {
		return false, fmt.Sprintf("player %d won.", winner)
	}

	if !b.IsAvailable() {
		return false, "Draw Game."
	} else {
		return true, ""
	}
}

// Description 現在の状態を可視化する
func (b *Board) Description() {
	fmt.Printf("\n%d\t%d\t%d\n%d\t%d\t%d\n%d\t%d\t%d\n", b.Mem[0], b.Mem[1], b.Mem[2], b.Mem[3], b.Mem[4], b.Mem[5], b.Mem[6], b.Mem[7], b.Mem[8])
}
