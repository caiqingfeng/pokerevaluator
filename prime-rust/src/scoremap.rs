// +--------+--------+--------+--------+
// |xxxbbbbb|bbbbbbbb|cdhsrrrr|xxpppppp|
// +--------+--------+--------+--------+
// xxxAKQJT 98765432 CDHSrrrr xxPPPPPP
// 00001000 00000000 01001011 00100101    King of Diamonds(Kd)
// 00000000 00001000 00010011 00000111    Five of Spades(5s)
// 00000010 00000000 10001001 00011101    Jack of Clubs(Jc)
// 00011111 11111111 11111101 00011111    赖子(Xn)
// p = prime number of rank (deuce=2,trey=3,four=5,...,ace=41)
// r = rank of card (deuce=0,trey=1,four=2,five=3,...,ace=12)
// cdhs = suit of card (bit turned on based on suit of card)
// b = bit turned on depending on rank of card
    // Straight Flush   10 
    // Four of a Kind   156      [(13 choose 2) * (2 choose 1)]
    // Full Houses      156      [(13 choose 2) * (2 choose 1)]
    // Flush            1277     [(13 choose 5) - 10 straight flushes]
    // Straight         10 
    // Three of a Kind  858      [(13 choose 3) * (3 choose 1)]
    // Two Pair         858      [(13 choose 3) * (3 choose 2)]
    // One Pair         2860     [(13 choose 4) * (4 choose 1)]
    // High Card      + 1277     [(13 choose 5) - 10 straights]
    // -------------------------
    // TOTAL            7462
        
    pub const straight_flush_max:u64 = 0 ;
    pub const four_of_akind_max:u64 = 10; 
    pub const full_house_max:u64 = 10+156; //166
    pub const flush_max:u64 = 10+156+156;  //322
    pub const straight_max:u64 = flush_max+1277; //1599
    pub const three_of_akind_max:u64  = straight_max+10; //1609
    pub const two_pair_max:u64 = three_of_akind_max+858; //2467
    pub const one_pair_max:u64 = two_pair_max+858; //3325
    pub const high_card_max:u64 = one_pair_max+2860; //6185
    
    use std::collections::HashMap;
    use std::vec::Vec;
    pub static primes:[u64; 14] = [2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43];

    lazy_static! {
        pub static ref ScoreMap: HashMap<u64, u64> = { //2019.12.16 rust 对overflow很严
            let mut m = HashMap::new();
            build_score_map(&mut m);
            m
        };
    
        pub static ref FlushMap: HashMap<u64, u64> = { //2019.12.16 rust 对overflow很严
            let mut m = HashMap::new();
            build_flush_map(&mut m);
            m
        };
    
        static ref FourMap: HashMap<u64, HashMap<u64, u64>> = {
            let mut offset:u64 = 0;
            let mut m = HashMap::new();
            let mut i:i32 = 12;
            while i>=0 {
                let mut m2: HashMap<u64, u64> = HashMap::new();
                let mut j:i32 = 12;
                while j>=0 {
                    if j==i { j=j-1; continue; }
                    m2.insert(primes[j as usize], offset.clone());
                    offset = offset+1;
                    j = j-1;
                }
                m.insert(primes[i as usize], m2);
                i = i-1;
            }
            m    
        };
    
        static ref ThreeMap: HashMap<u64, HashMap<u64, HashMap<u64, u64>>> = {
            let mut offset:u64 = 0;
            let mut m = HashMap::new();
            let mut i:i32 = 12;
            while i>=0 {
                let mut m1: HashMap<u64, HashMap<u64, u64>> = HashMap::new();
                let mut j:i32 = 12;
                while j>=0 {
                    if j == i {j=j-1; continue;}
                    let mut m2: HashMap<u64, u64> = HashMap::new();
                    m2.insert(0, offset.clone());
                    let mut x:i32 = j-1;
                    while x>=0 {
                        if x == i {x=x-1; continue;}
                        m2.insert(primes[x as usize], offset.clone());
                        offset = offset + 1;
                        x=x-1;
                    }
                    m1.insert(primes[j as usize], m2.clone());
                    m1.insert(0, m2.clone());
                    j=j-1;
                }
                m.insert(primes[i as usize], m1);
                i=i-1;
            }
            m
        };
    
        static ref TwoMap: HashMap<u64, HashMap<u64, HashMap<u64, u64>>> = {
            let mut offset:u64 = 0;
            let mut m = HashMap::new();    
            let mut i:i32 = 12;
            while i>=0 {
                let mut m1: HashMap<u64, HashMap<u64, u64>> = HashMap::new();
                let mut j:i32 = i-1;
                while j>=0 {
                    let mut m2: HashMap<u64, u64> = HashMap::new();
                    m2.insert(0, offset.clone());
                    let mut x:i32 = 12;
                    while x>=0 {
                        if x == i || x == j {x=x-1; continue;}
                        m2.insert(primes[x as usize], offset);
                        offset =  offset+1;
                        x=x-1;
                    }
                    m1.insert(primes[j as usize], m2.clone());
                    m1.insert(0, m2.clone());
                    j = j-1;
                }
                m.insert(primes[i as usize], m1);
                i=i-1;
            }
            m
        };
    
        static ref OneMap: HashMap<u64, HashMap<u64, HashMap<u64, HashMap<u64, u64>>>> = {
            let mut m = HashMap::new();    
            let mut offset:u64 = 0;
            let mut i:i32 = 12;
            while i>=0 {
                let mut m1: HashMap<u64, HashMap<u64, HashMap<u64, u64>>> = HashMap::new();
                let mut j:i32 = 12;
                while j>=0 {
                    let mut m2: HashMap<u64, u64> = HashMap::new();
                    if j == i {j=j-1; continue;}
                    let mut m2: HashMap<u64, HashMap<u64, u64>> = HashMap::new();
                    let mut x:i32 = j-1; 
                    while x>=0 {
                        if x == i {x=x-1; continue;}
                        let mut m3: HashMap<u64, u64> = HashMap::new();
                        let mut y:i32 = x-1;
                        m3.insert(0, offset);
                        while y>=0 {
                            if y == i {y=y-1; continue;}
                            m3.insert(primes[y as usize], offset);
                            offset = offset+1;
                            y = y-1;
                        }
                        m2.insert(primes[x as usize], m3.clone());
                        m2.insert(0, m3.clone());
                        x=x-1;
                    }
                    m1.insert(primes[j as usize], m2.clone());
                    m1.insert(0, m2.clone());
                    j=j-1;
                }
                m.insert(primes[i as usize], m1);
                i=i-1;
            }
            m
        };
    }
    
    fn add_key(m: &mut HashMap<u64, u64>, key: &u64, value: &u64) {
        // if *key == 41u64*43u64*31u64*29u64*19u64 { //41*41*41*41*43 //41*37*31*29*19
        //     println!("key = {}, v={}", *key, *value);
        // }    
        m.entry(*key).or_insert(value.clone());
    }
    
    fn offset_four(p1: &u64, p2: &u64) -> u64 {
        FourMap.get(&p1).unwrap().get(&p2).unwrap().clone()
    }
    
    fn offset_three(p1: &u64, p2: &u64, p3: &u64) -> u64 {
        ThreeMap.get(&p1).unwrap().get(&p2).unwrap().get(&p3).unwrap().clone()
    }
    
    fn offset_two(p1: &u64, p2: &u64, p3: &u64) -> u64 {
        TwoMap.get(&p1).unwrap().get(&p2).unwrap().get(&p3).unwrap().clone()
    }
    
    fn offset_one(p1: &u64, p2: &u64, p3: &u64, p4: &u64) -> u64 {
        match OneMap.get(&p1) {
            Some(ref m1) => {
                match m1.get(&p2) {
                    Some(ref m2) => {
                        match m2.get(&p3) {
                            Some(ref m3) => {
                                match m3.get(&p4) {
                                    Some(ref v) => return **v,
                                    _ => return 0,
                                }
                            }
                            _ => return 0,
                        }
                    }
                    _ => return 0,
                }
            }
            _ => return 0,
        }
        // OneMap.get(&p1).unwrap().get(&p2).unwrap().get(&p3).unwrap().get(&p4).unwrap().clone()
    }
    
    fn build_four_of_akind(sMap: &mut HashMap<u64, u64>) {
        //从A开始，A是primes[12]=41
        //首先计算4+1，4+2，4+3的各种情况
        let mut i:i32 = 12;
        while i >= 0 {
            // println!("i={}", i);
            let base = primes[i as usize]*primes[i as usize]*primes[i as usize]*primes[i as usize];
            //首先计算4+Xn，这里把Xn替换成最大的那个踢脚算其offset
            let key2:u64 = base * 43 as u64; 
            let fix_offset:u64 = (12-i as u64 )*12;
            add_key(sMap, &key2, &(fix_offset+four_of_akind_max)); //如果第一次已经加了(k,v)，不会覆盖
            let mut j:i32 = 12;
            while j >=0 {
                // println!("jj={}", j);
                if j == i {j=j-1;continue;}
                //计算4+1
                let key = base * primes[j as usize];
                let offset = offset_four(&primes[i as usize], &primes[j as usize]);
                add_key(sMap, &key, &(offset+four_of_akind_max));
                //然后计算6张牌的情况：4+1+1
                let mut x:i32 = j;
                while x>=0 {
                    //x 不能等于i
                    if x == i {x=x-1;continue;}
                    let key3 =  base * primes[j as usize] * primes[x as usize];
                    add_key(sMap, &key3, &(offset+four_of_akind_max)); 
                    //然后计算7张牌 4+1+1+1
                    let mut y:i32 = j;
                    while y>=0 {
                        //y 不能等于i
                        if y == i {y=y-1; continue;}
                        let key5 = key3 * primes[y as usize];
                        add_key(sMap, &key5, &(offset+four_of_akind_max)); 
                        y = y-1;
                    }
                    //然后计算7张牌 4+1+1+Xn
                    let key6 = key3 * 43  as u64;
                    add_key(sMap, &key6, &(fix_offset+four_of_akind_max));
                    x=x-1;
                }
                //然后计算4+1+Xn
                let key4 = base  * primes[j as usize] * 43  as u64;
                add_key(sMap, &key4, &(fix_offset+four_of_akind_max));
                j=j-1;
            }
            i=i-1;
        }
        //计算3+Xn+1，3+Xn+2，3+Xn+3的各种情况
        //这里Xn必须要替换成3条中的一个
        i = 12;
        while i>=0 {
            // println!("iii={}", i);
            let base = primes[i as usize]*primes[i as usize]*primes[i as usize]*43 as u64;
            let mut j:i32 = 12;
            while j>=0 {
                // println!("jjjj={}", j);
                if j == i {j=j-1;continue;}
                //计算4+1
                let key = base * primes[j as usize];
                let offset = offset_four(&primes[i as usize], &primes[j as usize]);
                add_key(sMap, &key, &(offset+four_of_akind_max));
                //然后计算6张牌的情况：4+1+1
                let mut x:i32 = j;
                while x>=0 {
                    //x 不能等于i
                    if x == i {x=x-1;continue;}
                    let key3 =  key * primes[x as usize];
                    add_key(sMap, &key3, &(offset+four_of_akind_max)); 
                    //然后计算7张牌 4+1+1+1
                    let mut y:i32 = j;
                    while y>=0 {
                        //y 不能等于i
                        if y == i {y=y-1;continue;}
                        let key5 = key3 * primes[y as usize];
                        add_key(sMap, &key5, &(offset+four_of_akind_max)); 
                        y = y-1;
                    }
                    x = x-1;
                }
                j = j-1;
            }
            i = i-1;
        }
    }
    
    fn build_full_house(sMap: &mut HashMap<u64, u64>) {
        //第一种情况：3条+1对+任意2张牌（不可能有癞子，否则就是4条了）
        //从A开始，A是primes[12]=41
        let mut i:i32 = 12;
        while i>=0 {
            // println!("i={}", i);
            let mut base = primes[i as usize]*primes[i as usize]*primes[i as usize];
            let mut j:i32 = 12;
            while j>=0 {
                // println!("j={}", j);
                if j == i {j=j-1;continue;}
                //计算3+2
                let mut key = base * primes[j as usize]*primes[j as usize];
                let mut offset = offset_four(&primes[i as usize], &primes[j as usize]);
                add_key(sMap, &key, &(offset+full_house_max));
                //然后计算6张牌的情况：3+2+1
                let mut x:i32 = 12;
                while x>=0 {
                    //x 不能等于i
                    // println!("x={}", x);
                    if x == i {x=x-1; continue;}
                    let mut key3 =  base * primes[j as usize] * primes[j as usize] * primes[x as usize];
                    add_key(sMap, &key3, &(offset+full_house_max)); 
                    //然后计算7张牌 3+2+1+1
                    let mut y:i32 = 12;
                    while y>=0 {
                        //y 不能等于i
                        // println!("y={}", y);
                        if y == i {y=y-1; continue;}
                        let mut key5 = base * primes[j as usize] * primes[j as usize] * primes[x as usize] * primes[y as usize];
                        add_key(sMap, &key5, &(offset+full_house_max)); 
                        y=y-1;
                    }
                    x=x-1;
                }
                j=j-1;
            }
            i=i-1;
        }
        //第二种情况：2对+1个癞子+任意2张牌（不能有3条，可以是3对）
        i = 12;
        while i>=0 {
            let mut base = primes[i as usize]*primes[i as usize]*43 as u64;
            // println!("ii={}", i);
            let mut j:i32 = i-1;
            while j>=0 {
                //计算2+2+Xn
                // println!("jj={}", j);
                let mut key = base * primes[j as usize] * primes[j as usize];
                let mut offset = offset_four(&primes[i as usize], &primes[j as usize]);
                add_key(sMap, &key, &(offset+full_house_max));
                //然后计算6张牌的情况：2+2+Xn+1
                let mut x:i32 = 12;
                while x>=0 {
                    // println!("xx={}", x);
                    //x 不能等于i or j
                    if x == i || x == j {x=x-1; continue;}
                    let mut key3 =  base * primes[j as usize] * primes[j as usize] * primes[x as usize];
                    add_key(sMap, &key3, &(offset+full_house_max)); 
                    //然后计算7张牌 2+2+Xn+1+1
                    let mut y:i32 = 12;
                    while y>=0 {
                        // println!("yy={}", y);
                        //y 不能等于i or j
                        if y == i || y == j {y=y-1; continue;}
                        let mut key5 = base * primes[j as usize] * primes[j as usize] * primes[x as usize] * primes[y as usize];
                        add_key(sMap, &key5, &(offset+full_house_max)); 
                        y = y-1;
                    }
                    x=x-1;
                }
                j=j-1;
            }
            i=i-1;
        }
    }
    
    //这个算法非常巧妙
    //
    fn build_straight(m: &mut HashMap<u64, u64>, base_offset: u64) {
        let mut straight: Vec<Vec<u64>> = Vec::new();
        let mut i:i32 = 12;
        while i>=4 {
            let mut s: Vec<u64> = Vec::new();
            s.push(primes[i as usize]);
            s.push(primes[(i-1)  as usize]);
            s.push(primes[(i-2) as usize]);
            s.push(primes[(i-3) as usize]);
            s.push(primes[(i-4) as usize]);
            s.push(43);
            straight.push(s);
            i = i-1;
        }
        //对12345的顺子做类似操作
        let mut s: Vec<u64> = Vec::new();
        s.push(primes[3]);
        s.push(primes[2]);
        s.push(primes[1]);
        s.push(primes[0]);
        s.push(41 as u64);
        s.push(43 as u64);
        straight.push(s);
    
        let mut offset:u64 = 0;
        for it in straight.iter() {
            let mut s: Vec<u64> = it.clone();
            let mut t:u64 = s[0]*s[1]*s[2]*s[3]*s[4]*s[5];
            let mut score:u64 = base_offset + offset;
            offset = offset + 1;
            for it2 in s.iter() {
                let mut key:u64 = t/it2;
                // std::cout<<"key = " << key << " score =" << score <<std::endl;
                add_key(m, &key, &score); //5张牌
                let mut x:i32 = 12;
                while x>=0 {
                    //added 2019.08.04 支持按6张牌来找
                    let mut key2 = key * primes[x as usize];
                    add_key(m, &key2, &score);
                    let mut y:i32 = 12;
                    while y>=0 {
                        let mut key1 = key * primes[x as usize] * primes[y as usize];
                        add_key(m, &key1, &score); //7张牌
                        y = y-1;
                    }
                    x=x-1;
                }
            }
        }
    }
    
    fn build_three_of_akind(m: &mut HashMap<u64, u64>) {
        //第一种情况：3条+任意4张散牌（不可能有癞子，和其它对子，否则就是4条或者full了）
        //从A开始，A是primes[12]=41
        let mut i:i32 = 12;
        while i>=0 {
            let mut base:u64 = primes[i as usize]*primes[i as usize]*primes[i as usize];
            //2019-08-25支持3张牌（菠萝头道）
            let mut offset1 = offset_three(&primes[i as usize], &0, &0);
            add_key(m, &base, &offset1);
            let mut j:i32 = 12;
            while j>=0 {
                if j == i {j=j-1; continue;}
                //计算3+1
                let mut key = base * primes[j as usize];
                let mut offset = offset_three(&primes[i as usize], &primes[j as usize], &0);
                add_key(m, &key, &(offset+three_of_akind_max));
                //然后计算5张牌的情况：3+1+1
                let mut x:i32 = j-1;
                while x>=0 {
                    //x 不能等于i
                    if x == i {x=x-1; continue;}
                    let mut key3 =  base * primes[j as usize] * primes[x as usize];
                    let mut fix_offset = offset_three(&primes[i as usize], &primes[j as usize], &primes[x as usize]);
                    add_key(m, &key3, &(fix_offset+three_of_akind_max)); 
                    //然后计算6张牌 3+1+1+1
                    let mut y:i32 = x-1;
                    while y>=0 {
                        //y 不能等于i
                        if y == i {y=y-1; continue;}
                        let mut key5 = base * primes[j as usize] * primes[x as usize] * primes[y as usize];
                        add_key(m, &key5, &(fix_offset+three_of_akind_max));
                        //7张牌的情况
                        let mut z:i32 = y-1;
                        while z>=0 {
                            if z == i {z=z-1; continue;}
                            let mut key6 = base * primes[j as usize] * primes[x as usize] * primes[y as usize] * primes[z as usize];
                            add_key(m, &key6, &(fix_offset+three_of_akind_max));
                            z=z-1;
                        } 
                        y=y-1;
                    }
                    x=x-1;
                }
                j=j-1;
            }
            i=i-1;
        }
        //第二种情况：1对+1个癞子+任意4张牌（不能有3条、2对）
        let mut i:i32 = 12;
        while i>=0 {
            let mut base = primes[i as usize]*primes[i as usize]*43;
            let mut offset1 = offset_three(&primes[i as usize], &0, &0);
            //2019-08-25 支持头道3张牌，1对+1个癞子
            add_key(m, &base, &(offset1+three_of_akind_max));
            let mut j:i32 = 12;
            while j>=0 {
                if j == i {j=j-1; continue;}
                //计算2+Xn+1
                let mut key = base * primes[j as usize];
                let mut offset = offset_three(&primes[i as usize], &primes[j as usize], &0);
                add_key(m, &key, &(offset+three_of_akind_max));
                //然后计算5张牌的情况：2+Xn+1+1
                let mut x:i32 = j-1;
                while x>=0 {
                    //x 不能等于i or j
                    if x == i {x=x-1; continue;}
                    let mut key3 =  base * primes[j as usize] * primes[x as usize];
                    let mut fix_offset = offset_three(&primes[i as usize], &primes[j as usize], &primes[x as usize]);
                    add_key(m, &key3, &(fix_offset+three_of_akind_max)); 
                    //然后计算6张牌 2+Xn+1+1+1
                    let mut y:i32 = x-1;
                    while y>=0 {
                        //y 不能等于i or j
                        if y == i {y=y-1; continue;}
                        let mut key5 = base * primes[j as usize] * primes[x as usize] * primes[y as usize];
                        add_key(m, &key5, &(fix_offset+three_of_akind_max)); 
                        //计算7张牌的情况
                        let mut z:i32 = y-1;
                        while z>=0 {
                            if z == i {z=z-1; continue;}
                            let mut key6 = base * primes[j as usize] * primes[x as usize] * primes[y as usize] * primes[z as usize];
                            add_key(m, &key6, &(fix_offset+three_of_akind_max));
                            z = z-1;
                        } 
                        y = y-1;
                    }
                    x=x-1;
                }
                j=j-1;
            }
            i=i-1;
        }
    }
    
    fn build_two_pair(m: &mut HashMap<u64, u64>) {
        //Two Pair，不可能有癞子，只要对普通2对+3建表
        //2对+任意3张散牌
        //从A开始，A是primes[12]=41
        let mut i:i32 = 12;
        while i>=1 {
            let mut base = primes[i as usize]*primes[i as usize];
            let mut j:i32 = i-1;
            while j>=0 {
                //至此选出了大小两对
                let mut key = base * primes[j as usize] * primes[j as usize];
                let mut offset = offset_two(&primes[i as usize], &primes[j as usize], &0);
                add_key(m, &key, &(offset+two_pair_max)); //支持4张牌的情况
                //然后计算5张牌的情况：2+2+1
                let mut x:i32 = 12;
                while x>=0 {
                    //x 不能等于i
                    if x == i || x == j {x=x-1; continue;}
                    let mut key3 =  base * primes[j as usize] * primes[j as usize] * primes[x as usize];
                    let mut fix_offset = offset_two(&primes[i as usize], &primes[j as usize], &primes[x as usize]);
                    add_key(m, &key3, &(fix_offset+two_pair_max)); 
                    //然后计算6张牌 2+2+1+1
                    let mut y:i32 = x;
                    while y>=0 {
                        //y 不能等于i/j
                        if y == i || y == j {y=y-1; continue;}
                        let mut key5 = base * primes[j as usize] * primes[j as usize] * primes[x as usize] * primes[y as usize];
                        add_key(m, &key5, &(fix_offset+two_pair_max));
                        //7张牌的情况
                        let mut z:i32 = y;
                        while z>=0 {
                            if z == i || z == j {z=z-1; continue;}
                            let mut key6 = base * primes[j as usize] * primes[j as usize] * primes[x as usize] * primes[y as usize] * primes[z as usize];
                            add_key(m, &key6, &(fix_offset+two_pair_max));
                            z = z-1;
                        } 
                        y=y-1;
                    }
                    x=x-1;
                }
                j=j-1;
            }
            i=i-1;
        }
    }
    
    fn build_one_pair(m: &mut HashMap<u64, u64>) {
        //第一种情况，1对+任意5张散牌，没有癞子
        //从A开始，A是primes[12]=41
        let mut i:i32 = 12;
        while i>=0 {
            let mut base = primes[i as usize]*primes[i as usize];
            let mut j:i32 = 12;
            while j>=0 {
                if j == i {j=j-1; continue;}
                //至此选出了1对+1
                let mut  key = base * primes[j as usize];
                let mut  offset = offset_one(&primes[i as usize], &primes[j as usize], &0, &0);
                add_key(m, &key, &(offset+one_pair_max));//支持3张牌的情况
                //然后计算4张牌的情况：2+1+1
                let mut x:i32 = j-1;
                while x>=0 {
                    //x 不能等于i
                    if x == i {x=x-1; continue; }
                    let mut key3 =  base * primes[j as usize] * primes[x as usize];
                    let mut offset2 = offset_one(&primes[i as usize], &primes[j as usize], &primes[x as usize], &0);
                    add_key(m, &key3, &(offset2+one_pair_max));
                    //然后计算5张牌 2+1+1+1
                    let mut y:i32 = x-1;
                    while y>=0 {
                        //y 不能等于i/j
                        if y == i {y=y-1; continue;}
                        let mut  key5 = base * primes[j as usize] * primes[x as usize] * primes[y as usize];
                        let mut  fix_offset = offset_one(&primes[i as usize], &primes[j as usize], &primes[x as usize], &primes[y as usize]);
                        add_key(m, &key5, &(fix_offset+one_pair_max));
                        //6张牌的情况
                        let mut z:i32 = y-1;
                        while z>=0 {
                            if z == i {z=z-1; continue;}
                            let mut  key6 = base * primes[j as usize] * primes[x as usize] * primes[y as usize] * primes[z as usize];
                            add_key(m, &key6, &(fix_offset+one_pair_max));
                            //7张牌的情况
                            let mut t:i32 = z-1;
                            while t>=0 {
                                if t == i {t=t-1; continue; }
                                let mut  key7 = base * primes[j as usize] * primes[x as usize] * primes[y as usize] * primes[z as usize] * primes[t as usize];
                                add_key(m, &key7, &(fix_offset+one_pair_max));
                                t=t-1;
                            }
                            z=z-1;
                        } 
                        y=y-1;
                    }
                    x=x-1;
                }
                j=j-1;
            }
            i=i-1;
        }
        //第2种情况，1癞子+任意6张散牌
        //从A开始，A是primes[12]=41
        i = 12;
        while i>=0 {
            let mut  base = 43*primes[i as usize];
            let mut j:i32 = i-1;
            while j>=0 {
                //至此选出了1对+1
                let mut  key = base * primes[j as usize];
                let mut  offset = offset_one(&primes[i as usize], &primes[j as usize], &0, &0);
                add_key(m, &key, &(offset+one_pair_max));//支持3张牌的情况
                //然后计算4张牌的情况：2+1+1
                let mut x:i32 = j-1;
                while x>=0 {
                    //x 不能等于i
                    if x == i {x=x-1; continue;}
                    let mut  key3 =  base * primes[j as usize] * primes[x as usize];
                    let mut  offset2 = offset_one(&primes[i as usize], &primes[j as usize], &primes[x as usize], &0);
                    add_key(m, &key3, &(offset2+one_pair_max));
                    //然后计算5张牌 2+1+1+1
                    let mut y:i32 = x-1;
                    while y>=0 {
                        //y 不能等于i/j
                        if y == i {y=y-1; continue;}
                        let mut  key5 = base * primes[j as usize] * primes[x as usize] * primes[y as usize];
                        let mut  fix_offset = offset_one(&primes[i as usize], &primes[j as usize], &primes[x as usize], &primes[y as usize]);
                        add_key(m, &key5, &(fix_offset+one_pair_max));
                        //6张牌的情况
                        let mut z:i32 = y-1;
                        while z>=0 {
                            if z == i { z=z-1; continue;}
                            let mut  key6 = base * primes[j as usize] * primes[x as usize] * primes[y as usize] * primes[z as usize];
                            add_key(m, &key6, &(fix_offset+one_pair_max));
                            //7张牌的情况
                            let mut t:i32 = z-1;
                            while t>=0 {
                                if t == i {t=t-1; continue;}
                                let mut  key7 = base * primes[j as usize] * primes[x as usize] * primes[y as usize] * primes[z as usize] * primes[t as usize];
                                add_key(m, &key7, &(fix_offset+one_pair_max));
                                t = t-1;
                            }
                            z = z-1;
                        } 
                        y=y-1;
                    }
                    x=x-1;
                }
                j=j-1;
            }
            i=i-1;
        }
    }
    
    fn build_high_card(m: &mut HashMap<u64, u64>, base_offset: u64) {
        //只有一种情况，任意5-7张散牌，没有癞子
        //modified 2019-08-26 直接对7张牌建表，避免排序
        //跟前面一样，这里调用addKey，不会覆盖之前的k,v
        //从A开始，A是primes[12]=41
        let mut offset:u64 = 0;
        let mut i:i32 = 12;
        while i>=4 {
            let mut base = primes[i as usize];
            let mut j:i32 = i-1;
            while j>=3 {
                //至此选出了1+1
                let mut key = base * primes[j as usize];
                add_key(m, &key, &(offset+base_offset)); //支持2张牌的情况
                //然后计算3张牌的情况：1+1+1
                let mut x:i32 = j-1;
                while x>=0 {
                    let mut key3 =  base * primes[j as usize] * primes[x as usize];
                    add_key(m, &key3, &(offset+base_offset)); 
                    //然后计算4张牌 1+1+1+1
                    let mut y:i32 = x-1;
                    while y>=0 {
                        let mut key5 = base * primes[j as usize] * primes[x as usize] * primes[y as usize];
                        add_key(m, &key5, &(offset+base_offset));
                        //5张牌的情况
                        let mut z:i32 = y-1;
                        while z>=0 {
                            let mut key6 = base * primes[j as usize] * primes[x as usize] * primes[y as usize] * primes[z as usize];
                                // std::cout << "kkkkk:" << key6 << offset << std::endl;
                            let mut fix_offset =  offset;
                            if !m.contains_key(&key6) {
                                add_key(m, &key6, &(fix_offset+base_offset));
                                offset = offset+1;
                                // std::cout << key6 << offset << std::endl;
                                // 6张牌的情况
                                let mut t:i32 = z-1;
                                while t>=0 {
                                    let mut key7 = key6 * primes[t as usize];
                                    add_key(m, &key7, &(fix_offset+base_offset));
                                    //7张牌的情况
                                    let mut v:i32 = t-1;
                                    while v>=0 {
                                        let mut key8 = key7 * primes[v as usize];
                                        add_key(m, &key8, &(fix_offset+base_offset));
                                        v = v-1;
                                    }
                                    t=t-1;
                                }
                            }
                            z=z-1;
                        } 
                        y=y-1;
                    }
                    x=x-1;
                }
                j=j-1;
            }
            i=i-1;
        }
    }
    
    // 得到二维的数组
    // 递归第一步get vector<vector<12>, vector<11>, ..., vector<1>, vector<0>>
    // 第二步vector<vector<12, 11>, vector<11, 10>, ..., vector<1, 0>>
    fn get_comb(comb: &mut Vec<Vec<u64>>, level:u32) {
        let mut orgSet:Vec<Vec<u64>> = comb.clone();
    
        if orgSet.len() == 0 {
            let mut i:i32 = 12;
            while i>=0 {
                let mut s: Vec<u64> = Vec::new();
                s.push(i as u64);
                comb.push(s);
                i=i-1;
            }
            get_comb(comb, level-1);
            return;
        }
    
        if level == 0 { return; }
        comb.clear();
        for it in orgSet.iter() {
            let mut n: &Vec<u64> = it;
            let mut last:i32 = n[n.len()-1] as i32;
            let mut i:i32 = last-1;
            while i>=0 {
                let mut x: Vec<u64> = n.clone();
                x.push(i as u64);
                comb.push(x);
                i = i-1;
            }        
        }
        get_comb(comb, level-1);
    }
    
    fn gen_ghost_score(m: &mut HashMap<u64, u64>, comb: &Vec<Vec<u64>>) {
        for it in comb.iter() {
            let mut key:u64 = 1;
            let mut maxCardIndex:u64 = 12;
            for nit in it.iter() {
                if maxCardIndex == *nit {maxCardIndex = maxCardIndex-1;};
                key = key * primes[*nit as usize];
            }
            //现在得到了任意4-5-6张牌+1个癞子的组合
            // if key == 41u64*31u64*29u64*19u64 {  //test case FlushMap[&(41u64*43u64*31u64*29u64*19u64)] 报错no entry found for key
            //     println!("k={}", key);
            // }
            let scoreKey = key * primes[maxCardIndex as usize];
            match m.get(&scoreKey) {
                Some(&v) => add_key(m, &(key*43u64), &v),
                _ => (),
            }
        }
    }
    
    fn build_flush_map(m: &mut HashMap<u64, u64>) {
        //针对任意5-7张散牌，有癞子和无赖子，一次性查表解决问题
        //modified 2019-08-26 直接对7张牌建表，避免排序和替换
        //从A开始，A是primes[12]=41
        //先把顺子、fullhouse、四条 建起来
        build_four_of_akind(m);
        build_full_house(m);
        build_straight(m, straight_flush_max);
    
        //然后再建立7张散牌，支持带癞子的情况
        // todo : 怎么和buildHighCard代码合并？
        // 7张无癞子的情况
        build_high_card(m, flush_max);
        let mut comb: Vec<Vec<u64>> = Vec::new();
        // //13 选4张牌+癞子
        get_comb(&mut comb, 4);
        // // std::cout << "len=" << comb.size() << std::endl;
        gen_ghost_score(m, &comb);
        //13 选5张牌+癞子
        comb.clear();
        get_comb(&mut comb, 5);
        gen_ghost_score(m, &comb);
        // //13 选6张牌+癞子
        comb.clear();
        get_comb(&mut comb, 6);
        gen_ghost_score(m, &comb);
    }
    
    fn build_score_map(m: &mut HashMap<u64, u64>) {
        // println!("hello, world");
        build_four_of_akind(m);
        build_full_house(m);
        build_straight(m, straight_max);
        build_three_of_akind(m);
        build_two_pair(m);
        build_one_pair(m);
        build_high_card(m, high_card_max);
    }
    
    #[cfg(test)]
    mod tests {
        #[test]
        fn score_map() {
            // super::build_score_map();
            // assert_eq!(super::ScoreMap.len(), 12909); //after four of akind
            // assert_eq!(super::ScoreMap.len(), 30667); //after full house
            // assert_eq!(super::ScoreMap.len(), 34355); //after straight
            // assert_eq!(super::ScoreMap.len(), 53360); //after three of akind
            // assert_eq!(super::ScoreMap.len(), 74502); //after two pair
            // assert_eq!(super::ScoreMap.len(), 97538); //after one pair
            assert_eq!(super::ScoreMap.len(), 102964); //after high card
            assert_eq!(super::FlushMap.len(), 42324); //
            assert_eq!(super::FourMap.len(), 13);
    
            //c++ buildFourOfAKind
            assert_eq!(super::ScoreMap.get(&(41*41*41*41*37 as u64)).unwrap(), &super::four_of_akind_max);
            assert_eq!(super::ScoreMap.get(&(2*2*2*2*3 as u64)).unwrap(), &(super::four_of_akind_max+155));    
            assert_eq!(super::ScoreMap.get(&(41*41*41*41*43 as u64)).unwrap(), &super::four_of_akind_max);
            assert_eq!(super::ScoreMap.get(&(2*2*2*2*43 as u64)).unwrap(), &(super::four_of_akind_max+144));    
            assert_eq!(super::ScoreMap.get(&(41*41*41*41*37*2 as u64)).unwrap(), &super::four_of_akind_max);
            assert_eq!(super::ScoreMap.get(&(2*2*2*2*31*41 as u64)).unwrap(), &(super::four_of_akind_max+144));
            assert_eq!(super::ScoreMap.get(&(41*41*41*41*37*43 as u64)).unwrap(), &super::four_of_akind_max);
            assert_eq!(super::ScoreMap.get(&(2*2*2*2*31*41*43 as u64)).unwrap(), &(super::four_of_akind_max+144));
            
            // c++ test3_xn
            assert_eq!(super::ScoreMap[&(41*41*41*43*37 as u64)], super::four_of_akind_max);    
            assert_eq!(super::ScoreMap[&(2*2*2*43*3 as u64)], super::four_of_akind_max+155);    
            assert_eq!(super::ScoreMap[&(41*41*41*43*37*37 as u64)], super::four_of_akind_max);    
            assert_eq!(super::ScoreMap[&(2*2*2*43*31*41 as u64)], super::four_of_akind_max+144);    
            assert_eq!(super::ScoreMap[&(2*2*2*43*41*41*41 as u64)], super::four_of_akind_max+11);    
    
            // c++ buildFullHouse
            assert_eq!(super::ScoreMap[&(41*41*41*37*37 as u64)], super::full_house_max);    
            assert_eq!(super::ScoreMap[&(2*2*2*3*3 as u64)], super::full_house_max+155);    
            assert_eq!(super::ScoreMap[&(41*41*43*37*37 as u64)], super::full_house_max);    
            assert_eq!(super::ScoreMap[&(2*2*43*3*3 as u64)], super::full_house_max+144-1);    
    
            // c++ buildThree
            assert_eq!(super::ScoreMap[&(41*41*41*37*31 as u64)], super::three_of_akind_max);    
            assert_eq!(super::ScoreMap[&(2*2*2*5*3 as u64)], super::three_of_akind_max+857);    
            assert_eq!(super::ScoreMap[&(41*41*43*37*31 as u64)], super::three_of_akind_max);    
            assert_eq!(super::ScoreMap[&(2*2*43*5*3 as u64)], super::three_of_akind_max+857);    
            assert_eq!(super::ScoreMap[&(41*41*41*37*31*2*3 as u64)], super::three_of_akind_max);    
            assert_eq!(super::ScoreMap[&(2*2*2*5*3*41*37 as u64)], super::three_of_akind_max+858-66);    
            assert_eq!(super::ScoreMap[&(41*41*43*37*31*2*3 as u64)], super::three_of_akind_max);    
            assert_eq!(super::ScoreMap[&(2*2*43*41*37 as u64)], super::three_of_akind_max+858-12*11/2);    
            assert_eq!(super::ScoreMap[&(2*2*43*41*37*31 as u64)], super::three_of_akind_max+858-12*11/2);    
            assert_eq!(super::ScoreMap[&(2*2*43*5*41*37 as u64)], super::three_of_akind_max+858-12*11/2);    
            assert_eq!(super::ScoreMap[&(2*2*43*5*41*37*31 as u64)], super::three_of_akind_max+858-12*11/2);    
            assert_eq!(super::ScoreMap[&(2*2*43*5*41*37*11 as u64)], super::three_of_akind_max+858-12*11/2);    
            
            // c++ buildTwo
            assert_eq!(super::ScoreMap[&(41*41*37*37*31 as u64)], super::two_pair_max);    
            assert_eq!(super::ScoreMap[&(2*2*3*5*3 as u64)], super::two_pair_max+857);    
            assert_eq!(super::ScoreMap[&(41*41*37*37*31*2*3 as u64)], super::two_pair_max);    
            assert_eq!(super::ScoreMap[&(2*2*3*5*3*41*37 as u64)], super::two_pair_max+858-11);    
            
            // c++ buildOne
            assert_eq!(super::ScoreMap[&(41*41*37*31*29 as u64)], super::one_pair_max);    
            assert_eq!(super::ScoreMap[&(2*2*3*5*7 as u64)], super::one_pair_max+2859);    
            assert_eq!(super::ScoreMap[&(41*43*37*31*19 as u64)], super::one_pair_max+2);    
            assert_eq!(super::ScoreMap[&(2*43*3*5*13 as u64)], super::one_pair_max+12*11*10/6*8-1);    
    
            // c++ buildHigh
            assert_eq!(super::ScoreMap[&(41*19*37*31*29 as u64)], super::high_card_max);    
            assert!(super::ScoreMap[&(41*37*13*7*11 as u64)] > super::high_card_max);    
            assert!(super::ScoreMap[&(41*37*13*7*11*5 as u64)] > super::high_card_max);    
            assert!(super::ScoreMap[&(41*37*13*7*11*3 as u64)] > super::high_card_max);    
            assert!(super::ScoreMap[&(41*37*13*7*11*2 as u64)] > super::high_card_max);    
            assert!(super::ScoreMap[&(41*37*13*11*7*3*2 as u64)] > super::high_card_max);    
            assert!(super::ScoreMap[&(41*37*13*7*11*2 as u64)] > super::high_card_max);    
        }
    
        #[test]
        fn flush_map() {
            let mut comb: Vec<Vec<u64>> = Vec::new();
            super::get_comb(&mut comb, 1);
            assert_eq!(comb.len(), 13);
    
            assert!(super::FlushMap.len() > 0);
            assert_eq!(super::FlushMap[&(41u64*37u64*31u64*29u64*19u64 as u64)], super::flush_max);
            assert_eq!(super::FlushMap[&(41u64*43u64*31u64*29u64*19u64)], super::flush_max);
        }
    }
    