package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func parseLine(line string) (string, string, string) {
	tmp := strings.Split(strings.Replace(strings.Replace(line, "(", "", -1), ")", "", -1), " = ")
	start := tmp[0]
	next := strings.Split(tmp[1], ", ")
	return start, next[0], next[1]
}

func traverseHuman(instructions []string, maps map[string]map[string]string) int {
	currentNode := "AAA"
	endNode := "ZZZ"

	i := 0
	for currentNode != endNode {
		inst := instructions[i%len(instructions)]
		currentNode = maps[currentNode][inst]
		i++
	}
	return i
}

func findCycles(instructions []string, maps map[string]map[string]string, currentNode string) []int {
	inst := 0
	firstZ := ""
	cycle := []int{}
	for {
		for inst == 0 || currentNode[2] != 'Z' {
			currentNode = maps[currentNode][instructions[inst%len(instructions)]]
			inst++
		}

		cycle = append(cycle, inst)
		if firstZ == "" {
			firstZ = currentNode
		} else if currentNode == firstZ {
			break
		}
	}
	return cycle
}

func traverseGhost(instructions []string, maps map[string]map[string]string, currentNodes []string) int {
	cycles := make([][]int, len(currentNodes))
	for i, node := range currentNodes {
		cycles[i] = findCycles(instructions, maps, node)
	}
	l := cycles[0][0]
	for i := 1; i < len(cycles); i++ {
		l = lcm(l, cycles[i][0])
	}
	return l
}

func main() {
	file, err := os.Open("day8/in.txt")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	instructions := strings.Split(scanner.Text(), "")
	scanner.Scan()

	maps := map[string]map[string]string{}
	currentNodes := []string{}
	for scanner.Scan() {
		start, left, right := parseLine(scanner.Text())
		maps[start] = map[string]string{
			"L": left,
			"R": right,
		}
		if start[2] == 'A' {
			currentNodes = append(currentNodes, start)
		}
	}

	//fmt.Println(traverseHuman(instructions, maps))
	fmt.Println(traverseGhost(instructions, maps, currentNodes))
}
