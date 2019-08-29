//
//  LYCardHelpers.h
//
//  Created by 蔡 庆丰 on 14-10-25.
//  Copyright (c) 2013年 蔡 庆丰. All rights reserved.
//
#ifndef _LY_CARD_HELPERS_H
#define _LY_CARD_HELPERS_H

#include <vector>
#include <string>
#include "LYPokerConstants.h"

class LYCardHelpers {
public:
	//2019-08-26 added for better performance
	static int scanHandString2(const std::string& cs, int& key, int& suitKey, LYSuit& suit);
};

#endif
