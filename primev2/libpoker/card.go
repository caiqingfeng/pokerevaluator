package libpoker

var faces = []string {
	"A", "2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A",
}
var primes = map[string]uint32 {
	"2":2,
	"3":3,
	"4":5,
	"5":7,
	"6":11,
	"7":13,
	"8":17,
	"9":19,
	"T":23,
	"J":29,
	"Q":31,
	"K":37,
	"A":41,
	"X":43,
}

var n_ranks = map[string]uint32 {
	"2": 0,
	"3": 1,
	"4": 2,
	"5": 3,
	"6": 4,
	"7": 5,
	"8": 6,
	"9": 7,
	"T": 8,
	"J": 9,
	"Q": 10,
	"K": 11,
	"A": 12,
	"X": 13,
}

func toNRank(card string) uint32 {
	return n_ranks[card[0:1]]
}

func getCardFaceByOffset(face string, offset int) string {
	rank := int(n_ranks[face]) + offset
	switch rank {
	case 0:
		return "2"
	case 1:
		return "3"
	case 2:
		return "4"
	case 3:
		return "5"
	case 4:
		return "6"
	case 5:
		return "7"
	case 6:
		return "8"
	case 7:
		return "9"
	case 8:
		return "T"
	case 9:
		return "J"
	case 10:
		return "Q"
	case 11:
		return "K"
	case 12:
		return "A"
	}
	return ""
}

func getCardFaceByPrime(p uint32) string {
	switch p {
	case 2:
		return "2"
	case 3:
		return "3"
	case 5:
		return "4"
	case 7:
		return "5"
	case 11:
		return "6"
	case 13:
		return "7"
	case 17:
		return "8"
	case 19:
		return "9"
	case 23:
		return "T"
	case 29:
		return "J"
	case 31:
		return "Q"
	case 37:
		return "K"
	case 41:
		return "A"
	case 43:
		return "X"
	}
	return ""
}

var suits = map[string]uint32 {
	"s": 1,
	"h": 2,
	"d": 4,
	"c": 8,
}

