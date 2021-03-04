/* Advent of Code 2020 Day 9 Part 1 */
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
	fmt.Println("The invalid number is", f[i], "and it is at line", i + 1, "not counting 0 index")
}