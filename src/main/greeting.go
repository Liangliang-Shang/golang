package main

import (
    "fmt"

    "lib/greet"
    "lib/hello"
    "lib/world"
)

func main() {
    fmt.Printf(greet.Now() + " Hello? or ")
    fmt.Printf(greet.Greet() + "\n")

    fmt.Println(hello.Hello())
    fmt.Println(world.World())
}
