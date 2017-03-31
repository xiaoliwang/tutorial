package basic

import (
    "math/cmplx"
)

// 基础类型
// bool
// string
// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr
// float32 float64
// complex64 complex128 complex为复数

// byte : alias for uint8
// rune : alias for int32, represents a Unicode code point

// 默认值
// 0 for numeric types,
// false for the boolean type, and
// "" (the empty string) for strings.

var (
    ToBe bool = false
    MaxInt uint64 = 1 << 64 - 1
    z complex128 = cmplx.Sqrt(-5 + 12i)
)

const Pi = 3.14

// 高级类型
// struct
type Vertex struct {
    X, Y int
}

var (
    v = Vertex {1, 2}
    vp = &v
    v1 = Vertex {X: 1}
)

// array
var a [10]int

// slice
var ss = []struct{
    i int
    b bool
} {
    {2, true},
    {3, false},
}

// slices are like references to arrays
func create_slices() {
    primes := [6]int{2, 3, 5, 7, 11, 13}
    var s[]int = primes[1:4]
    s[1] = 109
    fmt.Println(primes)
    fmt.Println(s)

    //var nss []int // nil slices
    for i, v := range s {
        fmt.Printf("i is %d, v is %d\n", i, v)
    }
}

// map

var m1 = map[string]Vertex {
    "Bell Labs": Vertex {
        1, -7,
    },
    "Google": {
        3, -1,
    },
}

var m2 map[string]Vertex

func create_map() {
    m2 = make(map[string]Vertex)
    m2["Bell Labs"] = Vertex {
        4, -7,
    }
    fmt.Println(m2)
}