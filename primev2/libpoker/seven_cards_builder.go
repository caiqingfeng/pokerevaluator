package libpoker

import "sort"

func buildSevenCardsForFullHousePlus() {
	//four of  a kind
	l := len(faces)
	var _s uint32 = 0
	for i:=l-1; i>0; i-- {
		t_base := primes[faces[i]]
		t_base_four := t_base * t_base * t_base * t_base
		score := uint32(four_of_akind_max + _s)

		//有一个癞子，所以score都相同，等于最大的那个踢脚A或者K（如果是4条K的话）
		for j:=l-1; j>0; j-- {
			if i == j {
				continue
			}
			for x := l - 1; x > 0; x-- {
				//不能等于i，但可以等于j
				if x == i {
					continue
				}
				key2 := t_base_four * 43 * primes[faces[j]] * primes[faces[x]]
				addToSoreTbl(key2, score)
			}
		}

		//另外三张牌任选，没有癞子，只要不等于4条中的一张牌就可以
		for j:=l-1; j>0; j-- {
			if i == j {
				continue
			}
			for x:=l-1; x>0; x--{
				//不能等于i，但可以等于j
				if x == i {
					continue
				}
				for y:=l-1; y>0; y--{
					if y == i {
						continue
					}
					//在这层循环里就得到了4+3，另外3张牌的face可能相同，但都不会等于4条中的一个
					key1 := t_base_four * primes[faces[j]] * primes[faces[x]] * primes[faces[y]]
					//算分的时候，取一张最大的单牌即可
					offset := countOffset(n_ranks[faces[i]], n_ranks[faces[j]],
						n_ranks[faces[x]], n_ranks[faces[y]])
					addToSoreTbl(key1, score+offset)
				}
			}
		}

		//接下来要计算3条+癞子+其它任意3张牌的情况
		t_base_three_xn := t_base * t_base * t_base * 43
		//另外三张牌任选，没有癞子，只要不等于4条中的一张牌就可以
		for j:=l-1; j>0; j-- {
			if i == j {
				continue
			}
			for x:=l-1; x>0; x--{
				//不能等于i，但可以等于j
				if x == i {
					continue
				}
				for y:=l-1; y>0; y--{
					if y == i {
						continue
					}
					//在这层循环里就得到了4+3，另外3张牌的face可能相同，但都不会等于4条中的一个
					key1 := t_base_three_xn * primes[faces[j]] * primes[faces[x]] * primes[faces[y]]
					//算分的时候，取一张最大的单牌即可
					//这里有可能出现222+x+555这种情况，因为已经在555+x+222的时候生成了score，那次的score比本次的大
					//调用addToScoreTbl就直接返回了，不会覆盖
					offset := countOffset(n_ranks[faces[i]], n_ranks[faces[j]],
						n_ranks[faces[x]], n_ranks[faces[y]])
					addToSoreTbl(key1, score+offset)
				}
			}
		}
	}

	//full house
	//第一种情况：3条+1对+任意2张牌（不可能有癞子，否则就是4条了）
	_s = 0
	score := uint32(full_house_max + _s)
	for i:=l-1; i>0; i-- {
		t_base := primes[faces[i]]
		t_base_three := t_base * t_base * t_base
		for j:=l-1; j>0; j-- {
			if i == j {
				continue
			}
			_base := t_base_three * primes[faces[j]] * primes[faces[j]]
			//选好了3条+1对，再另选任意2张牌，只要不构成4条即可
			for x:=l-1; x>0; x-- {
				//同上，x可以等于j，但不能等于i
				if x == i {
					continue
				}
				for y:=l-1; y>0; y-- {
					//同上，y可以等于j,x，但不能等于i
					if y == i {
						continue
					}
					//在这层循环里就得到了3+2+2，另外2+2张牌的face可能相同2,3,4个，但都不会等于3条中的一个
					key1 := _base * primes[faces[x]] * primes[faces[y]]
					//这里有可能出现222+55+55这种情况，因为已经在555+5+222的时候生成了score，那次的score比本次的大
					//另外一种情况222+55+77的时候，因为已经加了222+77+55，也不会覆盖
					//调用addToScoreTbl就直接返回了，不会覆盖
					addToSoreTbl(key1, score+_s)
				}
			}
			_s ++
		}
	}
	//第二种情况：2对+1个癞子+任意2张牌（不能有3条，可以是3对）
	_s = 0
	score = uint32(full_house_max + _s)
	for i:=l-1; i>0; i-- {
		t_base := primes[faces[i]]
		t_base_two_xn := t_base * t_base * 43
		//循环里的每个score都是一样的，key不同
		//第一手将是AA+KK+X+QQ
		for j:=i-1; j>0; j-- {
			//第二对必须第一对小，否则直接选第二对作为三条
			_s = countOffset2(uint32(i), uint32(j))
			_base := t_base_two_xn * primes[faces[j]] * primes[faces[j]]
			for x:=l-1; x>0; x-- {
				//x不能等于i或者j
				if x == i || x == j{
					continue
				}
				for y:=l-1; y>0; y-- {
					//同上
					if y == i || y == j{
						continue
					}
					//在这层循环里就得到了2+x+2+2，最后2张牌的face可能相同2，但都不会等于前面对子中的任一个
					key1 := _base * primes[faces[x]] * primes[faces[y]]
					//这里有可能出现88+x+55+66这种情况，因为已经在88+x+66+55的时候生成了score，那次的score比本次的大
					//调用addToScoreTbl就直接返回了，不会覆盖
					addToSoreTbl(key1, score+_s)
				}
			}
		}
	}

}

//AAAAK->0, AAAAQ->1, ... AAAA2->11
//22223->155
func countOffset(base uint32, k1 uint32, k2 uint32, k3 uint32) uint32 {
	arr := []int {int(k1), int(k2), int(k3)}
	sort.Ints(arr)
	r := (12 - int(base))*12
	if arr[2] > int(base) {
		r += 12 - arr[2]
	} else {
		r += 11 - arr[2]
	}
	return uint32(r)
}

//3条最大的两个踢脚
//每种3条有12*11/2=66种不同的踢脚
func countOffset3(base uint32, k1 uint32, k2 uint32, k3 uint32, k4 uint32) uint32 {
	arr := []int {int(k1), int(k2), int(k3), int(k4)}
	sort.Ints(arr)
	r := (13 - int(base))*66
	for i:=13; i>arr[3]; i-- {
		if i != int(base) {
			r += i-2
		}
	}
	r += arr[3] - arr[2] - 1
	if int(base) < arr[3] && int(base) > arr[2] {
		r --
	}
	return uint32(r)
}

//2对最大的踢脚
//每种2对11种不同的踢脚
//为方便这里只计算pair2<pair1
func countOffset4(pair1 int, pair2 int, k1 int, k2 int, k3 int) uint32 {
	arr := []int {k1, k2, k3}
	sort.Ints(arr)
	r := 0
	for i:=13; i>pair1; i-- {
		r += (12 - (13 - i)) * 11 //第一对等于AA时，有12对可以第二对；等于KK就只计算11对第二对
	}
	r += (pair1 - pair2 - 1) * 11
	r += 13 - arr[2]
	if arr[2] < pair2 {
		r -= 2
	} else if arr[2] < pair1 {
		r --
	}
	return uint32(r)
}

//1对的踢脚
//每种1对有(12里选3种)不同的踢脚，即12*11*10/(3*2*1)=220
//AAKQJ = 0, 22345 = 2859
func countOffset5(pair1 int, k1 int, k2 int, k3 int, k4 int, k5 int) uint32 {
	arr := []int {k1, k2, k3, k4, k5}
	sort.Ints(arr)
	r := (13 - pair1) * 220
	//确定了pair，对于每个给定的kicker1有从2到kicker1个kicker2选择，同时有2到kicker2个k3
	//比如计算88Q75,首先计算88Axx, 88Kxx的个数
	//所以计算到kicker1的偏移量公式：r=11*10+10*9+...+
	for i:=13; i>arr[4]; i-- {
		if i == pair1 {
			continue
		}
		if  i > pair1 {
			r += (11 - (13 - i)) * (10 - (13 - i)) / 2 //kicker1等于A时，有11*10其它散牌可做kicker
		} else if i < pair1 {
			r += (11 - (13 - i) + 1) * (10 - (13 - i) + 1) / 2 //kicker1等于A时，有11*10其它散牌可做kicker
		}
	}

	//确定了pair，kicker1时，对于每个给定的kicker2有从2到k2个k3选择
	//比如计算88Q75的偏移量时，首先88QJx共有 8个，88QTx共有7, 88Q9x共有6
	//所以计算偏移量公式：r+=10+9+...
	for i:=arr[4]-1; i>arr[3]; i-- { //循环计算88Q[JT9]
		//offset = i - 1(例如k2=j时，K3可以是2到T共有9个）- 1 （减掉8）
		if arr[3] > pair1 {
			r += i - 1 - 1
		} else {
			r += i - 1
		}
	}

	//确定了pair，kicker1，kicker2时,再计算k3的偏移相对简单
	//比如计算88Q75的偏移量时，88Q7x只有1个88Q76在88Q75前面
	//只需要判断pair是不是落在<arr[0], arr[1]>区间
	if arr[2] > pair1 || arr[3] < pair1{
		r += arr[3] - 1 - arr[2]
	} else {
		r += arr[3] - 1 - arr[2] - 1
	}

	return uint32(r)
}

func countOffset2(base uint32, k1 uint32) uint32{
	r := uint32((13-base)*12)
	if k1 > base {
		r += 13 - k1
	} else {
		r += 12 - k1
	}
	return r
}

//针对7张牌里的顺子建表
func buildSevenCardsForStraight() {
	//第一种情况：5张连牌+2张任意牌型，没有癞子
	//第二种情况：4张连牌+1个癞子+2张任意牌型
	l := len(faces)
	//straight
	//这个加map的做法非常巧妙
	//把癞子和5张联牌组成了6选5的排列组合
	com := make([][]uint32, 0)
	for i:=l-1; i>3; i-- {
		v := make([]uint32, 6)
		v[0] = primes[faces[i]]
		v[1] = primes[faces[i-1]]
		v[2] = primes[faces[i-2]]
		v[3] = primes[faces[i-3]]
		v[4] = primes[faces[i-4]]
		v[5] = 43
		com = append(com, v)
	}
	// fmt.Println(com, len(com))
	for x, v := range com {
		t := v[0] * v[1] * v[2] * v[3] * v[4] * v[5]
		var y = uint32(x)
		for i:=0; i<6; i++ {
			s := t/v[i]
			score := uint32(straight_max + y)
			for x:=l-1; x>0; x-- {
				//added 2019.08.04 支持按6张牌来找
				key2 := s * primes[faces[x]]
				addToSoreTbl(key2, score)
				for y:=l-1; y>0; y--{
					key1 := s * primes[faces[x]] * primes[faces[y]]
					addToSoreTbl(key1, score)
				}
			}
		}
		// fmt.Println(v, y, t)
	}

}

//针对7张牌里的3条建表
func buildSevenCardsForThreeOfAKind() {
	//Three of A Kind
	//第一种情况：3条+任意4张散牌（不可能有癞子，和其它对子，否则就是4条了）
	_s := uint32(0)
	l := len(faces)
	score := uint32(three_of_akind_max + _s)
	for i:=l-1; i>0; i-- {
		t_base := primes[faces[i]]
		t_base_three := t_base * t_base * t_base
		for j:=l-1; j>0; j-- {
			if i == j {
				continue
			}
			_base := t_base_three * primes[faces[j]]
			//选好了3条+1散，再另选任意3张牌
			for x:=j-1; x>0; x-- {
				//同上，x不能等于i，j
				if x == i || x == j{
					continue
				}
				for y:=x-1; y>0; y-- {
					//同上，y不能等于i，j，x
					if y == i || y == j || y == x{
						continue
					}
					for z:=y-1; z>0; z--{
						//同上，z不能等于i，j，x, y
						if y == i || y == j || y == x{
							continue
						}
						//在这层循环里就得到了3+1+1+1+1，另外4张牌的face都不同并且都不会等于3条中的一个
						key1 := _base * primes[faces[x]] * primes[faces[y]] * primes[faces[z]]
						offset := countOffset3(uint32(i), uint32(j),
							uint32(x), uint32(y), uint32(z))
						addToSoreTbl(key1, score+offset)
					}
				}
			}
		}
	}
	//第二种情况：1对+1个癞子+任意4张牌（不能有3条、2对）
	_s = 0
	score = uint32(three_of_akind_max + _s)
	for i:=l-1; i>0; i-- {
		t_base := primes[faces[i]]
		t_base_two_xn := t_base * t_base * 43
		//第一手将是AA+X+KQJT
		for j:=l-1; j>0; j-- {
			_base := t_base_two_xn * primes[faces[j]]
			for x:=j-1; x>0; x-- {
				//x不能等于i或者j其实只需要判断x==i
				if x == i || x == j{
					continue
				}
				for y:=x-1; y>0; y-- {
					//同上其实只需要判断y==i
					if y == i || y == j{
						continue
					}
					for z:=y-1; z>0; z-- {
						//同上,其实只需要判断z==i
						if z == i || z == j || z == y {
							continue
						}
						//在这层循环里就得到了2+x+1+1+1+1
						key1 := _base * primes[faces[x]] * primes[faces[y]] * primes[faces[z]]
						offset := countOffset3(uint32(i), uint32(j),
							uint32(x), uint32(y), uint32(z))
						addToSoreTbl(key1, score+offset)
					}
				}
			}
		}
	}

}

//针对7张牌里的2对建表
func buildSevenCardsForTwoPair() {
	//Two Pair，不可能有癞子，只要对普通2对+3建表
	//2对+任意3张散牌
	l := len(faces)
	for i:=l-1; i>1; i-- {
		for j:=i-1; j>0; j-- {
			//至此选出了大小两对
			_base := primes[faces[i]]*primes[faces[i]]*primes[faces[j]]*primes[faces[j]]
			for x:=l-1; x>0; x-- {
				if x == i || x == j {
					continue
				}
				for y:=l-1; y>0; y--{
					if y == i || y == j {
						continue
					}
					for z:=l-1; z>0; z-- {
						if z == i || z == j {
							continue
						}
						key1 := _base * primes[faces[x]] * primes[faces[y]] * primes[faces[z]]
						offset := countOffset4(i, j, x, y, z)
						score := uint32(two_pair_max + offset)
						addToSoreTbl(key1, score)
					}
				}
			}
		}
	}
}

//针对7张牌里的1对建表
func buildSevenCardsForOnePair() {
	//OnePair
	//第一种情况，1对+任意5张散牌，没有癞子
	l := len(faces)
	for i:=l-1; i>0; i-- { //一对
		_base := primes[faces[i]] * primes[faces[i]]
		for j:=l-1; j>0; j-- { //kicker1
			if j == i {
				continue
			}
			for x:=j-1; x>0; x-- { //kicker2
				if x == i {
					continue
				}
				for y:=x-1; y>0; y--{ //kicker3
					if y == i {
						continue
					}
					for z:=y-1; z>0; z-- { //kicker4
						if z == i {
							continue
						}
						for t:=z-1; t>0; t-- { //kicker5
							key1 := _base * primes[faces[j]] * primes[faces[x]] * primes[faces[y]] * primes[faces[z]] * primes[faces[t]]
							offset := countOffset5(i, j, x, y, z, t)
							score := uint32(one_pair_max + offset)
							addToSoreTbl(key1, score)
						}
					}
				}
			}
		}
	}
	//第2种情况，1癞子+任意6张散牌
	_base := uint32(43)
	for i:=l-1; i>0; i-- { //k1
		_base = 43 * primes[faces[i]]
		for j:=i-1; j>0; j-- { //k2
			for x:=j-1; x>0; x-- { //k3
				for y:=x-1; y>0; y--{ //k4
					for z:=y-1; z>0; z-- { //k5
						for t:=z-1; t>0; t-- { //k6
							key1 := _base * primes[faces[j]] * primes[faces[x]] * primes[faces[y]] * primes[faces[z]] * primes[faces[t]]
							offset := countOffset5(i, j, x, y, z, t)
							score := uint32(one_pair_max + offset)
							addToSoreTbl(key1, score)
						}
					}
				}
			}
		}
	}
}