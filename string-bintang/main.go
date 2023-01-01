package main

import "fmt"

func generateOutput(n int) {
  for i := 1; i <= n; i++ {
    fmt.Printf("%d**%s\n", i, "".join(str(j) for j in range(n+1, n+9-i)))
  }
}

func main() {
  // input n = 5
  generateOutput(5)

  // input n = 4
  generateOutput(4)
}
