#include <map>
#include <gtest/gtest.h>

#include "../libpoker/LYOffset.h"

extern std::map<int, std::map<int, int>> fourMap;
extern std::map<int, std::map<int, std::map<int, int>>> threeMap;
extern std::map<int, std::map<int, std::map<int, int>>> twoMap;
extern std::map<int, std::map<int, std::map<int, std::map<int, int>>>> oneMap;

class LYOffset_tests : public ::testing::Test
{
protected:

    void SetUp()
    {
    }
    void TearDown()
    {
   }
};

TEST_F(LYOffset_tests, initFourMap)
{
    initFourMap();
	ASSERT_EQ(fourMap[41][37], 0);
	ASSERT_EQ(fourMap[2][3], 155);
}

TEST_F(LYOffset_tests, initThreeMap)
{
    initThreeMap();
    ASSERT_NE(threeMap.size(), 0);
	ASSERT_EQ(threeMap[41][37][31], 0);
	ASSERT_EQ(threeMap[2][5][3], 857);
	ASSERT_EQ(threeMap[2][41][37], 858-66);
	ASSERT_EQ(threeMap[41][37][0], 0);
	ASSERT_EQ(threeMap[41][31][0], 11);
}

TEST_F(LYOffset_tests, initTwoMap)
{
    initTwoMap();
    ASSERT_NE(twoMap.size(), 0);
	ASSERT_EQ(twoMap[41][37][31], 0);
	ASSERT_EQ(twoMap[3][2][5], 857);
	ASSERT_EQ(twoMap[3][2][41], 858-11);
	ASSERT_EQ(twoMap[41][37][0], 0);
	ASSERT_EQ(twoMap[41][31][0], 11);
}

TEST_F(LYOffset_tests, initOneMap)
{
    initOneMap();
    ASSERT_NE(oneMap.size(), 0);
	ASSERT_EQ(oneMap[41][37][31][29], 0);
	ASSERT_EQ(oneMap[2][7][5][3], 2859);
}