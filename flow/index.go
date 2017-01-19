package main

import (
    "fmt"
    "math"
    "runtime"
    "time"
)

func sqrt(x float64) float64 {
    i, y := 0, x / 2
    for i < 10{
        i++
        y = y - (y * y - x) / (2 * y)
    }
    return y
}

/*func sqrt(x float64) string {
    if x < 0 {
        return sqrt(-x) + "i"
    }
    return fmt.Sprint(math.Sqrt(x))
}*/

func pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
        return v
    } else {
        fmt.Printf("%g >= %g\n", v, lim)
    }
    return lim
}

func c() (i int) {
    defer func() { i++ }()
    return 1
}

func main() {
    defer fmt.Println("the end")
    sum := 0
    for i := 0; i < 10; i++ {
        sum += i
    }
    fmt.Println(sum)

    sum = 1
    for sum < 1000 {
        sum += sum
    }
    fmt.Println(sum)

    // for {} forever loop

    fmt.Println(sqrt(2), sqrt(-4))

    // Both calls to pow are executed and 
    // return before the call to fmt.Println
    fmt.Println(
        pow(3, 2, 10),
        pow(3, 3, 20),
    )

    switch os := runtime.GOOS; os {
    case "darwin": 
        fmt.Println("OS X.")
    case "linux":
        fmt.Println("Linux.")
    default:
        fmt.Printf("%s.\n", os)
    }

    today := time.Now().Weekday()
    switch time.Saturday {
    case today + 0:
        fmt.Println("TOday.")
    case today + 1:
        fmt.Println("Tomorrow.")
    case today + 2:
        fmt.Println("In two days.")
    default:
        fmt.Println("Too far away.") 
    }

    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("Good morning!")
    case t.Hour() < 17:
        fmt.Println("Good afternoon.")
    default:
        fmt.Println("Good evening.")
    }

    fmt.Println("counting")

    for i := 0; i < 10; i++ {
        defer fmt.Println(i)
    }

    fmt.Println("done")

    fmt.Println(c())
}

