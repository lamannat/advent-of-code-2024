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
		if recordIsSafe(records[i]) {
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
	fixedRecord := dumpUnsafeLevel(record, unsafeLevel)
	unsafeLevel = findUnsafeLevel(fixedRecord)
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

}
