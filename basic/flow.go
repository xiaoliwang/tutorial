package type

import (
    "math"
    "runtime"
    "time"
    "fmt"
    "math/rand"
)

func add(x, y int) int {
    return x + y
}

func swap(x, y string) (string, string) {
    return y, x
}

func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

func initializers() {
    c, python, java := true, false, "no!"
    fmt.Println(c, python, java)
    fmt.Println(split(16))
}

func favoriteNumber() {
    fmt.Println("My favorite number is", rand.Intn(10));
}

// for {} forever loop

func sqrt(x float64) float64 {
    if x < 0 {
        return sqrt(-x)
    }
    i, y := 0, x / 2
    for i < 10 {
        i++
        y = y - (y * y - x) / (2 * y)
    }
    return y
}

func pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
        return v
    }
    return lim
}

// switch
func getOs() string {
    var myos string
    switch os := runtime.GOOS; os {
    case "darwin":
        myos = "OS X."
    case "linux":
        myos = "linux."
    default:
        myos = os + ".\n"
    }
    return myos;
}

// switch
func getSaturday() (date string) {
    today := time.Now().Weekday()
    switch time.Saturday {
    case today + 0:
        date = "TOday."
    case today + 1:
        date = "Tomorrow."
    case today + 2:
        date = "In two days."
    default:
        date = "Too far away."
    }
    return 
}