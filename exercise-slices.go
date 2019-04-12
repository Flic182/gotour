package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	outPic := make([][]uint8, dy)

	for outer := range outPic {
		outPic[outer] = make([]uint8, dx)

		for inner := range outPic[outer] {
			outPic[outer][inner] = uint8(inner*outer)
		}
	}

	return outPic
}

func main() {
	pic.Show(Pic)
}
