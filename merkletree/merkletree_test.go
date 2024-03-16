package merkletree

import (
	"testing"
)

func TestProofEvenBlocks(t *testing.T) {
	dataBlocks := [][]byte{
		[]byte("Block 1"),
		[]byte("Block 2"),
		[]byte("Block 3"),
		[]byte("Block 4"),
	}
	tree, err := NewMerkleTree(dataBlocks)
	if err != nil {
		t.Errorf(err.Error())
	}

	proof, err := tree.GenerateProof([]byte("Block 1"))
	if err != nil {
		t.Errorf(err.Error())
	}

	valid, err := tree.VerifyProof(proof, []byte("Block 1"), tree.Root.Hash)
	if err != nil {
		t.Errorf(err.Error())
	}

	if !valid {
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
	tree, err := NewMerkleTree(dataBlocks)
	if err != nil {
		t.Errorf(err.Error())
	}

	proof, err := tree.GenerateProof([]byte("Block 5"))
	if err != nil {
		t.Errorf(err.Error())
	}

	valid, err := tree.VerifyProof(proof, []byte("Block 5"), tree.Root.Hash)
	if err != nil {
		t.Errorf(err.Error())
	}

	if !valid {
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
	tree, err := NewMerkleTree(dataBlocks)
	if err != nil {
		t.Errorf(err.Error())
	}

	proof, err := tree.GenerateProof([]byte("Block 5"))
	if err != nil {
		t.Errorf(err.Error())
	}

	valid, err := tree.VerifyProof(proof, []byte("Block 4"), tree.Root.Hash)
	if err != nil {
		t.Errorf(err.Error())
	}

	if valid {
		t.Errorf("VerifyProof returned true for invalid proof")
	}
}
