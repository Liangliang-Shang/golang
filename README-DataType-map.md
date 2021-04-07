## DataType

### map - implements a hash table

* a `nil` map

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
> Q: map and slice are the same reference types, why can we append new elements
>    to nil slice?
>
> A: nil map doesn't point to an initialized map. Assigning value won't re-
>    allocate point address. The append function appends the elements x to
>    the end of the slice s, If the backing array of s is too small to fit all
>    the given values a bigger array will be allocated. The returned slice will
>    point to the newly allocated array.

* make a map or refer to a map literal

```Go
// make an empty map 
//
//      m := make(map[string]float64)
//
// or even with a pre-allocated cap
//
//      m := make(map[string]float64, 10)
//
//  or refer to a map literal
//
//      m := map[string]float64 {
//          "e": 2.71828,
//          "pi": 3.14159,              // the last , is required!!!
//      }
//
m := map[string]float64{}               // m := make(map[string]float64)
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
