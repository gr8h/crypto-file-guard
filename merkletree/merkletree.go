package merkletree

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
)

// hashData computes the SHA-256 hash of input data
func hashData(data []byte) ([]byte, error) {

	if data == nil {
		return nil, errors.New("merkletree: data is empty")
	}

	sum := sha256.Sum256(data)
	return sum[:], nil
}

// MerkleTree represents the structure of a Merkle tree
type MerkleTree struct {
	Root   *MerkleNode
	Leaves [][]byte
}

// MerkleNode represents a node in the Merkle tree
type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Hash  []byte
}

func NewMerkleNode(left, right *MerkleNode, hash []byte) (*MerkleNode, error) {
	if left == nil && right == nil && hash == nil {
		return nil, errors.New("merkletree: leaf node must have a hash")
	}
	node := &MerkleNode{}

	if left == nil && right == nil {
		// Leaf
		node.Hash = hash
	} else {
		prevHashes := append([]byte(left.Hash), right.Hash...)
		hashedData, err := hashData(prevHashes)
		if err != nil {
			return nil, fmt.Errorf("failed to hash node data: %w", err)
		}
		node.Hash = hashedData
	}

	node.Left = left
	node.Right = right

	return node, nil
}

func NewMerkleTree(dataBlocks [][]byte) (*MerkleTree, error) {
	if len(dataBlocks) == 0 {
		return nil, errors.New("merkletree: dataBlocks cannot be empty")
	}

	var hashes [][]byte
	for _, data := range dataBlocks {

		hashedData, err := hashData(data)
		if err != nil {
			return nil, fmt.Errorf("failed to hash node data: %w", err)
		}

		hashes = append(hashes, hashedData)
	}

	var nodes []*MerkleNode
	var leaves [][]byte
	for _, hash := range hashes {
		node, err := NewMerkleNode(nil, nil, hash)
		if err != nil {
			return nil, fmt.Errorf("failed to create new merkle node: %w", err)
		}
		nodes = append(nodes, node)
		leaves = append(leaves, hash)
	}

	for len(nodes) > 1 {
		var level []*MerkleNode

		for i := 0; i < len(nodes); i += 2 {
			left := nodes[i]

			var right *MerkleNode
			if i+1 < len(nodes) {
				right = nodes[i+1]
			} else {
				right, _ = NewMerkleNode(nil, nil, nodes[len(nodes)-1].Hash)
			}

			parentNode, err := NewMerkleNode(left, right, nil)
			if err != nil {
				return nil, fmt.Errorf("failed to create a parent Merkle node: %w", err)
			}
			level = append(level, parentNode)
		}

		nodes = level
	}

	return &MerkleTree{Root: nodes[0], Leaves: leaves}, nil
}

func (t *MerkleTree) GenerateProof(dataBlock []byte) ([][]byte, error) {
	targetHash, err := hashData(dataBlock)
	if err != nil {
		return nil, fmt.Errorf("failed to hash data block: %w", err)
	}
	var proof [][]byte

	var generate func(node *MerkleNode) bool
	generate = func(node *MerkleNode) bool {
		if node == nil {
			return false
		}
		if bytes.Equal(node.Hash, targetHash) {
			return true
		}

		if generate(node.Left) {
			if node.Right != nil {
				proof = append(proof, node.Right.Hash)
			}
			return true
		} else if generate(node.Right) {
			if node.Left != nil {
				proof = append(proof, node.Left.Hash)
			}
			return true
		}
		return false
	}

	if generate(t.Root) {
		return proof, nil
	}
	return nil, errors.New("merkletree: failed to generate proof")
}

func (t *MerkleTree) GetLeafIndex(dataBlock []byte) (int, error) {
	targetHash, err := hashData(dataBlock)
	if err != nil {
		return -1, fmt.Errorf("failed to hash data block: %w", err)
	}

	for i, hash := range t.Leaves {
		if bytes.Equal(hash, targetHash) {
			return i, nil
		}
	}
	return -1, errors.New("merkletree: data block not found")
}

func (t *MerkleTree) VerifyProof(proof [][]byte, dataBlock []byte, rootHash []byte) (bool, error) {
	targetHash, err := hashData(dataBlock)
	if err != nil {
		return false, fmt.Errorf("failed to hash data block: %w", err)
	}
	currentHash := targetHash
	index, err := t.GetLeafIndex(dataBlock)
	if err != nil {
		return false, fmt.Errorf("failed to get leaf index: %w", err)
	}

	for _, hash := range proof {
		var dataToHash []byte

		// Determine the order of concatenation
		if index%2 == 0 {
			dataToHash = append([]byte(currentHash), []byte(hash)...)
		} else {
			dataToHash = append([]byte(hash), []byte(currentHash)...)
		}

		index = index / 2

		currentHash, err = hashData(dataToHash)
		if err != nil {
			return false, fmt.Errorf("failed to hash data: %w", err)
		}
	}

	return bytes.Equal(currentHash, rootHash), nil
}

func (t *MerkleTree) PrintTree() {
	level := 0
	nodes := []*MerkleNode{t.Root}
	for len(nodes) > 0 {
		count := len(nodes)
		fmt.Printf("Level %d: ", level)
		for i := 0; i < count; i++ {
			node := nodes[0]
			nodes = nodes[1:]
			fmt.Printf("[%s] ", hex.EncodeToString(node.Hash)[:10])
			if node.Left != nil {
				nodes = append(nodes, node.Left)
			}
			if node.Right != nil {
				nodes = append(nodes, node.Right)
			}
		}
		fmt.Println()
		level++
	}
}
