// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: fileguard.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Hash represents the hash data. Assuming hash is a byte array.
type Hash struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Hash) Reset() {
	*x = Hash{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fileguard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hash) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hash) ProtoMessage() {}

func (x *Hash) ProtoReflect() protoreflect.Message {
	mi := &file_fileguard_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hash.ProtoReflect.Descriptor instead.
func (*Hash) Descriptor() ([]byte, []int) {
	return file_fileguard_proto_rawDescGZIP(), []int{0}
}

func (x *Hash) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

// MerkleNode represents a node in the Merkle tree.
type MerkleNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Left  *MerkleNode `protobuf:"bytes,1,opt,name=left,proto3" json:"left,omitempty"`
	Right *MerkleNode `protobuf:"bytes,2,opt,name=right,proto3" json:"right,omitempty"`
	Hash  *Hash       `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *MerkleNode) Reset() {
	*x = MerkleNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fileguard_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MerkleNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MerkleNode) ProtoMessage() {}

func (x *MerkleNode) ProtoReflect() protoreflect.Message {
	mi := &file_fileguard_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MerkleNode.ProtoReflect.Descriptor instead.
func (*MerkleNode) Descriptor() ([]byte, []int) {
	return file_fileguard_proto_rawDescGZIP(), []int{1}
}

func (x *MerkleNode) GetLeft() *MerkleNode {
	if x != nil {
		return x.Left
	}
	return nil
}

func (x *MerkleNode) GetRight() *MerkleNode {
	if x != nil {
		return x.Right
	}
	return nil
}

func (x *MerkleNode) GetHash() *Hash {
	if x != nil {
		return x.Hash
	}
	return nil
}

// MerkleTree represents the structure of a Merkle tree.
type MerkleTree struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Root   *MerkleNode `protobuf:"bytes,1,opt,name=root,proto3" json:"root,omitempty"`
	Leaves []*Hash     `protobuf:"bytes,2,rep,name=leaves,proto3" json:"leaves,omitempty"`
}

func (x *MerkleTree) Reset() {
	*x = MerkleTree{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fileguard_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MerkleTree) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MerkleTree) ProtoMessage() {}

func (x *MerkleTree) ProtoReflect() protoreflect.Message {
	mi := &file_fileguard_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MerkleTree.ProtoReflect.Descriptor instead.
func (*MerkleTree) Descriptor() ([]byte, []int) {
	return file_fileguard_proto_rawDescGZIP(), []int{2}
}

func (x *MerkleTree) GetRoot() *MerkleNode {
	if x != nil {
		return x.Root
	}
	return nil
}

func (x *MerkleTree) GetLeaves() []*Hash {
	if x != nil {
		return x.Leaves
	}
	return nil
}

// Server represents the server managing files with a Merkle tree.
type Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files       []*Hash     `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	Tree        *MerkleTree `protobuf:"bytes,2,opt,name=tree,proto3" json:"tree,omitempty"`
	StoragePath string      `protobuf:"bytes,3,opt,name=storagePath,proto3" json:"storagePath,omitempty"`
}

func (x *Server) Reset() {
	*x = Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fileguard_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_fileguard_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_fileguard_proto_rawDescGZIP(), []int{3}
}

func (x *Server) GetFiles() []*Hash {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *Server) GetTree() *MerkleTree {
	if x != nil {
		return x.Tree
	}
	return nil
}

func (x *Server) GetStoragePath() string {
	if x != nil {
		return x.StoragePath
	}
	return ""
}

// Request and Response messages for each service method
type NewServerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoragePath string `protobuf:"bytes,1,opt,name=storagePath,proto3" json:"storagePath,omitempty"`
}

func (x *NewServerRequest) Reset() {
	*x = NewServerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fileguard_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewServerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewServerRequest) ProtoMessage() {}

func (x *NewServerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fileguard_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewServerRequest.ProtoReflect.Descriptor instead.
func (*NewServerRequest) Descriptor() ([]byte, []int) {
	return file_fileguard_proto_rawDescGZIP(), []int{4}
}

func (x *NewServerRequest) GetStoragePath() string {
	if x != nil {
		return x.StoragePath
	}
	return ""
}

type NewServerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server *Server `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
}

func (x *NewServerResponse) Reset() {
	*x = NewServerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fileguard_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewServerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewServerResponse) ProtoMessage() {}

func (x *NewServerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fileguard_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewServerResponse.ProtoReflect.Descriptor instead.
func (*NewServerResponse) Descriptor() ([]byte, []int) {
	return file_fileguard_proto_rawDescGZIP(), []int{5}
}

func (x *NewServerResponse) GetServer() *Server {
	if x != nil {
		return x.Server
	}
	return nil
}

type AddFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *AddFileRequest) Reset() {
	*x = AddFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fileguard_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFileRequest) ProtoMessage() {}

func (x *AddFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fileguard_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFileRequest.ProtoReflect.Descriptor instead.
func (*AddFileRequest) Descriptor() ([]byte, []int) {
	return file_fileguard_proto_rawDescGZIP(), []int{6}
}

func (x *AddFileRequest) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type AddFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddFileResponse) Reset() {
	*x = AddFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fileguard_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFileResponse) ProtoMessage() {}

func (x *AddFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fileguard_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFileResponse.ProtoReflect.Descriptor instead.
func (*AddFileResponse) Descriptor() ([]byte, []int) {
	return file_fileguard_proto_rawDescGZIP(), []int{7}
}

type ConstructMerkleTreeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConstructMerkleTreeRequest) Reset() {
	*x = ConstructMerkleTreeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fileguard_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConstructMerkleTreeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConstructMerkleTreeRequest) ProtoMessage() {}

func (x *ConstructMerkleTreeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fileguard_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConstructMerkleTreeRequest.ProtoReflect.Descriptor instead.
func (*ConstructMerkleTreeRequest) Descriptor() ([]byte, []int) {
	return file_fileguard_proto_rawDescGZIP(), []int{8}
}

type ConstructMerkleTreeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RootHash *Hash `protobuf:"bytes,1,opt,name=rootHash,proto3" json:"rootHash,omitempty"`
}

func (x *ConstructMerkleTreeResponse) Reset() {
	*x = ConstructMerkleTreeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fileguard_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConstructMerkleTreeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConstructMerkleTreeResponse) ProtoMessage() {}

func (x *ConstructMerkleTreeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fileguard_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConstructMerkleTreeResponse.ProtoReflect.Descriptor instead.
func (*ConstructMerkleTreeResponse) Descriptor() ([]byte, []int) {
	return file_fileguard_proto_rawDescGZIP(), []int{9}
}

func (x *ConstructMerkleTreeResponse) GetRootHash() *Hash {
	if x != nil {
		return x.RootHash
	}
	return nil
}

type GetProofRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index int32 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
}

func (x *GetProofRequest) Reset() {
	*x = GetProofRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fileguard_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProofRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProofRequest) ProtoMessage() {}

func (x *GetProofRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fileguard_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProofRequest.ProtoReflect.Descriptor instead.
func (*GetProofRequest) Descriptor() ([]byte, []int) {
	return file_fileguard_proto_rawDescGZIP(), []int{10}
}

func (x *GetProofRequest) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

type GetProofResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Proof []*Hash `protobuf:"bytes,1,rep,name=proof,proto3" json:"proof,omitempty"`
}

func (x *GetProofResponse) Reset() {
	*x = GetProofResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fileguard_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProofResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProofResponse) ProtoMessage() {}

func (x *GetProofResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fileguard_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProofResponse.ProtoReflect.Descriptor instead.
func (*GetProofResponse) Descriptor() ([]byte, []int) {
	return file_fileguard_proto_rawDescGZIP(), []int{11}
}

func (x *GetProofResponse) GetProof() []*Hash {
	if x != nil {
		return x.Proof
	}
	return nil
}

type GetFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index int32 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
}

func (x *GetFileRequest) Reset() {
	*x = GetFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fileguard_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileRequest) ProtoMessage() {}

func (x *GetFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fileguard_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileRequest.ProtoReflect.Descriptor instead.
func (*GetFileRequest) Descriptor() ([]byte, []int) {
	return file_fileguard_proto_rawDescGZIP(), []int{12}
}

func (x *GetFileRequest) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

type GetFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileContent []byte `protobuf:"bytes,1,opt,name=fileContent,proto3" json:"fileContent,omitempty"`
}

func (x *GetFileResponse) Reset() {
	*x = GetFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fileguard_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileResponse) ProtoMessage() {}

func (x *GetFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fileguard_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileResponse.ProtoReflect.Descriptor instead.
func (*GetFileResponse) Descriptor() ([]byte, []int) {
	return file_fileguard_proto_rawDescGZIP(), []int{13}
}

func (x *GetFileResponse) GetFileContent() []byte {
	if x != nil {
		return x.FileContent
	}
	return nil
}

type VerifyProofRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Proof      []*Hash `protobuf:"bytes,1,rep,name=proof,proto3" json:"proof,omitempty"`
	TargetHash *Hash   `protobuf:"bytes,2,opt,name=targetHash,proto3" json:"targetHash,omitempty"`
	RootHash   *Hash   `protobuf:"bytes,3,opt,name=rootHash,proto3" json:"rootHash,omitempty"`
}

func (x *VerifyProofRequest) Reset() {
	*x = VerifyProofRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fileguard_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyProofRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyProofRequest) ProtoMessage() {}

func (x *VerifyProofRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fileguard_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyProofRequest.ProtoReflect.Descriptor instead.
func (*VerifyProofRequest) Descriptor() ([]byte, []int) {
	return file_fileguard_proto_rawDescGZIP(), []int{14}
}

func (x *VerifyProofRequest) GetProof() []*Hash {
	if x != nil {
		return x.Proof
	}
	return nil
}

func (x *VerifyProofRequest) GetTargetHash() *Hash {
	if x != nil {
		return x.TargetHash
	}
	return nil
}

func (x *VerifyProofRequest) GetRootHash() *Hash {
	if x != nil {
		return x.RootHash
	}
	return nil
}

type VerifyProofResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valid bool `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
}

func (x *VerifyProofResponse) Reset() {
	*x = VerifyProofResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fileguard_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyProofResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyProofResponse) ProtoMessage() {}

func (x *VerifyProofResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fileguard_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyProofResponse.ProtoReflect.Descriptor instead.
func (*VerifyProofResponse) Descriptor() ([]byte, []int) {
	return file_fileguard_proto_rawDescGZIP(), []int{15}
}

func (x *VerifyProofResponse) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

var File_fileguard_proto protoreflect.FileDescriptor

var file_fileguard_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x66, 0x69, 0x6c, 0x65, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x1c, 0x0a, 0x04, 0x48, 0x61, 0x73, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x6b, 0x0a, 0x0a, 0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a,
	0x04, 0x6c, 0x65, 0x66, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4d, 0x65,
	0x72, 0x6b, 0x6c, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6c, 0x65, 0x66, 0x74, 0x12, 0x21,
	0x0a, 0x05, 0x72, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x72, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x19, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x05, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x4c, 0x0a, 0x0a,
	0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x54, 0x72, 0x65, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x72, 0x6f,
	0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4d, 0x65, 0x72, 0x6b, 0x6c,
	0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x12, 0x1d, 0x0a, 0x06, 0x6c,
	0x65, 0x61, 0x76, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x48, 0x61,
	0x73, 0x68, 0x52, 0x06, 0x6c, 0x65, 0x61, 0x76, 0x65, 0x73, 0x22, 0x68, 0x0a, 0x06, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x12, 0x1f, 0x0a, 0x04, 0x74, 0x72, 0x65, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x54, 0x72, 0x65, 0x65, 0x52, 0x04, 0x74, 0x72,
	0x65, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x50, 0x61, 0x74,
	0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x50, 0x61, 0x74, 0x68, 0x22, 0x34, 0x0a, 0x10, 0x4e, 0x65, 0x77, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x50, 0x61, 0x74, 0x68, 0x22, 0x34, 0x0a, 0x11, 0x4e, 0x65,
	0x77, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1f, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x07, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x22, 0x2a, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x11, 0x0a, 0x0f,
	0x41, 0x64, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1c, 0x0a, 0x1a, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4d, 0x65, 0x72, 0x6b,
	0x6c, 0x65, 0x54, 0x72, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x40, 0x0a,
	0x1b, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65,
	0x54, 0x72, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x08,
	0x72, 0x6f, 0x6f, 0x74, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05,
	0x2e, 0x48, 0x61, 0x73, 0x68, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x74, 0x48, 0x61, 0x73, 0x68, 0x22,
	0x27, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x2f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x05,
	0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x48, 0x61,
	0x73, 0x68, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x22, 0x26, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x22, 0x33, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x7b, 0x0a, 0x12, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x05,
	0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x48, 0x61,
	0x73, 0x68, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x25, 0x0a, 0x0a, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e,
	0x48, 0x61, 0x73, 0x68, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x21, 0x0a, 0x08, 0x72, 0x6f, 0x6f, 0x74, 0x48, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x05, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x74, 0x48,
	0x61, 0x73, 0x68, 0x22, 0x2b, 0x0a, 0x13, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x50, 0x72, 0x6f,
	0x6f, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x32, 0xe4, 0x02, 0x0a, 0x09, 0x46, 0x69, 0x6c, 0x65, 0x47, 0x75, 0x61, 0x72, 0x64, 0x12, 0x34,
	0x0a, 0x09, 0x4e, 0x65, 0x77, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x4e, 0x65,
	0x77, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x4e, 0x65, 0x77, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12,
	0x0f, 0x2e, 0x41, 0x64, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x10, 0x2e, 0x41, 0x64, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x54, 0x72, 0x65, 0x65, 0x12, 0x1b, 0x2e, 0x43, 0x6f,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x54, 0x72, 0x65,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x54, 0x72, 0x65, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x6f, 0x66, 0x12, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6f,
	0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0b, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x13, 0x2e, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fileguard_proto_rawDescOnce sync.Once
	file_fileguard_proto_rawDescData = file_fileguard_proto_rawDesc
)

func file_fileguard_proto_rawDescGZIP() []byte {
	file_fileguard_proto_rawDescOnce.Do(func() {
		file_fileguard_proto_rawDescData = protoimpl.X.CompressGZIP(file_fileguard_proto_rawDescData)
	})
	return file_fileguard_proto_rawDescData
}

var file_fileguard_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_fileguard_proto_goTypes = []interface{}{
	(*Hash)(nil),                        // 0: Hash
	(*MerkleNode)(nil),                  // 1: MerkleNode
	(*MerkleTree)(nil),                  // 2: MerkleTree
	(*Server)(nil),                      // 3: Server
	(*NewServerRequest)(nil),            // 4: NewServerRequest
	(*NewServerResponse)(nil),           // 5: NewServerResponse
	(*AddFileRequest)(nil),              // 6: AddFileRequest
	(*AddFileResponse)(nil),             // 7: AddFileResponse
	(*ConstructMerkleTreeRequest)(nil),  // 8: ConstructMerkleTreeRequest
	(*ConstructMerkleTreeResponse)(nil), // 9: ConstructMerkleTreeResponse
	(*GetProofRequest)(nil),             // 10: GetProofRequest
	(*GetProofResponse)(nil),            // 11: GetProofResponse
	(*GetFileRequest)(nil),              // 12: GetFileRequest
	(*GetFileResponse)(nil),             // 13: GetFileResponse
	(*VerifyProofRequest)(nil),          // 14: VerifyProofRequest
	(*VerifyProofResponse)(nil),         // 15: VerifyProofResponse
}
var file_fileguard_proto_depIdxs = []int32{
	1,  // 0: MerkleNode.left:type_name -> MerkleNode
	1,  // 1: MerkleNode.right:type_name -> MerkleNode
	0,  // 2: MerkleNode.hash:type_name -> Hash
	1,  // 3: MerkleTree.root:type_name -> MerkleNode
	0,  // 4: MerkleTree.leaves:type_name -> Hash
	0,  // 5: Server.files:type_name -> Hash
	2,  // 6: Server.tree:type_name -> MerkleTree
	3,  // 7: NewServerResponse.server:type_name -> Server
	0,  // 8: ConstructMerkleTreeResponse.rootHash:type_name -> Hash
	0,  // 9: GetProofResponse.proof:type_name -> Hash
	0,  // 10: VerifyProofRequest.proof:type_name -> Hash
	0,  // 11: VerifyProofRequest.targetHash:type_name -> Hash
	0,  // 12: VerifyProofRequest.rootHash:type_name -> Hash
	4,  // 13: FileGuard.NewServer:input_type -> NewServerRequest
	6,  // 14: FileGuard.AddFile:input_type -> AddFileRequest
	8,  // 15: FileGuard.ConstructMerkleTree:input_type -> ConstructMerkleTreeRequest
	10, // 16: FileGuard.GetProof:input_type -> GetProofRequest
	12, // 17: FileGuard.GetFile:input_type -> GetFileRequest
	14, // 18: FileGuard.VerifyProof:input_type -> VerifyProofRequest
	5,  // 19: FileGuard.NewServer:output_type -> NewServerResponse
	7,  // 20: FileGuard.AddFile:output_type -> AddFileResponse
	9,  // 21: FileGuard.ConstructMerkleTree:output_type -> ConstructMerkleTreeResponse
	11, // 22: FileGuard.GetProof:output_type -> GetProofResponse
	13, // 23: FileGuard.GetFile:output_type -> GetFileResponse
	15, // 24: FileGuard.VerifyProof:output_type -> VerifyProofResponse
	19, // [19:25] is the sub-list for method output_type
	13, // [13:19] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_fileguard_proto_init() }
func file_fileguard_proto_init() {
	if File_fileguard_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fileguard_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hash); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fileguard_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MerkleNode); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fileguard_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MerkleTree); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fileguard_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fileguard_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewServerRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fileguard_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewServerResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fileguard_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFileRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fileguard_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFileResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fileguard_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConstructMerkleTreeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fileguard_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConstructMerkleTreeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fileguard_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProofRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fileguard_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProofResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fileguard_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fileguard_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fileguard_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyProofRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fileguard_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyProofResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_fileguard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_fileguard_proto_goTypes,
		DependencyIndexes: file_fileguard_proto_depIdxs,
		MessageInfos:      file_fileguard_proto_msgTypes,
	}.Build()
	File_fileguard_proto = out.File
	file_fileguard_proto_rawDesc = nil
	file_fileguard_proto_goTypes = nil
	file_fileguard_proto_depIdxs = nil
}