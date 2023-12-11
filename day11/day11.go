package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type galaxy struct {
	X int64
	Y int64
}

func newGalaxy(x, y int64) *galaxy {
	return &galaxy{
		X: x,
		Y: y,
	}
}

func (g *galaxy) minDist(g2 *galaxy) int64 {
	return g.X - g2.X + g.Y - g2.Y
}

func parseLines(lines []string, ex2 bool) int64 {
	lineLen := 0
	galaxies := []*galaxy{}
	rowsGalaxies := map[int64]int64{}
	columnsGalaxies := map[int64]int64{}
	for i, line := range lines {
		lineLen = len(line)
		for j, char := range strings.Split(line, "") {
			if char == "#" {
				galaxies = append(galaxies, newGalaxy(int64(i), int64(j)))
				rowsGalaxies[int64(i)]++
				columnsGalaxies[int64(j)]++
			}
		}
	}

	rowsToExpend := []int64{}
	for i := 0; i < lineLen; i++ {
		if _, ok := rowsGalaxies[int64(i)]; !ok {
			rowsToExpend = append(rowsToExpend, int64(i))
		}
	}

	columnsToExpend := []int64{}
	for i := 0; i < len(lines); i++ {
		if _, ok := columnsGalaxies[int64(i)]; !ok {
			columnsToExpend = append(columnsToExpend, int64(i))
		}
	}
	sort.Slice(rowsToExpend, func(i, j int) bool { return rowsToExpend[i] < rowsToExpend[j] })
	sort.Slice(columnsToExpend, func(i, j int) bool { return columnsToExpend[i] < columnsToExpend[j] })

	return 0
}

func main() {
	file, err := os.Open("day11/test.txt")
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
	fmt.Println(parseLines(lines, false))
}
