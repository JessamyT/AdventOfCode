/* Advent of Code 2020 Day 8 Part 2 */
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
	file, err := os.Open("aocInputs/input8.txt")
	check(err)
	f := []string{}  // Slice to hold file contents
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		f = append(f, line)
	}
	found := false // Have we found the error yet?
	j := 0 // Index of the line we will change
	for ! found {
		i := 0 // Current index within file f
		list := [200000]int{} // Slice of all the indices that have been seen thus far.
		list[0] = -2  // Initialize index 0 with nonzero so it doesn't already end program.
		dup := false // Existence of duplicate (indicative of not having found error yet)
		max := 0 // Maximum i value
		accum := 0  // Accumulator value
		origline := f[j] // So we can change it back if it isn't the erroneous line
		if f[j] != "nop +0" { // To avoid creating an infinite jmp +0 situation
			if f[j][:3] == "jmp" {
				f[j] = "nop" + f[j][3:]
			} else if f[j][:3] == "nop" { // So I don't change it right back
				f[j] = "jmp" + f[j][3:]
			}
		}
		for ! dup {
			if list[i] == i && i != 0 {
				dup = true
			}
			if i > max {
				max = i
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
			if i >= len(f) - 1 {
				found = true
				break
			} 
		} 
		if found {
			break // To get out of the ! found loop without changing f[j] back
		}
		f[j] = origline
		j++
	}
	// Now we get the accumulator value we want, with the modified file
	k := 0
	finala := 0 // Final accumulator
	for k < len(f) { // change back to <=
		instr := f[k][:3] // Instruction
		val, e := strconv.Atoi(f[k][4:])  // e.g. -99 or +2
		check(e)
		if instr == "acc" {
			finala += val
			k++
		}
		if instr == "jmp" {
			k += val
		}
		if instr == "nop" {
			k++
		}
	}
	fmt.Println("The answer is:", finala)
}