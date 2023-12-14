package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

  // The input for this one is not really regular.
  // I don't see a problem yet massaging the input in to separate files.
  // We'll see if that bites me in 5b.
	for scanner.Scan() {
		//line := []rune(scanner.Text())
		line := scanner.Text()
    fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func parse(l string){
  
}
