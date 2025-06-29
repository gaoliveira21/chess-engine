package board

type BoardFile uint8

const (
	FileA BoardFile = iota
	FileB
	FileC
	FileD
	FileE
	FileF
	FileG
	FileH
	FileNone
)

func (f BoardFile) String() string {
	if f == FileNone {
		return " "
	}

	return string('A' + f)
}
