package main

import (
  "os"
  "fmt"
  "strconv"
)

func calculateGCD(n1 int, n2 int) (int, error) {
  var a, b int
  if n1 > n2 {
    a = n1
    b = n2
  } else {
    a = n2
    b = n1
  }
  for b != 0 {
    r := a % b
    a = b
    b = r
  }
  if a == 1 {
    return 0, fmt.Errorf("%d and %d are relatively prime", n1, n2)
  }
  return a, nil
}

func main() {
  if len(os.Args) < 2 {
    fmt.Println("give two integers")
    return
  }
  n1, err := strconv.Atoi(os.Args[1]); if err != nil {
    fmt.Printf("%s is not integer\n", os.Args[1])
    return
  }
  n2, err := strconv.Atoi(os.Args[2]); if err != nil {
    fmt.Printf("%s is not integer\n", os.Args[2])
    return
  }
  i, err := calculateGCD(n1, n2); if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(i)
  }
}
