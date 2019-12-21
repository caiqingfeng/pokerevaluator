use crate::scoremap::primes;  //for rust 2018 https://doc.rust-lang.org/edition-guide/rust-2018/module-system/path-clarity.html
use std::collections::HashMap;

#[derive(Debug, PartialEq)]
pub enum LYSuit {
    Nosuit = 0, Clubs = 1, Diamonds = 2, Hearts = 3, Spades = 4
}
// impl PartialEq for LYSuit {
//     fn eq(&self, other: &LYSuit) -> bool {
//         self == other
//     }
// }
 
enum LYFace {
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
}

lazy_static! {
    static ref PrimesMap: HashMap<char, u64> = { 
        let mut m = HashMap::new();
        m.insert('X', 43);
        m.insert('A', 41);
        m.insert('K', 37);
        m.insert('Q', 31);
        m.insert('J', 29);
        m.insert('T', 23);
        m.insert('9', 19);
        m.insert('8', 17);
        m.insert('7', 13);
        m.insert('6', 11);
        m.insert('5', 7);
        m.insert('4', 5);
        m.insert('3', 3);
        m.insert('2', 2);
        m
    };
}

pub fn scan_hand_string(cs: &String, key: &mut u64, suitKey: &mut u64, suit: &mut LYSuit) ->u64 {
	let l = cs.len();
	*key = 1u64;
	*suitKey = 1u64;
	let mut hasGhost = false;
    let mut clubKey = 1u64;
    let mut club = 0u64; 
    let mut diamondKey = 1u64;
    let mut diamond = 0u64;  
    let mut heartKey = 1u64;
    let mut heart = 0u64;  
    let mut spadeKey = 1u64;
    let mut spade = 0u64;  
	for i in 0..l/2 {
		let mut p = PrimesMap[&(cs.as_bytes()[2*i] as char)];
        *key *= p;
        match cs.as_bytes()[2*i+1] as char {
			'c' => {
				clubKey *= p;
                club += 1;
            },
			'd' => {
				diamondKey *= p;
                diamond += 1;
            },
			'h' => {
				heartKey *= p;
                heart += 1;
            },
			's' => {
				spadeKey *= p;
				spade += 1;
            },
			_ => {
				hasGhost = true;
				club += 1;
				diamond += 1;
				heart += 1;
				spade += 1;
            },
		}
	}
	if club >= 5 {
		*suit = LYSuit::Clubs;
		*suitKey = clubKey;
		if hasGhost {
			*suitKey *= 43u64;
		}
	} else if diamond >= 5 {
		*suit = LYSuit::Diamonds;
		*suitKey = diamondKey;
		if hasGhost {
			*suitKey *= 43u64;
		}
	} else if heart >= 5 {
		*suit = LYSuit::Hearts;
		*suitKey = heartKey;
		if hasGhost {
			*suitKey *= 43u64;
		}
	} else if spade >= 5 {
		*suit = LYSuit::Spades;
		*suitKey = spadeKey;
		if hasGhost {
			*suitKey *= 43u64;
		}
	}
	key.clone()
}

#[cfg(test)]
mod tests {
    #[test]
    fn scan_hand_string() {
        let mut key=0u64;
        let mut suitKey=0u64;
        let mut suit = super::LYSuit::Clubs;
        let _ = super::scan_hand_string(&String::from("AsXnQsJs9s"), &mut key, &mut suitKey, &mut suit);
        assert_eq!(suit, super::LYSuit::Spades);
        assert_eq!(suitKey, 41u64*43u64*31u64*29u64*19u64);
    
        let _ = super::scan_hand_string(&String::from("5c2d5dXnKd7d6d"), &mut key, &mut suitKey, &mut suit);
        assert_eq!(suit, super::LYSuit::Diamonds);
        assert_eq!(suitKey, 43u64*37u64*13u64*11u64*7u64*2u64);
    }
}