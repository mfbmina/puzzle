package ui

import (
	"embed"
	"fmt"
	"image"
	_ "image/png"
	"log"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/mfbmina/puzzle/core"
)

//go:embed assets/*
var assets embed.FS

// distance between points a and b.
func distance(xa, ya, xb, yb int) float64 {
	x := math.Abs(float64(xa - xb))
	y := math.Abs(float64(ya - yb))
	return math.Sqrt(x*x + y*y)
}

type touch struct {
	originX, originY int
	currX, currY     int
}

type pan struct {
	id ebiten.TouchID

	prevX, prevY     int
	originX, originY int
}

// UI implements ebiten.UI interface.
type UI struct {
	Play *core.Play

	touchIDs []ebiten.TouchID
	touches  map[ebiten.TouchID]*touch
	pan      *pan

	currX, currY int

	lastTouchTime time.Time
}

func NewUI() *UI {
	return &UI{Play: core.NewPlay(), touches: map[ebiten.TouchID]*touch{}, lastTouchTime: time.Time{}}
}

func (u *UI) resolveTouches() {
	released := false

	// What touches have just ended?
	for id := range u.touches {
		if inpututil.IsTouchJustReleased(id) {
			released = true

			defer func() {
				if u.pan != nil && id == u.pan.id {
					u.pan = nil
				}

				delete(u.touches, id)
			}()
		}
	}

	// What touches are new in this frame?
	u.touchIDs = inpututil.AppendJustPressedTouchIDs(u.touchIDs[:0])
	for _, id := range u.touchIDs {
		x, y := ebiten.TouchPosition(id)
		u.touches[id] = &touch{
			originX: x, originY: y,
			currX: x, currY: y,
		}
	}

	u.touchIDs = ebiten.AppendTouchIDs(u.touchIDs[:0])

	// Update the current position of any touches that have
	// neither begun nor ended in this frame.
	for _, id := range u.touchIDs {
		t := u.touches[id]
		t.currX, t.currY = ebiten.TouchPosition(id)
	}

	if len(u.touches) == 1 && len(u.touchIDs) == 1 {
		id := u.touchIDs[0]
		t := u.touches[id]
		if u.pan == nil {

			diff := math.Abs(distance(t.originX, t.originY, t.currX, t.currY))
			if diff > 1 {
				u.pan = &pan{
					id:      id,
					originX: t.originX,
					originY: t.originY,
					prevX:   t.originX,
					prevY:   t.originY,
				}
			}
		}
	}

	// Trigger game moves when any pan touch has ended.
	if u.pan != nil {
		if !released {
			u.currX, u.currY = ebiten.TouchPosition(u.pan.id)
		}

		deltaX, deltaY := u.currX-u.pan.prevX, u.currY-u.pan.prevY

		switch {
		case deltaX > 0 && abs(deltaX) >= abs(deltaY) && released:
			u.Play.Left()
		case deltaX < 0 && abs(deltaX) >= abs(deltaY) && released:
			u.Play.Right()
		case deltaY > 0 && abs(deltaY) >= abs(deltaX) && released:
			u.Play.Up()
		case deltaY < 0 && abs(deltaY) >= abs(deltaX) && released:
			u.Play.Down()
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Update proceeds the UI state.
// Update is called every tick (1/60 [s] by default).
func (u *UI) Update() error {
	// Write your UI's logical update.
	u.resolveTouches()

	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		return fmt.Errorf("Quit")
	}
	if u.Play.IsWin() {
		return nil
	}
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

	return nil
}

// Draw draws the UI screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (u *UI) Draw(screen *ebiten.Image) {
	instructions := loadImage("instructions")
	opInstructions := &ebiten.DrawImageOptions{}
	opInstructions.GeoM.Translate(1900, 600)
	opInstructions.GeoM.Scale(0.5, 0.5)

	screen.DrawImage(instructions, opInstructions)

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
		img := loadImage("win")
		screen.DrawImage(img, nil)
	}
}

// Layout takes the outside size (e.u., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (u *UI) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1300, 900
}

func (u *UI) Render() {
	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(1300, 900)
	ebiten.SetWindowTitle("Puzzle Game")
	// Call ebiten.RunUI to start your UI loop.
	if err := ebiten.RunGame(u); err != nil {
		log.Fatal(err)
	}
}

func loadImage(name string) *ebiten.Image {
	fName := fmt.Sprintf("assets/%s.png", name)
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
