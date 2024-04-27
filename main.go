package main

import "github.com/mfbmina/puzzle/ui/stdout"

func main() {
	u := stdout.NewStdout()
	u.Render()
}
