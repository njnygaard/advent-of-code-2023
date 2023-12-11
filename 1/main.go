package main

import (
  "bufio"
  "fmt"
  "strconv"
  "log"
  "os"
  "unicode"
)

func main() {
  file, err := os.Open("./input")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  sum := 0
  // optionally, resize scanner's capacity for lines over 64K, see next example
  for scanner.Scan() {
    line := scanner.Text()
    fmt.Println(line)

    // I want to try and do this in one pass...
    //var m map[string][]int
    //m = make(map[string][]int)

    first := -1
    last := -1

    for _,r := range line {
      if unicode.IsDigit(r) {
        if first == -1 {
          i, err := strconv.Atoi(string(r))
          if err != nil {
            fmt.Printf("%q is not a digit I recognize", r)
          }
          first = i
          continue
        }
        //last = int(r)
        i, err := strconv.Atoi(string(r))
        if err != nil {
          fmt.Printf("%q is not a digit I recognize", r)
        }
        last = i
        //fmt.Printf("%q looks like a number.\n", r)
        //fmt.Printf("First: %d   Last: %d\n", first, last)
      }
    }
    if last == -1 {
      last = first
    }
    combined := strconv.Itoa(first) + strconv.Itoa(last)
    c, err := strconv.Atoi(combined)
    if err != nil {
      fmt.Printf("Something Happened")
    }
    sum += c
    fmt.Printf("First: %d   Last: %d\n", first, last)
  }

  fmt.Printf("Sum: %d\n", sum)

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
}
