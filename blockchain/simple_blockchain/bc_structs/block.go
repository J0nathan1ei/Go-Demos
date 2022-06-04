package bc_structs

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	PrevHash	[]byte
	Hash		[]byte
	Data		[]byte
}

/*
@description Genesis Block
@param
@return Genesis Block with empty PrevHash
*/
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}


/*
@description init a Block's Hash
@param
@return
*/
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	// This will join our previous block's relevant info with the new Blocks
	hash := sha256.Sum256(info)
	//This performs the actual hashing algorithm
	b.Hash = hash[:]
	//If this ^ doesn't make sense, you can look up slice defaults
}

/*
@description create one Block
@param data & prev block's hash
@return
*/
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{prevHash, []byte{}, []byte(data)}
	//If this is gibberish to you look up
	// pointer syntax in go
	block.DeriveHash()
	return block
}
