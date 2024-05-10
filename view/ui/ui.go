package ui

import (
	"embed"
	"fmt"
	"image"
	_ "image/jpeg"

	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/mfbmina/puzzle/core"
)

//go:embed assets/*
var assets embed.FS

// UI implements ebiten.UI interface.
type UI struct {
	Play *core.Play
}

func NewUI() *UI {
	return &UI{Play: core.NewPlay()}
}

// Update proceeds the UI state.
// Update is called every tick (1/60 [s] by default).
func (u *UI) Update() error {
	// Write your UI's logical update.
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		u.Play.Up()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		u.Play.Down()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		u.Play.Right()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		u.Play.Left()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		return fmt.Errorf("Quit")
	}

	return nil
}

// Draw draws the UI screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (u *UI) Draw(screen *ebiten.Image) {
	for x, row := range u.Play.Table {
		for y, value := range row {
			if value == 0 {
				continue
			}

			img := loadImage(fmt.Sprint(value))
			op := &ebiten.DrawImageOptions{}
			fX := float64(x)
			fY := float64(y)
			op.GeoM.Translate(300*fY, 300*fX)

			screen.DrawImage(img, op)
		}
	}

	if u.Play.IsWin() {
		screen.Clear()
		img := loadImage("congratz")
		screen.DrawImage(img, nil)
	}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (u *UI) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 900, 900
}

func (u *UI) Render() {
	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(900, 900)
	ebiten.SetWindowTitle("Puzzle Game")
	// Call ebiten.RunUI to start your UI loop.
	if err := ebiten.RunGame(u); err != nil {
		log.Fatal(err)
	}
}

func loadImage(name string) *ebiten.Image {
	fName := fmt.Sprintf("assets/%s.jpg", name)
	// Write your UI's rendering.
	f, err := assets.Open(fName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}
