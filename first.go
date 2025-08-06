package main

import "fmt"

func main() {

	for x := range 10 {
		if x%2 == 0 {
			fmt.Println("Value of x : ", x)
		} else {
			fmt.Println("Odd value of x : ", x)
		}
	}

	num := [4]int{1, 2, 3, 4}

	modifyArr(&num)

	fmt.Println(num)

	len := 10
	bred := 20

	result := AreaRact(len, bred)
	fmt.Println(result)
}

func modifyArr(arr *[4]int) {
	arr[0] = 10
}

func AreaRact(a, b int) int {
	area := a * b
	return area
}
