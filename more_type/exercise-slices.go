package main

import "fmt"

//import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	ary := make([][]uint8, dx)
	ary2 := make([]uint8, dy)

	for y := range ary {
		for x := range ary2 {
			ary2[x] = uint8((x + y) / 2)
		}
		ary[y] = ary2
	}

	return ary
}

func main() {
	//pic.Show(Pic(5, 5))
	fmt.Printf("%v", Pic(5, 5))
}
