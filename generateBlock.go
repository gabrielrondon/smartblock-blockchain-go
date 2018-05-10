package main

import "time"

func generateBlock(oldBlock Block, smartBlock int) (Block, error) {

	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.smartBlock = smartBlock
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock, nil
}
