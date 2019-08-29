package libpoker

import "sort"

// +--------+--------+--------+--------+
// |xxxbbbbb|bbbbbbbb|cdhsrrrr|xxpppppp|
// +--------+--------+--------+--------+
// xxxAKQJT 98765432 CDHSrrrr xxPPPPPP
// 00001000 00000000 01001011 00100101    King of Diamonds(Kd)
// 00000000 00001000 00010011 00000111    Five of Spades(5s)
// 00000010 00000000 10001001 00011101    Jack of Clubs(Jc)
// 00011111 11111111 11111101 00011111    赖子(Xn)
// p = prime number of rank (deuce=2,trey=3,four=5,...,ace=41)
// r = rank of card (deuce=0,trey=1,four=2,five=3,...,ace=12)
// cdhs = suit of card (bit turned on based on suit of card)
// b = bit turned on depending on rank of card
    // Straight Flush   10 
    // Four of a Kind   156      [(13 choose 2) * (2 choose 1)]
    // Full Houses      156      [(13 choose 2) * (2 choose 1)]
    // Flush            1277     [(13 choose 5) - 10 straight flushes]
    // Straight         10 
    // Three of a Kind  858      [(13 choose 3) * (3 choose 1)]
    // Two Pair         858      [(13 choose 3) * (3 choose 2)]
    // One Pair         2860     [(13 choose 4) * (4 choose 1)]
    // High Card      + 1277     [(13 choose 5) - 10 straights]
    // -------------------------
    // TOTAL            7462

const straight_flush_max uint32 = 0 
const four_of_akind_max uint32 = 10 
const full_house_max uint32 = 10+156 //166
const flush_max uint32 = 10+156+156  //322
const straight_max uint32 = flush_max+1277 //1599
const three_of_akind_max uint32 = straight_max+10 //1609
const two_pair_max uint32 = three_of_akind_max+858 //2467
const one_pair_max uint32 = two_pair_max+858 //3325
const high_card_max uint32 = one_pair_max+2860 //6185


const ghost_ptn string = "Xn"
// 43, 41*37*31*29*23: AKQJT
// 43, 37*31*29*23*19
var scoreTbl = map[uint32] uint32 {}

func addToSoreTbl(k, v uint32) {
    //fmt.Println("key=", k, ", score=", v)
    _, existed := scoreTbl[k]
    if !existed {
        scoreTbl[k] = v
    }
}

//5张牌里有一个Xn，把Xn替换成合适的牌，让该手牌成为straightFlush
func makeStraightFlush(hand string) (string, bool) {
    cardNum := len(hand)/2
    if cardNum > 7 || cardNum < 5 {
        return hand, false
    }

    key := uint32(1)
    suit :=  ""
    newHand := ""
    var cardPrimes = []int{}

    for i:=0; i<cardNum; i++ {
        key *= primes[hand[2*i:2*i+1]]
        if hand[2*i+1] != 'n' {
            suit = hand[2*i+1:2*i+2]
            newHand += hand[2*i:2*i+2]
            cardPrimes = append(cardPrimes, int(primes[hand[2*i:2*i+1]]))
        }
    }
    if key%43 != 0 { //没有癞子
        return hand, false
    }
    key /= 43
    newCard := ""
    genFourCardsMap()
    sort.Ints(cardPrimes)

    for i:=len(cardPrimes)-1; i>=3; i-- {
        for j:=i-1; j>=2; j-- {
            for x:=j-1; x>=1; x-- {
                for y:=x-1; y>=0; y-- {
                    k := uint32(cardPrimes[i] * cardPrimes[j] * cardPrimes[x] * cardPrimes[y])
                    if straightKey, found := fourCardsToBeStraight[k]; found {
                        p := straightKey/k
                        newCard = getCardFaceByPrime(p)+suit
                        newHand = newCard+getCardFaceByPrime(uint32(cardPrimes[i]))+suit+
                            getCardFaceByPrime(uint32(cardPrimes[j]))+suit+
                            getCardFaceByPrime(uint32(cardPrimes[x]))+suit+
                            getCardFaceByPrime(uint32(cardPrimes[y]))+suit
                        return newHand, true
                    }
                }
            }
        }
    }
    return newHand, false
}

func EvaluateHandStr(hand string) uint32 {
    //fmt.Println(hand)
    processHand := hand
    if s, found := FastDetector(processHand); found{
        if s < straight_max {
            return s
        }
        //否则得到了3条和顺子、2对、1对，要继续判断是否同花顺或者同花
        // todo
        if flushHand, isFlush := FastIsFlush(processHand, s); isFlush {
            return FastFlushDetector(flushHand)
        }
        return s
    } else {
        //连1对都不是，那只可能是high cards，取5张最大的
        processHand = maxFace(hand)
        if flushHand, isFlush := FastIsFlush(hand, s); isFlush {
           processHand = flushHand
           return FastFlushDetector(processHand)
        }

        score, _ := FastDetector(processHand)
        return score
    }

    return 0
}

func Score2str(score uint32) string {
    switch {
    case score == straight_flush_max:
        return "皇家同花顺"
    case score < four_of_akind_max && score > straight_flush_max:
        return "同花顺"
    case score < full_house_max && score >= four_of_akind_max:
        return "四条"
    case score < flush_max && score >= full_house_max:
        return "葫芦"
    case score < straight_max && score >= flush_max:
        return "同花"
    case score < three_of_akind_max && score >= straight_max:
        return "顺子"
    case score < two_pair_max && score >= three_of_akind_max:
        return "三条"
    case score < one_pair_max && score >= two_pair_max:
        return "两对"
    case score < high_card_max && score >= one_pair_max:
        return "一对"
    default:
        return "散牌"
    }
    return "散牌"
}

func hasCard(hand string, card string) bool {
    handLen := len(hand)
    for i:=0; i<handLen/2; i++ {
        if hand[2*i:2*i+2] == card {
            return true
        }
    }
    return false
}

