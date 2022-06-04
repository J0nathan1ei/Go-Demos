package bc_structs

type BcInterface interface {
	Genesis() *Block
	CreateBlock(data string, prevHash []byte) *Block
}
