#include <gtest/gtest.h>

#include "../libpoker/LYCardHelpers.h"
#include <boost/lexical_cast.hpp>

class LYCardHelpers_tests : public ::testing::Test
{
protected:

    void SetUp()
    {
    }
    void TearDown()
    {
   }

};

//int primes[] = {2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43};
TEST_F(LYCardHelpers_tests, scanHandString2)
{
	int key, suitKey;
	LYSuit suit;
	LYCardHelpers::scanHandString2("AsXnQsJs9s", key, suitKey, suit);
	ASSERT_EQ(suit, Spades);
	ASSERT_EQ(suitKey, 41*43*31*29*19);

	LYCardHelpers::scanHandString2("5c2d5dXnKd7d6d", key, suitKey, suit);
	ASSERT_EQ(suit, Diamonds);
	ASSERT_EQ(suitKey, 43*37*13*11*7*2);
}