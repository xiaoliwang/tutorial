package slice

import (
    "fmt"
    "testing"
)

type path []byte

var buffer [256]byte

func TestSliceHead(t *testing.T) {
    slice := buffer[10: 20]
    for i := 0; i < len(slice); i++ {
        slice[i] = byte(i)
    }

    fmt.Println("before", slice)
    AddOneToEachElement(slice)
    fmt.Println("after", slice)
}

func TestSubSlice(t *testing.T) {
    slice := buffer[10: 20]
    fmt.Println("Before: len(slice) =", len(slice))
    newSlice := SubtractOneFromLength(slice)
    fmt.Println("After:  len(slice) =", len(slice))
    fmt.Println("After:  len(newSlice) =", len(newSlice))
}

func TestPtrSubSlice(t *testing.T) {
    slice := buffer[10: 20]
    fmt.Println("Before: len(slice) =", len(slice))
    PtrSubtractOneFromLength(&slice)
    fmt.Println("After:  len(slice) =", len(slice))
}

func TestAtFinalSlash(t *testing.T) {
    pathName := path("/usr/bin/tso")
    pathName.TruncateAtFinalSlash()
    fmt.Printf("%s\n", pathName)
}

func TestExtend(t *testing.T) {
    var iBuffer [10]int
    slice := iBuffer[0:0]
    for i := 0; i < 10; i++ {
        slice = Extend(slice, i)
        fmt.Println(slice)
    }
}