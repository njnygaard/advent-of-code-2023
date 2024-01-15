package main

import (
	"bufio"
  "os"
  "fmt"
  "strings"
  "log"
)

func main() {

	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
    s := strings.Split(line, " ") 
    for i := range s{
      //num, err := strconv.Atoi(s[i])
      //if err != nil {
      //  fmt.Printf("Error parsing seed %s", s[i])
      //}
      fmt.Printf("%s",s[i])
    }
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

