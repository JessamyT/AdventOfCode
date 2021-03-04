/* Advent of Code 2020 Day 9 Part 2 */
package main
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func number(f []string, i int) int { // f is the file as a slice. i is the index of the line
	n, e := strconv.Atoi(f[i]) // n is the integer in line i
	check(e)
	return n
}
func sumr(f []string, i int, sum int, goal int, start int) (int, int) { // This is the recursive function 
	// that sums contiguous numbers. start keeps track of our starting place in the file. Returns start and end i.
	if i >= len(f)  {
		return sumr(f, start + 1, 0, goal, start + 1) // Resets sum if it gets to the end of the file
	}
	newsum := sum + number(f, i)
	if newsum == goal {
		return start, i
	} else { // else is redundant with return
		return sumr(f, i + 1, newsum, goal, start)
	}
	return 0, 0
}
func main() {
	file, err := os.Open("aocInputs/input9.txt")
	check(err)
	f := []string{}  // Slice to hold file contents
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		f = append(f, line)
	}
	valid := true // True means that so far, all numbers have been valid
	i := 25 // Line in file f
	for valid == true {
		valid = false // Then the for loop will make it true again if it's true
		for j := i - 25; j < i; j++ {
			for k := j + 1; k < i; k++ {
				if number(f, j) + number(f, k) == number(f, i) {
					valid = true
				}
			}
		}
		i++
	}
	i -= 1 // Because we ++ it once more before the loop ended
	fmt.Println("The invalid number is", f[i], "and it is at line", i + 1, "not counting the 0 index.")
	goal := number(f, i) // The number we're trying to find as a sum of contiguous numbers
	j := 0
	startpos, endpos := sumr(f, j, 0, goal, 0)
	// Find max and min within this range
	max := 0
	min := 90000000000000000
	for k := startpos; k <= endpos; k++ {
		n := number(f,k)
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
	}
	fmt.Println("Answer to Part 2:", max + min)
}