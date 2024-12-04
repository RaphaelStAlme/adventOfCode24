package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func count[T any](slice []T, f func(T) bool) int {
	count := 0
	for _, s := range slice {
		if f(s) {
			count++
		}
	}
	return count
}

func getDistanceDifferent(locationsID1 []int, locationsID2 []int) int {
	//sort array in order asc
	sort.Ints(locationsID1)
	sort.Ints(locationsID2)

	var totalDistance int

	for i := range locationsID1 {
		if locationsID1[i] < locationsID2[i] {
			totalDistance += locationsID2[i] - locationsID1[i]
		} else {
			totalDistance += locationsID1[i] - locationsID2[i]
		}
	}

	return totalDistance
}

func getSimilarityScore(locationsID1 []int, locationsID2 []int) int {
	var similarityScore int

	for _, locationID1 := range locationsID1 {
		similarityScore += locationID1 * count(
			locationsID2,
			func(locationID2 int) bool {
				return locationID2 == locationID1
			})
	}

	return similarityScore
}

func main() {
	file, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	var locationsID1, locationsID2 []int

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		for i, field := range fields {
			if intField, err := strconv.Atoi(field); err == nil {
				if i != 0 {
					locationsID2 = append(locationsID2, intField)
				} else {
					locationsID1 = append(locationsID1, intField)
				}
			}
		}
	}

	// Part 1
	// fmt.Println(getDistanceDifferent(locationsID1, locationsID2), "TotalDistance")

	// Part 2
	fmt.Println(getSimilarityScore(locationsID1, locationsID2), "Similarity score")

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
