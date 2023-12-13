package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	schematic := make([][]rune, 0)
	var sum uint64
	sum = 0

	// Could refactor but I decided to do the padding while scanning.
	// Instead of scanning first and then padding later.
	var pad []rune
	firstLine := true
	for scanner.Scan() {
		line := []rune(scanner.Text())
		// Pad with a line of '.'
		if firstLine {
			pad = make([]rune, 0)
			for i := 0; i < len(line)+2; i++ {
				pad = append(pad, '.')
			}
			schematic = append(schematic, pad)
			firstLine = false
		}
		// Pad each line.
		padded := make([]rune, 0)
		padded = append(padded, '.')
		padded = append(padded, line...)
		padded = append(padded, '.')

		// Pad with another line of '.'
		schematic = append(schematic, padded)
	}
	schematic = append(schematic, pad)

	//for i := range schematic {
	//  for j := range schematic[i] {
	//    fmt.Printf("%c", schematic[i][j])
	//  }
	//  fmt.Printf("\n")
	//}

	for i := 1; i < len(schematic)-1; i++ {
		fmt.Printf("Line %d\n", i)
		for j := 1; j < len(schematic[i])-1; j++ {
			if schematic[i][j] == '*' {
				if i == 45 {
					fmt.Printf("line45\n")
				}
				if res, ratio := checkSurroundings(i, j, schematic); res {
					fmt.Printf("Check Surroundings hit at i:%d and j:%d\n", i, j)
					fmt.Printf("Ratio: %d\n", ratio)
					sum += uint64(ratio)
				}
			}
		}
		fmt.Printf("\n")
	}

	fmt.Printf("Sum: %d\n", sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func extractNumber(i int, j int, m [][]rune) int {
	runes := make([]rune, 0)

	runes = append(runes, m[i][j])

	jb := j - 1
	for unicode.IsNumber(m[i][jb]) {
		runes = append([]rune{m[i][jb]}, runes...)
		jb--
	}
	jf := j + 1
	for unicode.IsNumber(m[i][jf]) {
		runes = append(runes, m[i][jf])
		jf++
	}

	//fmt.Printf("runes: %v\n", string(runes))
	num, err := strconv.Atoi(string(runes))
	if err != nil {
		fmt.Printf("fuck\n")
	}
	//fmt.Printf("num: %d\n", num)
	return num
}

func checkSurroundings(i int, j int, m [][]rune) (bool, int) {

	gears := make([]int, 0)
	// top
	if isNumber(m[i-1][j-1]) {
		gears = append(gears, extractNumber(i-1, j-1, m))
		goto Middle
	}
	if isNumber(m[i-1][j]) {
		gears = append(gears, extractNumber(i-1, j, m))
		goto Middle
	}
	if isNumber(m[i-1][j+1]) {
		gears = append(gears, extractNumber(i-1, j+1, m))
		goto Middle
	}

Middle:
	// middle
	if isNumber(m[i][j-1]) {
		gears = append(gears, extractNumber(i, j-1, m))
	}
	if isNumber(m[i][j+1]) {
		gears = append(gears, extractNumber(i, j+1, m))
	}

	// bottom
	if isNumber(m[i+1][j-1]) {
		gears = append(gears, extractNumber(i+1, j-1, m))
		goto Processing
	}
	if isNumber(m[i+1][j]) {
		gears = append(gears, extractNumber(i+1, j, m))
		goto Processing
	}
	if isNumber(m[i+1][j+1]) {
		gears = append(gears, extractNumber(i+1, j+1, m))
		goto Processing
	}
	//// beginning
	//if isNumber(m[i-1][j-1]) {
	//	gears = append(gears, extractNumber(i-1, j-1, m))
	//}
	//if isNumber(m[i][j-1]) {
	//	gears = append(gears, extractNumber(i, j-1, m))
	//}
	//if isNumber(m[i+1][j-1]) {
	//	gears = append(gears, extractNumber(i+1, j-1, m))
	//}

	//// middle
	//if isNumber(m[i-1][j]) {
	//	gears = append(gears, extractNumber(i-1, j, m))
	//}
	//if isNumber(m[i+1][j]) {
	//	gears = append(gears, extractNumber(i+1, j, m))
	//}

	//// end
	//if isNumber(m[i-1][j+1]) {
	//	gears = append(gears, extractNumber(i-1, j+1, m))
	//}
	//if isNumber(m[i][j+1]) {
	//	gears = append(gears, extractNumber(i, j+1, m))
	//}
	//if isNumber(m[i+1][j+1]) {
	//	gears = append(gears, extractNumber(i+1, j+1, m))
	//}

Processing:
	ratio := 0
	fmt.Printf("Non deduped gears: %v\n", gears)
	// ToDo: Removing duplicats like 44*44 gear on line 45 of input.
	//gears = removeDuplicateInt(gears)
	if len(gears) > 1 {
		fmt.Printf("Gears: %v\n", gears)
		ratio = gears[0] * gears[1]
	}
	return ratio > 1, ratio
}

func removeDuplicateInt(intSlice []int) []int {
	allKeys := make(map[int]bool)
	list := []int{}
	for _, item := range intSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func findBound(n []rune, x int) (l int) {
	l = 1
	ending := false
	for !ending {
		if n[x+l] == '.' || isSymbol(n[x+l]) {
			ending = true
		} else {
			l++
		}
	}
	return l
}

func isNumber(r rune) (f bool) {
	return unicode.IsNumber(r)
}

func isSymbol(r rune) (f bool) {
	// &*#@/=$+-%
	return r == '*'
}
