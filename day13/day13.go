package main

import (
	"bufio"
	"fmt"
	"os"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func isPalindrome(s string, i int) bool {
	for j := 0; i-j >= 0 && i+j+1 < len(s); j++ {
		if s[i-j] != s[i+j+1] {
			return false
		}
	}
	return true
}

func horizontalMatch(pattern []string, i int) bool {
	for j := 0; i-j >= 0 && i+j+1 < len(pattern); j++ {
		if pattern[i-j] != pattern[i+j+1] {
			return false
		}
	}

	return true
}

func findRelfections(pattern []string) int {
	vertRef := map[int]int{}
	horRef := map[int]int{}

	for i, row := range pattern {
		for j := range row {
			if j < len(row)-1 && isPalindrome(row, j) {
				vertRef[1+j]++
			}
		}
		if i < len(pattern)-1 && horizontalMatch(pattern, i) {
			horRef[1+i]++
		}
	}

	num := 0
	for key, val := range vertRef {
		if val == len(pattern) {
			num += key
		}
	}

	for key, val := range horRef {
		if val > 0 {
			num += key * 100
		}
	}
	return num
}

func main() {
	file, err := os.Open("day13/in.txt")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	i := 0
	patterns := [][]string{
		{},
	}

	for scanner.Scan() {
		s := scanner.Text()
		if s == "\n" || s == "" {
			i++
			patterns = append(patterns, []string{})
			continue
		}
		patterns[i] = append(patterns[i], s)
	}
	n := 0
	for _, pattern := range patterns {
		n += findRelfections(pattern)
	}
	fmt.Println(n)
}
