package stdout

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mfbmina/puzzle/core"
)

type Stdout struct {
	Play *core.Play
}

func NewStdout() *Stdout {
	return &Stdout{Play: core.NewPlay()}
}

func (s Stdout) Render() {
	k := ""
	w := false

	for !w && k != s.Play.Quit() {
		s.Play.PrintTable()

		k = getMove()
		err := s.Play.Move(k)
		if err != nil {
			fmt.Println(err)
		}

		w = s.Play.IsWin()
	}

	if w {
		fmt.Println("You win!")
	}
}

func getMove() string {
	reader := bufio.NewReader(os.Stdin)
	t, _ := reader.ReadString('\n')
	return strings.TrimSuffix(t, "\n")
}
