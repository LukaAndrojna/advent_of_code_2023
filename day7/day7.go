package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

var cardScores = map[string]int64{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

type camelCardsHand struct {
	Hand         []int64
	Types        map[int64]int64
	HandStrength int64
	Bet          int64
	Jokers       int64
}

func (cch *camelCardsHand) Less(cch2 *camelCardsHand) bool {
	if cch.HandStrength == cch2.HandStrength {
		for i := range cch.Hand {
			if cch.Hand[i] == cch2.Hand[i] {
				continue
			}
			return cch.Hand[i] < cch2.Hand[i]
		}
	}
	return cch.HandStrength < cch2.HandStrength
}

func newCamelCardsHand(hand []string, bet string, ex2 bool) *camelCardsHand {
	types := map[int64]int64{}
	betParsed, _ := strconv.ParseInt(bet, 10, 64)
	handParsed := []int64{}
	jokers := int64(0)
	for _, card := range hand {
		score := cardScores[card]
		if ex2 && score == 11 {
			jokers++
			score = 1
		}
		handParsed = append(handParsed, score)
		types[cardScores[card]]++
	}
	strength := int64(1)
	for _, num := range types {
		if len(types) == 1 {
			strength = 7
		} else if len(types) == 2 {
			if jokers > 0 {
				strength = 7
				break
			}
			if num == 3 || num == 2 {
				strength = 5
				break
			} else {
				strength = 6
				break
			}
		} else if len(types) == 3 {
			if num == 2 {
				strength = 3
				if jokers == 2 {
					strength = 6
				} else if jokers == 1 {
					strength = 5
				}
				break
			} else if num == 3 {
				strength = 4
				if jokers > 0 {
					strength = 6
				}
				break
			}
		} else if len(types) == 4 {
			strength = 2
			if jokers > 0 {
				strength = 4
			}
			break
		} else if len(types) == 5 {
			strength += jokers
			break
		}
	}

	return &camelCardsHand{
		Hand:         handParsed,
		Types:        types,
		HandStrength: strength,
		Bet:          betParsed,
		Jokers:       jokers,
	}
}

func parseLine(line string, ex2 bool) *camelCardsHand {
	tmp := strings.Split(line, " ")
	hand := strings.Split(tmp[0], "")
	return newCamelCardsHand(hand, tmp[1], ex2)
}

func main() {
	file, err := os.Open("day7/in.txt")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	players := []*camelCardsHand{}
	for scanner.Scan() {
		players = append(players, parseLine(scanner.Text(), true))
	}

	sort.Slice(players, func(i, j int) bool {
		return players[i].Less(players[j])
	})

	slices.Reverse(players)

	s := int64(0)
	for i, player := range players {
		s += player.Bet * int64(len(players)-i)
		fmt.Println(player.Hand, "|", player.HandStrength, "|", player.Bet, "|", len(players)-i)
	}

	fmt.Println(s)
}
