/* Advent of Code 2020 Day 8 Part 1 */
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
func main() {
	file, err := os.Open("input8.txt")
	check(err)
	f := []string{}  // Slice to hold file contents
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		f = append(f, line)
	}
	i := 0 // Current index within file f
	list := [200000]int{} // Slice of all the indices that have been seen thus far.
	list[0] = -2  // Initialize index 0 with nonzero so it doesn't already end program.
	dup := false // Existence of duplicate
	accum := 0  // Accumulator value
	for ! dup {
		if list[i] == i && i != 0 {
			fmt.Println("The value in the accumulator is", accum)
			dup = true
		}
		list[i] = i
		instr := f[i][:3] // Instruction
		val, e := strconv.Atoi(f[i][4:])  // e.g. -99 or +2
		check(e)
		if instr == "acc" {
			accum += val
			i++
		}
		if instr == "jmp" {
			i += val
		}
		if instr == "nop" {
			i++
		}
	}
}