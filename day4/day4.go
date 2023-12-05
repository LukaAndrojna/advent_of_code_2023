package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func parseLine(line string, ex2 bool) int64 {
	tmp := strings.Split(line, ": ")
	numbers := strings.Split(tmp[1], " | ")
	winning := strings.Split(numbers[0], " ")
	pulled := strings.Split(numbers[1], " ")
	num := float64(-1)

	winningNumbers := map[string]struct{}{}
	for _, wn := range winning {
		if wn == "" {
			continue
		}
		winningNumbers[wn] = struct{}{}
	}

	for _, pn := range pulled {
		if _, ok := winningNumbers[pn]; ok {
			num++
		}
	}

	if ex2 {
		return int64(num) + 1
	}

	if num == -1 {
		return 0
	}
	return int64(math.Pow(2, num))
}

func main() {
	ex2 := true
	file, err := os.Open("day4/day4.txt")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	s := int64(0)
	lines := []string{}
	winnings := []int64{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		winnings = append(winnings, 0)
	}

	for i := len(lines) - 1; i >= 0; i-- {
		n := parseLine(lines[i], ex2)
		if ex2 {
			winnings[i] = n
			for j := i + 1; j <= i+int(n); j++ {
				winnings[i] += winnings[j]
			}
		} else {
			s += n
		}
	}
	if ex2 {
		s = int64(len(lines))
		for _, w := range winnings {
			s += w
		}
	}
	fmt.Println(s)
}
