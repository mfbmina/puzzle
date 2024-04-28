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
	p := core.NewPlay()

	assert.False(t, p.IsWin())
}

func Test_IsWin_TrueWhenPlayTableIsDefaultTable(t *testing.T) {
	p := core.NewPlay()
	p.Table = core.DEFAULT_TABLE

	assert.True(t, p.IsWin())
}
