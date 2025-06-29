package board_test

import (
	"testing"

	"github.com/gaoliveira21/chess-engine/pkg/board"
	"github.com/stretchr/testify/assert"
)

func TestBoardFileString(t *testing.T) {
	assert.Equal(t, "A", board.FileA.String(), "Expected FileA to return 'A'")
	assert.Equal(t, "B", board.FileB.String(), "Expected FileB to return 'B'")
	assert.Equal(t, "C", board.FileC.String(), "Expected FileC to return 'C'")
	assert.Equal(t, "D", board.FileD.String(), "Expected FileD to return 'D'")
	assert.Equal(t, "E", board.FileE.String(), "Expected FileE to return 'E'")
	assert.Equal(t, "F", board.FileF.String(), "Expected FileF to return 'F'")
	assert.Equal(t, "G", board.FileG.String(), "Expected FileG to return 'G'")
	assert.Equal(t, "H", board.FileH.String(), "Expected FileH to return 'H'")
	assert.Equal(t, " ", board.RankNone.String(), "Expected RankNone to return ' '")
}
