#include <map>
#include <vector>
#include <iostream>

#include "LYOffset.h"
#include "LYScoreMap.h"
#include "LYCardHelpers.h"

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

std::map<int, int> scoreMap;
std::map<int, int> flushScoreMap; //flush 要单独建立一个map，例如AsXnQsJs9s,
                            //如果只是对face进行建表（scoreMap干的）,就无法正确的得到AsKsQsJs9s
                            //会跟AAQJ9冲突，所以另外建立一张表，没有癞子其实也不需要
int primes[] = {2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43};

int getScore(const std::string& hs)
{
    if (scoreMap.size() == 0) {
        buildScoreMap();
    }
    
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

void addKey(std::map<int, int>& sMap, int k, int v)
{
    if (sMap.find(k) == sMap.end()) {
        sMap[k] = v;
    }
}

void addKey(int k, int v)
{
    if (scoreMap.find(k) == scoreMap.end()) {
        scoreMap[k] = v;
    }
}

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

void buildFullHousePlus(std::map<int, int>&sMap)
{
    buildFourOfAKind(sMap);
    buildFullHouse(sMap);
}

void buildFourOfAKind(std::map<int, int>&sMap)
{
    initFourMap();
    //从A开始，A是primes[12]=41
    //首先计算4+1，4+2，4+3的各种情况
    for (int i=12; i>=0; i--) {
        int base = primes[i]*primes[i]*primes[i]*primes[i];
        //首先计算4+Xn，这里把Xn替换成最大的那个踢脚算其offset
        int key2 = base * 43; 
        int fix_offset = (12-i)*12;
        addKey(sMap, key2, fix_offset+four_of_akind_max); //如果第一次已经加了(k,v)，不会覆盖
        for (int j=12; j>=0; j--) {
            if (j == i) continue;
            //计算4+1
            int key = base * primes[j];
            int offset = offsetFour(primes[i], primes[j]);
            addKey(sMap, key, offset+four_of_akind_max);
            //然后计算6张牌的情况：4+1+1
            for (int x=j; x>=0; x--) {
                //x 不能等于i
                if (x == i) continue;
                int key3 =  base * primes[j] * primes[x];
                addKey(sMap, key3, offset+four_of_akind_max); 
                //然后计算7张牌 4+1+1+1
                for (int y=j; y>=0; y--) {
                    //y 不能等于i
                    if (y == i) continue;
                    int key5 = base * primes[j] * primes[x] * primes[y];
                    addKey(sMap, key5, offset+four_of_akind_max); 
                }
                //然后计算7张牌 4+1+1+Xn
                int key6 = base * primes[j] * primes[x] * 43;
                addKey(sMap, key6, fix_offset+four_of_akind_max);
            }
            //然后计算4+1+Xn
            int key4 = base * primes[j] * 43;
            addKey(sMap, key4, fix_offset+four_of_akind_max);
        }
    }
    //计算3+Xn+1，3+Xn+2，3+Xn+3的各种情况
    //这里Xn必须要替换成3条中的一个
    for (int i=12; i>=0; i--) {
        int base = primes[i]*primes[i]*primes[i]*43;
        for (int j=12; j>=0; j--) {
            if (j == i) continue;
            //计算4+1
            int key = base * primes[j];
            int offset = offsetFour(primes[i], primes[j]);
            addKey(sMap, key, offset+four_of_akind_max);
            //然后计算6张牌的情况：4+1+1
            for (int x=j; x>=0; x--) {
                //x 不能等于i
                if (x == i) continue;
                int key3 =  base * primes[j] * primes[x];
                addKey(sMap, key3, offset+four_of_akind_max); 
                //然后计算7张牌 4+1+1+1
                for (int y=j; y>=0; y--) {
                    //y 不能等于i
                    if (y == i) continue;
                    int key5 = base * primes[j] * primes[x] * primes[y];
                    addKey(sMap, key5, offset+four_of_akind_max); 
                }
            }
        }
    }
}

void buildFullHouse(std::map<int, int>&sMap)
{
    initThreeMap();
	//第一种情况：3条+1对+任意2张牌（不可能有癞子，否则就是4条了）
    //从A开始，A是primes[12]=41
    for (int i=12; i>=0; i--) {
        int base = primes[i]*primes[i]*primes[i];
        for (int j=12; j>=0; j--) {
            if (j == i) continue;
            //计算3+2
            int key = base * primes[j]*primes[j];
            int offset = offsetFour(primes[i], primes[j]);
            addKey(sMap, key, offset+full_house_max);
            //然后计算6张牌的情况：3+2+1
            for (int x=12; x>=0; x--) {
                //x 不能等于i
                if (x == i) continue;
                int key3 =  base * primes[j] * primes[j] * primes[x];
                addKey(sMap, key3, offset+full_house_max); 
                //然后计算7张牌 3+2+1+1
                for (int y=12; y>=0; y--) {
                    //y 不能等于i
                    if (y == i) continue;
                    int key5 = base * primes[j] * primes[j] * primes[x] * primes[y];
                    addKey(sMap, key5, offset+full_house_max); 
                }
            }
        }
    }
	//第二种情况：2对+1个癞子+任意2张牌（不能有3条，可以是3对）
    for (int i=12; i>=0; i--) {
        int base = primes[i]*primes[i]*43;
        for (int j=i-1; j>=0; j--) {
            //计算2+2+Xn
            int key = base * primes[j] * primes[j];
            int offset = offsetFour(primes[i], primes[j]);
            addKey(sMap, key, offset+full_house_max);
            //然后计算6张牌的情况：2+2+Xn+1
            for (int x=12; x>=0; x--) {
                //x 不能等于i or j
                if (x == i || x == j) continue;
                int key3 =  base * primes[j] * primes[j] * primes[x];
                addKey(sMap, key3, offset+full_house_max); 
                //然后计算7张牌 2+2+Xn+1+1
                for (int y=12; y>=0; y--) {
                    //y 不能等于i or j
                    if (y == i || y == j) continue;
                    int key5 = base * primes[j] * primes[j] * primes[x] * primes[y];
                    addKey(sMap, key5, offset+full_house_max); 
                }
            }
        }
    }
}

void buildStraight()
{
    buildStraight(scoreMap);
}

//这个算法非常巧妙
//
void buildStraight(std::map<int, int>& sMap, const int base_offset)
{
    std::vector<std::vector<int>> straight;
    for (int i=12; i>=4; i--) {
        std::vector<int> s;
        s.push_back(primes[i]);
        s.push_back(primes[i-1]);
        s.push_back(primes[i-2]);
        s.push_back(primes[i-3]);
        s.push_back(primes[i-4]);
        s.push_back(43);
        straight.push_back(s);
    }
    //对12345的顺子做类似操作
    std::vector<int> s;
    s.push_back(primes[3]);
    s.push_back(primes[2]);
    s.push_back(primes[1]);
    s.push_back(primes[0]);
    s.push_back(41);
    s.push_back(43);
    straight.push_back(s);

    std::vector<std::vector<int>>::iterator it = straight.begin();
    for (int offset=0; it!=straight.end(); it++) {
        std::vector<int> s = *it;
        std::vector<int>::iterator it2 = s.begin();
        int t = s[0]*s[1]*s[2]*s[3]*s[4]*s[5];
        int score = base_offset + (offset++);
        for(; it2!=s.end(); it2++) {
            int key = t/(*it2);
            // std::cout<<"key = " << key << " score =" << score <<std::endl;
            addKey(sMap, key, score); //5张牌
            for (int x=12; x>=0; x--) {
                //added 2019.08.04 支持按6张牌来找
                int key2 = key * primes[x];
                addKey(sMap, key2, score);
                for (int y=12; y>=0; y--) {
                    int key1 = key * primes[x] * primes[y];
                    addKey(sMap, key1, score); //7张牌
                }
            }
        }
    }
}

void buildThreeOfAKind()
{
	//第一种情况：3条+任意4张散牌（不可能有癞子，和其它对子，否则就是4条或者full了）
    //从A开始，A是primes[12]=41
    for (int i=12; i>=0; i--) {
        int base = primes[i]*primes[i]*primes[i];
        //2019-08-25支持3张牌（菠萝头道）
        int offset1 = offsetThree(primes[i], 0);
        addKey(base, offset1);
        for (int j=12; j>=0; j--) {
            if (j == i) continue;
            //计算3+1
            int key = base * primes[j];
            int offset = offsetThree(primes[i], primes[j], 0);
            addKey(key, offset+three_of_akind_max);
            //然后计算5张牌的情况：3+1+1
            for (int x=j-1; x>=0; x--) {
                //x 不能等于i
                if (x == i) continue;
                int key3 =  base * primes[j] * primes[x];
                int fix_offset = offsetThree(primes[i], primes[j], primes[x]);
                addKey(key3, fix_offset+three_of_akind_max); 
                //然后计算6张牌 3+1+1+1
                for (int y=x-1; y>=0; y--) {
                    //y 不能等于i
                    if (y == i) continue;
                    int key5 = base * primes[j] * primes[x] * primes[y];
                    addKey(key5, fix_offset+three_of_akind_max);
                    //7张牌的情况
                    for (int z=y-1; z>=0; z--) {
                        if (z == i) continue;
                        int key6 = base * primes[j] * primes[x] * primes[y] * primes[z];
                        addKey(key6, fix_offset+three_of_akind_max);
                    } 
                }
            }
        }
    }
	//第二种情况：1对+1个癞子+任意4张牌（不能有3条、2对）
    for (int i=12; i>=0; i--) {
        int base = primes[i]*primes[i]*43;
        int offset1 = offsetThree(primes[i], 0);
        //2019-08-25 支持头道3张牌，1对+1个癞子
        addKey(base, offset1+three_of_akind_max);
        for (int j=12; j>=0; j--) {
            if (j == i) continue;
            //计算2+Xn+1
            int key = base * primes[j];
            int offset = offsetThree(primes[i], primes[j]);
            addKey(key, offset+three_of_akind_max);
            //然后计算5张牌的情况：2+Xn+1+1
            for (int x=j-1; x>=0; x--) {
                //x 不能等于i or j
                if (x == i) continue;
                int key3 =  base * primes[j] * primes[x];
                int fix_offset = offsetThree(primes[i], primes[j], primes[x]);
                addKey(key3, fix_offset+three_of_akind_max); 
                //然后计算6张牌 2+Xn+1+1+1
                for (int y=x-1; y>=0; y--) {
                    //y 不能等于i or j
                    if (y == i) continue;
                    int key5 = base * primes[j] * primes[x] * primes[y];
                    addKey(key5, fix_offset+three_of_akind_max); 
                    //计算7张牌的情况
                    for (int z=y-1; z>=0; z--) {
                        if (z == i) continue;
                        int key6 = base * primes[j] * primes[x] * primes[y] * primes[z];
                        addKey(key6, fix_offset+three_of_akind_max);
                    } 
                }
            }
        }
    }
}

void buildTwoPair()
{
    initTwoMap();
	//Two Pair，不可能有癞子，只要对普通2对+3建表
	//2对+任意3张散牌
    //从A开始，A是primes[12]=41
    for (int i=12; i>=1; i--) {
        int base = primes[i]*primes[i];
        for (int j=i-1; j>=0; j--) {
			//至此选出了大小两对
            int key = base * primes[j] * primes[j];
            int offset = offsetTwo(primes[i], primes[j], 0);
            addKey(key, offset+two_pair_max); //支持4张牌的情况
            //然后计算5张牌的情况：2+2+1
            for (int x=12; x>=0; x--) {
                //x 不能等于i
                if (x == i || x == j) continue;
                int key3 =  base * primes[j] * primes[j] * primes[x];
                int fix_offset = offsetTwo(primes[i], primes[j], primes[x]);
                addKey(key3, fix_offset+two_pair_max); 
                //然后计算6张牌 2+2+1+1
                for (int y=x; y>=0; y--) {
                    //y 不能等于i/j
                    if (y == i || y == j) continue;
                    int key5 = base * primes[j] * primes[j] * primes[x] * primes[y];
                    addKey(key5, fix_offset+two_pair_max);
                    //7张牌的情况
                    for (int z=y; z>=0; z--) {
                        if (z == i || z == j) continue;
                        int key6 = base * primes[j] * primes[j] * primes[x] * primes[y] * primes[z];
                        addKey(key6, fix_offset+two_pair_max);
                    } 
                }
            }
        }
    }
}

void buildOnePair()
{
    initOneMap();
	//第一种情况，1对+任意5张散牌，没有癞子
    //从A开始，A是primes[12]=41
    for (int i=12; i>=0; i--) {
        int base = primes[i]*primes[i];
        for (int j=12; j>=0; j--) {
            if (j == i) continue;
			//至此选出了1对+1
            int key = base * primes[j];
            int offset = offsetOne(primes[i], primes[j], 0, 0);
            addKey(key, offset+one_pair_max); //支持3张牌的情况
            //然后计算4张牌的情况：2+1+1
            for (int x=j-1; x>=0; x--) {
                //x 不能等于i
                if (x == i) continue;
                int key3 =  base * primes[j] * primes[x];
                int offset2 = offsetOne(primes[i], primes[j], primes[x], 0);
                addKey(key3, offset2+one_pair_max); 
                //然后计算5张牌 2+1+1+1
                for (int y=x-1; y>=0; y--) {
                    //y 不能等于i/j
                    if (y == i) continue;
                    int key5 = base * primes[j] * primes[x] * primes[y];
                    int fix_offset = offsetOne(primes[i], primes[j], primes[x], primes[y]);
                    addKey(key5, fix_offset+one_pair_max);
                    //6张牌的情况
                    for (int z=y-1; z>=0; z--) {
                        if (z == i) continue;
                        int key6 = base * primes[j] * primes[x] * primes[y] * primes[z];
                        addKey(key6, fix_offset+one_pair_max);
                        //7张牌的情况
                        for (int t=z-1; t>=0; t--) {
                            if (t == i) continue;
                            int key7 = base * primes[j] * primes[x] * primes[y] * primes[z] * primes[t];
                            addKey(key7, fix_offset+one_pair_max);
                        }
                    } 
                }
            }
        }
    }
	//第2种情况，1癞子+任意6张散牌
    //从A开始，A是primes[12]=41
    for (int i=12; i>=0; i--) {
        int base = 43*primes[i];
        for (int j=i-1; j>=0; j--) {
			//至此选出了1对+1
            int key = base * primes[j];
            int offset = offsetOne(primes[i], primes[j], 0, 0);
            addKey(key, offset+one_pair_max); //支持3张牌的情况
            //然后计算4张牌的情况：2+1+1
            for (int x=j-1; x>=0; x--) {
                //x 不能等于i
                if (x == i) continue;
                int key3 =  base * primes[j] * primes[x];
                int offset2 = offsetOne(primes[i], primes[j], primes[x], 0);
                addKey(key3, offset2+one_pair_max); 
                //然后计算5张牌 2+1+1+1
                for (int y=x-1; y>=0; y--) {
                    //y 不能等于i/j
                    if (y == i) continue;
                    int key5 = base * primes[j] * primes[x] * primes[y];
                    int fix_offset = offsetOne(primes[i], primes[j], primes[x], primes[y]);
                    addKey(key5, fix_offset+one_pair_max);
                    //6张牌的情况
                    for (int z=y-1; z>=0; z--) {
                        if (z == i) continue;
                        int key6 = base * primes[j] * primes[x] * primes[y] * primes[z];
                        addKey(key6, fix_offset+one_pair_max);
                        //7张牌的情况
                        for (int t=z-1; t>=0; t--) {
                            if (t == i) continue;
                            int key7 = base * primes[j] * primes[x] * primes[y] * primes[z] * primes[t];
                            addKey(key7, fix_offset+one_pair_max);
                        }
                    } 
                }
            }
        }
    }
}

void buildHighCard(std::map<int, int>& sMap, const int base_offset)
{
	//只有一种情况，任意5-7张散牌，没有癞子
    //modified 2019-08-26 直接对7张牌建表，避免排序
    //跟前面一样，这里调用addKey，不会覆盖之前的k,v
    //从A开始，A是primes[12]=41
    int offset = 0;
    for (int i=12; i>=4; i--) {
        int base = primes[i];
        for (int j=i-1; j>=3; j--) {
			//至此选出了1+1
            int key = base * primes[j];
            addKey(sMap, key, offset+base_offset); //支持2张牌的情况
            //然后计算3张牌的情况：1+1+1
            for (int x=j-1; x>=0; x--) {
                int key3 =  base * primes[j] * primes[x];
                addKey(sMap, key3, offset+base_offset); 
                //然后计算4张牌 1+1+1+1
                for (int y=x-1; y>=0; y--) {
                    int key5 = base * primes[j] * primes[x] * primes[y];
                    addKey(sMap, key5, offset+base_offset);
                    //5张牌的情况
                    for (int z=y-1; z>=0; z--) {
                        int key6 = base * primes[j] * primes[x] * primes[y] * primes[z];
                            // std::cout << "kkkkk:" << key6 << offset << std::endl;
                        int fix_offset =  offset;
                        if (sMap.find(key6) == sMap.end()) {
                            addKey(sMap, key6, fix_offset+base_offset);
                            offset++;
                            // std::cout << key6 << offset << std::endl;
                            // 6张牌的情况
                            for (int t=z-1; t>=0; t--) {
                                int key7 = key6 * primes[t];
                                addKey(sMap, key7, fix_offset+base_offset);
                                //7张牌的情况
                                for (int v=t-1; v>=0; v--) {
                                    int key8 = key7 * primes[v];
                                    addKey(sMap, key8, fix_offset+base_offset);
                                }
                            }
                        }
                    } 
                }
            }
        }
    }
}

// 得到二维的数组
// 递归第一步get vector<vector<12>, vector<11>, ..., vector<1>, vector<0>>
// 第二步vector<vector<12, 11>, vector<11, 10>, ..., vector<1, 0>>
void getComb(std::vector<std::vector<int>>& set, const int level)
{
    // std::cout << "set.size=" << set.size() << " level=" << level << std::endl;
    std::vector<std::vector<int>> orgSet = set;

    if (orgSet.size() == 0) {
        for (int i=12; i>=0; i--) {
            std::vector<int> s;
            s.push_back(i);
            set.push_back(s);
        }
        getComb(set, level-1);
        return;
    }

    if (level == 0) return;
    set.clear();
    std::vector<std::vector<int>>::iterator it = orgSet.begin();
    for (; it!=orgSet.end(); it++) {
        std::vector<int> n = *it;
        int last = n[n.size()-1];
        for (int i=last-1; i>=0; i--) {
            std::vector<int> x = n;
            x.push_back(i);
            set.push_back(x);
        }        
    }
    getComb(set, level-1);
}

void genGhostScore(std::vector<std::vector<int>>& comb)
{
    std::vector<std::vector<int>>::iterator it = comb.begin();
    for (; it!=comb.end(); it++) {
        std::vector<int>::iterator nit = (*it).begin();
        int key = 1;
        int maxCardIndex = 12;
        for (; nit!=(*it).end(); nit++) {
            if (maxCardIndex == *nit) maxCardIndex--;
            key *= primes[*nit];
        }
        //现在得到了任意4-5-6张牌+1个癞子的组合
        int scoreKey = key * primes[maxCardIndex];
        addKey(flushScoreMap, key*43, flushScoreMap[scoreKey]);
    }
}

void buildFlushMap()
{
	//针对任意5-7张散牌，有癞子和无赖子，一次性查表解决问题
    //modified 2019-08-26 直接对7张牌建表，避免排序和替换
    //从A开始，A是primes[12]=41
    //先把顺子、fullhouse、四条 建起来
    buildFullHousePlus(flushScoreMap);
    buildStraight(flushScoreMap, straight_flush_max);

    //然后再建立7张散牌，支持带癞子的情况
    // todo : 怎么和buildHighCard代码合并？
    // 7张无癞子的情况
    buildHighCard(flushScoreMap, flush_max);
    std::vector<std::vector<int>> comb;
    //13 选4张牌+癞子
    getComb(comb, 4);
    // std::cout << "len=" << comb.size() << std::endl;
    genGhostScore(comb);
    //13 选5张牌+癞子
    comb.clear();
    getComb(comb, 5);
    genGhostScore(comb);
    //13 选6张牌+癞子
    comb.clear();
    getComb(comb, 6);
    genGhostScore(comb);
}