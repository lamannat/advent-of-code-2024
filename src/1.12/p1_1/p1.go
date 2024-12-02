package p1_1

import (
	"advent-of-code-2024/src/1.12"
	"fmt"
	"math"
	"slices"
)

func P1() {
	err, c1, c2 := __1.ReadInputFile("src/1.12/input")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	slices.Sort(c1)
	slices.Sort(c2)
	distance := 0
	for i := range c1 {
		distance += int(math.Abs(float64(c1[i] - c2[i])))
	}
	fmt.Println("Distance:", distance)
}
