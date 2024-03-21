package client

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gr8h/crypto-file-guard/merkletree"
	"github.com/gr8h/crypto-file-guard/server"
)

type Client struct {
	RootHash    []byte
	Server      *server.Server
	StoragePath string
}

// NewClient initializes a new client instance.
func NewClient(server *server.Server) *Client {
	return &Client{
		Server:      server,
		StoragePath: "./client/files",
	}
}

// ConstructMerkleTree uploads the given files to the server by reading each file's content,
// calling the server's StoreFileContent method for each one, and then
// instructing the server to construct the Merkle tree.
func (c *Client) ConstructMerkleTree(files []string) error {
	for _, file := range files {
		file = filepath.Join(c.StoragePath, file)
		file, err := filepath.Abs(file)
		if err != nil {
			return fmt.Errorf("error getting absolute path for file %s: %v", file, err)
		}
		content, err := os.ReadFile(file)

		if err != nil {
			return fmt.Errorf("error reading file %s: %v", file, err)
		}

		err = c.Server.AddFile(content)
		if err != nil {
			return fmt.Errorf("error uploading file %s: %v", filepath.Base(file), err)
		}
	}

	if err := c.Server.ConstructMerkleTree(); err != nil {
		return fmt.Errorf("error constructing Merkle tree: %v", err)
	}

	rootHash, err := c.Server.Tree.GetRootHash()
	if err != nil {
		return fmt.Errorf("error retrieving Merkle tree root hash: %v", err)
	}

	c.RootHash = rootHash

	// os.RemoveAll(c.StoragePath)

	return nil
}

// RequestFile requests a file and its proof from the server, then verifies the file's integrity.
func (c *Client) GetProof(index int) ([]merkletree.Hash, error) {
	proof, err := c.Server.GetProof(index)
	if err != nil {
		return nil, fmt.Errorf("error requesting file: %v", err)
	}

	return proof, nil
}

// VerifyProof verifies the given proof for the given data block and root hash.
func (c *Client) VerifyProof(proof []merkletree.Hash, fileHash merkletree.Hash, rootHash []byte) (bool, error) {
	return c.Server.Tree.VerifyProof(proof, fileHash, rootHash)
}
