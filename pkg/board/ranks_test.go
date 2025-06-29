package board_test

import (
	"testing"

	"github.com/gaoliveira21/chess-engine/pkg/board"
	"github.com/stretchr/testify/assert"
)

func TestBoardRankString(t *testing.T) {
	assert.Equal(t, "1", board.Rank1.String(), "Expected Rank1 to return '1'")
	assert.Equal(t, "2", board.Rank2.String(), "Expected Rank2 to return '2'")
	assert.Equal(t, "3", board.Rank3.String(), "Expected Rank3 to return '3'")
	assert.Equal(t, "4", board.Rank4.String(), "Expected Rank4 to return '4'")
	assert.Equal(t, "5", board.Rank5.String(), "Expected Rank5 to return '5'")
	assert.Equal(t, "6", board.Rank6.String(), "Expected Rank6 to return '6'")
	assert.Equal(t, "7", board.Rank7.String(), "Expected Rank7 to return '7'")
	assert.Equal(t, "8", board.Rank8.String(), "Expected Rank8 to return '8'")
	assert.Equal(t, " ", board.RankNone.String(), "Expected RankNone to return ' '")
}
