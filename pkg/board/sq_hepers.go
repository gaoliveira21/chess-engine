package board

import "fmt"

func Sq64ToRankAndFile(sq64 int) (BoardFile, BoardRank) {
	f := BoardFile(sq64 % 8)
	r := BoardRank(sq64 / 8)
	return f, r
}

func Sq120ToRankAndFile(sq120 int) (BoardFile, BoardRank) {
	if sq120 < 21 || sq120 >= 99 {
		return FileNone, RankNone
	}

	lastDigit := sq120 % 10
	if lastDigit == 0 || lastDigit == 9 {
		return FileNone, RankNone
	}

	f := BoardFile((sq120 - 21) % 10)
	r := BoardRank((sq120 - 21) / 10)
	return f, r
}

func GetSq64Index(f BoardFile, r BoardRank) int {
	return int(r)*8 + int(f)
}

func GetSq120Index(f BoardFile, r BoardRank) int {
	return (21 + int(f)) + (10 * int(r))
}

func Sq64ToSq120(sq64 int) int {
	f, r := Sq64ToRankAndFile(sq64)

	return GetSq120Index(f, r)
}

func Sq120ToSq64(sq120 int) (int, error) {
	f, r := Sq120ToRankAndFile(sq120)
	if f == FileNone || r == RankNone {
		return 0, fmt.Errorf("invalid square: %d", sq120)
	}

	return GetSq64Index(f, r), nil
}
