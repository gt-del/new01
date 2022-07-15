package main

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	PrevHash []byte
	Hash     []byte
	Data     []byte
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		prevBlockHash,
		[]byte{},
		[]byte(data),
	}
	block.SetHash()
	return &block
}
func (block *Block) SetHash() {

	blockInfo := append(block.PrevHash, block.Data...)
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}

//引入区块链
type BlockChain struct {
	blocks []*Block
}
func NewBlockChain() *BlockChain {
	return &BlockChain{
		blocks: ([]*Block),
	}
}
func main() {
	block := NewBlock("sdasdasdas", []byte{})
	fmt.Printf("父区块哈希： %x\n", block.PrevHash)
	fmt.Printf("当前区块哈希： %x\n", block.Hash)
	fmt.Printf("当前区块数据： %s\n", block.Data)
}
