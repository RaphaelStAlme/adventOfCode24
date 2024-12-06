package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func part1(fileData string) int {
	mulRegex := regexp.MustCompile(`mul\((\d+),\s*(\d+)\)`)

	var results int
	matches := mulRegex.FindAllStringSubmatch(fileData, -1)

	fmt.Println(matches)

	if len(matches) > 0 {
		for _, match := range matches {
			a, _ := strconv.Atoi(match[1])
			b, _ := strconv.Atoi(match[2])
			results += a * b
		}
		fmt.Println("")
	}

	return results
}

func part2(fileData string) int {
	//dontRegex := regexp.MustCompile(`don't\(\)`)
	//doRegex := regexp.MustCompile(``)
	var results int
	var shouldMultiply = true

	mulRegx := regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d+),\s*(\d+)\)`)

	matches := mulRegx.FindAllStringSubmatch(fileData, -1)

	if len(matches) > 0 {
		for _, match := range matches {
			matchedText := match[0]

			if matchedText == "do()" {
				shouldMultiply = true
			} else if matchedText == "don't()" {
				shouldMultiply = false
			}
			if shouldMultiply {
				a, _ := strconv.Atoi(match[1])
				b, _ := strconv.Atoi(match[2])

				results += a * b

			}
		}
	}
	return results
}

func main() {
	fileData, err := os.ReadFile("day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	partChoosen := 2

	var results int = 0

	if partChoosen == 1 {
		results += part1(string(fileData))
	} else {
		results += part2(string(fileData))
	}

	fmt.Println(results)
}
