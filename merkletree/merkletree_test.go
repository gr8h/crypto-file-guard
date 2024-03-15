package merkletree

import (
	"testing"
)

func TestNewMerkleNode(t *testing.T) {
	left := NewMerkleNode(nil, nil, []byte("left"))
	right := NewMerkleNode(nil, nil, []byte("right"))
	node := NewMerkleNode(left, right, nil)

	if node.Left != left || node.Right != right {
		t.Errorf("NewMerkleNode did not correctly set left and right nodes")
	}
}

func TestProofEvenBlocks(t *testing.T) {
	dataBlocks := [][]byte{
		[]byte("Block 1"),
		[]byte("Block 2"),
		[]byte("Block 3"),
		[]byte("Block 4"),
	}
	tree := NewMerkleTree(dataBlocks)
	proof := tree.GenerateProof([]byte("Block 1"))

	if !tree.VerifyProof(proof, []byte("Block 1"), tree.Root.Hash, 0) {
		t.Errorf("VerifyProof returned false for valid proof")
	}
}

func TestProofOddBlocks(t *testing.T) {
	dataBlocks := [][]byte{
		[]byte("Block 1"),
		[]byte("Block 2"),
		[]byte("Block 3"),
		[]byte("Block 4"),
		[]byte("Block 5"),
	}
	tree := NewMerkleTree(dataBlocks)
	proof := tree.GenerateProof([]byte("Block 5"))

	if !tree.VerifyProof(proof, []byte("Block 5"), tree.Root.Hash, 4) {
		t.Errorf("VerifyProof returned false for valid proof")
	}
}

func TestWrongProof(t *testing.T) {
	dataBlocks := [][]byte{
		[]byte("Block 1"),
		[]byte("Block 2"),
		[]byte("Block 3"),
		[]byte("Block 4"),
		[]byte("Block 5"),
	}
	tree := NewMerkleTree(dataBlocks)
	proof := tree.GenerateProof([]byte("Block 5"))

	if tree.VerifyProof(proof, []byte("Block 4"), tree.Root.Hash, 3) {
		t.Errorf("VerifyProof returned true for invalid proof")
	}
}
