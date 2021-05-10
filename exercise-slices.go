package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	a := make([][]uint8, dy)
	for i := range a {
		a[i] = make([]uint8, dx)
	}
	for i := range a {
		for j := range a {		
			a[i][j]=uint8(j * 1000 / (i + 1))
		}
	}
	return a
}

func main() {
	pic.Show(Pic)
}