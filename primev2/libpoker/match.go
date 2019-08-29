package libpoker

type Anb struct {
	Alice string
	AliceRank uint32
	Bob string
	BobRank uint32
	Result int
}

//1: alice排位(rank) 靠前，即牌力大
//2: alice排位(rank) 靠后，即牌力小
//0: 二者相等牌力
func ProcessMatch(match* Anb)  {
	//fmt.Println(match.Alice, match.Bob)
	match.AliceRank = EvaluateHandStr(match.Alice)
	match.BobRank = EvaluateHandStr(match.Bob)
	if match.AliceRank  < match.BobRank {
		match.Result = 1
	} else if match.AliceRank  > match.BobRank  {
		match.Result = 2
	}
}

