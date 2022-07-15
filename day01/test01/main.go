package main

import "fmt"

func main() {
	s := make([]int, 5, 5)
	Modify(s)
	fmt.Println(s)
}
func Modify(sli []int) {
	for i := 0; i < 5; i++ {
		sli = append(sli, i)
	}
	fmt.Println("sli:", sli)
} //使用append函数给切片扩容时，当切片
