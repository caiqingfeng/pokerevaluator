#include "../libpoker/LYScoreMap.h"
#include "../libpoker/LYCardHelpers.h"

#include <chrono>
#include <iostream>
#include <boost/property_tree/ptree.hpp>
#include <boost/property_tree/json_parser.hpp>

// Short alias for this namespace
namespace pt = boost::property_tree;

using namespace std::chrono;

class LYMatch 
{
public:
    std::string alice;
    int aliceScore;
    std::string bob;
    int bobScore;
    int result;
};

int main() {
    buildScoreMap();
    std::cout << "len of scoreTable " << scoreMap.size() << std::endl;
    std::cout << "len of flushScoreTable " << flushScoreMap.size() << std::endl;
    // Create a root
    pt::ptree root;
    std::vector<LYMatch*> matches;

    // Load the json file in this ptree
    pt::read_json("../alice_vs_bob/seven_cards_with_ghost.json", root);
    // pt::read_json("../alice_vs_bob/match.json", root);
    // Iterator over all matches
    for (pt::ptree::value_type &match : root.get_child("matches")) {
        LYMatch *m = new LYMatch();
        m->alice = match.second.get_child("alice").data();
        m->bob = match.second.get_child("bob").data();
        matches.push_back(m);
        // matches[count++] = m;
    }
    // std::cout << "has " << matches.size() << " matches" << std::endl;

    milliseconds started_at = duration_cast< milliseconds >(
        system_clock::now().time_since_epoch()
    );

    for (int i=0; i<matches.size(); i++) {
        LYMatch* m = matches[i];
        // std::cout << "Match: " << m->alice << " vs " << m->bob << std::endl;
        m->aliceScore = getScore(m->alice);
        m->bobScore = getScore(m->bob);
        if (m->aliceScore > m->bobScore) {
            m->result = 2;
        } else if (m->aliceScore < m->bobScore) {
            m->result = 1;
        } else {
            m->result = 0;
        }
    }

    milliseconds ended_at = duration_cast< milliseconds >(
        system_clock::now().time_since_epoch()
    );
    std::cout << "counting " << matches.size() << " hands, spent " << ended_at.count()-started_at.count() << " milliseconds" << std::endl;

    pt::ptree root2;
    std::vector<LYMatch*> resultMatches;
    pt::read_json("../alice_vs_bob/seven_cards_with_ghost.result.json", root2);
    // pt::read_json("../alice_vs_bob/result.json", root2);
    // Iterator over all matches
    for (pt::ptree::value_type &match : root2.get_child("matches")) {
        LYMatch *m = new LYMatch();
        m->alice = match.second.get_child("alice").data();
        m->bob = match.second.get_child("bob").data();
        m->result = match.second.get<int>("result");
        resultMatches.push_back(m);
    }
    for (int i=0; i<matches.size(); i++) {
        if (matches[i]->result == resultMatches[i]->result) continue;
        std::cout << "alice:" << matches[i]->alice << std::endl;
        std::cout << "bob:" << matches[i]->bob << std::endl;
        std::cout << "result:" << matches[i]->result << std::endl;
        std::cout << "aliceScore:" << matches[i]->aliceScore << std::endl;
        std::cout << "bobScore:" << matches[i]->bobScore << std::endl;
        std::cout << "expected:" << resultMatches[i]->result << std::endl;
        break;
    }


    for (int i=0; i<matches.size(); i++) {
        delete matches[i];
    }
    for (int i=0; i<resultMatches.size(); i++) {
        delete resultMatches[i];
    }
}