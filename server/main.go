package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/gr8h/crypto-file-guard/pkg/merkletree"
	pb "github.com/gr8h/crypto-file-guard/server/pb"
	server "github.com/gr8h/crypto-file-guard/server/pkg"
)

type fileGuardServer struct {
	pb.UnimplementedFileGuardServer
	Server *server.Server
}

func (fgs *fileGuardServer) NewSession(ctx context.Context, req *pb.NewSessionRequest) (*pb.NewSessionResponse, error) {

	sessionID := fgs.Server.NewSession()

	return &pb.NewSessionResponse{SessionId: &pb.SessionID{Value: sessionID}}, nil
}

func (fgs *fileGuardServer) ConstructMerkleTree(ctx context.Context, req *pb.ConstructMerkleTreeRequest) (*pb.ConstructMerkleTreeResponse, error) {
	rootHash, err := fgs.Server.ConstructMerkleTree(req.GetSessionId())
	if err != nil {
		return nil, err
	}

	// Convert the rootHash to the protobuf format
	protobufRootHash := &pb.Hash{Value: rootHash[:]}

	return &pb.ConstructMerkleTreeResponse{RootHash: protobufRootHash}, nil
}

func (fgs *fileGuardServer) GetFile(ctx context.Context, req *pb.GetFileRequest) (*pb.GetFileResponse, error) {
	content, err := fgs.Server.GetFile(req.GetSessionId(), int(req.GetIndex()))
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

	// Convert the proof to the protobuf format if necessary
	// This is a simplified representation
	var proofBytes []*pb.Hash
	for _, p := range proof {
		proofBytes = append(proofBytes, &pb.Hash{Value: p[:]})
	}

	return &pb.GetProofResponse{Proof: proofBytes}, nil
}

func (fgs *fileGuardServer) AddFile(ctx context.Context, req *pb.AddFileRequest) (*pb.AddFileResponse, error) {
	err := fgs.Server.AddFile(req.GetSessionId(), req.GetContent())
	if err != nil {
		return nil, err
	}
	return &pb.AddFileResponse{}, nil
}

func (fgs *fileGuardServer) VerifyProof(ctx context.Context, req *pb.VerifyProofRequest) (*pb.VerifyProofResponse, error) {
	proof := make(merkletree.Proof, len(req.GetProof()))
	for i, p := range req.GetProof() {
		proof[i] = p.GetValue()
	}

	valid, err := fgs.Server.VerifyProof(req.GetSessionId(), proof, req.GetTargetHash().GetValue(), req.GetRootHash().GetValue())
	if err != nil {
		return nil, err
	}
	return &pb.VerifyProofResponse{Valid: valid}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	// Initialize your Server struct here.
	serverInstance := server.NewServer("")

	pb.RegisterFileGuardServer(s, &fileGuardServer{Server: serverInstance})

	log.Printf("Server listening at %v", listener.Addr())

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
