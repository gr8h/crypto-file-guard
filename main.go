package main

import (
	"encoding/hex"
	"fmt"

	"github.com/gr8h/crypto-file-guard/internal/client"
	"github.com/gr8h/crypto-file-guard/internal/server"
)

func main() {

	files := []string{"0.txt", "1.txt", "2.txt", "3.txt", "4.txt"}

	server := server.NewServer("test")

	client := client.NewClient(server)

	client.ConstructMerkleTree(files)

	server.Tree.PrintTree()

	fmt.Println("Root hash main:", hex.EncodeToString(client.RootHash))

	proof, _ := client.GetProof(2)

	fmt.Print("Proof: ")
	for _, p := range proof {
		fmt.Printf("[%s] ", hex.EncodeToString(p)[:10])
	}
	fmt.Println()

	fileHash, _ := server.GetFile(2)

	isValid, err := client.VerifyProof(proof, fileHash, client.RootHash)
	if err != nil {
		fmt.Println("Error verifying proof: ", err)
	}
	fmt.Println("Is valid: ", isValid)
}
