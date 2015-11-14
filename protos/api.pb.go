// Code generated by protoc-gen-go.
// source: api.proto
// DO NOT EDIT!

/*
Package protos is a generated protocol buffer package.

It is generated from these files:
	api.proto
	chaincode.proto
	devops.proto
	openchain.proto
	server_admin.proto

It has these top-level messages:
	BlockchainInfo
	BlockNumber
	BlockCount
	ChainletID
	ChainletMessage
	ChaincodeInvocation
	ChainletSpec
	ChainletDeploymentSpec
	ChainletIdentifier
	ChainletRequestContext
	ChainletExecutionContext
	ChaincodeMessage
	DevopsResponse
	BuildResult
	Transaction
	TransactionBlock
	Block
	OpenchainMessage
	ServerStatus
*/
package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "google/protobuf"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Contains information about the blockchain ledger such as height, current
// block hash, and previous block hash.
type BlockchainInfo struct {
	Height            uint64 `protobuf:"varint,1,opt,name=height" json:"height,omitempty"`
	CurrentBlockHash  []byte `protobuf:"bytes,2,opt,name=currentBlockHash,proto3" json:"currentBlockHash,omitempty"`
	PreviousBlockHash []byte `protobuf:"bytes,3,opt,name=previousBlockHash,proto3" json:"previousBlockHash,omitempty"`
}

func (m *BlockchainInfo) Reset()         { *m = BlockchainInfo{} }
func (m *BlockchainInfo) String() string { return proto.CompactTextString(m) }
func (*BlockchainInfo) ProtoMessage()    {}

// Specifies the block number to be returned from the blockchain.
type BlockNumber struct {
	Number uint64 `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
}

func (m *BlockNumber) Reset()         { *m = BlockNumber{} }
func (m *BlockNumber) String() string { return proto.CompactTextString(m) }
func (*BlockNumber) ProtoMessage()    {}

// Specifies the current number of blocks in the blockchain.
type BlockCount struct {
	Count uint64 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
}

func (m *BlockCount) Reset()         { *m = BlockCount{} }
func (m *BlockCount) String() string { return proto.CompactTextString(m) }
func (*BlockCount) ProtoMessage()    {}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Openchain service

type OpenchainClient interface {
	// GetBlockchainInfo returns information about the blockchain ledger such as
	// height, current block hash, and previous block hash.
	GetBlockchainInfo(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*BlockchainInfo, error)
	// GetBlockByNumber returns the data contained within a specific block in the
	// blockchain. The genesis block is block zero.
	GetBlockByNumber(ctx context.Context, in *BlockNumber, opts ...grpc.CallOption) (*Block, error)
	// GetBlockCount returns the current number of blocks in the blockchain data
	// structure.
	GetBlockCount(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*BlockCount, error)
}

type openchainClient struct {
	cc *grpc.ClientConn
}

func NewOpenchainClient(cc *grpc.ClientConn) OpenchainClient {
	return &openchainClient{cc}
}

func (c *openchainClient) GetBlockchainInfo(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*BlockchainInfo, error) {
	out := new(BlockchainInfo)
	err := grpc.Invoke(ctx, "/protos.Openchain/GetBlockchainInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openchainClient) GetBlockByNumber(ctx context.Context, in *BlockNumber, opts ...grpc.CallOption) (*Block, error) {
	out := new(Block)
	err := grpc.Invoke(ctx, "/protos.Openchain/GetBlockByNumber", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openchainClient) GetBlockCount(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*BlockCount, error) {
	out := new(BlockCount)
	err := grpc.Invoke(ctx, "/protos.Openchain/GetBlockCount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Openchain service

type OpenchainServer interface {
	// GetBlockchainInfo returns information about the blockchain ledger such as
	// height, current block hash, and previous block hash.
	GetBlockchainInfo(context.Context, *google_protobuf1.Empty) (*BlockchainInfo, error)
	// GetBlockByNumber returns the data contained within a specific block in the
	// blockchain. The genesis block is block zero.
	GetBlockByNumber(context.Context, *BlockNumber) (*Block, error)
	// GetBlockCount returns the current number of blocks in the blockchain data
	// structure.
	GetBlockCount(context.Context, *google_protobuf1.Empty) (*BlockCount, error)
}

func RegisterOpenchainServer(s *grpc.Server, srv OpenchainServer) {
	s.RegisterService(&_Openchain_serviceDesc, srv)
}

func _Openchain_GetBlockchainInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(OpenchainServer).GetBlockchainInfo(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Openchain_GetBlockByNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(BlockNumber)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(OpenchainServer).GetBlockByNumber(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Openchain_GetBlockCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(OpenchainServer).GetBlockCount(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Openchain_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Openchain",
	HandlerType: (*OpenchainServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBlockchainInfo",
			Handler:    _Openchain_GetBlockchainInfo_Handler,
		},
		{
			MethodName: "GetBlockByNumber",
			Handler:    _Openchain_GetBlockByNumber_Handler,
		},
		{
			MethodName: "GetBlockCount",
			Handler:    _Openchain_GetBlockCount_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
