package libpoker

import "testing"

func TestFullHouse(t *testing.T) {
	BuildScoreTbl()
	if len(scoreTbl) == 0{
		t.Error("no build")
	}
	//fmt.Println("len=", len(scoreTbl))
	if score, found := FastDetector("AhAcKsKcQsQhXn"); !found ||score != full_house_max {
		t.Error("should be AAAKK", score, full_house_max)
	}

	if score, found := FastDetector("AhAcKsKcQsJhXn"); !found || score != full_house_max{
		t.Error("should be AAAKK", score, full_house_max)
	}

	if score, found := FastDetector("AhAcAs2c2sJhQs"); !found || score != full_house_max+11{
		t.Error("should be AAA22", score, full_house_max+11)
	}

	if score, found := FastDetector("AhAcAs2c2s2h3s"); !found || score != full_house_max+11{
		t.Error("should be AAA22", score, full_house_max+11)
	}

	if score, found := FastDetector("KsKcKdAcAh"); !found || score != full_house_max+12{
		t.Error("should be KKKAA", score, full_house_max+12)
	}

	if score, found := FastDetector("KsKcKdAcAhJcJs"); !found || score != full_house_max+12{
		t.Error("should be KKKAA", score, full_house_max+12)
	}

	if score, found := FastDetector("QsQcQdAcAhJcJs"); !found || score != full_house_max+12*2{
		t.Error("should be QQQAA", score, full_house_max+12*2)
	}

	if score, found := FastDetector("JsJcJdAcAhKcKs"); !found || score != full_house_max+12*3{
		t.Error("should be JJJAA", score, full_house_max+12*3)
	}

	if score, found := FastDetector("TsTcTdAcAhKcKs"); !found || score != full_house_max+12*4{
		t.Error("should be TTTAA", score, full_house_max+12*4)
	}

	if score, found := FastDetector("TsTcTd9c9hKcQs"); !found || score != full_house_max+12*4+4{
		t.Error("should be TTTAA", score, full_house_max+12*4+4)
	}

	if score, found := FastDetector("AhAc2s2cQsJhXn"); !found || score != full_house_max+11{
		t.Error("should be AAA22", score, full_house_max+11)
	}

	if score, found := FastDetector("KhKc2s2cQsJhXn"); !found || score != full_house_max+11+12{
		t.Error("should be KKK22", score, full_house_max+12+11)
	}

	if score, found := FastDetector("ThTc9s9cQsJhXn"); !found || score != full_house_max+12*4+4{
		s, _ := FastDetector("ThTc9s9cTsQsJh")
		t.Error("should be TTT99", score, full_house_max+12*4+4, s)
	}

	if score, found := FastDetector("9h6c6s9c3s8hXn"); !found || score != full_house_max+12*5+7{
		t.Error("should be 99966", score, full_house_max+12*5+7)
	}
}

func TestFourOfAKind(t *testing.T) {
	scoreTbl = make(map[uint32] uint32)
	BuildScoreTbl()

	score, found := FastDetector("AcAdAsXnKc")
	if score != four_of_akind_max || !found{
		t.Error("failed", score)
	}

	score, found = FastDetector("AcAdAsXnKcKdKs")
	if score != four_of_akind_max || !found{
		t.Error("failed", score)
	}

	score, found = FastDetector("2c2d2s2h3c3d3s")
	if score != four_of_akind_max+155 || !found{
		t.Error("failed", score)
	}
}

func TestFlush(t *testing.T) {
	BuildScoreTbl()
	if flushHand, isFlush := FastIsFlush("6cAh6h2h6d9h8h", 0); !isFlush{
		t.Error("failed", flushHand)
	}

	if flushHand, isFlush := FastIsFlush("5hTdXn9cKd3d2d", 0); !isFlush{
		t.Error("failed", flushHand)
	}

}

func TestStraight(t *testing.T) {
	BuildScoreTbl()
	if score, found := FastDetector("AcKdQhJsTc6c4s"); score != straight_max || !found{
		t.Error("failed", score, straight_max)
	}

	if score, found := FastDetector("AcKdQhJsXn6c4s"); score != straight_max || !found{
		t.Error("failed", score, straight_max)
	}

	if score, found := FastDetector("KdQhJsTcXn6c4s"); score != straight_max || !found{
		t.Error("failed", score, straight_max)
	}

	if score, found := FastDetector("QdQhJsTcXn9c4s"); score != straight_max+1 || !found{
		t.Error("failed", score, straight_max)
	}

	if score, found := FastDetector("2d2h3s4cXn9cAs"); score != straight_max+9 || !found{
		t.Error("failed", score, straight_max)
	}

	if score, found := FastDetector("2d2h3s4cXnAc5s"); score != straight_max+8 || !found{
		t.Error("failed", score, straight_max)
	}

	if score, found := FastDetector("2c2d2h5s3c6c4s"); score != straight_max+8 || !found{
		t.Error("failed", score, straight_max+8)
	}
}

func TestThreeOfAKind(t *testing.T) {
	BuildScoreTbl()
	if score, found := FastDetector("AcAdAsQcKc"); score != three_of_akind_max || !found{
		t.Error("failed", score)
	}

	if score, found := FastDetector("2c2d2sQcKc"); score != three_of_akind_max+12*66+11 || !found{
		t.Error("failed", score, three_of_akind_max+12*66+11)
	}

	if score, found := FastDetector("2c2d2sQcKcJsTh"); score != three_of_akind_max+12*66+11 || !found{
		t.Error("failed", score, three_of_akind_max+12*66+11)
	}

	if score, found := FastDetector("Xn3h6cQdQsTdKd"); score != three_of_akind_max+2*66+11+1 || !found{
		t.Error("failed", score, three_of_akind_max+2*66+11+1)
	}

	if score, found := FastDetector("6cQdQsTdKd3dQh"); score != three_of_akind_max+2*66+11+1 || !found{
		t.Error("failed", score, three_of_akind_max+2*66+11+1)
	}

	if score, found := FastDetector("QdQsTdKdQh"); score != three_of_akind_max+2*66+11+1 || !found{
		t.Error("failed", score, three_of_akind_max+2*66+11+1)
	}
}

func TestOnePair(t *testing.T) {
	BuildScoreTbl()
	if score, found := FastDetector("2s2hKh8h9s5s4h"); score != one_pair_max+12*220+11*10/2+9+8+7 || !found {
		t.Error("failed", score, one_pair_max+12*220+11*10/2+9+8+7)
	}
}

func TestHighCard(t *testing.T) {
	BuildScoreTbl()
	if score, found := FastDetector("AcKsQsJs9s"); score != high_card_max || !found {
		t.Error("failed", score, high_card_max)
	}

	//AKT76的offset=AKQJ[9..2]+AKQT[9..2]+...+AKQ32+AKJ[T..3][9..2]+AKT8[7..2]
	if score, found := FastDetector("AcKsTs7s6s"); score != high_card_max+8+9*8/2+9*8/2+7+6 || !found {
		t.Error("failed", score, high_card_max+8+9*8/2+9*8/2+7+6)
	}
}
