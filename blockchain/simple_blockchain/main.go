package main

import(
	"fmt"
	"simple_blockchain/bc_structs"
)

func main(){
	root := bc_structs.InitBlockChain()
	root.AddBlock("hello")
	fmt.Println(root.Blocks[0].Data)
	fmt.Println(root.Blocks[1].Data)

	genesis := bc_structs.Genesis()
	genesis.DeriveHash()
	d := bc_structs.CreateBlock("helloWorld", genesis.Hash)
	fmt.Println(d.Data)
}
