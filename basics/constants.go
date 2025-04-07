package basics

import "fmt"

const pi = 3.14
const GRAVITY = 9.81

func main() {

	const day int = 06
	const month string = "april"
	const year int = 2025

	const (
		monday          = 1
		tuesday         = 2
		wednesday int32 = 3
		thursday        = 4
		friday          = 5
		saturday  int   = 6
		sunday          = 7
	)

	fmt.Println(day, month, year)

	fmt.Println(pi)
	fmt.Println(GRAVITY)
}
