package libpoker

import "sort"

var fourCardsToBeStraight = map[uint32]uint32 {}

func genFourCardsMap() {
	if len(fourCardsToBeStraight) > 0 {
		return
	}

	straightKeys := []uint32 {41*37*31*29*23, 37*31*29*23*19, 31*29*23*19*17, 29*23*19*17*13,
		23*19*17*13*11, 19*17*13*11*7, 17*13*11*7*5, 13*11*7*5*3,
		11*7*5*3*2, 7*5*3*2*41}
	for i:=0; i<len(straightKeys); i++ {
		maxFace := getCardFaceByOffset("A", i*(-1))
		for j:=0; j<4; j++ {
			key := straightKeys[i]/primes[getCardFaceByOffset(maxFace, j*(-1))]
			fourCardsToBeStraight[key] = straightKeys[i]
		}
	}
}

// 不管是7张牌还是5张牌，可以快速探测是否是full house或者four of a kind
// 如果是full house或者four of a kind，就不可能是flush or straight
// 更不可能是straight flush

func FastDetector(hand string) (uint32, bool) {
	cardNum := len(hand)/2
	key := uint32(1)
	for i:=0; i<cardNum; i++ {
		key *= primes[hand[2*i:2*i+1]]
	}
	if score, found := scoreTbl[key]; found {
		return score, found
	}
	return 0, false
}

func maxSuit(hand string) (string, uint32){
	club := uint32(0)
	clubHand := ""
	diamond := uint32(0)
	diamondHand := ""
	heart := uint32(0)
	heartHand := ""
	spades := uint32(0)
	spadesHand := ""
	l := len(hand)
	withGhost := false
	for i:=0; i<l/2; i++ {
		switch hand[2*i+1:2*i+2] {
		case "s":
			spades++
			spadesHand += hand[2*i:2*i+2]
		case "h":
			heart++
			heartHand += hand[2*i:2*i+2]
		case "d":
			diamond++
			diamondHand += hand[2*i:2*i+2]
		case "c":
			club++
			clubHand += hand[2*i:2*i+2]
		case "n":
			spades++
			heart++
			diamond++
			club++
			withGhost =  true
		}
	}
	arr := []int {int(club), int(diamond), int(heart), int(spades)}
	sort.Ints(arr)
	newHand := ""
	suit := uint32(0)
	switch uint32(arr[3]) {
	case club:
		newHand, suit = clubHand, club
	case diamond:
		newHand, suit = diamondHand, diamond
	case heart:
		newHand, suit = heartHand, heart
	case spades:
		newHand, suit = spadesHand, spades
	}
	if withGhost {
		//replace ghost with a max card
		//此处最多有6张同花+1张癞子,最少4张同花
		l := len(newHand)/2
		if l > 6 || l < 4 {
			return newHand+"Xn", suit
		}

		if hand2, isStraightFlush := makeStraightFlush(newHand+"Xn"); isStraightFlush {
			return hand2, 5
		}
		for i:=13; i>0; i-- {
			card := faces[i] + newHand[1:2]
			if !hasCard(newHand, card) {
				newHand += card
				break
			}
		}
	}
	return newHand, suit
}

//在已经判断是3条或者straight之上的前提下，判断是否flush or straightflush
func FastIsFlush(hand string, score uint32) (string, bool) {
	flushHand, flushCount := maxSuit(hand)
	if flushCount < 5 {
		return hand, false
	}

	return flushHand, true
}

//取5张最大face的牌
func maxFace(hand string) string {
	l := len(hand)
	_f := []int{}
	for i:=0; i<l/2; i++ {
		_f = append(_f, int(toNRank(hand[2*i:2*i+2])))
	}
	sort.Ints(_f)

	newHand := ""
	for i:=len(_f)-1; i>=len(_f)-5; i-- {
		face := _f[i]
		for j:=0; j<l/2; j++ {
			if int(n_ranks[hand[2*j:2*j+1]]) == face {
				newHand += hand[2*j:2*j+2]
				break
			}
		}
	}
	return newHand
}
