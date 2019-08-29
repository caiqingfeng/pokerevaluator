## 德州扑克算法
德州扑克系列扑克游戏，最多5张公共牌，2张或4张手牌，进行组合，判断最大的牌型。其算法实现属于看起来很简单，但是做的好却非常有挑战。之前开源过一个libpoker的C++代码，https://github.com/caiqingfeng/libpoker.  里面用正则表达式对牌型进行判断，代码可读性很好，非常容易理解，但是效率不高。一个入门级别的后台工程师基本都可以独立实现德州扑克算法，但是如果要想做到效率非常高，可扩展性强，需要克服以下几个难点。

### 7张牌快速得到最大牌型
关键点是需要避免7选5，21次循环

### 算法可以适应2张到9张牌
对应德州扑克游戏、大菠萝游戏、奥马哈游戏等等不同玩法，比如单纯比较手牌、三张头道牌、4张手牌选2+5张公共牌选3的奥马哈等等。

### 可以支持癞子
癞子即百搭牌，可以把四张同花升级为5张同花，升级为顺子等等，这种玩法在国内大菠萝手机游戏上小范围流行。

### Cactus Kev's Poker Hand Evaluator
这篇 http://www.suffecool.net/poker/evaluator.html （链接地址好像失效了）给出了一个非常好的算法。用32bits来表示牌面（face)和花色（suit），其中最有创意的是用13个素数来表示face，计算牌面的素数乘积，很容易就可以检测出来对子及三条等组合。

```
+--------+--------+--------+--------+
|xxxbbbbb|bbbbbbbb|cdhsrrrr|xxpppppp|
+--------+--------+--------+--------+
p = prime number of rank (deuce=2,trey=3,four=5,...,ace=41)
r = rank of card (deuce=0,trey=1,four=2,five=3,...,ace=12)
cdhs = suit of card (bit turned on based on suit of card)
b = bit turned on depending on rank of card


Hand Value	Unique	Distinct
Straight Flush	40	10
Four of a Kind	624	156
Full Houses	3744	156
Flush	5108	1277
Straight	10200	10
Three of a Kind	54912	858
Two Pair	123552	858
One Pair	1098240	2860
High Card	1302540	1277
TOTAL	2598960	7462

```
对任一手牌，通过计算7462个不同的rank。这个算法很巧妙，有很多基于它的实现。原算法没有给出7张牌，更没有支持癞子的说明。

### 已有的优秀算法实现A - 52bits位运算
https://github.com/HenryRLee/PokerHandEvaluator/blob/master/Documentation/Algorithm.md

这个算法用一个52bits来map 52张牌，然后利用位运算，首先判断是否同花，然后根据7张牌，如果是同花就不可能是fullhouse或者four of a kind，再计算13bits的5元组（quinary)；

 |   Spades   |   Hearts   |  Diamonds  |   Clubs   |
 23456789TJQKA23456789TJQKA23456789TJQKA23456789TJQKA
 0001010010000000100000000000000000010000010000000001


算法对于百搭牌（癞子）的适应性不好，比如这手牌："Xn9cKd3d2d9d9s",其中Xn表示百搭牌，其最大的牌型应该是4条9，采用这个算法可能误判为同花。

### 已有的优秀算法实现B - binary search
https://www.johnbelthoff.com/web-programming/poker-project/cactus-kev.aspx

实现了上文中Cactus Kev的7张牌算法，比较特别的是，用binary search来解决4888个distinct rank，据作者自己声称，已被一个加拿大poker academy购买，速度是之前的3倍，并且击败了所有他知道的算法。

### 性能和直观简洁并存的实现
#### 每张牌用prime来表示

Cactus Kev的最给人启发的一点就是把每张牌面用一个唯一的素数prime来表示，这样根据乘积就可以在一个lookup table中迅速判断是否对子、3条甚至full house。
这里需要对癞子进行扩展，癞子用43来代表，建立这样的lookup table，这里唯一需要担心的问题是，如果最大的乘积43*41*41*41*37*37*37=150115382759，是无符号32位整数上限4294967296的34倍，因此提供的代码适用于64位CPU。（32位没有测试，不过貌似也没有什么问题。）


#### 7张牌直接建立lookup table，不需要7选5的21次循环
```
void buildScoreMap()
{
    if (scoreMap.size() > 0) return;
    buildFullHousePlus(scoreMap);
    buildStraight(scoreMap);
    buildThreeOfAKind();
    buildTwoPair();
    buildOnePair();
    buildHighCard(scoreMap, high_card_max);

    buildFlushMap();
}

```

#### Flush的处理
没有用位运算，而是直接用计数器，并且在扫描的过程中一次性把手牌的lookup key、flush的lookup key都找出来。 </br>
以下是scanHandString2的实现：
```
	if (club >= 5) {
		suit = Clubs;
		suitKey = clubKey;
		if (hasGhost) {
			suitKey *= 43;
		}
	} else if (diamond >= 5) {
		suit = Diamonds;
		suitKey = diamondKey;
		if (hasGhost) {
			suitKey *= 43;
		}
	} 
    ...

```


#### 癞子带来的挑战一
如果没有癞子，getScore的实现大概是：（前文两个算法基本也都是先判断是否flush）
```
首先判断是否flush
如果是，则只针对flush cards查表，查出来high card的score，减去high card的偏移，即为相对最大的flush rank。
若否，则对所有手牌的primes乘积查表。
```

问题在于引入了癞子后，对于这样的手牌："Xn9cKd3d2d9d9s"，既是4条又是同花的情况，上述算法需要调整下先后顺序：</br>
```
首先则对所有手牌的primes乘积查表，判断是否比顺子大，
如果是，直接返回。
若否，则再判断一次是否flush，并且针对flush cards查表（其中可能包括癞子的prime即43），若还查出来是straight，则为straight flush，否则查出来high card的score，减去high card的偏移，即为相对最大的flush rank。
```

#### 癞子带来的挑战二
上文中，如果考虑这手牌 "XnKdQd2d9d4d5d"，直接在之前建的lookup table中，只会找到"KKQ95"的rank，而我们需要的是"AKQ95"，解决的思路可以是在扫描过程中，直接把flush key的癞子key替换成一个最大的牌面，本例中即为Ace。那导致getScore的流程会修改为：
```
首先则对所有手牌的primes乘积查表，判断是否比顺子大，
如果是，直接返回。
若否，则再判断一次是否flush，并且针对flush cards查表（其中可能包括癞子的prime即43），若还查出来是straight，则为straight flush，否则再对替换后的prime key查一次表，查出来high card的score，减去high card的偏移，即为相对最大的flush rank。
需要多一次查表过程。
```

#### 优化：针对flush 另外建立flushScoreMap
如果对flush hand建立一个独立的表，用于判断至少是同花的牌型，前文中的primes乘积依然可以作为判断是否straight, fullhouse, four of a kind等的手段。
Lookup即根据手牌的primes乘积作为Key，取到rank score，其实非常简洁：

```
int getScore(const std::string& hs)
{
    int k=1, suitKey=1, keyWithoutXn=1;
    LYSuit suit = Nosuit;

    LYCardHelpers::scanHandString2(hs, k, suitKey, suit);
    int score = scoreMap[k];
    if (score < straight_max || hs.size() < 10 || suit == Nosuit) {
        return score;
    }
    if (suit != Nosuit) {
        return flushScoreMap[suitKey];
    }
    return score;
}
```
### 性能测试
在MacBook Pro (15-inch, 2017)，2.9 GHz Intel Core i7，7张带癞子的可以在9ms里处理完2万手。（单线程）</br>
```
score lookup table的size：len of scoreTable 102964
flush score lookup table的size：len of flushScoreTable 42324
其内存共占用在800KB(64bits)左右。
```

### 安装使用（c++ 实现）
#### pre-require
1. boost, need to be instal manually
2. cmake, need to be installed manually
3. gtest (will be installed automatically)

#### git clone, build, test and run
$ git clone git@github.com:caiqingfeng/pokerevaluator 
$ cd primev3-cpp-clean
$ mkdir release
$ cd release 
$ cmake -DCMAKE_BUILD_TYPE=Release ..
$ make all test
$ cd .. && release/pokercpp

### API
Just Call the API getScore(handString);

### 安装使用（golang 实现）
golang的单线程处理2万手大概是20ms，在查表过程没有优化单独建立flushScoreMap。

```
$ cd $GOPATH/src/github.com && mkdir -p caiqingfeng && cd caiqingfeng
$ git clone git@github.com:caiqingfeng/pokerevaluator
$ cd $GOPATH/src/github.com/caiqingfeng/pokerevaluator/primev2
$ go run poker.go
```
