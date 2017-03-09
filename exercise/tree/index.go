package main

import (
    "fmt"
    "math/rand"
)

type Tree struct {
    Left *Tree
    Value int
    Right *Tree
}

func New(k int) *Tree {
    var t *Tree
    // rand.Perm(n int) []int 返回一个长度为n的随机数组
    for _, v := range rand.Perm(10) {
        t = insert(t, (1 + v) * k)
    }
    return t
}

func insert(t *Tree, v int) *Tree {
    if t == nil {
        return &Tree{nil, v, nil}
    }
    if v < t.Value {
        t.Left = insert(t.Left, v)
    } else {
        t.Right = insert(t.Right, v)
    }
    return t
}

func (t *Tree) String() string {
    if t == nil {
        return "()"
    }
    s := ""
    if t.Left != nil {
        s += t.Left.String() + " "
    }
    s += fmt.Sprint(t.Value)
    if t.Right != nil {
        s += " " + t.Right.String()
    }
    return "(" + s + ")"
}

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
    t1 := New(3);
    t2 := New(4);
    same := Same(t1, t2);
    fmt.Println(same);
}