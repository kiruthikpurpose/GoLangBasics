package main

import "fmt"

func main() {
    // Conditional statement with type switch
    var value interface{} = "hello"
    switch v := value.(type) {
    case int:
        fmt.Println("It's an integer:", v)
    case string:
        fmt.Println("It's a string:", v)
    default:
        fmt.Println("Unknown type")
    }
}
