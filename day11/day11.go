package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

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

func (g *galaxy) MinDist(g2 *galaxy) int64 {
	return abs(g.X-g2.X) + abs(g.Y-g2.Y)
}

func parseLines(lines []string, expansion int64) int64 {
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

	for i, row := range rowsToExpend {
		for _, gal := range galaxies {
			if gal.X >= row+int64(i)*expansion {
				gal.X += int64(expansion)
			}
		}
	}

	for i, col := range columnsToExpend {
		for _, gal := range galaxies {
			if gal.Y >= col+int64(i)*expansion {
				gal.Y += int64(expansion)
			}
		}
	}

	total := int64(0)
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			total += galaxies[i].MinDist(galaxies[j])
		}
	}

	return total
}

func main() {
	file, err := os.Open("day11/in.txt")
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
	fmt.Println(parseLines(lines, 999999))
}
