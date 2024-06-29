package main

import "fmt"

func main() {
    // Arrays, slices, and maps
    arr := [...]int{1, 2, 3, 4, 5}
    fmt.Println("Array:", arr)

    slice := arr[1:4]
    fmt.Println("Slice:", slice)

    // Map example
    person := make(map[string]int)
    person["age"] = 30
    person["height"] = 180
    fmt.Println("Person:", person)
}
