package main

import "fmt"

func main() {
	fmt.Println("Hello Aakash")

	for x := range 10 {
		if x%2 == 0 {
			fmt.Println("Value of x : ", x)
		}
	}
}
