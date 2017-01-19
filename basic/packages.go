package main

import (
    "fmt"
    "math/rand"
)

func add(x, y int) int {
    return x + y
}

func swap(x, y string) (string, string) {
    return y, x
}

func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

var i int = 10

func main() {
    c, python, java := true, false, "no!"
    fmt.Println(i, c, python, java)
    fmt.Println(split(16))
    a, b := swap("hello", "world")
    fmt.Println(a, b)
    fmt.Println(add(42, 13))
    fmt.Println("My favorite number is", rand.Intn(10));
}