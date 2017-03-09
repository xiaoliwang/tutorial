package main
import "fmt"

// slices
var ss = []struct{
    i int
    b bool
} {
    {2, true},
    {3, false},
}

// slices are like references to arrays
func control() {
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