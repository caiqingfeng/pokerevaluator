package libpoker

func BuildScoreTbl() {
	if len(scoreTbl) > 0 {
		return
	}

	//AsKsXnQhJc,也是一个顺子，它的score=straight_max
	//key=41*37*31*29*43,scoreTbl[41*37*31*29*43] = straight_max
	buildFiveCards()
	//AsKsTsQhJc,是一个顺子，5张牌它的(k,v) = (41*37*31*29*23, straight_max)
	//7张牌，AsKsTsQhJc+2c6c,这手牌的score = straight_max，但是key=41*37*31*29*23*2*7
	//scoreTbl[41*37*31*29*23*2*7]=straight_max
	//scoreTbl[41*37*31*29*23] = straight_max
	buildSevenCardsForFullHousePlus()
	buildSevenCardsForStraight()
	buildSevenCardsForThreeOfAKind()
	buildSevenCardsForTwoPair()
	buildSevenCardsForOnePair()
}

func LenOfScoreTbl() int {
	return len(scoreTbl)
}