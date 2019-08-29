#include <map>
#include <iostream>

std::map<int, std::map<int, int>> fourMap;
std::map<int, std::map<int, std::map<int, int>>> threeMap;
std::map<int, std::map<int, std::map<int, int>>> twoMap;
std::map<int, std::map<int, std::map<int, std::map<int, int>>>> oneMap;

extern int primes[];

void initFourMap()
{
    if (fourMap.size() > 0) return;
    int offset = 0;
    for (int i=12; i>=0; i--) {
        std::map<int, int> m;
        for (int j=12; j>=0; j--) {
            if (j == i) continue;
            m[primes[j]] = offset++;
        }
        fourMap[primes[i]] = m;
    }
}

int offsetFour(int p1, int p2)
{
    return fourMap[p1][p2];
}

//
void initThreeMap()
{
    if (threeMap.size() > 0) return;
    int offset = 0;
    for (int i=12; i>=0; i--) {
        std::map<int, std::map<int, int>> m1;
        for (int j=12; j>=0; j--) {
            if (j == i) continue;
            std::map<int, int> m2;
            m2[0] = offset;
            for (int x=j-1; x>=0; x--) {
                if (x == i) continue;
                m2[primes[x]] = offset++;
            }
            m1[primes[j]] = m2;
            m1[0] = m2;
        }
        threeMap[primes[i]] = m1;
    }
}

void initTwoMap()
{
    if (twoMap.size() > 0) return;
    int offset = 0;
    for (int i=12; i>=0; i--) {
        std::map<int, std::map<int, int>> m1;
        for (int j=i-1; j>=0; j--) {
            std::map<int, int> m2;
            m2[0] = offset;
            for (int x=12; x>=0; x--) {
                if (x == i || x == j) continue;
                m2[primes[x]] = offset++;
            }
            m1[primes[j]] = m2;
            m1[0] = m2;
        }
        twoMap[primes[i]] = m1;
    }
}

void initOneMap()
{
    if (oneMap.size() > 0) return;
	//第一种情况，1对+任意5张散牌，没有癞子
    int offset = 0;
    for (int i=12; i>=0; i--) {
        std::map<int, std::map<int, std::map<int, int>>> m1;
        for (int j=12; j>=0; j--) {
            if (j == i) continue;
            std::map<int, std::map<int, int>> m2;
            for (int x=j-1; x>=0; x--) {
                if (x == i) continue;
                std::map<int, int> m3;
                for (int y=x-1; y>=0; y--) {
                    if (y == i) continue;
                    m3[primes[y]] = offset++;
                }
                m2[primes[x]] = m3;
                m2[0] = m3;
            }
            m1[primes[j]] = m2;
            m1[0] = m2;
        }
        oneMap[primes[i]] = m1;
    }
}

int offsetThree(int p1, int p2, int p3)
{
    return threeMap[p1][p2][p3];
}

int offsetThree(int p1, int p2)
{
    return threeMap[p1][p2][0];
}

int offsetTwo(int p1, int p2, int p3)
{
    return twoMap[p1][p2][p3];
}

int offsetTwo(int p1, int p2)
{
    return twoMap[p1][p2][0];
}

int offsetOne(int p1, int p2, int p3, int p4)
{
    return oneMap[p1][p2][p3][p4];
}