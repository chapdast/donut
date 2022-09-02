package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

func main() {
	width, err := strconv.Atoi(os.Getenv("WIDTH"))
	if err != nil {
		width = 80
	}
	height, err := strconv.Atoi(os.Getenv("HEIGHT"))
	if err != nil {
		height = 24
	}
	// width, height := 80, 24
	hw, hh := float64(width/2), float64(height/2)

	fraction := 50
	size := width * height

	var (
		A, B, i, j float64
	)
	// chars := []rune{'.', ',', '-', '~', ':', ';', '=', '!', '*', '#', '$', '@'}
	chars := []rune{'ه', 'ن', 'ا', 'م', 'ی', 'پ', 'د', 'م', 'ح', 'م', 'ن', 'م'}
	z := make([]float64, size)
	b := make([]rune, size)

	fmt.Printf("\x1b[2J")
	ticker := time.NewTicker(time.Duration(fraction) * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			for i := range b {
				b[i] = ' '
			}
			for i := range z {
				z[i] = 0
			}
			for j = 0; 6.28 > j; j += 0.07 {
				for i = 0; 6.28 > i; i += 0.02 {
					si := math.Sin(i)
					ci := math.Cos(i)

					sj := math.Sin(j)
					cj := math.Cos(j)
					cjj := cj + 2

					sA := math.Sin(A)
					cA := math.Cos(A)

					sB := math.Sin(B)
					cB := math.Cos(B)

					mess := 1 / (si*cjj*sA + sj*cA + 5)
					t := si*cjj*cA - sj*sA
					x := int(hw + (hw-hh/3)*mess*(ci*cjj*cB-t*sB))
					y := int(hh + (hw/2-hh/3)*mess*(ci*cjj*sB+t*cB))
					o := x + width*y
					N := 8 * ((sj*sA-si*cj*cA)*cB - si*cj*sA - sj*cA - ci*cj*sB)
					if y < height && y > 0 && x < width && x > 0 && mess > z[o] {
						z[o] = mess
						b[o] = chars[int(math.Abs(N))]
					}
					// fmt.Printf(" %f %f %d %d %d %f \n", mess, t, x, y, o, N)
				}
			}
			// fmt.Printf("\x1b[2J")
			fmt.Printf("%v", string(b))

			A += 0.04
			B += 0.02
		}

	}

}
