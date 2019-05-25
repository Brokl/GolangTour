package main

import "golang.org/x/tour/pic"
import "math"

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)

	for y := range pic {
		pic = append(pic, []uint8{})

		for x := range pic {
			pic[y] = append(pic[y], uint8(int(math.Ceil(float64(x)))))
		}
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
