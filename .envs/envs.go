package envs

const (
	LENASCII            = uint16(len(ASCIILIST) - 1)
	MAXGRAYCOLOR uint16 = 65535
)

var ASCIILIST = [...]rune{'.', ':', '*', '+', '%', '#', '@'}