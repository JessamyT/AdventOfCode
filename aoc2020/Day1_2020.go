/* Advent of Code 2020 Day 1 in Go */
package main
import (
	"fmt"
	"bufio"
	"strconv"
	"os"
)
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	file, err := os.Open("input1.txt")
	check(err)
	scanner := bufio.NewScanner(file)
	var inp []int  /* Array to hold input */
	for scanner.Scan() {  /* Moves one line forward in the file, and is bool so this is basically while */ 
		lineInt, err := strconv.Atoi(scanner.Text())
		inp = append(inp, lineInt)
		check(err)
	}
/* Part 1 */ 
	for index1, value1 := range inp {
		for _, value2 := range inp[index1:] { // underscore is placeholder in Go
			if value1 + value2 == 2020 {
				ans := value1 * value2
				fmt.Println("The answer to Par 1 is ", ans)
				break
			}
		}
	}
/* Part 2 */
	for index1, value1 := range inp {
		for index2, value2 := range inp[index1:] {
			for _, value3 := range inp[index2:] {
				if value1 + value2 + value3 == 2020 {
					ans := value1 * value2 * value3
					fmt.Println("The answer to Part 2 is ", ans)
					break
				}
			}
		}
	}
}