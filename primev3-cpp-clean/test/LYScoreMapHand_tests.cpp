#include <map>
#include <gtest/gtest.h>

#include "../libpoker/LYScoreMap.h"

class LYScoreMapHand_tests : public ::testing::Test
{
protected:

    void SetUp()
    {
    }
    void TearDown()
    {
   }
};

TEST_F(LYScoreMapHand_tests, getScoreThreeCards)
{
    //todo
    // buildScoreMap();
}

TEST_F(LYScoreMapHand_tests, getScoreFourCards)
{
    //todo
    // buildScoreMap();
}

TEST_F(LYScoreMapHand_tests, getScoreSixCards)
{
    //todo
    // buildScoreMap();
}

TEST_F(LYScoreMapHand_tests, getScoreFiveCards)
{
    buildScoreMap();
	ASSERT_EQ(getScore("AsKsQsJsTs"), straight_flush_max);
	ASSERT_EQ(getScore("AsAcAhAdKs"), four_of_akind_max);
	ASSERT_EQ(getScore("AsAcAhKcKs"), full_house_max);
	ASSERT_EQ(getScore("AsKsQsJs9s"), flush_max);
	ASSERT_EQ(getScore("AsKcQsJsTs"), straight_max);
	ASSERT_EQ(getScore("AsAcAhQdKs"), three_of_akind_max);
	ASSERT_EQ(getScore("AsAcKhQdKs"), two_pair_max);
	ASSERT_EQ(getScore("AsAcQsKsJs"), one_pair_max);
	ASSERT_EQ(getScore("AsKsQcJc9c"), high_card_max);
}

TEST_F(LYScoreMapHand_tests, getScoreFiveWithGhost)
{
    buildScoreMap();
	ASSERT_EQ(getScore("AsXnQsJsTs"), straight_flush_max);
	ASSERT_EQ(getScore("AsAcAhXnKs"), four_of_akind_max);
	ASSERT_EQ(getScore("AsAcXnKcKs"), full_house_max);
	ASSERT_EQ(getScore("AsXnQsJs9s"), flush_max);
	ASSERT_EQ(getScore("AsKcXnJsTs"), straight_max);
	ASSERT_EQ(getScore("AsAcXnQdKs"), three_of_akind_max);
	ASSERT_EQ(getScore("AsXnQsKs9c"), one_pair_max+2);
	ASSERT_EQ(getScore("AsXnQsKs9s"), flush_max);
}

TEST_F(LYScoreMapHand_tests, getScoreSevenCards)
{
    buildScoreMap();
	ASSERT_EQ(getScore("AsKsQsJsTs2s3s"), straight_flush_max);
	ASSERT_EQ(getScore("AsAcAhAdKs2c3c"), four_of_akind_max);
	ASSERT_EQ(getScore("AsAcAhKcKsQsQh"), full_house_max);
	ASSERT_EQ(getScore("AsKsQsJs9s2c3c"), flush_max);
	ASSERT_EQ(getScore("AsKcQsJsTs2c3c"), straight_max);
	ASSERT_EQ(getScore("AsAcAhQdKs2c3c"), three_of_akind_max);
	ASSERT_EQ(getScore("AsAcKhQdKs2c3c"), two_pair_max);
	ASSERT_EQ(getScore("AsAcQsKsJs2c3c"), one_pair_max);
	ASSERT_EQ(getScore("AsKsQcJc9c5s4s"), high_card_max);
}

TEST_F(LYScoreMapHand_tests, getScoreSevenWithGhost)
{
    buildScoreMap();
	ASSERT_EQ(getScore("AsXnQsJsTs2s3s"), straight_flush_max);
	ASSERT_EQ(getScore("AsXnAhAdKs2c3c"), four_of_akind_max);
	ASSERT_EQ(getScore("AsXnAhKcKsQsQh"), full_house_max);
	ASSERT_EQ(getScore("AsXnQsJs9s2c3c"), flush_max);
	ASSERT_EQ(getScore("AsKcXnJsTs2c3c"), straight_max);
	ASSERT_EQ(getScore("AsXnAhQdKs2c3c"), three_of_akind_max);
	ASSERT_EQ(getScore("AsXnQsKs9c2c3c"), one_pair_max+2);
}

TEST_F(LYScoreMapHand_tests, fromMatches)
{
    buildScoreMap();
	ASSERT_GT(getScore("5dXnKd7d6d6sQc"), flush_max);
	// std::cout << "kkkkk" << std::endl;
	ASSERT_GT(getScore("5cAd2d5dKd7d6d"), flush_max);
	ASSERT_GT(getScore("5c2d5dXnKd7d6d"), flush_max);
	ASSERT_LT(getScore("Xn9cKd3d2d9d9s"), four_of_akind_max+155);
	ASSERT_LT(getScore("5hTdXn9cKd3d2d"), flush_max+1276);
	ASSERT_LT(getScore("As2h7sXn5hTs6s"), flush_max+1276);
	ASSERT_LT(getScore("6h3hAs2h7sXn5h"), straight_flush_max+10);
}