package merkletree

import (
	"bytes"
	"testing"
)

func TestEmptyTree(t *testing.T) {
	_, err := NewMerkleTree([][]byte{})
	if err == nil {
		t.Errorf("Expected error for empty data blocks, got nil")
	}
}

func TestSingleNodeTree(t *testing.T) {
	dataBlocks := [][]byte{[]byte("Single Block")}
	tree, err := NewMerkleTree(dataBlocks)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	rootHash, err := tree.GetRootHash()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	expectedHash, _ := hashData([]byte("Single Block"))
	if !bytes.Equal(rootHash, expectedHash) {
		t.Errorf("Root hash does not match expected single block hash")
	}
}

func TestProofForNonexistentBlock(t *testing.T) {
	dataBlocks := [][]byte{
		[]byte("Block 1"),
		[]byte("Block 2"),
	}
	tree, err := NewMerkleTree(dataBlocks)
	if err != nil {
		t.Errorf(err.Error())
	}

	_, err = tree.GenerateProof([]byte("Nonexistent Block"))
	if err == nil {
		t.Errorf("Expected error when generating proof for nonexistent block, got nil")
	}
}
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

	targetHash, err := tree.GetLeafByIndex(0)
	if err != nil {
		t.Errorf(err.Error())
	}

	proof, err := tree.GenerateProof(targetHash)
	if err != nil {
		t.Errorf(err.Error())
	}

	valid, err := tree.VerifyProof(proof, dataBlocks[0], tree.Root.Hash)
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

	targetHash, err := tree.GetLeafByIndex(4)
	if err != nil {
		t.Errorf(err.Error())
	}

	proof, err := tree.GenerateProof(targetHash)
	if err != nil {
		t.Errorf(err.Error())
	}

	valid, err := tree.VerifyProof(proof, dataBlocks[4], tree.Root.Hash)
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

	targetHash, err := tree.GetLeafByIndex(4)
	if err != nil {
		t.Errorf(err.Error())
	}

	proof, err := tree.GenerateProof(targetHash)
	if err != nil {
		t.Errorf(err.Error())
	}

	valid, err := tree.VerifyProof(proof, dataBlocks[3], tree.Root.Hash)
	if err != nil {
		t.Errorf(err.Error())
	}

	if valid {
		t.Errorf("VerifyProof returned true for invalid proof")
	}
}
