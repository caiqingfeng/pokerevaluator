#ifndef _LY_SCOREMAP_H
#define _LY_SCOREMAP_H
#include "LYPokerConstants.h"
#include <map>
#include <vector>

int getScore(const std::string& hs);

void buildScoreMap();
void buildFullHousePlus(std::map<int, int>&sMap);
void buildFourOfAKind(std::map<int, int>&sMap);
void buildFullHouse(std::map<int, int>&sMap);
void buildStraight();
void buildStraight(std::map<int, int>&sMap, const int base_offset=straight_max);
void buildThreeOfAKind();
void buildTwoPair();
void buildOnePair();
void buildHighCard(std::map<int, int>& sMap, const int base_offset);
void buildFlushMap();

void getComb(std::vector<std::vector<int>>& set, const int level);

extern std::map<int, int> scoreMap;
extern std::map<int, int> flushScoreMap; 

#endif //_LY_SCOREMAP_H