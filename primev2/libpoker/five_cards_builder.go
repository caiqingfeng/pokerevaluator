package libpoker

func buildFiveCards() {
	l := len(faces)
	com := make([][]uint32, 0)

	//four of  a kind
	var _s uint32 = 0
	for i:=l-1; i>0; i-- {
		for j:=l-1; j>0; j-- {
			if i == j {
				continue
			}
			v := make([]uint32, 6)
			v[0] = primes[faces[i]]
			v[1] = primes[faces[j]]
			t := v[0] * v[1] * v[0] * v[0] * v[0]
			score := uint32(four_of_akind_max + _s)
			addToSoreTbl(t, score)
			s := v[0] * v[1] * v[0] * v[0] * 43
			addToSoreTbl(s, score)
			_s++
		}
	}

	//full house
	_s = 0
	for i:=l-1; i>0; i-- {
		for j:=l-1; j>0; j-- {
			if i == j {
				continue
			}
			v := make([]uint32, 6)
			v[0] = primes[faces[i]]
			v[1] = primes[faces[j]]
			t := v[0] * v[1] * v[0] * v[1] * v[0]
			score := uint32(full_house_max + _s)
			addToSoreTbl(t, score)
			s := v[0] * v[1] * v[0] * v[1] * 43
			addToSoreTbl(s, score)
			_s++
		}
	}

	//straight
	com = make([][]uint32, 0)
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
		var y uint32 = uint32(x)
		for i:=0; i<6; i++ {
			s := t/v[i]
			score := uint32(straight_max + y)
			addToSoreTbl(s, score)
		}
		// fmt.Println(v, y, t)
	}
	// fmt.Println(scoreTbl, len(scoreTbl), straight_max)

	// three of a kind
	_s = 0
	com = make([][]uint32, 0)
	for i:=l-1; i>0; i-- {
		for j:=l-1; j>1; j-- {
			if i == j {
				continue
			}
			for x:=j-1; x>0; x-- {
				if x == i {
					continue
				}
				v := make([]uint32, 6)
				v[0] = primes[faces[i]]
				v[1] = primes[faces[j]]
				v[2] = primes[faces[x]]
				s := v[0] * v[0] * v[1] * v[2]
				score := uint32(three_of_akind_max + _s)
				addToSoreTbl(s*v[0], score)
				s = s * 43
				addToSoreTbl(s, score)
				_s++
			}
		}
	}

	//有赖子，就没有2对，所以这里只需要计算普通的2对即可
	//从13张牌里选出3张来，组成2对
	_s = 0
	for i:=l-1; i>1; i-- {
		for j:=i-1; j>0; j-- {
			for x:=l-1; x>0; x-- {
				if x == i || x == j {
					continue
				}
				v := make([]uint32, 6)
				v[0] = primes[faces[i]]
				v[1] = primes[faces[j]]
				v[2] = primes[faces[x]]
				s := v[0] * v[0] * v[1] * v[1] * v[2]
				score := uint32(two_pair_max + _s)
				addToSoreTbl(s, score)
				_s++
			}
		}
	}

	//one pair
	_s = 0
	for i:=l-1; i>0; i-- {
		for j:=l-1; j>2; j-- {
			if j==i {
				continue
			}
			for x:=j-1; x>1; x-- {
				if x==i {
					continue
				}
				for y:=x-1; y>0; y-- {
					if y==i {
						continue
					}
					v := make([]uint32, 6)
					v[0] = primes[faces[i]]
					v[1] = primes[faces[j]]
					v[2] = primes[faces[x]]
					v[3] = primes[faces[y]]
					s := v[0] * v[1] * v[2] * v[3] * v[0]
					score := uint32(one_pair_max + _s)
					addToSoreTbl(s, score)
					s = v[0] * v[1] * v[2] * v[3] * 43
					addToSoreTbl(s, score)
					_s++
				}
			}
		}
	}

	//high card
	_s = 0
	for i:=l-1; i>4; i-- {
		for j:=i-1; j>3; j-- {
			for x:=j-1; x>2; x-- {
				for y:=x-1; y>1; y-- {
					for z:=y-1; z>0; z-- {
						v := make([]uint32, 6)
						v[0] = primes[faces[i]]
						v[1] = primes[faces[j]]
						v[2] = primes[faces[x]]
						v[3] = primes[faces[y]]
						v[4] = primes[faces[z]]
						t := v[0] * v[1] * v[2] * v[3] * v[4]
						score := uint32(high_card_max + _s)
						_, existed := scoreTbl[t]
						if existed {
							continue
						}
						addToSoreTbl(t, score)
						_s++
					}
				}
			}
		}
	}
}
