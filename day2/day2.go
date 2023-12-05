package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseLine(line string, ex2 bool) int64 {
	tmp := strings.Split(line, ": ")
	game := strings.Split(tmp[0], "ame ")[1]
	pulls := strings.Split(tmp[1], "; ")

	var num, r, g, b int64
	for _, pull := range pulls {
		for _, color := range strings.Split(pull, ", ") {
			numCol := strings.Split(color, " ")
			if numCol[1] == "red" {
				rT, _ := strconv.ParseInt(numCol[0], 10, 64)
				if rT > r {
					r = rT
				}
			} else if numCol[1] == "green" {
				gT, _ := strconv.ParseInt(numCol[0], 10, 64)
				if gT > g {
					g = gT
				}
			} else if numCol[1] == "blue" {
				bT, _ := strconv.ParseInt(numCol[0], 10, 64)
				if bT > b {
					b = bT
				}
			}
		}
	}

	if ex2 {
		return r*b*g
	}

	if r <= 12 && g <= 13 && b <= 14 {
		num, _ = strconv.ParseInt(game, 10, 64)
	}
	return num
}

func main() {
	file, err := os.Open("day2/day2.txt")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	s := int64(0)
	for scanner.Scan() {
		s += parseLine(scanner.Text(), true)
	}
	fmt.Println(s)
}
