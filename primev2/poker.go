package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/caiqingfeng/pokerevaluator/primev2/libpoker"
)
import "time"

type Fhand struct {
	Hand  string
	Score uint32
	Rank  string
}

type FileMatches struct {
	Matches []libpoker.Anb
}

type FileHands struct {
	Hands []Fhand
}

var matches FileMatches
var hands FileHands
var cores = 1 // runtime.NumCPU() //* 4

func readfile(file string) {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		panic("wrong file")
	}
	json.Unmarshal(dat, &matches)
	// fmt.Println(matches)
}

func match2json(match libpoker.Anb) string {
	s := ""
	s += "{\"alice\":\"" + match.Alice + "\","
	s += "\"bob\":\"" + match.Bob + "\","
	s += "\"result\":" + strconv.Itoa(match.Result) + "}"
	return s
}

func doParallel(core int, c chan int) {
	for i := 0; i < len(matches.Matches)/cores; i++ {
		// fmt.Println(match)
		libpoker.ProcessMatch(&matches.Matches[i*cores+core])
		// fmt.Println(match)
	}
	m := len(matches.Matches) % cores
	b := len(matches.Matches) - m
	for i := 0; i < m; i++ {
		libpoker.ProcessMatch(&matches.Matches[b+i])
	}
	c <- 1
}

func processMatches(file string) {
	readfile(file)

	now := time.Now().UTC().UnixNano()
	var chans = []chan int{}
	for i := 0; i < cores; i++ {
		c := make(chan int)
		chans = append(chans, c)
		go doParallel(i, c)
	}
	for _, c := range chans {
		<-c
	}

	final := time.Now().UTC().UnixNano()
	fmt.Println("total ", len(matches.Matches), "hands, spending", (final-now)/int64(time.Millisecond), "milli seconds, using ", cores, "threads")
	//fmt.Print("{\"matches\": [")
	//for i:=0; i<len(matches.Matches)-1; i++ {
	//  match := matches.Matches[i]
	//  fmt.Println(match2json(match), ",")
	//}
	//match := matches.Matches[len(matches.Matches)-1]
	//fmt.Println(match2json(match)+"]}")
}

func main() {
	libpoker.BuildScoreTbl()
	fmt.Println("len of score table:", libpoker.LenOfScoreTbl())
	//processMatches("../alice_vs_bob/match.json")
	//processMatches("../alice_vs_bob/seven_cards.json")
	//processMatches("../alice_vs_bob/five_cards_with_ghost.json")
	processMatches("../alice_vs_bob/seven_cards_with_ghost.json")
	//rankHands("rank_hands.json")
	// rankHands("rank_hands_no_ghost.json")
}
