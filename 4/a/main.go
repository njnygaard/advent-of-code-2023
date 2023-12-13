package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
  "regexp"
  "math"
)

type Game struct{
  id int
  winners []int
  yours []int
  hits int
  points int
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

  sum := 0

	scanner := bufio.NewScanner(file)

  var games []Game
  games = make([]Game, 0)

	for scanner.Scan() {
		line := []rune(scanner.Text())
    games = append(games, parseInput(line))
	}

  for i:= range games {
    h := check(games[i])
    games[i].hits = h
    games[i].points = powInt(2, games[i].hits-1)
    sum += games[i].points
    fmt.Printf("ID: %d\nYours: %v\nWinners: %v\nHits: %d\n\n", games[i].id, games[i].yours, games[i].winners, games[i].hits)
  }

  fmt.Printf("Sum: %d\n", sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func powInt(x, y int) int {
    return int(math.Pow(float64(x), float64(y)))
}

func check(g Game)(hits int){
  for i:= range g.yours {
    for j:= range g.winners {
      if g.yours[i] == g.winners[j] {
        hits++
      }
    }
  }
  return hits
}

func parseInput(line []rune)(g Game){
  parsed := string(line)
  space := regexp.MustCompile(`\s+`)
  cleaned := space.ReplaceAllString(parsed, " ")
  colon := strings.Split(cleaned, ":")
  gameId := strings.Split(colon[0], " ")
  id, _ := strconv.Atoi(gameId[1])

  numbers := strings.Split(colon[1], "|")
  winners := strings.Split(strings.Trim(numbers[0], " "), " ")
  yours := strings.Split(strings.Trim(numbers[1], " "), " ")

  //fmt.Printf("Winners: %v\n", winners)
  //fmt.Printf("Yours: %v\n", yours)
  winnerInts := make([]int, 0)
  yourInts := make([]int, 0)
  for i := range winners {
    iT,_ := strconv.Atoi(winners[i])
    winnerInts = append(winnerInts, iT)
  }
  for i := range yours {
    iT,_ := strconv.Atoi(yours[i])
    yourInts = append(yourInts, iT)
  }

  g.id = id
  g.winners = winnerInts
  g.yours = yourInts
  return g
}
