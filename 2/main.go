package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strconv"
  "strings"
  "unicode/utf8"
)

func trimFirstRune(s string) string {
  _, i := utf8.DecodeRuneInString(s)
  return s[i:]
}

func tookenize(input string)(output []int){

 return output
}

func tokenize(input string)(output []int){

  var tokens map[string]int
  tokens = make(map[string]int)

  tokens["0"] = 0
  tokens["1"] = 1
  tokens["2"] = 2
  tokens["3"] = 3
  tokens["4"] = 4
  tokens["5"] = 5
  tokens["6"] = 6
  tokens["7"] = 7 
  tokens["8"] = 8
  tokens["9"] = 9
  tokens["zero"] = 0
  tokens["one"] = 1
  tokens["two"] = 2
  tokens["three"] = 3
  tokens["four"] = 4
  tokens["five"] = 5
  tokens["six"] = 6
  tokens["seven"] = 7
  tokens["eight"] = 8
  tokens["nine"] = 9

  Outer:
  for i,_ := range input {
    for k,v := range tokens {
      if strings.HasPrefix(input[i:], k) {
        output = append(output,v) 
        continue Outer
      }
    }
  }

  return output
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
    fmt.Println(line)
    tokens := tokenize(line)
    fmt.Println(tokens)
   
    var err error
    c := 0

    switch len(tokens) {
    case 0:
      c = 0
      break
    case 1:
      combined := strconv.Itoa(tokens[0]) + strconv.Itoa(tokens[0])
      c, err = strconv.Atoi(combined)
      if err != nil {
        fmt.Printf("Something Happened")
      }
      break
    default:
      combined := strconv.Itoa(tokens[0]) + strconv.Itoa(tokens[len(tokens)-1])
      c, err = strconv.Atoi(combined)
      if err != nil {
        fmt.Printf("Something Happened")
      }
      break
    }

    fmt.Printf("Outer Sum: %d\n", c)
    sum += c
    fmt.Printf("Total Sum: %d\n", sum)
  }

  fmt.Printf("Sum: %d\n", sum)

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
}
