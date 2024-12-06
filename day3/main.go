package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func part1(mulRegex *regexp.Regexp, line string) int {
	var results int
	matches := mulRegex.FindAllStringSubmatch(line, -1)

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

func part2(line string, lineNumber int) int {
	//dontRegex := regexp.MustCompile(`don't\(\)`)
	//doRegex := regexp.MustCompile(``)
	var results int
	var shouldMultiply = false

	if lineNumber == 0 {
		shouldMultiply = true
	}

	mulRegx := regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d+),\s*(\d+)\)`)

	matches := mulRegx.FindAllStringSubmatch(line, -1)

	if len(matches) > 0 {
		for _, match := range matches {
			matchedText := match[0]
			fmt.Println(matchedText)

			if matchedText == "do()" {
				fmt.Println("GO IN DO()")
				shouldMultiply = true
			} else if matchedText == "don't()" {
				fmt.Println("GO IN DON'T()")
				shouldMultiply = false
			}
			if shouldMultiply {
				fmt.Println("GO HERE")
				a, _ := strconv.Atoi(match[1])
				b, _ := strconv.Atoi(match[2])

				results += a * b

			}
		}
	}
	//response is between 59000000 and 69247082
	return results
}

func main() {
	file, err := os.Open("day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	//mulRegx := regexp.MustCompile(`mul\((\d+),\s*(\d+)\)`)

	var linesText []string

	for scanner.Scan() {
		line := scanner.Text()
		linesText = append(linesText, line)
	}

	var results int = 0

	for i, line := range linesText {
		//results += part1(mulRegx, line)
		fmt.Println(i, "line number")
		results += part2(line, i)
	}

	fmt.Println(results)
}
