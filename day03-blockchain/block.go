package main

import (
	"bytes"
	"encoding/binary"
	"log"
	"time"
)

type Block struct {
	//1.版本号
	Version uint64
	//2.前区块哈希
	PrevHash []byte
	//3.默克尔根
	MerkleRoot []byte
	//4.时间戳
	TimeStamp uint64
	//5.难度值
	Difficulty uint64
	//6.随机数
	Nonce uint64
	//a.当前区块哈希
	Hash []byte
	//b.数据
	Data []byte
}

// Uint64ToByte 辅助函数，将uint64转为[]byte类型
func Uint64ToByte(num uint64) []byte {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buf.Bytes()
}
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		Version:    0.0, // version
		PrevHash:   prevBlockHash,
		MerkleRoot: []byte{},
		TimeStamp:  uint64(time.Now().Unix()),
		Difficulty: 0,
		Nonce:      0,
		Hash:       []byte{},
		Data:       []byte(data),
	}
	//block.SetHash()
	pow := NewProofOfWork(&block)
	hash, nonce := pow.Run()
	block.Hash = hash
	block.Nonce = nonce
	return &block
}

/*func (block *Block) SetHash() {
	tmp := [][]byte{
		Uint64ToByte(block.Version),
		block.PrevHash,
		block.MerkleRoot,
		Uint64ToByte(block.TimeStamp),
		Uint64ToByte(block.Difficulty),
		Uint64ToByte(block.Nonce),
		block.Data,
	}
	//将二维数组的切片数据连接起来，返回一个一维的切片
	blockInfo := bytes.Join(tmp, []byte{})
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}*/
