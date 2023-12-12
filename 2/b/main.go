package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// 12 red cubes, 13 green cubes, and 14 blue cubes
const (
	RED_LIMIT   int = 12
	GREEN_LIMIT int = 13
	BLUE_LIMIT  int = 14
)

type Game struct {
	red   int
	green int
	blue  int
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		// Split out Game ID
		game := strings.Split(line, ": ")
		//idString := strings.Split(game[0], " ")[1]
		//id, err := strconv.Atoi(idString)
		//if err != nil {
		//	fmt.Printf("%s doesn't look like a number\n", idString)
		//}

		// Split Games
		games := strings.Split(game[1], ";")

		max_red := 0
		max_green := 0
		max_blue := 0

		for i := range games {
			// 1 red, 2 green, 6 blue
			colors := strings.Split(games[i], ",")
			for j := range colors {
				// 1 red
				t := strings.Trim(colors[j], " ")
				qc := strings.Split(t, " ")
				q, err := strconv.Atoi(qc[0])
				if err != nil {
					fmt.Printf("%s doesn't look like a number\n", qc[1])
				}
				switch qc[1] {
				case "red":
					if q > max_red {
						max_red = q
					}
				case "green":
					if q > max_green {
						max_green = q
					}
				case "blue":
					if q > max_blue {
						max_blue = q
					}
				default:
				}
			}
		}

		sum += max_red * max_blue * max_green

	}

	fmt.Printf("Power Sum: %d\n", sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
