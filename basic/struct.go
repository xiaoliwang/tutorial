package main

type Vertex struct {
    X, Y int
}

// structs
var (
    v = Vertex {1, 2}
    vp = &v
    v1 = Vertex {X: 1}
)