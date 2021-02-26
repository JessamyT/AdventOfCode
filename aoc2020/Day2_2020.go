/* Advent of Code 2020 Day 2 in Go */
package main
import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	file, err := os.Open("input2.txt")
	check(err)
	scanner := bufio.NewScanner(file)
	counter1 := 0
	counter2 := 0
	for scanner.Scan() {   
		line := scanner.Text()
		splitcol := strings.Split(line, ": ") // Results in 2 strings i.e. "11-14 j" and "abcdcba"
		split1stSpace := strings.Split(splitcol[0], " ") // Split at first space i.e. "11-14" and "j"
		splitpoln := strings.Split(split1stSpace[0], "-") // Policy number range i.e. "11" and "14"
		poln1, err := strconv.Atoi(splitpoln[0]) // beginning of range
		check(err)
		poln2, err := strconv.Atoi(splitpoln[1]) // end of range
		check(err)
		charArray := split1stSpace[1] // Letter at end of policy i.e. "j"
		charStr := fmt.Sprint(charArray) // Convert []string to string
		inst := strings.Count(splitcol[1], charStr) // Instances of the character in the rest of the line
		if inst >= poln1 && inst <= poln2 {
			counter1++
		}
		if string(splitcol[1][poln1 -1 ]) == charStr || string(splitcol[1][poln2 - 1]) == charStr {
			if string(splitcol[1][poln1 - 1]) != charStr || string(splitcol[1][poln2 -1]) != charStr {
				counter2++
			}
		} 
	}
	fmt.Println("The answer to Part 1 is", counter1)
	fmt.Println("The answer to Part 1 is", counter2)
}