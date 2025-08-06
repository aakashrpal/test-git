package main

import "fmt"

func main() {
	fmt.Println("Hello Aakash")

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
}

func modifyArr(arr *[4]int) {
	arr[0] = 10
}
