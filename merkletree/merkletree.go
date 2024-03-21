package merkletree

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
)

type Data []byte
type Hash []byte
type Proof []Hash

// MerkleTree represents the structure of a Merkle tree
type MerkleTree struct {
	Root   *MerkleNode
	Leaves []Hash
}

// MerkleNode represents a node in the Merkle tree
type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Hash  Hash
}

// NewMerkleNode creates a new Merkle node with the given left and right children and hash.
func NewMerkleNode(left, right *MerkleNode, hash Hash) (*MerkleNode, error) {
	if left == nil && right == nil && hash == nil {
		return nil, errors.New("merkletree: leaf node must have a hash")
	}
	node := &MerkleNode{}

	if left == nil && right == nil {
		// Leaf
		node.Hash = hash
	} else {
		prevHashes := append([]byte(left.Hash), right.Hash...)
		hashedData, err := HashData(prevHashes)
		if err != nil {
			return nil, fmt.Errorf("failed to hash node data: %w", err)
		}
		node.Hash = hashedData
	}

	node.Left = left
	node.Right = right

	return node, nil
}

// NewMerkleTree creates a new Merkle tree from the given data blocks.
func NewMerkleTree(dataHashes []Hash) (*MerkleTree, error) {
	if len(dataHashes) == 0 {
		return nil, errors.New("merkletree: dataBlocks cannot be empty")
	}

	var nodes []*MerkleNode
	var leaves []Hash
	for _, hash := range dataHashes {
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

// GenerateProof returns the Merkle proof for the given data block.
func (t *MerkleTree) GenerateProof(targetHash Hash) ([]Hash, error) {
	var proof []Hash

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

// GetLeafByData returns the index of the leaf node that contains the given data block.
func (t *MerkleTree) GetLeafIndex(targetHash Hash) (int, error) {
	for i, hash := range t.Leaves {
		if bytes.Equal(hash, targetHash) {
			return i, nil
		}
	}
	return -1, errors.New("merkletree: data block not found")
}

// GetLeafByIndex returns the leaf node at the given index.
func (t *MerkleTree) GetLeafByIndex(index int) (Hash, error) {
	if index < 0 || index >= len(t.Leaves) {
		return nil, errors.New("merkletree: index out of bounds")
	}
	return t.Leaves[index], nil
}

// GetRootHash returns the hash of the root of the Merkle Tree.
// If the root is nil, it returns nil and an error.
func (t *MerkleTree) GetRootHash() (Hash, error) {
	if t.Root == nil {
		return nil, errors.New("merkletree: the tree is empty or not initialized")
	}
	return t.Root.Hash, nil
}

// VerifyProof verifies the given proof for the given data block and root hash.
func (t *MerkleTree) VerifyProof(proof []Hash, targetHash Hash, rootHash Hash) (bool, error) {
	currentHash := targetHash
	index, err := t.GetLeafIndex(currentHash)
	if err != nil {
		return false, fmt.Errorf("failed to get leaf index: %w", err)
	}

	for _, hash := range proof {
		var dataToHash []byte

		// Determine the order of concatenation
		if index%2 == 0 {
			dataToHash = append(Hash(currentHash), Hash(hash)...)
		} else {
			dataToHash = append(Hash(hash), Hash(currentHash)...)
		}

		index = index / 2

		currentHash, err = HashData(dataToHash)
		if err != nil {
			return false, fmt.Errorf("failed to hash data: %w", err)
		}
	}

	return bytes.Equal(currentHash, rootHash), nil
}

// PrintTree prints the Merkle tree level by level.
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
