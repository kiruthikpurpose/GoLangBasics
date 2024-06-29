package main

import "fmt"

func main() {
    greeting := getGreeting()
    fmt.Println(greeting)
}

func getGreeting() string {
    return "Hello, World!"
}
