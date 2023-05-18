package main

import (
	"log"
	"taoguo/bolt"
)

type BlockChainIterator struct {
	db                 *bolt.DB
	currentHashPointer []byte
}

func (bc *BlockChain) NewIterator() *BlockChainIterator {
	return &BlockChainIterator{
		bc.db,
		bc.tail,
	}
}
func (it *BlockChainIterator) Next() *Block {
	var block Block
	it.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Panic("迭代器遍历的数据库为空，请检查")
		}
		blockTmp := bucket.Get(it.currentHashPointer)
		block = Desrialize(blockTmp)
		it.currentHashPointer = block.PrevHash
		return nil
	})
	return &block
}
