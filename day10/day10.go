package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func makeGrid(lines []string, ex2 bool) ([][]string, int, int) {
	grid := [][]string{}
	var x, y int
	for i, line := range lines {
		grid = append(grid, strings.Split(line, ""))
		for j, char := range grid[i] {
			if char == "S" {
				x = i
				y = j
			}
		}
	}
	return grid, x, y
}

func considerUp(char string) bool {
	pipes := map[string]struct{}{
		"S": {},
		"|": {},
		"J": {},
		"L": {},
	}
	_, ok := pipes[char]
	return ok

}
func considerDown(char string) bool {
	pipes := map[string]struct{}{
		"S": {},
		"|": {},
		"F": {},
		"7": {},
	}
	_, ok := pipes[char]
	return ok

}
func considerLeft(char string) bool {
	pipes := map[string]struct{}{
		"S": {},
		"-": {},
		"7": {},
		"J": {},
	}
	_, ok := pipes[char]
	return ok

}
func considerRight(char string) bool {
	pipes := map[string]struct{}{
		"S": {},
		"-": {},
		"F": {},
		"L": {},
	}
	_, ok := pipes[char]
	return ok

}
func goUp(char string) bool {
	pipes := map[string]struct{}{
		"S": {},
		"|": {},
		"F": {},
		"7": {},
	}
	_, ok := pipes[char]
	return ok

}

func goDown(char string) bool {
	pipes := map[string]struct{}{
		"S": {},
		"|": {},
		"L": {},
		"J": {},
	}
	_, ok := pipes[char]
	return ok

}

func goLeft(char string) bool {
	pipes := map[string]struct{}{
		"S": {},
		"-": {},
		"F": {},
		"L": {},
	}
	_, ok := pipes[char]
	return ok

}

func goRight(char string) bool {
	pipes := map[string]struct{}{
		"S": {},
		"-": {},
		"7": {},
		"J": {},
	}
	_, ok := pipes[char]
	return ok

}

func moveGrid(grid [][]string, x, y, xPrev, yPrev int) (int, int) {
	if considerUp(grid[x][y]) && !(xPrev == x-1 && yPrev == y) && x > 0 && goUp(grid[x-1][y]) {
		return x - 1, y
	}
	if considerDown(grid[x][y]) && !(xPrev == x+1 && yPrev == y) && x < len(grid)-1 && goDown(grid[x+1][y]) {
		return x + 1, y
	}
	if considerLeft(grid[x][y]) && !(xPrev == x && yPrev == y-1) && y > 0 && goLeft(grid[x][y-1]) {
		return x, y - 1
	}
	if considerRight(grid[x][y]) && !(xPrev == x && yPrev == y+1) && y < len(grid[x])-1 && goRight(grid[x][y+1]) {
		return x, y + 1
	}

	return -1, -1
}

func traversePipes(grid [][]string, x int, y int) int {
	pathLen := 0
	xPrev := x
	yPrev := y
	for pathLen == 0 || grid[x][y] != "S" {
		xTmp, yTmp := moveGrid(grid, x, y, xPrev, yPrev)
		xPrev = x
		yPrev = y
		x = xTmp
		y = yTmp

		pathLen++
	}
	if pathLen%2 == 0 {
		return pathLen / 2
	}
	return pathLen/2 + 1
}

func addBorder(grid [][]string) [][]string {
	for i, row := range grid {
		grid[i] = append(append([]string{"."}, row...), ".")
	}

	grid = append([][]string{strings.Split(
		strings.Repeat(".", len(grid[0])), "")}, grid...)
	grid = append(grid, strings.Split(
		strings.Repeat(".", len(grid[0])), ""))
	return grid
}

func expandRow(row []string) []string {
	expandedRow := []string{}
	for i, s := range row {
		expandedRow = append(expandedRow, s)
		if i < len(row)-1 {
			if considerRight(s) && considerLeft(row[i+1]) {
				expandedRow = append(expandedRow, "-")
			} else {
				expandedRow = append(expandedRow, "o")
			}

		}
	}
	return expandedRow
}

func expandGrid(grid [][]string, expandCols bool) [][]string {
	expandedGrid := [][]string{}
	
	for i, row := range grid {
		expandedGrid = append(expandedGrid, expandRow(row))
		if expandCols && i < len(grid)-1 {
			newRow := []string{}
			for j := range row {
				if considerDown(row[j]) && considerUp(grid[i+1][j]) {
					newRow = append(newRow, "|")
				} else {
					newRow = append(newRow, "o")
				}
			}
		}
	}

	return expandedGrid
}

func floodFill(grid *[][]string, i, j int) {
	(*grid)[i][j] = "x"
	if (*grid)[i][j] == "x" {
		if i > 0 && ((*grid)[i-1][j] == "." || (*grid)[i-1][j] == "o") {
			floodFill(grid, i-1, j)
		}
		if j > 0 && ((*grid)[i][j-1] == "." || (*grid)[i][j-1] == "o") {
			floodFill(grid, i, j-1)
		}
		if i < len(*grid)-1 && ((*grid)[i+1][j] == "." || (*grid)[i+1][j] == "o") {
			floodFill(grid, i+1, j)
		}
		if j < len((*grid)[i])-1 && ((*grid)[i][j+1] == "." || (*grid)[i][j+1] == "o") {
			floodFill(grid, i, j+1)
		}
	}
}

func countTiles(grid [][]string) int {
	num := 0
	for _, row := range grid {
		for _, col := range row {
			if col == "." {
				num++
			}
		}
	}
	return num
}

func main() {
	file, err := os.Open("day10/test2.txt")
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
	grid, x, y := makeGrid(lines, false)
	fmt.Println(traversePipes(grid, x, y))
	grid = expandGrid(addBorder(grid), false)
	floodFill(&grid, 0, 0)

	fmt.Println(countTiles(grid))
}
