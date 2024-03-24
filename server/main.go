package main

import (
	"context"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/gr8h/crypto-file-guard/server/pb"
	server "github.com/gr8h/crypto-file-guard/server/pkg"
)

type fileGuardServer struct {
	pb.UnimplementedFileGuardServer
	Server *server.Server
}

func (fgs *fileGuardServer) NewSession(ctx context.Context, req *pb.NewSessionRequest) (*pb.NewSessionResponse, error) {

	sessionId := fgs.Server.NewSession()

	return &pb.NewSessionResponse{SessionId: sessionId}, nil
}

func (fgs *fileGuardServer) ConstructMerkleTree(ctx context.Context, req *pb.ConstructMerkleTreeRequest) (*pb.ConstructMerkleTreeResponse, error) {
	rootHash, err := fgs.Server.ConstructMerkleTree(req.GetSessionId())
	if err != nil {
		return nil, err
	}

	protobufRootHash := &pb.Hash{Value: rootHash[:]}

	return &pb.ConstructMerkleTreeResponse{RootHash: protobufRootHash}, nil
}

func (fgs *fileGuardServer) GetFile(ctx context.Context, req *pb.GetFileRequest) (*pb.GetFileResponse, error) {
	content, err := fgs.Server.GetFileContent(req.GetSessionId(), int(req.GetIndex()))
	if err != nil {
		return nil, err
	}
	return &pb.GetFileResponse{FileContent: content[:]}, nil
}

func (fgs *fileGuardServer) GetProof(ctx context.Context, req *pb.GetProofRequest) (*pb.GetProofResponse, error) {
	proof, err := fgs.Server.GetProof(req.GetSessionId(), int(req.GetIndex()))
	if err != nil {
		return nil, err
	}

	return &pb.GetProofResponse{Proof: proof}, nil
}

func (fgs *fileGuardServer) AddFile(ctx context.Context, req *pb.AddFileRequest) (*pb.AddFileResponse, error) {
	err := fgs.Server.AddFile(req.GetSessionId(), req.GetContent())
	if err != nil {
		return nil, err
	}
	return &pb.AddFileResponse{}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	baseStoragePath := os.Getenv("BASESTORAGEPATH")

	serverInstance := server.NewServer(baseStoragePath)

	pb.RegisterFileGuardServer(s, &fileGuardServer{Server: serverInstance})

	log.Printf("Server listening at %v", listener.Addr())

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
