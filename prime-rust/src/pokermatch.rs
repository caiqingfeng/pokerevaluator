use std::string::String;
use crate::cardhelper::{LYSuit, scan_hand_string};
use crate::scoremap;

#[derive(Clone)]
pub struct LYMatch {
    pub alice: String,
    pub alice_score: u64,
    pub bob: String,
    pub bob_score: u64,
    pub result: i32
}

pub fn get_score(hs: &String) -> u64 {
    let mut k=1u64;
    let mut suitKey=1u64;
    let mut keyWithoutXn=1u64;
    let mut suit = LYSuit::Nosuit;

    let _ = scan_hand_string(hs, &mut k, &mut suitKey, &mut suit);
    let score = scoremap::ScoreMap[&k];
    if score < scoremap::straight_max || hs.len() < 10 || suit == LYSuit::Nosuit {
        return score;
    }
    if suit != LYSuit::Nosuit {
        return scoremap::FlushMap[&suitKey];
    }
    score
}

pub fn process_match(m: &mut LYMatch) {
    m.alice_score = get_score(&m.alice);
    m.bob_score = get_score(&m.bob);
    m.result = match m.alice_score>m.bob_score {
        true => 2,
        false => match m.alice_score<m.bob_score {
            true => 1,
            false => 0
        }
    }
}