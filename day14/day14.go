package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func turn(lines [][]string) [][]string {
	turnedLines := [][]string{}
	for i := range lines {
		for j := range lines[i] {
			if i == 0 {
				turnedLines = append(turnedLines, []string{})
			}
			turnedLines[j] = append([]string{lines[i][j]}, turnedLines[j]...)
		}
	}
	return turnedLines
}

func tilt(lines [][]string) [][]string {
	for i := range lines {
		if i == 0 {
			continue
		}
		for j := range lines[i] {
			if lines[i][j] == "O" {
				for k := i - 1; k >= 0; k-- {
					if k == 0 && lines[k][j] == "." {
						lines[i][j] = "."
						lines[k][j] = "O"
					} else if lines[k][j] == "O" || lines[k][j] == "#" {
						lines[i][j] = "."
						lines[k+1][j] = "O"
						break
					}
				}
			}
		}
	}
	return lines
}

func score(lines [][]string) int {
	n := 0

	for i, line := range lines {
		for _, char := range line {
			if char == "O" {
				n += len(lines) - i
			}
		}
	}
	return n
}

func main() {
	file, err := os.Open("day14/in.txt")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	lines := [][]string{}

	for scanner.Scan() {
		lines = append(lines, strings.Split(scanner.Text(), ""))
	}

	fmt.Println(score(tilt(lines)))

	for i := 0; i < 4000; i++ {
		lines = turn(tilt(lines))
	}
	fmt.Println(score(lines))
}
