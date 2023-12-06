package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getNumbers(line string, ex2 bool) []int64 {
	tmp := strings.Split(line, ": ")
	if ex2 {
		tmp[1] = strings.Replace(tmp[1], " ", "", -1)
	}
	nums := strings.Split(tmp[1], " ")
	numbers := []int64{}
	for _, num := range nums {
		if num == "" {
			continue
		}
		n, err := strconv.ParseInt(num, 10, 64)
		if err != nil {
			fmt.Println(err)
			return []int64{}
		}

		numbers = append(numbers, n)
	}
	return numbers
}

func wonRace(time, timeHeld, dist int64) bool {
	return timeHeld*(time-timeHeld) > dist
}

func runRaces(times []int64, dists []int64) int64 {
	n := int64(1)
	for i, time := range times {
		l := time
		h := int64(0)
		for j := int64(0); j < time; j++ {
			won := wonRace(time, time-j, dists[i])
			if won && l == time {
				l = j
			}
			if l < time && !won {
				h = j
				break
			}
		}
		n *= h-l
	}
	return n
}

func main() {
	ex2 := true
	file, err := os.Open("day6/day6.txt")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	times := getNumbers(scanner.Text(), ex2)

	scanner.Scan()
	dists := getNumbers(scanner.Text(), ex2)

	n := runRaces(times, dists)

	fmt.Println(n)
}
