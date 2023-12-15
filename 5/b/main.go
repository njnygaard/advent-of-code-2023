package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const MAX_INT_FOUND = 4288015283

type Entry struct {
	// Each line within a map contains three numbers: the destination range start, the source range start, and the range length.
	DestinationRangeStart int
	SourceRangeStart      int
	RangeLength           int
}

type Map struct {
	Name    string
	Entries []Entry
  Map []int
}

func main() {

	var seeds []int

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
    m.Map = make([]int, 0)
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
    s := strings.Split(line, " ") 
    for i := range s{
      num, err := strconv.Atoi(s[i])
      if err != nil {
        fmt.Printf("Error parsing seed %s", s[i])
      }

      if i % 2 == 0 {
        seeds = append(seeds, num)
      }
      if i % 2 == 1 {
        rangeStart := seeds[len(seeds)-1]+1
        for j := 0; j < num; j++ {
          seeds = append(seeds, rangeStart+j)
        }
      }
    }
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

  outputs := make([]int, 0)
  for i := range seeds {
    //fmt.Printf("Seed: %d\n", seeds[i])
    output := process(seeds[i], maps)
    fmt.Printf("Output: %d\n", output)
    outputs = append(outputs, output)
  }

  min := 0
  for i:= range outputs {
    if i == 0 {
      min = outputs[i]
    }
    if outputs[i] < min {
      min = outputs[i]
    }
  }

  fmt.Printf("Lowest: %d\n", min)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func process(s int, m []Map)(int){
  //output := 0
  for i := range m {
    for j := range m[i].Entries {
      //fmt.Printf("Seed: %d  SourceRangeStart: %d  RangeLength: %d\n", s, m[i].Entries[j].SourceRangeStart, m[i].Entries[j].RangeLength)
      if s >= m[i].Entries[j].SourceRangeStart && s < m[i].Entries[j].SourceRangeStart + m[i].Entries[j].RangeLength {
        //fmt.Printf("hit %d\n", s)
        offset := s - m[i].Entries[j].SourceRangeStart
        s = m[i].Entries[j].DestinationRangeStart + offset
        //fmt.Printf("tih %d\n", s)
        break
      }
    }
    //fmt.Printf("tih %d\n", s)
  }
  return s
} 

//func populateMap(m Map)(t []int){
//  t = make([]int, MAX_INT_FOUND)
//  for i := range t {
//    t[i] = i
//  }
//  for i := range m.Entries {
//    for j := m.Entries[i].SourceRangeStart; j < m.Entries[i].RangeLength; j ++ {
//      t[j] = m.Entries[i].DestinationRangeStart + j 
//    }
//  }
//  return t
//}

