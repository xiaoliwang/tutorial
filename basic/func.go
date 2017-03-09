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

func initializers() {
    c, python, java := true, false, "no!"
    fmt.Println(c, python, java)
    fmt.Println(split(16))
}

func favoriteNumber() {
    fmt.Println("My favorite number is", rand.Intn(10));
}