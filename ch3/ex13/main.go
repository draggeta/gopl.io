// Ex: write const declaration for KB through YB as compact as possible

package main

import "fmt"

const (
	// KB  MB GB etc. are SI orders of magnitude
	KB = 1000
	MB = KB * 1000
	GB = MB * 1000
	TB = GB * 1000
	PB = TB * 1000
	EB = PB * 1000
	ZB = EB * 1000 // overflows
	YB = ZB * 1000 // overflows
)

func main() {
	fmt.Println(KB, MB, GB, TB, PB, EB)
}
