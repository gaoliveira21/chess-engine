package board

const MAX_MOVES = 2048

type MoveHistory struct {
	Move        uint32
	PosKey      uint64
	EnPassant   uint8
	FiftyMove   uint8
	CastlePerms uint8
}

type MovesHistory [MAX_MOVES]MoveHistory
