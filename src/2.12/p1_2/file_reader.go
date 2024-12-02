package p1_2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInputFile(name string) (error, [][]int) {
	file, err := os.Open(name)
	if err != nil {
		return err, nil
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	var records [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, " ")
		tmp := make([]int, 0)
		for _, num := range nums {
			n, err := strconv.Atoi(num)
			if err != nil {
				return err, nil
			}
			tmp = append(tmp, n)
		}
		records = append(records, tmp)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return err, nil
	}
	return nil, records
}
