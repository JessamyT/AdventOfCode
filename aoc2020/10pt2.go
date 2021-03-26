/* Advent of Code 2020 Day 10 Part 2 */
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
func min(s map[int]int) (int, int) {  // s is a slice of numbers from which to find the minimum
	m := 100000000
	index := 0
	for n, i := range s {
		if n < m {
			m, index = n, i
		}
	}
	return m, index
}
func remove(s []int, j int) []int {
	s[len(s)-1], s[j] = s[j], s[len(s)-1]
	return s[:len(s)-1]
}
func ordremove(slice []int, s int) []int { // For when order does matter (part 2)
	var l []int
	for i, v := range(slice) {
		if i != s {
			l = append(l, v)
		}
	}
	return l
}
// To make more efficient, put valid indices in a map with whatever values
func arrange(L []int, q int, solutions int) int {
	for i := q; i < len(L) - 1; i++ {
		if L[i + 1] - L[i - 1] < 4 {
			newL := ordremove(L, i)
			solutions = arrange(newL, i, solutions + 1)
		}
	}
	return solutions
}

func main() {
	file, err := os.Open("aocInputs/input10.txt")
	check(err)
	f := []int{}  // Slice to hold file contents
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line, e := strconv.Atoi(scanner.Text())
		check(e)
		f = append(f, line)
	}
	totalmax := 0 // Where to end
	for i := 0; i < len(f); i++ {
		if f[i] > totalmax {
			totalmax = f[i]
		}
	}
	chain := []int{} // A slice containing all the joltages of the adapters in ascending order
	chain = append(chain, 0) // To include the outlet
	low := 0 // Current number
	length := len(f) // Before we change it
	for len(chain) < length + 1 { // +1 because of 0 outlet element
		holder := make(map[int]int) // Temporary map to hold all remaining elements of f that are within 3 of low, and their indices
		for j := 0; j < len(f); j++ {
			if f[j] <= low + 3 {
				holder[f[j]] = j
			}
		}
		addtochain, itoremove := min(holder) // Element of f to remove after adding it to chain
		chain = append(chain, addtochain)
		f = remove(f, itoremove)
		low, _ = min(holder)
	}
	chain = append(chain, totalmax + 3) // To include the device
	jolt1 := 0 // Number of 1-jolt differences
	jolt3 := 0 // Number of 3-jolt differences
	for k := 0; k < len(chain) - 1; k++ { // -1 because need two numbers before end
		diff := chain[k + 1] - chain[k]
		if diff == 1 {
			jolt1++
		}
		if diff == 3 {
			jolt3++
		}
	}
	fmt.Println("There are", jolt1, "1-jolt differences and", jolt3, "3-jolt differences.")
	fmt.Print("Their product is ", jolt1 * jolt3, ".\n")
	// Split the chain approximately in half for efficiency
	var chain1 []int
	var chain2 []int
	var splitHere int
	for i := len(chain) / 2; i < len(chain); i++ {
		if chain[i + 1] - chain[i] == 3 { // Good place to split since neither can be removed
			splitHere = i
			break
		}
	}
	for i, v := range(chain) {
		if i <= splitHere {
			chain1 = append(chain1, v)
		} else {
			chain2 = append(chain2, v)
		}
	}
	answer := (arrange(chain1, 1, 0) + 1) * (arrange(chain2, 1, 0) + 1) // Add + 1 for the unchanged chain
	fmt.Println("The answer to part 2 is:", answer)
}