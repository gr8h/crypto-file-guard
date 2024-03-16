package merkletree

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// hashData computes the SHA-256 hash of input data
func hashData(data []byte) []byte {
	sum := sha256.Sum256(data)
	return sum[:]
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

func NewMerkleNode(left, right *MerkleNode, hash []byte) *MerkleNode {
	node := &MerkleNode{}

	if left == nil && right == nil {
		// Leaf
		node.Hash = hash
	} else {
		prevHashes := append([]byte(left.Hash), right.Hash...)
		node.Hash = hashData(prevHashes)
	}

	node.Left = left
	node.Right = right

	return node
}

func NewMerkleTree(dataBlocks [][]byte) *MerkleTree {
	var hashes [][]byte
	for _, data := range dataBlocks {
		hashes = append(hashes, hashData(data))
	}

	var nodes []*MerkleNode
	var leaves [][]byte
	for _, hash := range hashes {
		nodes = append(nodes, NewMerkleNode(nil, nil, hash))
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
				right = NewMerkleNode(nil, nil, nodes[len(nodes)-1].Hash)
			}

			level = append(level, NewMerkleNode(left, right, nil))
		}

		nodes = level
	}

	return &MerkleTree{Root: nodes[0], Leaves: leaves}
}

func (t *MerkleTree) GenerateProof(dataBlock []byte) [][]byte {
	targetHash := hashData(dataBlock)
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
		return proof
	}
	return nil
}

func (t *MerkleTree) GetLeafIndex(dataBlock []byte) int {
	targetHash := hashData(dataBlock)
	for i, hash := range t.Leaves {
		if bytes.Equal(hash, targetHash) {
			return i
		}
	}
	return -1
}

func (t *MerkleTree) VerifyProof(proof [][]byte, dataBlock []byte, rootHash []byte) bool {
	targetHash := hashData(dataBlock)
	currentHash := targetHash
	index := t.GetLeafIndex(dataBlock)

	for _, hash := range proof {
		var dataToHash []byte

		// Determine the order of concatenation
		if index%2 == 0 {
			dataToHash = append([]byte(currentHash), []byte(hash)...)
		} else {
			dataToHash = append([]byte(hash), []byte(currentHash)...)
		}

		index = index / 2

		currentHash = hashData(dataToHash)
	}

	return bytes.Equal(currentHash, rootHash)
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
