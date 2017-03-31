package slice

import (
    "bytes"
)

/**
 * type sliceHeader struct {
 *   Length int
 *   Capacity int
 *   ZerothElement *byte
 * }
 *
 */
func AddOneToEachElement(slice []byte) {
    for i := range slice {
        slice[i]++
    }
}

func SubtractOneFromLength(slice []byte) [] byte {
    slice = slice[0: len(slice) - 1]
    return slice
}

func PtrSubtractOneFromLength(slicePtr *[]byte) {
    slice := *slicePtr
    *slicePtr = slice[0: len(slice) - 1]
}

func (p *path) TruncateAtFinalSlash() {
    i := bytes.LastIndex(*p, []byte("/"))
    if i >= 0 {
        *p = (*p)[0:i]
    }
}

func Extend(slice []int, element int) []int {
    n := len(slice)
    slice = slice[0: n+1]
    slice[n] = element
    return slice
}