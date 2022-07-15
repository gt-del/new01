package main

import "fmt"

func main() {
	v := [3]int{1, 2, 3}
	for i, value := range v {
		v[i] = 3
		fmt.Println("value:", value)
	}
	fmt.Println("v", v)
}
