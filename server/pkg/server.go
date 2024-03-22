package server

import (
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gr8h/crypto-file-guard/pkg/merkletree"
)

type Server struct {
	Files       []merkletree.Hash
	Tree        *merkletree.MerkleTree
	StoragePath string
}

// NewServer initializes a new server instance without any files.
func NewServer(storagePath string) *Server {

	storagePath = filepath.Join("./server/files/", storagePath, "/")

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
func (s *Server) ConstructMerkleTree() (merkletree.Hash, error) {
	s.Tree = nil
	var hashes []merkletree.Hash

	for _, fileHash := range s.Files {

		filePath := filepath.Join(s.StoragePath, hex.EncodeToString(fileHash)[:10])

		content, err := os.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to read file %s: %v", filePath, err)
		}

		hashedData, err := merkletree.HashData(content)
		if err != nil {
			return nil, fmt.Errorf("failed to hash node data: %w", err)
		}

		hashes = append(hashes, hashedData)
	}

	tree, err := merkletree.NewMerkleTree(hashes)
	if err != nil {
		return nil, fmt.Errorf("failed to construct/update Merkle tree: %v", err)
	}

	s.Tree = tree

	tree.PrintTree()
	s.Files = nil

	return tree.Root.Hash, nil
}

// GetProof returns the Merkle proof for the file at the given index.
func (s *Server) GetProof(index int) (merkletree.Proof, error) {
	if index < 0 || index >= len(s.Tree.Leaves) {
		return nil, fmt.Errorf("server: index out of bounds")
	}

	fileHash, err := s.Tree.GetLeafByIndex(index)
	if err != nil {
		return nil, err
	}

	proof, err := s.Tree.GenerateProof(fileHash)
	if err != nil {
		return nil, err
	}

	return proof, nil
}

// GetFile returns the content of a file given its index.
func (s *Server) GetFile(index int) (merkletree.Hash, error) {
	return s.Tree.GetLeafByIndex(index)
}

// VerifyProof verifies the given proof for the given data block and root hash.
func (s *Server) VerifyProof(proof merkletree.Proof, fileHash merkletree.Hash, rootHash merkletree.Hash) (bool, error) {
	return s.Tree.VerifyProof(proof, fileHash, rootHash)
}
