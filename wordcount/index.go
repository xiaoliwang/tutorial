package main

import (
    "fmt"
    "strings"
)

func WordCount(s string) map[string]int {
    ss := strings.Fields(s)
    m := make(map[string]int)
    for _, v := range ss {
        m[v]++
    }
    fmt.Println(m);
    return m
}

func main() {
    s := "Hello world Hello world"
    WordCount(s)
}
