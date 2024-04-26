package core

import (
	"fmt"
	"math/rand"
)

var DEFAULT_TABLE = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 0}}
var KEYS = map[string]string{
	"up":    "w",
	"left":  "a",
	"down":  "s",
	"right": "d",
	"quit":  "q",
}

type Play struct {
	Table    [3][3]int
	Keys     map[string]string
	EmptyRow int
	EmptyCol int
}

func NewPlay() *Play {
	t, x, y := generateRandomTable()
	return &Play{
		Table:    t,
		Keys:     KEYS,
		EmptyRow: x,
		EmptyCol: y,
	}
}

func generateRandomTable() ([3][3]int, int, int) {
	t := DEFAULT_TABLE
	s := 3
	xEmpty := 0
	yEmpty := 0

	for i, r := range t {
		for j := range r {
			x := rand.Intn(s)
			y := rand.Intn(s)
			t[i][j], t[x][y] = t[x][y], t[i][j]

			if t[i][j] == 0 {
				xEmpty = i
				yEmpty = j
			}

			if t[x][y] == 0 {
				xEmpty = x
				yEmpty = y
			}
		}
	}

	return t, xEmpty, yEmpty
}

func (p *Play) PrintTable() {
	for _, row := range p.Table {
		for _, col := range row {
			fmt.Printf("%d ", col)
		}
		fmt.Printf("\n")
	}
}

func (p *Play) IsWin() bool {
	return p.Table == DEFAULT_TABLE
}

func (p *Play) Quit() string {
	return p.Keys["quit"]
}

func (p *Play) Move(k string) error {
	switch k {
	case p.Keys["up"]:
		// return p.up()
	case p.Keys["left"]:
		// return p.left()
	case p.Keys["down"]:
		// return p.down()
	case p.Keys["right"]:
		// return p.right()
	case p.Keys["quit"]:
		return nil
	default:
		return fmt.Errorf("Invalid key. Play again.")
	}

	return nil
}

// func (p *Play) up() error {
// 	if p.EmptyRow == 0 {
// 		return t, fmt.Errorf("can't move up")
// 	}

// 	t[xEmpty][yEmpty], t[xEmpty-1][yEmpty] = t[xEmpty-1][yEmpty], t[xEmpty][yEmpty]
// 	xEmpty--
// 	return nil
// }
