package basics

import "fmt"

func main() {

	// i := 1
	// for i < 5 {
	// 	fmt.Println(i)
	// 	i++
	// }

	//Infinite loop
	// for {
	// 	fmt.Println("Hello")
	// }

	sum := 0
	for {
		if sum >= 50 {
			break
		}
		sum += 10
		fmt.Println(sum)
	}

	num := 1
	for num <= 10 {
		if num%2 == 0 {
			num++
			continue
		}
		fmt.Println(num)
		num++
	}
}
