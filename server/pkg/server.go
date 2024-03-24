package server

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/gr8h/crypto-file-guard/pkg/merkletree"
)

type ClientSession struct {
	Files       []merkletree.Hash
	Tree        *merkletree.MerkleTree
	StoragePath string
}

type Server struct {
	mu              sync.Mutex
	sessions        map[string]*ClientSession
	SessionId       string
	BaseStoragePath string
}

// generateSessionId generates a random session ID.
func generateSessionId() string {
	length := 16
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}

// createSession initializes a new session with the given session ID.
func (s *Server) createSession(sessionID string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	storagePath := filepath.Join(s.BaseStoragePath, sessionID)

	s.sessions[sessionID] = &ClientSession{
		Files:       []merkletree.Hash{},
		Tree:        nil,
		StoragePath: storagePath,
	}
}

// NewServer initializes a new server instance without any files.
func NewServer(baseStoragePath string) *Server {

	if baseStoragePath == "" {
		baseStoragePath = filepath.Join("./server/files/")
	}

	server := &Server{
		sessions:        make(map[string]*ClientSession),
		BaseStoragePath: baseStoragePath,
	}

	return server
}

// NewSession creates a new session and returns the session ID.
func (s *Server) NewSession() string {
	sessionId := generateSessionId()
	s.createSession(sessionId)
	s.SessionId = sessionId

	fmt.Println("Server: Number of sessions: ", len(s.sessions))

	return sessionId
}

func (s *Server) AddFile(sessionID string, content []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if sessionID == "" {
		return fmt.Errorf("server: session ID is empty")
	}

	storagePath := filepath.Join(s.BaseStoragePath, sessionID)
	if err := os.MkdirAll(storagePath, 0755); err != nil {
		return fmt.Errorf("server: failed to create storage directory: %v", err)
	}

	fileHash, err := merkletree.HashData(content)
	if err != nil {
		return fmt.Errorf("server: failed to hash file content: %v", err)
	}

	s.sessions[sessionID].Files = append(s.sessions[sessionID].Files, fileHash)

	fileIndex, err := getIndex(s.sessions[sessionID].Files, fileHash)
	if err != nil {
		return fmt.Errorf("server: file does not exist")
	}

	filePath := filepath.Join(storagePath, strconv.Itoa(fileIndex))

	if err := os.WriteFile(filePath, content, 0644); err != nil {
		return fmt.Errorf("server: failed to write file %s: %v", filePath, err)
	}
	return nil
}

// ConstructMerkleTree reads the content of all stored files and constructs/updates the Merkle tree.
func (s *Server) ConstructMerkleTree(sessionID string) (merkletree.Hash, error) {

	if sessionID == "" {
		return nil, fmt.Errorf("server: session ID is empty")
	}

	var hashes []merkletree.Hash

	for _, fileHash := range s.sessions[sessionID].Files {

		storagePath := filepath.Join(s.BaseStoragePath, sessionID)
		fileIndex, err := getIndex(s.sessions[sessionID].Files, fileHash)
		if err != nil {
			return nil, fmt.Errorf("server: file does not exists")
		}
		filePath := filepath.Join(storagePath, strconv.Itoa(fileIndex))

		content, err := os.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("server: failed to read file %s: %v", filePath, err)
		}

		hashedData, err := merkletree.HashData(content)
		if err != nil {
			return nil, fmt.Errorf("server: failed to hash node data: %w", err)
		}

		hashes = append(hashes, hashedData)
	}

	tree, err := merkletree.NewMerkleTree(hashes)
	if err != nil {
		return nil, fmt.Errorf("server: failed to construct/update Merkle tree: %v", err)
	}

	s.sessions[sessionID].Tree = tree

	s.sessions[sessionID].Files = nil

	tree.PrintTree()
	return tree.Root.Hash, nil
}

// GetProof returns the Merkle proof for the file at the given index.
func (s *Server) GetProof(sessionID string, index int) ([][]byte, error) {

	if sessionID == "" {
		return nil, fmt.Errorf("server: session ID is empty")
	}

	if index < 0 || index >= len(s.sessions[sessionID].Tree.Leaves) {
		return nil, fmt.Errorf("server: index out of bounds")
	}

	fileHash, err := s.sessions[sessionID].Tree.GetLeafByIndex(index)
	if err != nil {
		return nil, err
	}

	proof, err := s.sessions[sessionID].Tree.GenerateProof(fileHash)
	if err != nil {
		return nil, err
	}

	return proof, nil
}

// GetFile returns the content of a file given its index.
func (s *Server) GetFile(sessionID string, index int) (merkletree.Hash, error) {
	if sessionID == "" {
		return nil, fmt.Errorf("server: session ID is empty")
	}
	return s.sessions[sessionID].Tree.GetLeafByIndex(index)
}

// get file content
func (s *Server) GetFileContent(sessionID string, index int) ([]byte, error) {
	if sessionID == "" {
		return nil, fmt.Errorf("server: session ID is empty")
	}

	storagePath := filepath.Join(s.BaseStoragePath, sessionID)
	filePath := filepath.Join(storagePath, strconv.Itoa(index))

	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("server: failed to read file %s: %v", filePath, err)
	}
	return content, nil
}

// indexOf returns the index of the given item in the slice.
func getIndex(files []merkletree.Hash, item merkletree.Hash) (int, error) {
	for i, v := range files {
		if bytes.Equal(v, item) {
			return i, nil
		}
	}
	return -1, fmt.Errorf("server: item not found")
}
