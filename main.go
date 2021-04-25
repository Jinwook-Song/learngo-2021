package main

import "fmt"

func superAdd(numbers ...int) int{
	sumNum := 0
	for _, number:= range numbers {
		sumNum += number
	}
	return sumNum
}

func greatAdd(numbers ...int) int {
	sumNum := 0
	for i := 0; i < len(numbers); i++ {
		sumNum += numbers[i]
	}
	return sumNum
}

func main() {
	sum1 := superAdd(1214,421,41,412,12)
	sum2 := greatAdd(1,2,3,1231414,5,6)

	fmt.Println(sum1, sum2)
}
