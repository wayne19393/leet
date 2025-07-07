package main

import (
	"bufio"
	_ "bufio"
	"fmt"
	_ "fmt"
	"os"
	_ "os"
	"strconv"
	"strings"
	_ "strings"
)

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] += 1
			return digits
		}
		digits[i] = 0
	}
	//if we are here and didn't return anything it means all the numbers are 9
	return append([]int{1}, digits...)
}

// helper function to get input from user and create the array(testing purposes
func getDigits() []int {
	fmt.Println("Enter array of digits (space separated): ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fields := strings.Fields(input)
	nums := make([]int, 0, len(fields))
	for _, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			fmt.Printf("Invalid input: %s\n", field)
			continue
		}
		nums = append(nums, num)
	}
	return nums
}

func main() {
	digits := getDigits()
	fmt.Println("You entered:", digits)
	fmt.Println("plus one numbers:", plusOne(digits))
}
