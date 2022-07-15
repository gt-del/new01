package main

type BlockChain struct {
	blocks []*Block
}

func NewBlockChain() *BlockChain {
	genesisBlock := GenesisBlock()
	return &BlockChain{
		blocks: []*Block{genesisBlock},
	}
}

// GenesisBlock 定义一个创世块
func GenesisBlock() *Block {
	return NewBlock("一期创世块!", []byte{})
}

// AddBlock 添加区块
func (bc *BlockChain) AddBlock(data string) {
	lastBlock := bc.blocks[len(bc.blocks)-1]
	prevHash := lastBlock.Hash
	block := NewBlock(data, prevHash)
	bc.blocks = append(bc.blocks, block)
}
