package pieces

type PieceCode uint8

const (
	Empty PieceCode = iota
	WPawn
	WKnight
	WBishop
	WRook
	WQueen
	WKing
	BPawn
	BKnight
	BBishop
	BRook
	BQueen
	BKing
)
