package main

import "fmt"

func main() {
    // Simple calculation with pointers
    a, b := 5, 3
    sum := add(&a, &b)
    fmt.Printf("Sum of %d and %d is %d\n", a, b, sum)
}

func add(x, y *int) int {
    return *x + *y
}
