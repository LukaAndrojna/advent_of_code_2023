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

func extrapolate(readings []int64) (int64, int64) {
	if onlyZeroes(readings) {
		return 0, 0
	}

	deltas := []int64{}
	for i := 1; i < len(readings); i++ {
		deltas = append(deltas, readings[i]-readings[i-1])
	}

	h, f := extrapolate(deltas)

	return readings[0] - h, readings[len(readings)-1] + f
}

func parseLine(line string) (int64, int64) {
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

	h := int64(0)
	f := int64(0)
	for scanner.Scan() {
		lh, lf := parseLine(scanner.Text())
		h += lh
		f += lf
	}
	fmt.Println(h, f)
}
