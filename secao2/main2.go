package main

import (
	"fmt"
)

func todo() {
	//var arr []int
	arr := []int{1, 2, 3, 4}
	arr = append(arr, 5)
	fmt.Println(arr)
}

func main() {
	todo()
}
