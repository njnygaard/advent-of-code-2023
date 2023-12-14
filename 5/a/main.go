package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Entry struct {
	// Each line within a map contains three numbers: the destination range start, the source range start, and the range length.
	DestinationRangeStart int
	SourceRangeStart      int
	RangeLength           int
}

type Map struct {
	Name    string
	Entries []Entry
}

func main() {

	var seeds []string

	// The input for this one is not really regular.
	// I don't see a problem yet massaging the input in to separate files.
	// We'll see if that bites me in 5b.
	maps := make([]Map, 0)

	mapNames := []string{
		//"seeds",
		"seed-to-soil",
		"soil-to-fertilizer",
		"fertilizer-to-water",
		"water-to-light",
		"light-to-temperature",
		"temperature-to-humidity",
		"humidity-to-location",
	}

	for i := range mapNames {
		var m Map
		m.Entries = make([]Entry, 0)
		m.Name = mapNames[i]
		maps = append(maps, m)
	}

	// 'seeds' is of a different format, so handle it outside of the loop
	file, err := os.Open("./inputs/seeds")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		seeds = strings.Split(line, " ")
	}

	for i := range maps {
		file, err := os.Open("./inputs/" + maps[i].Name)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			var e Entry
			line := scanner.Text()
			tokens := strings.Split(line, " ")
			for i := range tokens {
				num, err := strconv.Atoi(tokens[i])
				if err != nil {
					fmt.Printf("Error parsing '%s' to int\n", tokens[i])
				}
				switch i {
				case 0:
					e.DestinationRangeStart = num
				case 1:
					e.SourceRangeStart = num
				case 2:
					e.RangeLength = num
				}
			}
			maps[i].Entries = append(maps[i].Entries, e)
		}
	}

	fmt.Println(seeds)

	for i := range maps {
		fmt.Printf("Name: %s\n", maps[i].Name)

		for j := range maps[i].Entries {
			fmt.Printf("\tDestinationRangeStart: %d\n\tSourceRangeStart: %d\n\tRangeLength: %d\n\n", maps[i].Entries[j].DestinationRangeStart, maps[i].Entries[j].SourceRangeStart, maps[i].Entries[j].RangeLength)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func parse(l string) {

}
