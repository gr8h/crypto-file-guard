package main

import (
	"encoding/hex"
	"fmt"

	"github.com/gr8h/crypto-file-guard/internal/client"
	"github.com/gr8h/crypto-file-guard/internal/server"
)

func main() {

	files := []string{"File1", "File2", "File3"}
	server := server.NewServer()
	client := client.NewClient()

	client.UploadFiles(server, files)

	server.Tree.PrintTree()

	fmt.Println("Root hash:", hex.EncodeToString(server.Tree.Root.Hash))
	fmt.Println("Root hash:", hex.EncodeToString(client.RootHash))

	file, isValid := client.RequestFile(server, 3)

	fmt.Println("File:", file)
	fmt.Println("Is valid:", isValid)

	// dataBlocks := [][]byte{
	// 	[]byte("data block 1"),
	// 	[]byte("data block 2"),
	// 	[]byte("data block 3"),
	// 	[]byte("data block 4"),
	// 	[]byte("data block 5"),
	// 	// []byte("data block 6"),
	// 	// []byte("data block 7"),
	// 	// []byte("data block 8"),
	// 	// []byte("data block 9"),
	// }
	// tree, _ := merkletree.NewMerkleTree(dataBlocks)
	// fmt.Println("Root hash:", tree.Root.Hash)

	// tree.PrintTree()

	// proof, _ := tree.GenerateProof([]byte("data block 5"))
	// fmt.Print("Proof: ")
	// for _, p := range proof {
	// 	fmt.Printf("[%s] ", hex.EncodeToString(p)[:10])
	// }
	// fmt.Println()

	// verify, _ := tree.VerifyProof(proof, []byte("data block 5"), tree.Root.Hash)
	// fmt.Println("Verify: ", verify)
}
