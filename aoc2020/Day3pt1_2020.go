/* Advent of Code 2020 Day 3 Part 1 in Go */
package main
import (
	"fmt"
	"bufio"
	"os"
)
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	file, err := os.Open("input3.txt")
	check(err)
	scanner := bufio.NewScanner(file)
	posX := 0  // X position (how far right)
	stepX := 3 // Step increment to the right
	treecount := 0
	for scanner.Scan() {   
		line := scanner.Text()
		if posX >= len(line) {
			posX -= len(line)
		}
		if string(line[posX]) == "#" {
			treecount++
		}
		posX += stepX
	}
	fmt.Print("The tree count for Part 1 is ", treecount, ".\n")
	
}