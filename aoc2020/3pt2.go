/* Advent of Code 2020 Day 3 Part 2 in Go */
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
	stepXarr := []int{1, 3, 5, 7, 1} // Step increment to the right
	treemult := 1 // All tree counts multiplied together
	for i, stepX := range stepXarr {
		file, err := os.Open("input3.txt")
		check(err)
		scanner := bufio.NewScanner(file)
		posX := 0  // X position (how far right)
		treecount := 0
		for scanner.Scan() {   
			line := scanner.Text()
			if posX >= len(line) {
				posX -= len(line)
			}
			if string(line[posX]) == "#" {
				treecount++
			}
			if i == 4 {
				scanner.Scan() // Based on the puzzle, we want to advance this one down 2 at a time.
			}
			posX += stepX
		}
		fmt.Print("The tree count for Part ", i + 1, " is ", treecount, ".\n")
		treemult = treemult * treecount
	}
	fmt.Println("The product of the tree counts multiplied together is: ", treemult)
	
	
}
