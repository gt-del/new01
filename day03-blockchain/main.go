package main

import "fmt"

func main() {
	bc := NewBlockChain()
	bc.AddBlock("sssss")
	bc.AddBlock("ssdss")
	for i, block := range bc.blocks {
		fmt.Printf("========当前区块高度：%d ==========\n", i)
		fmt.Printf("当前区块哈希： %x\n", block.Hash)
		fmt.Printf("父区块哈希： %x\n", block.PrevHash)
		fmt.Printf("区块数据：%s\n", block.Data)
	}

}
