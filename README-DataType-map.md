## DataType

### map

```Go
var m map[string]float64        // nil map of string-float64 pairs
m["pi"] = 3.14159               // panic: assignment to entry in nil map
```

> This variable m is a map of string keys to int values:
>
>       var m map[string]int
>
> Map types are reference types, like pointers or slices, and so the value of m
> above is nil; it doesn't point to an initialized map. A nil map behaves like 
> an empty map when reading, but attempts to write to a nil map will cause a 
> runtime panic; don't do that. To initialize a map, use the built in make 
> function:
>
>       m = make(map[string]int)
>

```Go
m := make(map[string]float64)           // empty map
m["pi"] = 3.14                          // add a new key-value pair
m["pi"] = 3.1415926                     // update value
v := m["pi"]                            // get value
v = m["pie"]                            // not found: v == 0 (zero value)

_, isin := m["pi"]                      // isin == true
_, isin = m["pie"]                      // isin == false

delete(m, "pi")                         // delete a key-value pair

fmt.Println(m)                          // print map: "map[]"

for key, val := range m {               // order not specified
    fmt.Println(key, val)
}
```

```Go
m2 := make(map[string]float64, 100)     // preallocate 100 cap
m3 := map[string]float64 {              // map literal
    "e": 2.71828,
    "pi": 3.14159,
}
```

```
// variable declaration
var str string              // with auto-inited zero value, e.g. "", 0, 0.0

// var pi float32 = 3.14    // type could be inferred by the init value 
var pi = 3.14               // even in shorthand: pi := 3.14

var (
    name    = "Good morning!"
    age     = 35
    height  = 1.60
) 

// short hand declaration
first_name, last_name, age := "John", "Doe", 35

var price float32
```

### []type{}, map[] && built-in

```go
var int_seq = []int{0, 1, 2, 3, 4, 5}
var sum int
for index, value := range int_seq {
    fmt.Print(index, " -> ", value, " ")
    sum = sum + value
    fmt.Println("Sum:", sum)
}

m := map[string]string{"a": "A", "b": "B"}
for k, v := range m {
    fmt.Println(k, "->", v)
}
```

## Control flow

### if/else

```Go
    a := [...]int{-1, 0, 1, 2}

    for _, v := range &a {
        if v < 0 {
            fmt.Printf("%2d negative\n", v)
        } else if v > 0 {
            if v % 2 == 0 {
                fmt.Printf("%2d positive and even\n", v)
            } else {
                fmt.Printf("%2d positive and odd\n", v)
            }
        } else {
            fmt.Printf("%2d 0\n", v)
        }
    }

/*
 * -1 negative
 *  0 0
 *  1 positive and odd
 *  2 positive and even
 *
 */
```

### switch/case

```Go
    switch time.Now().Weekday().String() {
        case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
            fmt.Println("It's time to learn some Go.")
        default:
            fmt.Println("It's weekend, time to rest!")
    }

    fmt.Println(time.Now().Weekday().String())

    letter := "i"
    switch letter {
        case "a", "e", "i", "o", "u":
            fmt.Printf("%s vowel\n", letter)
        default:
            fmt.Println("%s not a vowel\n", letter)
    }

```

## function

### anonymous

```Go
    // define a function and then call it with ()
    func(a, b int) {
        fmt.Printf("Sum of %d and %d: %d\n", a, b, a+b)
    }(99, 1)

    // define a function and assign it to a variable
    f := func(a, b int) {
        fmt.Printf("Sum of %d and %d: %d\n", a, b, a+b)
    }
    
    // call the function with the variable
    f(99, 1)
```

### user-defined function type

```Go
// function type, take two int and return an int
type add func(a, b int) int

    var f add = func(a, b int) int {
        return a + b
    }
    
    fmt.Println(f(99, 1))
```

### Higher-order function
  * takes one or more functions as arguments
  * returns a function as its result

```Go

```
