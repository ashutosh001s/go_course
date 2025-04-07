package basics

import (
	"fmt"
	"math"
)

func main() {
	//Variable declaration
	var a, b int = 10, 3
	var result int

	result = a + b
	fmt.Println(result)

	result = a - b
	fmt.Println(result)

	result = a * b
	fmt.Println(result)

	result = a / b
	fmt.Println(result)

	result = a % b
	fmt.Println(result)

	result = a << b
	fmt.Println(result)

	const pi float64 = 22.0 / 7.0
	fmt.Println(pi)

	//overflow signed integer

	var max8bitInt int8 = 127
	max8bitInt += 2
	fmt.Println(max8bitInt)

	//overflow un-signed integer
	var max8bitUInt uint8 = 255
	max8bitUInt++
	fmt.Println(max8bitUInt)

	//underflow

	//1.0e-323 is a very small number written in scientific notation.
	var smallFloat float64 = 1.0e-323
	fmt.Println(smallFloat)

	fmt.Println(math.MaxFloat64)

	//the resulting value is soo small that it cant be represented by float64 so it will result to 0
	smallFloat = smallFloat / math.MaxFloat64
	fmt.Println(smallFloat)

}
