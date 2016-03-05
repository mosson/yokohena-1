package ai

import (
	"fmt"
	"sort"
	"strconv"
	"yokohena-1/game"
)

const (
	nonce = 0
	plain = 2
	reach = 16
	block = 32
	bingo = 64
)

const player = 2

// AI is game engine
type AI struct {
	Board *game.Board
}

// New returns new AI game engine
func New() *AI {
	return &AI{Board: game.New()}
}

func (ai *AI) cloneMem() []uint8 {
	clone := make([]uint8, len(ai.Board.Mem))
	for i := 0; i < len(ai.Board.Mem); i++ {
		clone[i] = ai.Board.Mem[i]
	}
	return clone
}

func line(mem []uint8, a int, b int, c int) string {
	slice := []int{
		int(mem[a]),
		int(mem[b]),
		int(mem[c]),
	}
	sort.Ints(slice)
	var result string
	for _, n := range slice {
		result = result + strconv.Itoa(n)
	}
	return result
}

func (ai *AI) lines(index int) []string {
	mem := ai.cloneMem()
	mem[index] = player

	return []string{
		line(mem, 0, 1, 2),
		line(mem, 3, 4, 5),
		line(mem, 6, 7, 8),
		line(mem, 0, 3, 6),
		line(mem, 1, 4, 7),
		line(mem, 2, 5, 8),
		line(mem, 0, 4, 8),
		line(mem, 2, 4, 6),
	}
}

func (ai *AI) calc(index int) int {
	var score int

	lines := ai.lines(index)

	for _, str := range lines {
		if str == "002" {
			score = score + plain
		}
		if str == "012" {
			score = score + nonce
		}
		if str == "022" {
			score = score + reach
		}
		if str == "112" {
			score = score + block
		}
		if str == "122" {
			score = score + nonce
		}
		if str == "222" {
			score = score + bingo
		}
	}
	return score
}

func max(m map[int]int) int {
	var buf int
	var i int
	for k, v := range m {
		if buf < v {
			buf = v
			i = k
		}
	}
	return i
}

// Think is thinking better place and returns it
func (ai *AI) Think() int {
	scores := make(map[int]int)
	for i, n := range ai.Board.Mem {
		if n == 0 {
			scores[i] = ai.calc(i)
		}
	}

	return max(scores)
}

// Tick is player and ai play
func (ai *AI) Tick(index int) (bool, string) {
	isContinue, message := ai.Board.Tick(index - 1)
	if !isContinue {
		return isContinue, message
	}

	return ai.Board.Tick(ai.Think())
}

// BoardDescription displays board memory
func (ai *AI) BoardDescription() {
	fmt.Print("\033[H\033[2J")

	for i, n := range ai.Board.Mem {
		if n == 0 {
			fmt.Print(i + 1)
		} else if n == 1 {
			fmt.Print("o")
		} else {
			fmt.Print("x")
		}
		if i%3 == 2 {
			fmt.Print("\n")
		} else {
			fmt.Print("\t")
		}
	}
}

// Wait shows user interaction
func (ai *AI) Wait() {
	ai.BoardDescription()

	fmt.Print("\n")

	fmt.Print("Your Turn.\nPlease Enter Index > ")
}
