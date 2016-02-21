package lesson1

import (
	"fmt"
	"io"
	"log"
	"strconv"
)

// 三目並べ( tick-tack-toe )の手を入力とし、勝敗を出力する。
// 先攻がo、後攻がx
// すでに打ってある場所に打った場合、反則負け
// x が反則をした場合、「Foul : o won.」と出力
// 縦横斜めのいずれかで一列揃ったら、揃えた方の勝ち
// x が揃えた場合、「x won.」と出力
// 9マス埋まっても揃わなかったら引き分け
// 「Draw game.」と出力
// 勝敗が決した後の手は無視する
// 入力文字列は、先攻から順に打った位置を示す。盤上の位置と数の対応は下表を参照。
// 入力文字列が「91593」の場合、「oが9の位置、xが1の位置、oが5の位置、xが9の位置→xの反則負け(最後の3は無視)」となる。
// 以下の様なケースは考慮しなくてよい
// 入力が 1〜9 以外の文字を含んでいる
// 入力が不足していて、ゲームの勝敗が決まらない

//  1 2 3
//  4 5 6
//  7 8 9

const (
	blank = 0x0
	white = 0x1
	black = 0x10
)

const (
	foul = "Foul : %s won."
	won  = "x won."
	lose = "o won."
	draw = "Draw game."
)

var mark map[int]string = map[int]string{
	blank: "-",
	white: "o",
	black: "x",
}

func read(str string) []int {
	nums := make([]int, len(str))
	for i, rune := range str {
		buf, err := strconv.Atoi(string(rune))
		if err != nil {
			log.Fatal("read: ", err)
		}
		nums[i] = buf
	}
	return nums
}

type turn struct {
	player  int
	address int
}

type board struct {
	mem    map[int]int
	result string
}

func newBoard() *board {
	return &board{
		mem: map[int]int{
			1: blank,
			2: blank,
			3: blank,
			4: blank,
			5: blank,
			6: blank,
			7: blank,
			8: blank,
			9: blank,
		},
	}
}

func (b *board) lines() []int {
	return []int{
		b.mem[1] & b.mem[4] & b.mem[7],
		b.mem[2] & b.mem[5] & b.mem[8],
		b.mem[3] & b.mem[6] & b.mem[9],
		b.mem[1] & b.mem[2] & b.mem[3],
		b.mem[4] & b.mem[5] & b.mem[6],
		b.mem[7] & b.mem[8] & b.mem[9],
		b.mem[1] & b.mem[5] & b.mem[9],
		b.mem[3] & b.mem[5] & b.mem[7],
	}
}

func (b *board) drawCheck() {
	isDraw := true
	for _, v := range b.mem {
		if v == 0 {
			isDraw = false
			break
		}
	}

	if isDraw {
		b.result = draw
	}
}

func (b *board) check() {
	lines := b.lines()

	for _, line := range lines {
		if line > 0 {
			if line == white {
				b.result = lose
				return
			}
			if line == black {
				b.result = won
				return
			}
		}
	}

	b.drawCheck()
}

func (b *board) fault(player int) {
	var p int
	if player == 0 {
		p = black
	} else {
		p = white
	}
	b.result = fmt.Sprintf(foul, mark[p])
}

func (b *board) step(t *turn) bool {
	if b.mem[t.address] != blank {
		b.fault(t.player)
		return false
	}

	if t.player == 0 {
		b.mem[t.address] = white
	} else {
		b.mem[t.address] = black
	}

	return true
}

func (b *board) description() {
	fmt.Printf(`

    %v  %v  %v
    %v  %v  %v
    %v  %v  %v

    `,
		mark[b.mem[1]],
		mark[b.mem[2]],
		mark[b.mem[3]],
		mark[b.mem[4]],
		mark[b.mem[5]],
		mark[b.mem[6]],
		mark[b.mem[7]],
		mark[b.mem[8]],
		mark[b.mem[9]],
	)
}

func (b *board) solve(input_c []int, w io.Writer) {
	for i, n := range input_c {
		if ok := b.step(&turn{player: i % 2, address: n}); !ok {
			break
		}
	}

	b.check()

	b.description()

	w.Write([]byte(b.result))
}
