package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func replace_words(line string) string {
	dict := map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}
	for key, val := range dict {
		line = strings.Replace(line, key, val, -1)
	}
	return line
}

func parseLine(line string) int64 {
	r := regexp.MustCompile(`[0-9]`)
	digits := r.FindAllString(replace_words(line), -1)
	a, _ := strconv.ParseInt(digits[0], 10, 64)
	b, _ := strconv.ParseInt(digits[len(digits)-1], 10, 64)

	return a*10 + b
}

func main() {
	file, err := os.Open("day1.txt")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	s := int64(0)
	for scanner.Scan() {
		s += parseLine(scanner.Text())
	}
	fmt.Println(s)
}
