#include <map>
#include <gtest/gtest.h>

#include "../libpoker/LYScoreMap.h"

class LYScoreMap_tests : public ::testing::Test
{
protected:

    void SetUp()
    {
    }
    void TearDown()
    {
   }
};

TEST_F(LYScoreMap_tests, buildFourOfAKind)
{
    buildFourOfAKind(scoreMap);
	ASSERT_NE(scoreMap.size(), 0);
    ASSERT_EQ(scoreMap[41*41*41*41*37], four_of_akind_max);
    ASSERT_EQ(scoreMap[2*2*2*2*3], four_of_akind_max+155);

    ASSERT_EQ(scoreMap[41*41*41*41*43], four_of_akind_max);
    ASSERT_EQ(scoreMap[2*2*2*2*43], four_of_akind_max+144);

    ASSERT_EQ(scoreMap[41*41*41*41*37*2], four_of_akind_max);
    ASSERT_EQ(scoreMap[2*2*2*2*31*41], four_of_akind_max+144);

    ASSERT_EQ(scoreMap[41*41*41*41*37*43], four_of_akind_max);
    ASSERT_EQ(scoreMap[2*2*2*2*31*41*43], four_of_akind_max+144);
}

TEST_F(LYScoreMap_tests, test3_Xn)
{
    buildFourOfAKind(scoreMap);
    ASSERT_EQ(scoreMap[41*41*41*43*37], four_of_akind_max);
    ASSERT_EQ(scoreMap[2*2*2*43*3], four_of_akind_max+155);

    ASSERT_EQ(scoreMap[41*41*41*43*37*37], four_of_akind_max);
    ASSERT_EQ(scoreMap[2*2*2*43*31*41], four_of_akind_max+144);

    ASSERT_EQ(scoreMap[2*2*2*43*41*41*41], four_of_akind_max+11);
}

TEST_F(LYScoreMap_tests, buildFullHouse)
{
    buildFullHouse(scoreMap);
    ASSERT_EQ(scoreMap[41*41*41*37*37], full_house_max);
    ASSERT_EQ(scoreMap[2*2*2*3*3], full_house_max+155);

    ASSERT_EQ(scoreMap[41*41*43*37*37], full_house_max);
    ASSERT_EQ(scoreMap[2*2*43*3*3], full_house_max+144-1);
}

TEST_F(LYScoreMap_tests, buildStraight)
{
    buildStraight();
    ASSERT_EQ(scoreMap[41*37*31*29*23], straight_max);
    ASSERT_EQ(scoreMap[41*2*3*5*7], straight_max+9);

    ASSERT_EQ(scoreMap[41*37*31*29*23*43], straight_max);
    ASSERT_EQ(scoreMap[41*2*3*5*7*43], straight_max+8);

    ASSERT_EQ(scoreMap[41*37*31*29*2*43*3], straight_max);
    ASSERT_EQ(scoreMap[41*2*3*5*7*37*43], straight_max+8);
}

TEST_F(LYScoreMap_tests, buildThree)
{
    buildThreeOfAKind();
    ASSERT_EQ(scoreMap[41*41*41*37*31], three_of_akind_max);
    ASSERT_EQ(scoreMap[2*2*2*5*3], three_of_akind_max+857);

    ASSERT_EQ(scoreMap[41*41*43*37*31], three_of_akind_max);
    ASSERT_EQ(scoreMap[2*2*43*5*3], three_of_akind_max+857);

    ASSERT_EQ(scoreMap[41*41*41*37*31*2*3], three_of_akind_max);
    ASSERT_EQ(scoreMap[2*2*2*5*3*41*37], three_of_akind_max+858-66);

    ASSERT_EQ(scoreMap[41*41*43*37*31*2*3], three_of_akind_max);
    ASSERT_EQ(scoreMap[2*2*43*41*37], three_of_akind_max+858-12*11/2);
    ASSERT_EQ(scoreMap[2*2*43*41*37*31], three_of_akind_max+858-12*11/2);
    ASSERT_EQ(scoreMap[2*2*43*5*41*37], three_of_akind_max+858-12*11/2);
    ASSERT_EQ(scoreMap[2*2*43*5*41*37*31], three_of_akind_max+858-12*11/2);
    ASSERT_EQ(scoreMap[2*2*43*5*41*37*11], three_of_akind_max+858-12*11/2);
}

TEST_F(LYScoreMap_tests, buildTwo)
{
    buildTwoPair();
    ASSERT_EQ(scoreMap[41*41*37*37*31], two_pair_max);
    ASSERT_EQ(scoreMap[2*2*3*5*3], two_pair_max+857);

    ASSERT_EQ(scoreMap[41*41*37*37*31*2*3], two_pair_max);
    ASSERT_EQ(scoreMap[2*2*3*5*3*41*37], two_pair_max+858-11);
}

TEST_F(LYScoreMap_tests, buildOne)
{
    buildOnePair();
    ASSERT_EQ(scoreMap[41*41*37*31*29], one_pair_max);
    // ASSERT_EQ(scoreMap[2*2*3], one_pair_max+2859);
    ASSERT_EQ(scoreMap[2*2*3*5*7], one_pair_max+2859);

    ASSERT_EQ(scoreMap[41*43*37*31*19], one_pair_max+2);//AXKQ9
    ASSERT_EQ(scoreMap[2*43*3*5*13], one_pair_max+12*11*10/6*8-1);//7X234
}

//int primes[] = {2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43};
TEST_F(LYScoreMap_tests, buildHigh)
{
    scoreMap.clear();
    buildScoreMap();
    ASSERT_EQ(scoreMap[41*19*37*31*29], high_card_max);
    // ASSERT_EQ(scoreMap[2*2*3], one_pair_max+2859);
    // ASSERT_EQ(scoreMap[2*13*3*5*7], high_card_max+1277-1);
    //AK765 23
    ASSERT_GT(scoreMap[41*37*13*7*11], high_card_max); //AK765
    ASSERT_GT(scoreMap[41*37*13*7*11*5], high_card_max); //AK7654
    ASSERT_GT(scoreMap[41*37*13*7*11*3], high_card_max); //AK7653
    ASSERT_GT(scoreMap[41*37*13*7*11*2], high_card_max); //AK7652
    ASSERT_GT(scoreMap[41*37*13*11*7*3*2], high_card_max); //AK7654
    ASSERT_GT(scoreMap[41*37*13*7*11*2], high_card_max);
}

TEST_F(LYScoreMap_tests, getComb)
{
    std::vector<std::vector<int>> comb;
    getComb(comb, 1);
    ASSERT_EQ(comb.size(), 13); 
}

TEST_F(LYScoreMap_tests, buildFlushMap)
{
    scoreMap.clear();
    flushScoreMap.clear();
    buildScoreMap();
    ASSERT_GT(flushScoreMap.size(), 0);
    ASSERT_EQ(flushScoreMap[41*37*31*29*19], flush_max); //AsKsQsJs9s
    ASSERT_EQ(flushScoreMap[41*43*31*29*19], flush_max); //AsXnQsJs9s
}