# The Go programming language

## env 
### go && the basic directory structure
```bash
$ go version
go version go1.13.8 linux/amd64

#@    WorkSpace/Project: golang    @#
$ tree /home/lshang/pro/golang
/home/lshang/pro/golang
├── README.md
├── bin
└── src

2 directories, 1 file
```

### go run/build
```go
$ cat src/helloworld/main.go
package main

import "fmt"

func main() {
    fmt.Printf("Hello World!\n")
}

$ go run src/helloworld/main.go
Hello World!

$ go build src/helloworld/main.go   # build out main - an binary exe file

$ tree /home/lshang/pro/golang
/home/lshang/pro/golang
├── README.md
├── bin
├── main                            # newly built binary exe file
└── src
    └── helloworld
        └── main.go

3 directories, 3 files

$ /home/lshang/pro/golang/main
Hello World!
```

### go install in $GOBIN
```go
$ export GOBIN=/home/lshang/pro/golang/bin

$ go install src/helloworld/main.go

$ tree
.
├── README.md
├── bin
│   └── main                        # the binary built and installed into bin
└── src
    └── helloworld
        └── main.go

3 directories, 3 files

$ bin/main
Hello World!
```

### package/lib/module in $GOPATH
```go
$ export GOPATH=/home/lshang/pro/golang

$ tree
.
├── README.md
├── bin
│   └── main
└── src
    ├── helloworld
    │   └── main.go
    ├── lib
    │   ├── hello
    │   │   └── hello.go
    │   └── world
    │       └── world.go
    └── main
        └── helloworld.go

7 directories, 6 files

$ cat src/lib/hello/hello.go
package hello

func Hello() string {
    return "Hello from Hello() in package hello"
}

$ cat src/lib/world/world.go
package world

func World() string {
    return "World from World() in package world"
}

$ cat src/main/helloworld.go
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

$ go run src/main/helloworld.go
Hello from Hello() in package hello
World from World() in package world

$ go install src/main/helloworld.go

$ ls -l bin
total 3972
-rwxrwxr-x 1 lshang lshang 2038088 Mar 14 22:43 helloworld
-rwxrwxr-x 1 lshang lshang 2025490 Mar 14 22:18 main

$ bin/helloworld
Hello from Hello() in package hello
World from World() in package world
```

### diff *.go files in the same package
```go
$ export GOPATH=/home/lshang/pro/golang

$ cat src/lib/greet/morning.go
package greet

func Morning() string {
    return "Good morning!"
}

$ cat src/lib/greet/afternoon.go
package greet

func Afternoon() string {
    return "Good afternoon!"
}

$ cat src/lib/greet/evening.go
package greet

func Evening() string {
    return "Good evening!"
}

$ cat src/lib/greet/greet.go
package greet

import "time"

var t = time.Now()


func Now() string {
    return "Now it is " + t.Format("15:04:05") + "."
}


func Greet() string {

    switch {
        default:
            return Now() + " Hello?"

        case t.Hour() < 12:
            return Morning()

        case t.Hour() < 17:
            return Afternoon()

        case t.Hour() >= 17:
            return Evening()

    }
}

$ cat src/main/greeting.go
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

$ go run src/main/greeting.go
Now it is 21:36:54.
Hello? or Good evening!
Hello from Hello() in package hello
World from World() in package world
```
