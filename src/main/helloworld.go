package main

import (
    "lib/hello"     // $GOPATH/src/lib/hello/*.go with `package hello
    "lib/world"     // $GOPATH/src/lib/world/*.go with `package world

    "fmt"
    "strings"
)

func main() {
    fmt.Printf(strings.Join([]string{hello.Hello(), world.World()}, "\n") + "\n")
}
