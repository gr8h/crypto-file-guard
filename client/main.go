package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/gr8h/crypto-file-guard/client/pb"
)

const StoragePath = "./client/files"

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewFileGuardClient(conn)

	fmt.Println("Client connected to server")

	// add the current working directory to filePath
	rootHash, err := ConstructMerkleTree(client)
	if err != nil {
		log.Fatalf("error constructing Merkle tree: %v", err)
	}

	fmt.Println("Merkle tree constructed successfully")

	fmt.Printf("Root hash: %x\n", rootHash.GetRootHash().Value)

	// Get proof for the first file
	proof, err := client.GetProof(context.Background(), &pb.GetProofRequest{Index: 0})
	if err != nil {
		log.Fatalf("error getting proof: %v", err)
	}

	fileContenet, err := client.GetFile(context.Background(), &pb.GetFileRequest{Index: 0})
	if err != nil {
		log.Fatalf("error getting file content: %v", err)
	}

	fmt.Printf("File content: %x\n", fileContenet.GetFileContent())
	targetHashPB := &pb.Hash{Value: fileContenet.GetFileContent()}

	isValid, err := client.VerifyProof(context.Background(), &pb.VerifyProofRequest{Proof: proof.GetProof(), TargetHash: targetHashPB, RootHash: rootHash.GetRootHash()})
	if err != nil {
		log.Fatalf("error verifying proof: %v", err)
	}

	if isValid.GetValid() {
		fmt.Println("Proof is valid")
	} else {
		fmt.Println("Proof is invalid")
	}
}

func ConstructMerkleTree(client pb.FileGuardClient) (*pb.ConstructMerkleTreeResponse, error) {
	files := []string{"0.txt", "1.txt", "2.txt", "3.txt", "4.txt"}
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

		_, err = client.AddFile(context.Background(), &pb.AddFileRequest{Content: content})
		if err != nil {
			return nil, fmt.Errorf("error uploading file %s: %v", filepath.Base(file), err)
		}
	}

	rootHash, err := client.ConstructMerkleTree(context.Background(), &pb.ConstructMerkleTreeRequest{})
	if err != nil {
		return nil, fmt.Errorf("error constructing Merkle tree: %v", err)
	}
	return rootHash, nil
}
