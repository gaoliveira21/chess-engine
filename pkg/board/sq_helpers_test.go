package board_test

import (
	"testing"

	"github.com/gaoliveira21/chess-engine/pkg/board"
	"github.com/stretchr/testify/assert"
)

func TestSq64ToRankAndFile(t *testing.T) {
	f, r := board.Sq64ToRankAndFile(0)
	assert.Equal(t, board.Rank1, r, "Expected Rank1 for square 0")
	assert.Equal(t, board.FileA, f, "Expected FileA for square 0")

	f, r = board.Sq64ToRankAndFile(7)
	assert.Equal(t, board.Rank1, r, "Expected Rank1 for square 7")
	assert.Equal(t, board.FileH, f, "Expected FileH for square 7")

	f, r = board.Sq64ToRankAndFile(56)
	assert.Equal(t, board.Rank8, r, "Expected Rank8 for square 56")
	assert.Equal(t, board.FileA, f, "Expected FileA for square 56")
}

func TestSq120ToRankAndFile(t *testing.T) {
	f, r := board.Sq120ToRankAndFile(11)
	assert.Equal(t, board.RankNone, r, "Expected RankNone for square 11")
	assert.Equal(t, board.FileNone, f, "Expected FileNone for square 11")

	f, r = board.Sq120ToRankAndFile(20)
	assert.Equal(t, board.RankNone, r, "Expected RankNone for square 20")
	assert.Equal(t, board.FileNone, f, "Expected FileNone for square 20")

	f, r = board.Sq120ToRankAndFile(21)
	assert.Equal(t, board.Rank1, r, "Expected Rank1 for square 21")
	assert.Equal(t, board.FileA, f, "Expected FileA for square 21")

	f, r = board.Sq120ToRankAndFile(40)
	assert.Equal(t, board.RankNone, r, "Expected RankNone for square 40")
	assert.Equal(t, board.FileNone, f, "Expected FileNone for square 40")

	f, r = board.Sq120ToRankAndFile(42)
	assert.Equal(t, board.Rank3, r, "Expected Rank3 for square 42")
	assert.Equal(t, board.FileB, f, "Expected FileB for square 42")

	f, r = board.Sq120ToRankAndFile(49)
	assert.Equal(t, board.RankNone, r, "Expected RankNone for square 49")
	assert.Equal(t, board.FileNone, f, "Expected FileNone for square 49")

	f, r = board.Sq120ToRankAndFile(98)
	assert.Equal(t, board.Rank8, r, "Expected Rank8 for square 98")
	assert.Equal(t, board.FileH, f, "Expected FileH for square 98")

	f, r = board.Sq120ToRankAndFile(99)
	assert.Equal(t, board.RankNone, r, "Expected RankNone for square 99")
	assert.Equal(t, board.FileNone, f, "Expected FileNone for square 99")
}

func TestGetSq64Index(t *testing.T) {
	assert.Equal(t, 0, board.GetSq64Index(board.FileA, board.Rank1), "Expected index for FileA, Rank1 to be 0")
	assert.Equal(t, 7, board.GetSq64Index(board.FileH, board.Rank1), "Expected index for FileH, Rank1 to be 7")
	assert.Equal(t, 56, board.GetSq64Index(board.FileA, board.Rank8), "Expected index for FileA, Rank8 to be 56")
	assert.Equal(t, 63, board.GetSq64Index(board.FileH, board.Rank8), "Expected index for FileH, Rank8 to be 63")
}

func TestGetSq120Index(t *testing.T) {
	assert.Equal(t, 21, board.GetSq120Index(board.FileA, board.Rank1), "Expected index for FileA, Rank1 to be 21")
	assert.Equal(t, 42, board.GetSq120Index(board.FileB, board.Rank3), "Expected index for FileA, Rank1 to be 21")
	assert.Equal(t, 55, board.GetSq120Index(board.FileE, board.Rank4), "Expected index for FileA, Rank1 to be 21")
	assert.Equal(t, 98, board.GetSq120Index(board.FileH, board.Rank8), "Expected index for FileA, Rank1 to be 21")
}

func TestSq64ToSq120(t *testing.T) {
	assert.Equal(t, 21, board.Sq64ToSq120(0), "Expected square 0 to map to 21 in 120 board representation")
	assert.Equal(t, 28, board.Sq64ToSq120(7), "Expected square 7 to map to 28 in 120 board representation")
	assert.Equal(t, 91, board.Sq64ToSq120(56), "Expected square 56 to map to 91 in 120 board representation")
}

func TestSq120ToSq64(t *testing.T) {
	sq64, err := board.Sq120ToSq64(21)
	assert.NoError(t, err, "Expected no error for valid square 21")
	assert.Equal(t, 0, sq64, "Expected square 21 to map to 0 in 64 board representation")

	sq64, err = board.Sq120ToSq64(28)
	assert.NoError(t, err, "Expected no error for valid square 28")
	assert.Equal(t, 7, sq64, "Expected square 28 to map to 7 in 64 board representation")

	sq64, err = board.Sq120ToSq64(91)
	assert.NoError(t, err, "Expected no error for valid square 91")
	assert.Equal(t, 56, sq64, "Expected square 91 to map to 56 in 64 board representation")

	_, err = board.Sq120ToSq64(11)
	assert.Error(t, err, "Expected error for invalid square 11")
}
