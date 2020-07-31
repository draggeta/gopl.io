// Ex: write a version of rotate that operates in a single pass

package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6}
	r := 5
	fmt.Println(rotate(s, r))
}

func rotate(s []int, r int) []int {
	r = r % len(s)
	if r == 0 {
		return s
	}
	s = append(s[r:], s[:r]...)
	return s
}
