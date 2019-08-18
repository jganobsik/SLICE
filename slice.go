package main

import (
	"fmt"
	"sort"
	"strconv"
)

func askForNumber() string {
	fmt.Printf("Please enter a number:")
	input1 := ""
	fmt.Scanln(&input1)
	return input1
}
func getNumber(input *string) int {
	num, err := strconv.Atoi(*input)
	if err != nil {
		fmt.Println(err)

	}
	return num

}
func main() {
	counter := 0
	s := make([]int, 3)
	for {
		input1 := askForNumber()
		if input1 == "X" {
			break
		}
		n := getNumber(&input1)

		if counter > 2 {
			s = append(s, n)
		} else if s[0] == 0 {
			s[0] = n
		} else if s[1] == 0 {
			s[1] = n
		} else {
			s[2] = n
		}
		counter++
		sort.SliceStable(s, func(i, j int) bool { return s[i] < s[j] })
		fmt.Println(s)
	}
}
