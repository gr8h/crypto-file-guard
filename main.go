package main

import (
	"crypto-file-guard/merkletree"
	"encoding/hex"
	"fmt"
)

func main() {
	dataBlocks := [][]byte{
		[]byte("data block 1"),
		[]byte("data block 2"),
		[]byte("data block 3"),
		[]byte("data block 4"),
		[]byte("data block 5"),
		// []byte("data block 6"),
		// []byte("data block 7"),
		// []byte("data block 8"),
		// []byte("data block 9"),
	}

	tree := merkletree.NewMerkleTree(dataBlocks)
	fmt.Println("Root hash:", tree.Root.Hash)

	tree.PrintTree()

	var proof [][]byte = tree.GenerateProof([]byte("data block 5"))
	fmt.Print("Proof: ")
	for _, p := range proof {
		fmt.Printf("[%s] ", hex.EncodeToString(p)[:10])
	}
	fmt.Println()

	var verify bool = tree.VerifyProof(proof, []byte("data block 5"), tree.Root.Hash, 4)
	fmt.Println("Verify: ", verify)
}
