package board

import "strconv"

type BoardRank uint8

const (
	Rank1 BoardRank = iota
	Rank2
	Rank3
	Rank4
	Rank5
	Rank6
	Rank7
	Rank8
	RankNone
)

func (r BoardRank) String() string {
	if r == RankNone {
		return " "
	}

	return strconv.Itoa(int(r + 1))
}
