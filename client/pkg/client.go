package client

import (
	"bytes"
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	pb "github.com/gr8h/crypto-file-guard/client/pb"
)

const StoragePath = "./client/files"

func HashData(data []byte) ([]byte, error) {

	if data == nil {
		return nil, errors.New("data is empty")
	}

	sum := sha256.Sum256(data)
	return sum[:], nil
}

// write file
func WriteFile(fileName string, dir string, content []byte) error {

	filePath := filepath.Join(StoragePath, dir, fileName)
	fmt.Println("Writing file to: ", filePath)

	// create directory if it does not exist
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		return fmt.Errorf("client: failed to create directory: %v", err)
	}

	if err := os.WriteFile(filePath, content, 0644); err != nil {
		return fmt.Errorf("client: failed to write file %s: %v", filePath, err)
	}
	return nil
}

// read file
func ReadFile(fileName string, dir string) ([]byte, error) {
	filePath := filepath.Join(StoragePath, dir, fileName)

	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("client: failed to read file %s: %v", filePath, err)
	}
	return content, nil
}

// verify proof
func VerifyProof(proof [][]byte, targetHash []byte, rootHash []byte, index int) (bool, error) {
	currentHash := targetHash

	for _, hash := range proof {
		var dataToHash []byte

		if index%2 == 0 {
			dataToHash = append([]byte(currentHash), []byte(hash)...)
		} else {
			dataToHash = append([]byte(hash), []byte(currentHash)...)
		}

		index = index / 2

		var err error
		currentHash, err = HashData(dataToHash)
		if err != nil {
			return false, fmt.Errorf("failed to hash data: %w", err)
		}
	}

	return bytes.Equal(currentHash, rootHash), nil
}

// construct Merkle tree
func ConstructMerkleTree(server pb.FileGuardClient, sessionId string, files []string) (*pb.ConstructMerkleTreeResponse, error) {
	for _, file := range files {
		file = filepath.Join(StoragePath, file)
		file, err := filepath.Abs(file)
		if err != nil {
			return nil, fmt.Errorf("error getting absolute path for file %s: %v", file, err)
		}

		content, err := os.ReadFile(file)
		if err != nil {
			return nil, fmt.Errorf("error reading file %s: %v", file, err)
		}

		_, err = server.AddFile(context.Background(), &pb.AddFileRequest{Content: content, SessionId: sessionId})
		if err != nil {
			return nil, fmt.Errorf("error uploading file %s: %v", filepath.Base(file), err)
		}
	}

	rootHash, err := server.ConstructMerkleTree(context.Background(), &pb.ConstructMerkleTreeRequest{SessionId: sessionId})
	if err != nil {
		return nil, fmt.Errorf("error constructing Merkle tree: %v", err)
	}
	return rootHash, nil
}
