//
//  LYCardHelpers.cpp
//
//  Created by 蔡 庆丰 on 14-10-25.
//  Copyright (c) 2013年 蔡 庆丰. All rights reserved.
//

#include <vector>
#include <map>
#include "LYCardHelpers.h"
#include <boost/lexical_cast.hpp>
#include <boost/algorithm/string.hpp>
#include <iostream>

std::map<char, int> primesMap = {{'X', 43}, {'A', 41}, {'K', 37}, {'Q', 31}, {'J', 29}, {'T', 23}, 
							{'9', 19}, {'8', 17}, {'7', 13}, {'6', 11}, {'5', 7}, 
							{'4', 5}, {'3', 3}, {'2', 2}};
extern int primes[13];

//没有符号分割，直接AsKsJs...
int LYCardHelpers::scanHandString2(const std::string& cs, int& key, int& suitKey, LYSuit& suit)
{
	int l = cs.size();
	key = 1;
	suitKey = 1;
	bool hasGhost = false;
	int clubKey = 1, club = 0; 
	int diamondKey = 1, diamond = 0;  
	int heartKey = 1, heart = 0;  
	int spadeKey = 1, spade = 0;  
	for (int i=0; i<l/2; i++) {
		int p = primesMap[cs[2*i]];
		key *= p;
		switch (cs[2*i+1]) {
			case 'c':
				clubKey *= p;
				club++;
				break;
			case 'd':
				diamondKey *= p;
				diamond++;
				break;
			case 'h':
				heartKey *= p;
				heart++;
				break;
			case 's':
				spadeKey *= p;
				spade++;
				break;
			case 'n':
			default:
				hasGhost = true;
				club++;
				diamond++;
				heart++;
				spade++;
				break;
		}
	}
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
	} else if (heart >= 5) {
		suit = Hearts;
		suitKey = heartKey;
		if (hasGhost) {
			suitKey *= 43;
		}
	} else if (spade >= 5) {
		suit = Spades;
		suitKey = spadeKey;
		if (hasGhost) {
			suitKey *= 43;
		}
	}
	return key;
}
