//
//  LYCommonDefine.h
//  iBuddyHoldem
//
//  Created by 蔡 庆丰 on 13-2-28.
//  Copyright (c) 2013年 蔡 庆丰. All rights reserved.
//

#ifndef LYPOKER_CONSTANTS_H_
#define LYPOKER_CONSTANTS_H_

#include <string>

const std::string LUYUN_HOUSE = "luyun_house"; //table's owner default is house但如果支持用户占桌子成为桌主的话，会改变

enum LYFace
{
    NOFACE = 0,
    ACE = 14,
    TWO = 2,
    THREE = 3,
    FOUR = 4,
    FIVE = 5,
    SIX = 6,
    SEVEN = 7,
    EIGHT = 8,
    NINE = 9,
    TEN = 10,
    JACK = 11,
    QUEEN = 12,
    KING = 13,
    SMALL_GHOST = 24,
    BIG_GHOST = 25
};

enum LYSuit
{
    Nosuit = 0, Clubs = 1, Diamonds = 2, Hearts = 3, Spades = 4
};

const int straight_flush_max = 0 ;
const int four_of_akind_max = 10; 
const int full_house_max = 10+156; //166
const int flush_max = 10+156+156;  //322
const int straight_max = flush_max+1277; //1599
const int three_of_akind_max  = straight_max+10; //1609
const int two_pair_max = three_of_akind_max+858; //2467
const int one_pair_max = two_pair_max+858; //3325
const int high_card_max = one_pair_max+2860; //6185

#endif
