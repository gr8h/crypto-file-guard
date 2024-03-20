package server

import (
	"fmt"

	"github.com/gr8h/crypto-file-guard/internal/merkletree"
)

type Server struct {
	Files []string // List of files
	Tree  *merkletree.MerkleTree
}

// NewServer initializes a new server instance without any files.
func NewServer() *Server {
	return &Server{}
}

// ProcessFiles processes the given set of files and constructs the Merkle tree.
func (s *Server) ProcessFiles(files []string) error {
	var data [][]byte
	for _, file := range files {
		data = append(data, []byte(file))
	}
	tree, err := merkletree.NewMerkleTree(data)
	if err != nil {
		return err
	}

	s.Files = files
	s.Tree = tree
	return nil
}

// GetFileAndProof returns the file at the given index and its Merkle proof.
func (s *Server) GetFileAndProof(index int) (string, [][]byte, error) {
	if index < 0 || index >= len(s.Files) {
		return "", nil, fmt.Errorf("server: index out of bounds")
	}

	data, err := s.Tree.GetLeafByIndex(index)
	if err != nil {
		return "", nil, err
	}

	proof, err := s.Tree.GenerateProof(data)
	if err != nil {
		return "", nil, err
	}
	return s.Files[index], proof, nil
}
