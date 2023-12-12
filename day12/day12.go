package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toIntSlice(s []string) []int64 {
	nums := make([]int64, len(s))
	for i, num := range s {
		nums[i], _ = strconv.ParseInt(num, 10, 64)
	}
	return nums
}

func filterEmpty(s []string) []string {
	sf := []string{}
	for _, group := range s {
		if group != "" {
			sf = append(sf, group)
		}
	}
	return sf
}

func parseLine(line string, ex2 bool) int64 {
	num := int64(0)
	tmp := strings.Split(line, " ")
	condition := filterEmpty(strings.Split(tmp[0], "."))
	plan := toIntSlice(strings.Split(tmp[1], ","))
	fmt.Println(condition)
	fmt.Println(plan)

	return num
}

func main() {
	file, err := os.Open("day12/test.txt")
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
