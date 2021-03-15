package main

import (
    "fmt"

    "lib/greet"
    "lib/hello"
    "lib/world"
)

func main() {
    fmt.Println(greet.Now())

    fmt.Print("Hello? or ")
    fmt.Println(greet.Greet())

    fmt.Println(hello.Hello())
    fmt.Println(world.World())
}
