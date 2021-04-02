package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	pixels := make([][]uint8, dy)
	data := make([]uint8, dx)

	for i := range pixels {
		for j := range data {
			data[i] = uint8(i * j)
		}
		pixels[i] = data
	}
	return pixels
}

func main() {
	pic.Show(Pic)
}
