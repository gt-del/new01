package main

import (
	"log"
	"taoguo/bolt"
)

//blockChain是由数据库和最后一个区块的哈希组成
//数据库中存放的数据最后一个键值对也是最后一个块的哈希
type BlockChain struct {
	db   *bolt.DB
	tail []byte
}

const (
	blockChainDb = "blockchain.db"
	blockBucket  = "blockBucket"
)

func NewBlockChain() *BlockChain {
	/*genesisBlock := GenesisBlock()
	return &BlockChain{
		blocks: []*Block{genesisBlock},
	}*/
	var lastHash []byte
	db, err := bolt.Open(blockChainDb, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			bucket, err := tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				log.Panic("创建bucket err")
			}
			genesisBlock := GenesisBlock()
			bucket.Put(genesisBlock.Hash, genesisBlock.Serialize())
			bucket.Put([]byte("lastHashKey"), genesisBlock.Hash)
		} else {
			lastHash = bucket.Get([]byte("lastHashKey"))
		}
		return nil
	})
	//db.Close()此处如果关闭，添加区块时无法添加数据
	return &BlockChain{db, lastHash}
}

// GenesisBlock 定义一个创世块
func GenesisBlock() *Block {
	return NewBlock("一期创世块!", []byte{})
}

// AddBlock 添加区块
func (bc *BlockChain) AddBlock(data string) {
	/*
		lastBlock := bc.blocks[len(bc.blocks)-1]
		prevHash := lastBlock.Hash
		block := NewBlock(data, prevHash)
		bc.blocks = append(bc.blocks, block)*/
	db := bc.db
	lastHash := bc.tail
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Panic("请检查数据库")
		}
		block := NewBlock(data, lastHash)
		bucket.Put(block.Hash, block.Serialize())
		bucket.Put([]byte("LastHashKey"), block.Hash)
		bc.tail = block.Hash
		return nil
	})
}
