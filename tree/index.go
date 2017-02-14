package main

import (
    "code.google.com/p/go-tour/tree"
    "fmt"
)

func Walk(t *Tree, ch chan int) {
    if t != nil {
        Walk(t.Left, ch)
        ch <- t.Value
        Walk(t.Right, ch)
    }
}

func Same(t1, t2 *Tree) bool {
    ch1 := make(chan int)
    ch2 := make(chan int)
    go Walk(t1, ch1)
    go Walk(t2, ch2)
    for i := 0; i < 10; i++ {
        if <-ch1 != <-ch2 {
            return false
        }
    }
    return true
}

func main() {

}