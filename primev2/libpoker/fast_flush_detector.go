package libpoker

func FastFlushDetector(hand string) uint32{
	if score, found := FastDetector(hand); found {
		if score < straight_max + 10 && score >= straight_max {
			//straight flush
			score -= straight_max
			return score
		}
		//其它5张牌的flush都应该能一次在表里找到，所以直接返回
		score = score - high_card_max + flush_max
		return score
	}

	//6张牌或者7张牌里选5张最大的，再调用一次FastDetector
	maxHand := maxFace(hand)
	if score, found := FastDetector(maxHand); found {
		score = score - high_card_max + flush_max
		return score
	}

	return 0
}
