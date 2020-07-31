// Ex: write an in-place function to eliminate adjacent duplicates in a []string slice

package main

import "fmt"

func main() {
	s := []string{"hoi", "hoi", "ola", "no", "ola", "yes", "yes"}
	fmt.Println(dedup(s))

}

func dedup(s []string) []string {
	j := 1
	for range s {
		if s[j] == s[j-1] {
			copy(s[j:], s[j+1:])
			continue
		}
		j++
	}
	return s[:j]
}
