package main

import "fmt"

func main() {
    // Looping with a for range loop
    numbers := []int{1, 2, 3, 4, 5}
    for index, num := range numbers {
        fmt.Printf("Index: %d, Number: %d\n", index, num)
    }
}
