package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/gr8h/crypto-file-guard/client/pb"
	client "github.com/gr8h/crypto-file-guard/client/pkg"
)

func main() {
	serverBaseUrl := os.Getenv("SERVERBASEURL")
	if serverBaseUrl == "" {
		serverBaseUrl = "localhost:8080"
	}

	conn, err := grpc.Dial(serverBaseUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	defer conn.Close()

	server := pb.NewFileGuardClient(conn)

	// ----------------- Files ----------------- ///
	files := []string{"0.txt", "1.txt", "2.txt", "3.txt", "4.txt", "5.txt", "6.txt"}

	// ----------------- CMD ----------------- ///
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("> Enter the session ID or leave empty to generate a new one:")
	sessionId, _ := reader.ReadString('\n')
	sessionId = strings.TrimSpace(sessionId)

	// ----------------- Session ----------------- ///
	if sessionId == "" {
		s, err := server.NewSession(context.Background(), &pb.NewSessionRequest{})
		if err != nil {
			log.Fatalf("error creating new session: %v", err)
		}
		sessionId = s.GetSessionId()
		fmt.Printf("New session ID: %s\n", sessionId)
	} else {
		fmt.Printf("Using existing session ID: %s\n", sessionId)
	}

	// ----------------- Action ----------------- ///
	fmt.Println("> Enter the action you want to perform: \n - construct \n - verify")
	action, _ := reader.ReadString('\n')
	action = strings.TrimSpace(action)

	switch action {
	case "construct":
		fmt.Println("> Please specify number of file(s) to add.")
		input, _ := reader.ReadString('\n')
		inputT := strings.TrimSpace(input)
		numberOfFiles, _ := strconv.Atoi(inputT)

		rootHash, err := client.ConstructMerkleTree(server, sessionId, files[:numberOfFiles])
		if err != nil {
			log.Fatalf("error constructing Merkle tree: %v", err)
		}

		err = client.WriteFile(sessionId, "roothash", rootHash.GetRootHash().Value)
		if err != nil {
			log.Fatalf("error writing root hash: %v", err)
		}

		fmt.Printf("Merkle tree constructed successfully - Root hash: %x\n", rootHash.GetRootHash().Value)

	case "verify":
		fmt.Println("> Please specify the file index to verify.")
		input, _ := reader.ReadString('\n')
		inputT := strings.TrimSpace(input)
		fileIndex, _ := strconv.Atoi(inputT)

		proof, err := server.GetProof(context.Background(), &pb.GetProofRequest{Index: int32(fileIndex), SessionId: sessionId})
		if err != nil {
			log.Fatalf("error getting proof: %v", err)
		}

		fileContent, err := server.GetFile(context.Background(), &pb.GetFileRequest{Index: int32(fileIndex), SessionId: sessionId})
		if err != nil {
			log.Fatalf("error getting file content: %v", err)
		}

		rootHash, err := client.ReadFile(sessionId, "roothash")
		if err != nil {
			log.Fatalf("error reading root hash: %v", err)
		}

		fileHash, err := client.HashData(fileContent.GetFileContent())
		if err != nil {
			log.Fatalf("error hashing file content: %v", err)
		}

		isValid, err := client.VerifyProof(proof.GetProof(), fileHash, rootHash, fileIndex)
		if err != nil {
			log.Fatalf("error verifying proof: %v", err)
		}

		if isValid {
			fmt.Println("Proof is valid")
		} else {
			fmt.Println("Proof is invalid")
		}

	default:
		fmt.Println("Invalid action. Use -action with 'addfiles', 'getproof', or 'verify'")
	}
}
