package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
    pic := make([][]uint8, dy)
    for i := 0; i < dx; i++ {
        pic[i] = make([]uint8, dx)
        for j := 0; j < dy; j++ {
            pic[i][j] = uint8(i^j+(i+j)/2)
        }
    }
    return pic
}

func main() {
    pic.Show(Pic)
}