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
	max := 0 // Maximum i value
	for ! dup {
		if list[i] == i && i != 0 {
			fmt.Println("The value in the accumulator is", accum, "and i after is", i)
			dup = true
		}
		if i > max {
			max = i
		}
		list[i] = i
		fmt.Println(f[i], "i:", i, "list[i] = ", list[i])
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
	// Part 2
	newlist := [200000]int{} 
	newlist[0] = -2
	fmt.Println("max:", max, "fmax:", f[max])
	if f[max][0:3] == "jmp" {
		f[max] = "nop" + f[max][3:]
	} else {
		f[max] = "jmp" + f[max][3:]
	}
	fmt.Println(f[max])
	dup = false
	j := 0
	accum = 0
	fmt.Println(len(f))
	for ! dup && j < len(f) {
		/*if newlist[j] == j && j != 0 {
			fmt.Println("The value in the accumulator is", accum, "and i after is", j)
			dup = true
		} */
		newlist[j] = j
//		fmt.Println(j)
//		fmt.Println(f[j], "j:", j, "list[j] = ", newlist[j])
		instr := f[j][:3] // Instruction
		val, e := strconv.Atoi(f[j][4:])  // e.g. -99 or +2
		check(e)
		if instr == "acc" {
			accum += val
			j++
		}
		if instr == "jmp" {
			j += val
		}
		if instr == "nop" {
			j++
		}
	}
	fmt.Println("Value in accumulator after normal termination:", accum)
} // Pt 1 answer: 1584. i after is 347 (that's the position it lands back at).
//   i before is 374.