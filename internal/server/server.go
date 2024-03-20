package server

import (
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gr8h/crypto-file-guard/internal/merkletree"
)

type Server struct {
	Files       []merkletree.Hash
	Tree        *merkletree.MerkleTree
	StoragePath string
}

// NewServer initializes a new server instance without any files.
func NewServer(storagePath string) *Server {

	storagePath = filepath.Join("./internal/server/files/", storagePath, "/")

	return &Server{
		StoragePath: storagePath,
	}
}

func (s *Server) AddFile(content []byte) error {
	// Ensure the storage directory exists
	if err := os.MkdirAll(s.StoragePath, 0755); err != nil {
		return fmt.Errorf("failed to create storage directory: %v", err)
	}

	// Write the content to a new file in the storage directory
	fileHash, err := merkletree.HashData(content)
	if err != nil {
		return fmt.Errorf("failed to hash file content: %v", err)
	}

	filePath := filepath.Join(s.StoragePath, hex.EncodeToString(fileHash)[:10])
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %v", filePath, err)
	}

	// Add the file name to the list to maintain order
	s.Files = append(s.Files, fileHash)

	return nil
}

// ConstructMerkleTree reads the content of all stored files and constructs/updates the Merkle tree.
func (s *Server) ConstructMerkleTree() error {
	var hashes []merkletree.Hash

	for _, fileHash := range s.Files {

		filePath := filepath.Join(s.StoragePath, hex.EncodeToString(fileHash)[:10])

		content, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %v", filePath, err)
		}

		hashedData, err := merkletree.HashData(content)
		if err != nil {
			return fmt.Errorf("failed to hash node data: %w", err)
		}

		hashes = append(hashes, hashedData)
	}

	tree, err := merkletree.NewMerkleTree(hashes)
	if err != nil {
		return fmt.Errorf("failed to construct/update Merkle tree: %v", err)
	}

	s.Tree = tree
	return nil
}

// GetFileAndProof returns the file at the given index and its Merkle proof.
func (s *Server) GetFileAndProof(index int) (merkletree.Hash, []merkletree.Hash, error) {
	if index < 0 || index >= len(s.Files) {
		return nil, nil, fmt.Errorf("server: index out of bounds")
	}

	data, err := s.Tree.GetLeafByIndex(index)
	if err != nil {
		return nil, nil, err
	}

	proof, err := s.Tree.GenerateProof(data)
	if err != nil {
		return nil, nil, err
	}
	return s.Files[index], proof, nil
}

// VerifyProof verifies the given proof for the given data block and root hash.
func (s *Server) VerifyProof(proof []merkletree.Hash, dataBlock []byte, rootHash merkletree.Hash) (bool, error) {
	return s.Tree.VerifyProof(proof, dataBlock, rootHash)
}

// GetFileContent returns the content of a file given its name.
func (s *Server) GetFileContent(fileHash merkletree.Hash) ([]byte, error) {
	filePath := filepath.Join(s.StoragePath, hex.EncodeToString(fileHash)[:10])

	fmt.Println("File path: ", filePath)
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %v", filePath, err)
	}
	return content, nil
}
