package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mfbmina/puzzle/core"
)

func main() {
	p := core.NewPlay()
	k := ""
	w := false

	for !w && k != p.Quit() {
		p.PrintTable()

		k = readPlay()
		err := p.Move(k)
		if err != nil {
			fmt.Println(err)
		}

		w = p.IsWin()
	}

	if w {
		fmt.Println("You win!")
	}
}

func readPlay() string {
	reader := bufio.NewReader(os.Stdin)
	t, _ := reader.ReadString('\n')
	return strings.TrimSuffix(t, "\n")
}
