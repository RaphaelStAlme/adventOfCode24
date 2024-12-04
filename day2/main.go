package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func checkLevelIsSafe(levels []int) int {
	if len(levels) <= 1 {
		return -1
	}
	increasing := levels[1] > levels[0] //know if level increase or decrease
	for i := 1; i < len(levels); i++ {
		diff := math.Abs(float64(levels[i] - levels[i-1]))                                                  // get diff between previous level and actual level
		isSequential := (increasing && levels[i] > levels[i-1]) || (!increasing && levels[i] < levels[i-1]) // know if it's sequential or not (i.e sequential =  1 2 4 6 4 vs non-sequential 1 2 6 4 5
		validDiff := 1 <= diff && diff <= 3                                                                 // check if level has acceptable difference
		if !isSequential || !validDiff {
			return i // give level index which will be removed
		}
	}
	return -1
}

func deleteLevelAt(index int, levels []int) []int {
	deleted := make([]int, len(levels)-1)
	copy(deleted[:index], levels[:index])   // we copy all good level before invalid level
	copy(deleted[index:], levels[index+1:]) //we continue to complete deleted slice by adding good level after invalid level if it exist
	return deleted                          // we return levels without invalid level
}

func part1_old(intFields []int, previousLevelDifference *int, isReportSafe *bool) {
	for i := range intFields {
		if i != len(intFields)-1 {
			levelDifference := intFields[i] - intFields[i+1]

			if i != 0 {
				if math.Signbit(float64(*previousLevelDifference)) != math.Signbit(float64(levelDifference)) {
					fmt.Println(intFields, "intFields before")
					*isReportSafe = false
					break
				}
			}

			*previousLevelDifference = levelDifference
			levelAbsDifference := math.Abs(float64(levelDifference))

			if !(levelAbsDifference > 0 && levelAbsDifference <= 3) {
				*isReportSafe = false
				break
			}
		}
	}
}

func part1(levels []int, isReportSafe *bool) {
	if checkLevelIsSafe(levels) == -1 {
		*isReportSafe = true
	}
}

func part2(levels []int, isReportSafe *bool) {
	unsafeAtIndex := checkLevelIsSafe(levels) // when index is different to -1 we get incorrect index of level

	if unsafeAtIndex == -1 {
		*isReportSafe = true
	} else {
		replacement1, replacement2 := unsafeAtIndex-1, unsafeAtIndex // we got previous level index and incorrect index level and we store in replacement variable

		change1Safe := checkLevelIsSafe(deleteLevelAt(replacement1, levels)) == -1 // We check if delete index - 1 fix level
		change2Safe := checkLevelIsSafe(deleteLevelAt(replacement2, levels)) == -1 // We check if delete index fix level
		if change1Safe || change2Safe {                                            // if one of them solve levels, so report is consider as safe
			*isReportSafe = true
		}
	}
}

func main() {
	file, err := os.Open("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var safeReportsCount int
	var reports [][]int

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		var levels []int

		for _, field := range fields {
			if level, err := strconv.Atoi(field); err == nil {
				levels = append(levels, level)
			}
		}
		reports = append(reports, levels)

		isReportSafe := false

		//var previousLevelDifference int
		//part1_old(levels, &previousLevelDifference, &isReportSafe)
		//part1(levels, &isReportSafe)
		part2(levels, &isReportSafe)

		if isReportSafe {
			safeReportsCount++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(safeReportsCount, "safe reports count")
}
