package main
import "fmt"

type Vertex struct {
    X int
    Y int
}

// structs
var (
    v = Vertex{1, 2}
    vp = &v
    v2 = Vertex{X: 1}
)

// slices
var ss = []struct{
    i int
    b bool
} {
    {2, true},
    {3, false},
}

func main() {
    i := 42
    p := &i
    fmt.Println(*p)
    *p = 21
    fmt.Println(i)

    vp.X = 1e9 //(*pv).X = 1e9
    fmt.Println(v)
    fmt.Println(v2)

    var a [2]string
    a[0] = "Hello"
    a[1] = "World"
    fmt.Println(a[0], a[1])
    fmt.Println(a)

    // slices are like references to arrays
    primes := [6]int{2, 3, 5, 7, 11, 13}
    var s[]int = primes[1:4]
    s[1] = 109
    fmt.Println(primes)
    fmt.Println(s)

    fmt.Println(ss);

    //var nss []int // nil slices
    for i, v := range s {
        fmt.Printf("i is %d, v is %d\n", i, v);
    }
}
