package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func onlyZeroes(s []int64) bool {
	if len(s) < 2 {
		return true
	}

	for _, n := range s {
		if n != 0 {
			return false
		}
	}

	return true
}

func extrapolate(readings []int64) int64 {
	if onlyZeroes(readings) {
		return 0
	}

	deltas := []int64{}
	for i := 1; i < len(readings); i++ {
		deltas = append(deltas, readings[i]-readings[i-1])
	}

	return readings[len(readings)-1] + extrapolate(deltas)
}

func parseLine(line string, ex2 bool) int64 {
	tmp := strings.Split(line, " ")
	readings := make([]int64, len(tmp))
	for i, n := range tmp {
		ni, _ := strconv.ParseInt(n, 10, 64)
		readings[i] = ni
	}

	return extrapolate(readings)
}

func main() {
	file, err := os.Open("day9/in.txt")
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
