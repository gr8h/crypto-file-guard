package main

import (
	"context"
	"log"
	"net"

	"github.com/heshamshabanah/crypto-file-guard/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type fileGuardServer struct {
	pb.UnimplementedFileGuardServer
	Server *server.Server
}

func (fgs *fileGuardServer) NewServer(ctx context.Context, req *pb.NewServerRequest) (*pb.NewServerResponse, error) {
	fgs.Server = server.NewServer(req.GetStoragePath())
	return &pb.NewServerResponse{}, nil
}

func (fgs *fileGuardServer) ConstructMerkleTree(ctx context.Context, req *pb.ConstructMerkleTreeRequest) (*pb.ConstructMerkleTreeResponse, error) {
	err := fgs.Server.ConstructMerkleTree()
	if err != nil {
		return nil, err
	}
	return &pb.ConstructMerkleTreeResponse{}, nil
}

func (fgs *fileGuardServer) GetFile(ctx context.Context, req *pb.GetFileRequest) (*pb.GetFileResponse, error) {
	content, err := fgs.Server.GetFile(int(req.GetIndex()))
	if err != nil {
		return nil, err
	}
	return &pb.GetFileResponse{FileContent: content[:]}, nil
}

func (fgs *fileGuardServer) GetProof(ctx context.Context, req *pb.GetProofRequest) (*pb.GetProofResponse, error) {
	proof, err := fgs.Server.GetProof(int(req.GetIndex()))
	if err != nil {
		return nil, err
	}

	// Convert the proof to the protobuf format if necessary
	// This is a simplified representation
	var proofBytes []*pb.Hash
	for _, p := range proof {
		proofBytes = append(proofBytes, &pb.Hash{Value: p[:]})
	}

	return &pb.GetProofResponse{Proof: proofBytes}, nil
}

func (fgs *fileGuardServer) AddFile(ctx context.Context, req *pb.AddFileRequest) (*pb.AddFileResponse, error) {
	err := fgs.Server.AddFile(req.GetContent())
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

	// Initialize your Server struct here.
	storagePath := "test"
	serverInstance := server.NewServer(storagePath)

	pb.RegisterFileGuardServer(s, &fileGuardServer{Server: serverInstance})

	log.Printf("Server listening at %v", listener.Addr())

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
