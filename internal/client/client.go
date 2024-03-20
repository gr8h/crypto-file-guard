package client

import (
	"encoding/hex"
	"fmt"

	"github.com/gr8h/crypto-file-guard/internal/server"
)

type Client struct {
	RootHash []byte
}

// NewClient initializes a new client instance.
func NewClient() *Client {
	return &Client{}
}

// UploadFiles uploads the given files to the server and stores the root hash.
func (c *Client) UploadFiles(server *server.Server, files []string) {
	err := server.ProcessFiles(files)
	if err != nil {
		fmt.Println("Error uploading files to server:", err)
		return
	}
	c.RootHash, _ = server.Tree.GetRootHash()
}

// RequestFile requests a file and its proof from the server, then verifies the file's integrity.
func (c *Client) RequestFile(server *server.Server, index int) (string, bool) {
	file, proof, err := server.GetFileAndProof(index)
	if err != nil {
		fmt.Println("Error requesting file:", err)
		return "", false
	}

	fmt.Print("Proof: ")
	for _, p := range proof {
		fmt.Printf("[%s] ", hex.EncodeToString(p)[:10])
	}
	fmt.Println()

	isValid, _ := server.Tree.VerifyProof(proof, []byte(file), c.RootHash)
	return file, isValid
}
