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
  sum := 0

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
    for j := 1; j < len(schematic[i])-1; j++ {
      if unicode.IsNumber(schematic[i][j]){
        fmt.Printf("Found number rune: %c\n", schematic[i][j])
        length := findBound(schematic[i],j)
        if checkSurroundings(i,j,schematic){

          fmt.Printf("checkSurroundings hit: %c\n", schematic[i][j])

          // We have a hit that this belongs in the sum
          num, err := strconv.Atoi(string(schematic[i][j:j+length]))
          if err != nil {
            fmt.Println("we tried to convert something wrong")
          }
          sum += num
        }
        // We found a number, we check the length, then continue scanning after it.
        j += length
      }
    }
    fmt.Printf("\n")
  }

  fmt.Printf("Sum: %d\n", sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func checkSurroundings(i int, j int, m [][]rune)(bool){

  length := findBound(m[i], j)
  fmt.Printf("Checking Surroundings for i: %d and j: %d with length: %d\n",i,j,length)

  for k := 0; k < length; k++ {
    // Just check everything in a redundant pattern...
    if k == 0 {
      // beginning
      if isSymbol(m[i-1][j-1]){
        return true
      }
      if isSymbol(m[i][j-1]){
        return true
      }
      if isSymbol(m[i+1][j-1]){
        return true
      }
    }

    // middle
    if isSymbol(m[i-1][j]){
      return true
    }
    if isSymbol(m[i+1][j]){
      return true
    }

    if k == length-1 {
      // end
      if isSymbol(m[i-1][j+1]){
        return true
      }
      if isSymbol(m[i][j+1]){
        return true
      }
      if isSymbol(m[i+1][j+1]){
        return true
      }
    }
  }
  return false
}

func findBound(n []rune, x int)(l int){
  l = 1
  ending := false
  for !ending {
    if n[x+l] == '.' {
      ending = true
    } else {
      l++ 
    }
  }
  return l
}

func isSymbol(r rune)(f bool){
  // &*#@/=$+-%
  return r == '&' || r == '*' || r == '#' || r == '@' || r == '/' || r == '=' || r == '$' || r == '+' || r == '-' || r == '%' 
}
