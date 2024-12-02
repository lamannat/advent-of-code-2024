package p2_1

import (
	__1 "advent-of-code-2024/src/1.12"
	"fmt"
	"slices"
	"sort"
)

func P2() {
	err, c1, c2 := __1.ReadInputFile("src/1.12/input")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	similarity := 0
	slices.Sort(c1)
	slices.Sort(c2)
	for i := range c1 {
		j := sort.SearchInts(c2, c1[i])
		if c2[j] != c1[i] {
			continue
		}
		k := 0
		for ; (k+j) < len(c2) && c2[k+j] == c1[i]; k++ {
		}
		similarity += k * c1[i]

	}

	fmt.Println("Similarity:", similarity)
}
