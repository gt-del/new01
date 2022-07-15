package main

import (
	"bufio"
	"fmt"
	"os"
)

//bufio包实现了有缓冲的I/O。
//它包装一个io.Reader或io.Writer接口对象，创建另一个也实现了该接口，且同时还提供了缓冲和一些文本I/O的帮助函数的对象。
func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin) //创建并返回一个从r读取数据的scanner，默认的分割函数是scanlines。
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
