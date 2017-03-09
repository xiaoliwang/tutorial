package main
import "fmt"

func main() {
    m2 = make(map[string]Vertex)
    m2["Bell Labs"] = Vertex{
        4, -7,
    }
    fmt.Println(m2)
    favoriteNumber()
    date := getSaturday()
    fmt.Println(date)
}