package __1

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInputFile(name string) (error, []int, []int) {
	file, err := os.Open(name)
	if err != nil {
		return err, nil, nil
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	var column1, column2 []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "   ")
		if len(parts) == 2 {
			num1, err1 := strconv.Atoi(parts[0])
			num2, err2 := strconv.Atoi(parts[1])
			if err1 == nil && err2 == nil {
				column1 = append(column1, num1)
				column2 = append(column2, num2)
			} else {
				return errors.New("invalid conversion to number"), nil, nil
			}
		} else {
			fmt.Println("Invalid line:", line)
			return errors.New("invalid line"), nil, nil
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return err, nil, nil
	}
	return nil, column1, column2
}
