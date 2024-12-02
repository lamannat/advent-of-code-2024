package p1_2

import (
	"fmt"
)

const (
	maxLevelJump = 3
)

func P1() {
	err, records := ReadInputFile("src/2.12/p1_2/input")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	rc := countSafe(records)
	fmt.Println("safe records:", rc)
}

func countSafe(records [][]int) int {
	safeCount := 0
	for i := 0; i < len(records); i++ {
		if recordIsSafeBruteForceAllLevels(records[i]) {
			if !recordIsSafe(records[i]) {
				fmt.Println("algorithm bug detected:", records[i])
			}
			safeCount++
		}
	}
	return safeCount
}

func recordIsSafe(record []int) bool {
	unsafeLevel := findUnsafeLevel(record)
	if unsafeLevel == -1 {
		return true
	}
	fixedRecord1 := dumpUnsafeLevel(record, unsafeLevel)
	fixedRecord2 := dumpUnsafeLevel(record, unsafeLevel-1)
	unsafeLevel = findUnsafeLevel(fixedRecord1)
	if unsafeLevel == -1 {
		return true
	}
	unsafeLevel = findUnsafeLevel(fixedRecord2)
	if unsafeLevel == -1 {
		return true
	}
	return false
}

func findUnsafeLevel(record []int) int {
	increasing := false
	for j := 1; j < len(record); j++ {
		if record[j] == record[j-1] {
			return j
		}
		if record[j-1] < record[j] {
			if j == 1 {
				increasing = true
			}
			if !increasing {
				return j
			}
			if record[j]-record[j-1] > maxLevelJump {
				return j
			}
		}
		if record[j-1] > record[j] {
			if increasing {
				return j
			}
			if record[j-1]-record[j] > maxLevelJump {
				return j
			}
		}
	}
	return -1
}

func dumpUnsafeLevel(r []int, level int) []int {
	tmp := make([]int, 0)
	if level == 0 {
		tmp = append(tmp, r[1:]...)
		return tmp
	}
	if level == len(r)-1 {
		tmp = append(tmp, r[:len(r)-1]...)
		return tmp
	}
	tmp = append(tmp, r[:level]...)
	tmp = append(tmp, r[level+1:]...)
	return tmp
}

func recordIsSafeBruteForceAllLevels(record []int) bool {
	unsafeLevel := findUnsafeLevel(record)
	if unsafeLevel == -1 {
		return true
	}
	for i := 0; i < len(record); i++ {
		fixedRecord := dumpUnsafeLevel(record, i)
		unsafeLevel = findUnsafeLevel(fixedRecord)
		if unsafeLevel == -1 {
			return true
		}
	}
	return false
}
