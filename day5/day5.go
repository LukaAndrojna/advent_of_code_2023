package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getRange(line string) []int64 {
	nums := strings.Split(line, " ")
	numbers := []int64{}
	for _, num := range nums {
		n, err := strconv.ParseInt(num, 10, 64)
		if err != nil {
			fmt.Println(err)
			return []int64{}
		}

		numbers = append(numbers, n)
	}
	return numbers
}

func getNext(fromMap [][]int64, num int64) int64 {
	for _, m := range fromMap {
		if num >= m[1] && num < m[1] + m[2] {
			return num - m[1] + m[0]
		}
	}
	return num
}

func getLoc(maps map[string][][]int64, seed int64, debug bool) int64 {
	soil := getNext(maps["seedToSoil"], seed)
	fert := getNext(maps["soilToFert"], soil)
	wat := getNext(maps["fertToWat"], fert)
	lig := getNext(maps["watToLig"], wat)
	temp := getNext(maps["ligToTemp"], lig)
	hum := getNext(maps["tempToHum"], temp)
	loc := getNext(maps["humToLoc"], hum)
	
	if debug {
		fmt.Printf("%d > %d > %d > %d > %d > %d > %d > %d\n", seed, soil, fert, wat, lig, temp, hum, loc)
	}
	return loc
}

func parseLines(lines []string, ex2 bool) int64 {
	var seeds []int64
	key := "first"
	keys := map[string]string{
		"seed-to-soil map:":           "seedToSoil",
		"soil-to-fertilizer map:":     "soilToFert",
		"fertilizer-to-water map:":    "fertToWat",
		"water-to-light map:":         "watToLig",
		"light-to-temperature map:":    "ligToTemp",
		"temperature-to-humidity map:": "tempToHum",
		"humidity-to-location map:":    "humToLoc",
	}
	maps := map[string][][]int64{
		"seedToSoil": {},
		"soilToFert": {},
		"fertToWat":  {},
		"watToLig":   {},
		"ligToTemp":  {},
		"tempToHum":  {},
		"humToLoc":   {},
	}

	for _, line := range lines {
		if key == "first" {
			tmp := strings.Split(line, ": ")
			seeds = getRange(tmp[1])
			key = ""
			continue
		}
		if line == lines[1] {
			continue
		}
		if keyTmp, ok := keys[line]; ok {
			key = keyTmp
			continue

		}
		maps[key] = append(maps[key], getRange(line))
	}

	closestLoc := int64(1000000000000)
	if !ex2 {
		for _, seed := range seeds {
			loc := getLoc(maps, seed, false)
			if loc < closestLoc {
				closestLoc = loc
			}
		}
	} else {
		for i, seed := range seeds {
			if i % 2 == 0 {
				for j := int64(0); j < seeds[i+1]; j++ {
					loc := getLoc(maps, seed+j, false)
					if loc < closestLoc {
						closestLoc = loc
					}
				}
			}
		}
	}

	return closestLoc
}

func main() {
	ex2 := true
	file, err := os.Open("day5/day5.txt")
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

	n := parseLines(lines, ex2)

	fmt.Println(n)
}
