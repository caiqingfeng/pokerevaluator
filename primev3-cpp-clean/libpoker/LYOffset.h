#ifndef _LY_OFFSET_H
#define _LY_OFFSET_H

void initFourMap();
void initThreeMap();
void initTwoMap();
void initOneMap();

int offsetFour(int p1, int p2);
int offsetThree(int p1, int p2);
int offsetThree(int p1, int p2, int p3);
int offsetTwo(int p1, int p2, int p3);
int offsetTwo(int p1, int p2);
int offsetOne(int p1, int p2, int p3, int p4);

#endif //_LY_OFFSET_H