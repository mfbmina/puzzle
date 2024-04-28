package core_test

import (
	"testing"

	"github.com/mfbmina/puzzle/core"
	"github.com/stretchr/testify/assert"
)

func Test_NewPlay_ReturnAValidPlay(t *testing.T) {
	p := core.NewPlay()

	assert.NotEqual(t, p.Table, core.DEFAULT_TABLE)
}

func Test_IsWin_FalseWhenPlayTableIsNotDefaultTable(t *testing.T) {
	p := core.Play{Table: [3][3]int{{3, 2, 1}, {4, 5, 6}, {7, 8, 0}}}

	assert.False(t, p.IsWin())
}

func Test_IsWin_TrueWhenPlayTableIsDefaultTable(t *testing.T) {
	p := core.Play{Table: core.DEFAULT_TABLE}

	assert.True(t, p.IsWin())
}

func Test_Up(t *testing.T) {
	p := core.Play{
		Table:    [3][3]int{{1, 2, 3}, {4, 0, 5}, {6, 7, 8}},
		EmptyRow: 1,
		EmptyCol: 1,
	}
	err := p.Up()
	assert.Nil(t, err)
	assert.Equal(t, [3][3]int{{1, 0, 3}, {4, 2, 5}, {6, 7, 8}}, p.Table)
	assert.Equal(t, 0, p.EmptyRow)
	assert.Equal(t, 1, p.EmptyCol)
}

func Test_Up_CantMoveZeroIfItIsAtRow0(t *testing.T) {
	p := core.Play{
		Table:    [3][3]int{{1, 0, 3}, {4, 2, 5}, {6, 7, 8}},
		EmptyRow: 0,
		EmptyCol: 1,
	}
	err := p.Up()
	assert.NotNil(t, err)
	assert.Equal(t, [3][3]int{{1, 0, 3}, {4, 2, 5}, {6, 7, 8}}, p.Table)
	assert.Equal(t, 0, p.EmptyRow)
	assert.Equal(t, 1, p.EmptyCol)
}

func Test_Down(t *testing.T) {
	p := core.Play{
		Table:    [3][3]int{{1, 2, 3}, {4, 0, 5}, {6, 7, 8}},
		EmptyRow: 1,
		EmptyCol: 1,
	}
	err := p.Down()
	assert.Nil(t, err)
	assert.Equal(t, [3][3]int{{1, 2, 3}, {4, 7, 5}, {6, 0, 8}}, p.Table)
	assert.Equal(t, 2, p.EmptyRow)
	assert.Equal(t, 1, p.EmptyCol)
}

func Test_Down_CantMoveZeroIfItIsAtRow2(t *testing.T) {
	p := core.Play{
		Table:    [3][3]int{{1, 2, 3}, {4, 7, 5}, {6, 0, 8}},
		EmptyRow: 2,
		EmptyCol: 1,
	}
	err := p.Down()
	assert.NotNil(t, err)
	assert.Equal(t, [3][3]int{{1, 2, 3}, {4, 7, 5}, {6, 0, 8}}, p.Table)
	assert.Equal(t, 2, p.EmptyRow)
	assert.Equal(t, 1, p.EmptyCol)
}

func Test_Left(t *testing.T) {
	p := core.Play{
		Table:    [3][3]int{{1, 2, 3}, {4, 0, 5}, {6, 7, 8}},
		EmptyRow: 1,
		EmptyCol: 1,
	}
	err := p.Left()
	assert.Nil(t, err)
	assert.Equal(t, [3][3]int{{1, 2, 3}, {0, 4, 5}, {6, 7, 8}}, p.Table)
	assert.Equal(t, 1, p.EmptyRow)
	assert.Equal(t, 0, p.EmptyCol)
}

func Test_Left_CantMoveZeroIfItIsAtCol0(t *testing.T) {
	p := core.Play{
		Table:    [3][3]int{{1, 2, 3}, {0, 4, 5}, {6, 7, 8}},
		EmptyRow: 1,
		EmptyCol: 0,
	}
	err := p.Left()
	assert.NotNil(t, err)
	assert.Equal(t, [3][3]int{{1, 2, 3}, {0, 4, 5}, {6, 7, 8}}, p.Table)
	assert.Equal(t, 1, p.EmptyRow)
	assert.Equal(t, 0, p.EmptyCol)
}

func Test_Right(t *testing.T) {
	p := core.Play{
		Table:    [3][3]int{{1, 2, 3}, {4, 0, 5}, {6, 7, 8}},
		EmptyRow: 1,
		EmptyCol: 1,
	}
	err := p.Right()
	assert.Nil(t, err)
	assert.Equal(t, [3][3]int{{1, 2, 3}, {4, 5, 0}, {6, 7, 8}}, p.Table)
	assert.Equal(t, 1, p.EmptyRow)
	assert.Equal(t, 2, p.EmptyCol)
}

func Test_Right_CantMoveZeroIfItIsAtCol2(t *testing.T) {
	p := core.Play{
		Table:    [3][3]int{{1, 2, 3}, {4, 5, 0}, {6, 7, 8}},
		EmptyRow: 1,
		EmptyCol: 2,
	}
	err := p.Right()
	assert.NotNil(t, err)
	assert.Equal(t, [3][3]int{{1, 2, 3}, {4, 5, 0}, {6, 7, 8}}, p.Table)
	assert.Equal(t, 1, p.EmptyRow)
	assert.Equal(t, 2, p.EmptyCol)
}
