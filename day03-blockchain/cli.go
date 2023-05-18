package main

import (
	"fmt"
	"os"
)

type CLI struct {
	bc *BlockChain
}

const Usage = `
	addBlock --data DATA     	 "add data to blockchain"
	printChain     				 "print all blockchain data"
`

func (cli *CLI) Run() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf(Usage)
		return
	}
	cmd := args[1]
	switch cmd {
	case "addBlock":
		fmt.Printf("添加区块\n")
		if len(args) == 4 && args[2] == "--data" {
			data := args[3]
			cli.AddBlock(data)
		} else {
			fmt.Printf("添加区块指令不当")
			fmt.Printf(Usage)
		}
	case "printChain":
		fmt.Printf("打印区块\n")
		cli.PrintBlockChain()
	default:
		fmt.Printf("无效的命令")
		fmt.Printf(Usage)
	}
}
