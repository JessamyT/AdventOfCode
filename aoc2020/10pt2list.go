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
func remove(s []int, j int) []int { // Faster when order doesn't matter (part 1)
	s[len(s)-1], s[j] = s[j], s[len(s)-1]
	return s[:len(s)-1]
}
func ordremove(slice []int, s int) []int { // For when order does matter (part 2)
    return append(slice[:s], slice[s+1:]...)
}
func Equal(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}
func arrange(chain []int, counter int, position int, list [][]int) ([]int, int, int, [][]int) {  // Returns modified chain and counter
//	fmt.Println(chain)
/*	if ! Equal(newchain, chain) {
		return arrange(chain, counter, position + 1, chain)
	} */
	max := 0
	for i := 0; i < len(chain) - 1; i++ {
		if chain[i + 1] - chain[i] > max {
			max = chain[i + 1] - chain[i]
		}
	}
//	fmt.Println(chain)
//	fmt.Println("count:", counter, "pos:", position, "len:", len(chain), "\n")
	for j := position; j < len(chain) - 1; j++ { // j starts at 1 because it's the outlet and we can't remove it
		if chain[j + 1] - chain[j - 1] < 4 { // Safe to remove element j
			newchain := ordremove(chain, j)
			fmt.Println(j, newchain, "newchain")
			_, newchaincounter, _, _ := arrange(newchain, 1, position, list) // Give newchain counter starting value 1 
			// because we just deleted an element making a new chain to count. Don't need to advance position because element deletion shortens chain.
			_, oldchaincounter, _, _ := arrange(chain, counter, position + 1, list)
			return arrange(chain, newchaincounter + oldchaincounter, position + 1, list)
		}
		return arrange(chain, counter, position + 1, list) // We need to return this AND the last thing
	}
//	fmt.Println("It gets to the bottom!!!!!!~~~~~~~~~~~~~~")
//	fmt.Println(chain, counter, position + 1)
	return chain, counter, position + 1, list // Not sure if this is right or will result in infinite loop
// It doesn't undo the chain deletion if it doesn't work. Is this taken care of with old and new?
}
func main() {
	file, err := os.Open("aocInputs/input10test.txt")
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
	list := make([][]int, 1)
	list[0] = chain
	fmt.Println(chain, list)
	chain = append(chain, totalmax + 3) // To include the device
	fmt.Println(chain, list)
	_, answer, _, finallist := arrange(chain, 0, 1, list)  // Should counter start at 1 to count the whole original chain? Or will it get counted?
	fmt.Println("The answer to  part 2 is:", answer)
	fmt.Println(finallist)
	fmt.Println(list)
	fmt.Println("The answer is also given by the length of the final list:", len(finallist))
} // Could use minimum difference. Could also use a built-in ordering function.