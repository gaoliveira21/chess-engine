/*
Board representation: 10x12
https://www.chessprogramming.org/10x12_Board

Rules:
  - Side to move: https://www.chessprogramming.org/Side_to_move
  - Fifty-move rule: https://www.chessprogramming.org/Fifty-move_Rule
  - Threefold repetition: https://www.chessprogramming.org/Repetitions
  - En passant: https://www.chessprogramming.org/En_passant
  - Ply: https://www.chessprogramming.org/Ply
  - Castling: https://www.chessprogramming.org/Castling
*/
package board

const (
	WKCA = 1
	WQCA = 2
	BKCA = 4
	BQCA = 8
)

type Board struct {
	squares [120]uint8

	sideToMove uint8
	enPassant  uint8
	fiftyMove  uint8

	ply    uint
	hisPly uint

	posKey uint64

	castlePerms uint8

	pceNum   [13]uint8
	bigPce   [3]uint8
	majorPce [3]uint8
	minorPce [3]uint8

	history MovesHistory
}

const (
	White = iota
	Black
	Both
)
