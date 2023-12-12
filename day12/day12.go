package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cache = map[string]int64{}

func toIntSlice(s []string) []int64 {
	nums := make([]int64, len(s))
	for i, num := range s {
		nums[i], _ = strconv.ParseInt(num, 10, 64)
	}
	return nums
}

func makeKey(condition string, plan []int64) string {
	return condition + fmt.Sprint(plan)
}

func countConf(condition string, plan []int64) int64 {
	num := int64(0)
	if condition == "" {
		if len(plan) == 0 {
			return 1
		}
		return 0
	}

	if len(plan) == 0 {
		if strings.Contains(condition, "#") {
			return 0
		}
		return 1
	}

	key := makeKey(condition, plan)
	if n, ok := cache[key]; ok {
		return n
	}

	if condition[0] == '.' || condition[0] == '?' {
		num += countConf(condition[1:], plan)
	}
	if condition[0] == '#' || condition[0] == '?' {
		if plan[0] <= int64(len(condition)) &&
		!strings.Contains(condition[:plan[0]], ".") {
			if plan[0] == int64(len(condition)) {
				num += countConf(condition[plan[0]:], plan[1:])
			} else if condition[plan[0]] != '#' {
				num += countConf(condition[plan[0]+1:], plan[1:])
			}
		}

	}

	cache[key] = num
	return num
}

func parseLine(line string, ex2 bool) int64 {
	tmp := strings.Split(line, " ")
	condition := tmp[0]
	plan := toIntSlice(strings.Split(tmp[1], ","))

	if ex2 {
		condition = strings.Repeat("?"+condition, 5)[1:]
		tmpPlan := []int64{}
		for i := 0; i < 5; i++ {
			tmpPlan = append(tmpPlan, plan...)
		}
		plan = tmpPlan
	}

	return countConf(condition, plan)
}

func main() {
	file, err := os.Open("day12/in.txt")
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
