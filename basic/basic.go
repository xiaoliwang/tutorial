package main

import (
    "math/cmplx"
)

// 基础类型
// bool
// string
// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr
// float32 float64
// complex64 complex128

// byte : alias for uint8
// rune : alias for int32, represents a Unicode code point

// 默认值
// 0 for numeric types,
// false for the boolean type, and
// "" (the empty string) for strings.

var (
    ToBe bool = false
    MaxInt uint64 = 1 << 64 -1
    z complex128 = cmplx.Sqrt(-5 + 12i)
)

const Pi = 3.14