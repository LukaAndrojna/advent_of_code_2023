package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isNumeric(char rune) bool {
	numSet := map[rune]struct{}{
		'0': {},
		'1': {},
		'2': {},
		'3': {},
		'4': {},
		'5': {},
		'6': {},
		'7': {},
		'8': {},
		'9': {},
	}

	_, ok := numSet[char]
	return ok
}

func isSymbol(char string) bool {
	symSet := map[string]struct{}{
		"=": {},
		"*": {},
		"+": {},
		"/": {},
		"&": {},
		"#": {},
		"-": {},
		"%": {},
		"$": {},
		"@": {},
	}

	_, ok := symSet[char]
	return ok
}

func engineParts(lines []string) int64 {
	partNumbers := []int64{}
	for i, line := range lines {
		firstIndex := -1
		lastIndex := -1
		isPart := false
		for j, char := range line {
			if isNumeric(char) {
				if firstIndex == -1 {
					firstIndex = j
					lastIndex = j

					if j > 0 && isSymbol(line[j-1:j]) ||
						j > 0 && i > 0 && isSymbol(lines[i-1][j-1:j]) ||
						j > 0 && i < len(lines)-1 && isSymbol(lines[i+1][j-1:j]) ||
						i > 0 && isSymbol(lines[i-1][j:j+1]) ||
						i < len(lines)-1 && isSymbol(lines[i+1][j:j+1]) ||
						j < len(line)-1 && isSymbol(line[j+1:j+2]) ||
						j < len(line)-1 && i > 0 && isSymbol(lines[i-1][j+1:j+2]) ||
						j < len(line)-1 && i < len(lines)-1 && isSymbol(lines[i+1][j+1:j+2]) {
						isPart = true
					}
				} else {
					lastIndex = j

					if j > 0 && isSymbol(line[j-1:j]) ||
						j > 0 && i > 0 && isSymbol(lines[i-1][j-1:j]) ||
						j > 0 && i < len(lines)-1 && isSymbol(lines[i+1][j-1:j]) ||
						i > 0 && isSymbol(lines[i-1][j:j+1]) ||
						i < len(lines)-1 && isSymbol(lines[i+1][j:j+1]) ||
						j < len(line)-1 && isSymbol(line[j+1:j+2]) ||
						j < len(line)-1 && i > 0 && isSymbol(lines[i-1][j+1:j+2]) ||
						j < len(line)-1 && i < len(lines)-1 && isSymbol(lines[i+1][j+1:j+2]) {
						isPart = true
					}
				}
			}
			if firstIndex > -1 && (!isNumeric(char) || j == len(line)-1) {
				if isPart {
					num, err := strconv.ParseInt(line[firstIndex:lastIndex+1], 10, 64)
					if err != nil {
						fmt.Println(err)
						return 0
					}
					partNumbers = append(partNumbers, num)
				}
				firstIndex = -1
				lastIndex = -1
				isPart = false
			}
		}
	}

	s := int64(0)
	for _, n := range partNumbers {
		s += n
	}
	return s
}

func readNumber(line string, i int) int64 {
	firstIndex := i
	lastIndex := i

	runes := []rune(line)
	for j := i; j >= 0; j-- {
		if isNumeric(runes[j]) {
			firstIndex = j
		} else {
			break
		}
	}

	for j := i; j < len(runes); j++ {
		if isNumeric(runes[j]) {
			lastIndex = j
		} else {
			break
		}
	}

	num, err := strconv.ParseInt(line[firstIndex:lastIndex+1], 10, 64)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return num
}

func charToRune(char string) rune {
	runes := []rune(char)
	return runes[0]
}

func gearRatios(lines []string) int64 {
	totalGearRatio := int64(0)
	for i, line := range lines {
		for j, char := range line {
			nums := []int64{}
			t := false
			b := false
			if char == '*' {
				if j > 0 && isNumeric(charToRune(line[j-1:j])) {
					nums = append(nums, readNumber(line, j-1))
				}
				if i > 0 && isNumeric(charToRune(lines[i-1][j:j+1])) {
					nums = append(nums, readNumber(lines[i-1], j))
					t = true
				}
				if i < len(lines)-1 && isNumeric(charToRune(lines[i+1][j:j+1])) {
					nums = append(nums, readNumber(lines[i+1], j))
					b = true
				}
				if !t && j > 0 && i > 0 && isNumeric(charToRune(lines[i-1][j-1:j])) {
					nums = append(nums, readNumber(lines[i-1], j-1))
				}
				if !b && j > 0 && i < len(lines)-1 && isNumeric(charToRune(lines[i+1][j-1:j])) {
					nums = append(nums, readNumber(lines[i+1], j-1))
				}
				if j < len(line)-1 && isNumeric(charToRune(line[j+1:j+2])) {
					nums = append(nums, readNumber(line, j+1))
				}
				if !t && j < len(line)-1 && i > 0 && isNumeric(charToRune(lines[i-1][j+1:j+2])) {
					nums = append(nums, readNumber(lines[i-1], j+1))
				}
				if !b && j < len(line)-1 && i < len(lines)-1 && isNumeric(charToRune(lines[i+1][j+1:j+2])) {
					nums = append(nums, readNumber(lines[i+1], j+1))
				}
			}
			if len(nums) == 2 {
				totalGearRatio += nums[0] * nums[1]
			}
		}
	}
	return totalGearRatio
}

func main() {
	file, err := os.Open("day3/day3.txt")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	fmt.Println(engineParts(lines))
	fmt.Println(gearRatios(lines))
}
