package main

import (
    "fmt"

    "golang.org/x/example/stringutil"
    "com.example/greetings"
)

func main() {
    fmt.Println(stringutil.Reverse("Hello"))
    var message string
    message = greetings.Hello("Gladys")
    fmt.Println(message)
}