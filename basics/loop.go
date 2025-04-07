package basics

import "fmt"

func main() {

	//Simple iteration over a range
	// for i := 0; i <= 5; i++ {
	// 	fmt.Println(i)
	// }

	// //iterate over collection
	// numbers := []string{"harry", "ron", "dutch", "tony", "steve"}

	// for index, value := range numbers {
	// 	if index == 3 {
	// 		break
	// 	}
	// 	fmt.Printf("Index : %v , Value : %v\n", index, value)
	// }

	// for i := 0; i < 10; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	if i == 7 {
	// 		break
	// 	}
	// 	fmt.Println(i)

	// }

	// rows := 5
	// //outer loop
	// for i := 1; i <= rows; i++ {
	// 	//inner lop
	// 	for j := 1; j <= rows-i; j++ {
	// 		fmt.Print(" ")
	// 	}
	// 	//inner loop for stars
	// 	for k := 1; k <= 2*i-1; k++ {
	// 		fmt.Print("*")

	// 	}
	// 	fmt.Println() // move to next line
	// }

	for i := range 10 {

		fmt.Println(10 - i)
	}

}
