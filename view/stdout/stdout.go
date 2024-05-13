package stdout

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mfbmina/puzzle/core"
)

var KEYS = map[string]string{
	"up":    "w",
	"left":  "a",
	"down":  "s",
	"right": "d",
	"quit":  "q",
}

type Stdout struct {
	Play *core.Play
}

func NewStdout() *Stdout {
	return &Stdout{Play: core.NewPlay()}
}

func (s *Stdout) Render() {
	k := ""
	w := false

	for !w && !isQuit(k) {
		s.printTable()

		k = getMove()
		err := s.move(k)
		if err != nil {
			fmt.Println(err)
		}

		w = s.Play.IsWin()
	}

	if w {
		fmt.Println("You win!")
	}
}

func (s *Stdout) move(k string) error {
	switch k {
	case KEYS["up"]:
		return s.Play.Up()
	case KEYS["left"]:
		return s.Play.Left()
	case KEYS["down"]:
		return s.Play.Down()
	case KEYS["right"]:
		return s.Play.Right()
	case KEYS["quit"]:
		return nil
	default:
		return fmt.Errorf("Invalid key. Play again.")
	}
}

func (s *Stdout) printTable() {
	for _, row := range s.Play.Table {
		for _, col := range row {
			fmt.Printf("%d ", col)
		}
		fmt.Printf("\n")
	}
}

func isQuit(k string) bool {
	return KEYS["quit"] == k
}

func getMove() string {
	reader := bufio.NewReader(os.Stdin)
	t, _ := reader.ReadString('\n')
	return strings.TrimSuffix(t, "\n")
}
