// Code generated by protoc-gen-go. DO NOT EDIT.
// source: raft.proto

package raftpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type EntryType int32

const (
	EntryType_EntryNormal     EntryType = 0
	EntryType_EntryConfChange EntryType = 1
)

var EntryType_name = map[int32]string{
	0: "EntryNormal",
	1: "EntryConfChange",
}

var EntryType_value = map[string]int32{
	"EntryNormal":     0,
	"EntryConfChange": 1,
}

func (x EntryType) String() string {
	return proto.EnumName(EntryType_name, int32(x))
}

func (EntryType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b042552c306ae59b, []int{0}
}

type MessageType int32

const (
	MessageType_MsgHup            MessageType = 0
	MessageType_MsgBeat           MessageType = 1
	MessageType_MsgProp           MessageType = 2
	MessageType_MsgApp            MessageType = 3
	MessageType_MsgAppResp        MessageType = 4
	MessageType_MsgVote           MessageType = 5
	MessageType_MsgVoteResp       MessageType = 6
	MessageType_MsgSnap           MessageType = 7
	MessageType_MsgHeartbeat      MessageType = 8
	MessageType_MsgHeartbeatResp  MessageType = 9
	MessageType_MsgUnreachable    MessageType = 10
	MessageType_MsgSnapStatus     MessageType = 11
	MessageType_MsgCheckQuorum    MessageType = 12
	MessageType_MsgTransferLeader MessageType = 13
	MessageType_MsgTimeoutNow     MessageType = 14
	MessageType_MsgReadIndex      MessageType = 15
	MessageType_MsgReadIndexResp  MessageType = 16
	MessageType_MsgPreVote        MessageType = 17
	MessageType_MsgPreVoteResp    MessageType = 18
)

var MessageType_name = map[int32]string{
	0:  "MsgHup",
	1:  "MsgBeat",
	2:  "MsgProp",
	3:  "MsgApp",
	4:  "MsgAppResp",
	5:  "MsgVote",
	6:  "MsgVoteResp",
	7:  "MsgSnap",
	8:  "MsgHeartbeat",
	9:  "MsgHeartbeatResp",
	10: "MsgUnreachable",
	11: "MsgSnapStatus",
	12: "MsgCheckQuorum",
	13: "MsgTransferLeader",
	14: "MsgTimeoutNow",
	15: "MsgReadIndex",
	16: "MsgReadIndexResp",
	17: "MsgPreVote",
	18: "MsgPreVoteResp",
}

var MessageType_value = map[string]int32{
	"MsgHup":            0,
	"MsgBeat":           1,
	"MsgProp":           2,
	"MsgApp":            3,
	"MsgAppResp":        4,
	"MsgVote":           5,
	"MsgVoteResp":       6,
	"MsgSnap":           7,
	"MsgHeartbeat":      8,
	"MsgHeartbeatResp":  9,
	"MsgUnreachable":    10,
	"MsgSnapStatus":     11,
	"MsgCheckQuorum":    12,
	"MsgTransferLeader": 13,
	"MsgTimeoutNow":     14,
	"MsgReadIndex":      15,
	"MsgReadIndexResp":  16,
	"MsgPreVote":        17,
	"MsgPreVoteResp":    18,
}

func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}

func (MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b042552c306ae59b, []int{1}
}

type ConfChangeType int32

const (
	ConfChangeType_ConfChangeAddNode        ConfChangeType = 0
	ConfChangeType_ConfChangeRemoveNode     ConfChangeType = 1
	ConfChangeType_ConfChangeUpdateNode     ConfChangeType = 2
	ConfChangeType_ConfChangeAddLearnerNode ConfChangeType = 3
)

var ConfChangeType_name = map[int32]string{
	0: "ConfChangeAddNode",
	1: "ConfChangeRemoveNode",
	2: "ConfChangeUpdateNode",
	3: "ConfChangeAddLearnerNode",
}

var ConfChangeType_value = map[string]int32{
	"ConfChangeAddNode":        0,
	"ConfChangeRemoveNode":     1,
	"ConfChangeUpdateNode":     2,
	"ConfChangeAddLearnerNode": 3,
}

func (x ConfChangeType) String() string {
	return proto.EnumName(ConfChangeType_name, int32(x))
}

func (ConfChangeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b042552c306ae59b, []int{2}
}

type Entry struct {
	Term                 uint64    `protobuf:"varint,2,opt,name=Term,proto3" json:"Term,omitempty"`
	Index                uint64    `protobuf:"varint,3,opt,name=Index,proto3" json:"Index,omitempty"`
	Type                 EntryType `protobuf:"varint,1,opt,name=Type,proto3,enum=raftpb.EntryType" json:"Type,omitempty"`
	Data                 []byte    `protobuf:"bytes,4,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Entry) Reset()         { *m = Entry{} }
func (m *Entry) String() string { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()    {}
func (*Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_b042552c306ae59b, []int{0}
}

func (m *Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entry.Unmarshal(m, b)
}
func (m *Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entry.Marshal(b, m, deterministic)
}
func (m *Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entry.Merge(m, src)
}
func (m *Entry) XXX_Size() int {
	return xxx_messageInfo_Entry.Size(m)
}
func (m *Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_Entry proto.InternalMessageInfo

func (m *Entry) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *Entry) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Entry) GetType() EntryType {
	if m != nil {
		return m.Type
	}
	return EntryType_EntryNormal
}

func (m *Entry) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type SnapshotMetadata struct {
	ConfState            *ConfState `protobuf:"bytes,1,opt,name=conf_state,json=confState,proto3" json:"conf_state,omitempty"`
	Index                uint64     `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	Term                 uint64     `protobuf:"varint,3,opt,name=term,proto3" json:"term,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SnapshotMetadata) Reset()         { *m = SnapshotMetadata{} }
func (m *SnapshotMetadata) String() string { return proto.CompactTextString(m) }
func (*SnapshotMetadata) ProtoMessage()    {}
func (*SnapshotMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_b042552c306ae59b, []int{1}
}

func (m *SnapshotMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnapshotMetadata.Unmarshal(m, b)
}
func (m *SnapshotMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnapshotMetadata.Marshal(b, m, deterministic)
}
func (m *SnapshotMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapshotMetadata.Merge(m, src)
}
func (m *SnapshotMetadata) XXX_Size() int {
	return xxx_messageInfo_SnapshotMetadata.Size(m)
}
func (m *SnapshotMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapshotMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_SnapshotMetadata proto.InternalMessageInfo

func (m *SnapshotMetadata) GetConfState() *ConfState {
	if m != nil {
		return m.ConfState
	}
	return nil
}

func (m *SnapshotMetadata) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *SnapshotMetadata) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

type Snapshot struct {
	Data                 []byte            `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Metadata             *SnapshotMetadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Snapshot) Reset()         { *m = Snapshot{} }
func (m *Snapshot) String() string { return proto.CompactTextString(m) }
func (*Snapshot) ProtoMessage()    {}
func (*Snapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_b042552c306ae59b, []int{2}
}

func (m *Snapshot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Snapshot.Unmarshal(m, b)
}
func (m *Snapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Snapshot.Marshal(b, m, deterministic)
}
func (m *Snapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Snapshot.Merge(m, src)
}
func (m *Snapshot) XXX_Size() int {
	return xxx_messageInfo_Snapshot.Size(m)
}
func (m *Snapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_Snapshot.DiscardUnknown(m)
}

var xxx_messageInfo_Snapshot proto.InternalMessageInfo

func (m *Snapshot) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Snapshot) GetMetadata() *SnapshotMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Message struct {
	Type                 MessageType `protobuf:"varint,1,opt,name=type,proto3,enum=raftpb.MessageType" json:"type,omitempty"`
	To                   uint64      `protobuf:"varint,2,opt,name=to,proto3" json:"to,omitempty"`
	Term                 uint64      `protobuf:"varint,4,opt,name=term,proto3" json:"term,omitempty"`
	LogTerm              uint64      `protobuf:"varint,5,opt,name=logTerm,proto3" json:"logTerm,omitempty"`
	Index                uint64      `protobuf:"varint,6,opt,name=index,proto3" json:"index,omitempty"`
	Entries              []*Entry    `protobuf:"bytes,7,rep,name=entries,proto3" json:"entries,omitempty"`
	Commit               uint64      `protobuf:"varint,8,opt,name=commit,proto3" json:"commit,omitempty"`
	Snapshot             *Snapshot   `protobuf:"bytes,9,opt,name=snapshot,proto3" json:"snapshot,omitempty"`
	Reject               bool        `protobuf:"varint,10,opt,name=reject,proto3" json:"reject,omitempty"`
	RejectHint           uint64      `protobuf:"varint,11,opt,name=rejectHint,proto3" json:"rejectHint,omitempty"`
	Context              []byte      `protobuf:"bytes,12,opt,name=context,proto3" json:"context,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_b042552c306ae59b, []int{3}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageType_MsgHup
}

func (m *Message) GetTo() uint64 {
	if m != nil {
		return m.To
	}
	return 0
}

func (m *Message) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *Message) GetLogTerm() uint64 {
	if m != nil {
		return m.LogTerm
	}
	return 0
}

func (m *Message) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Message) GetEntries() []*Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func (m *Message) GetCommit() uint64 {
	if m != nil {
		return m.Commit
	}
	return 0
}

func (m *Message) GetSnapshot() *Snapshot {
	if m != nil {
		return m.Snapshot
	}
	return nil
}

func (m *Message) GetReject() bool {
	if m != nil {
		return m.Reject
	}
	return false
}

func (m *Message) GetRejectHint() uint64 {
	if m != nil {
		return m.RejectHint
	}
	return 0
}

func (m *Message) GetContext() []byte {
	if m != nil {
		return m.Context
	}
	return nil
}

type HardState struct {
	Term                 uint64   `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	Vote                 uint64   `protobuf:"varint,2,opt,name=vote,proto3" json:"vote,omitempty"`
	Commit               uint64   `protobuf:"varint,3,opt,name=commit,proto3" json:"commit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HardState) Reset()         { *m = HardState{} }
func (m *HardState) String() string { return proto.CompactTextString(m) }
func (*HardState) ProtoMessage()    {}
func (*HardState) Descriptor() ([]byte, []int) {
	return fileDescriptor_b042552c306ae59b, []int{4}
}

func (m *HardState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HardState.Unmarshal(m, b)
}
func (m *HardState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HardState.Marshal(b, m, deterministic)
}
func (m *HardState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HardState.Merge(m, src)
}
func (m *HardState) XXX_Size() int {
	return xxx_messageInfo_HardState.Size(m)
}
func (m *HardState) XXX_DiscardUnknown() {
	xxx_messageInfo_HardState.DiscardUnknown(m)
}

var xxx_messageInfo_HardState proto.InternalMessageInfo

func (m *HardState) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *HardState) GetVote() uint64 {
	if m != nil {
		return m.Vote
	}
	return 0
}

func (m *HardState) GetCommit() uint64 {
	if m != nil {
		return m.Commit
	}
	return 0
}

type ConfState struct {
	Nodes                []uint64 `protobuf:"varint,1,rep,packed,name=nodes,proto3" json:"nodes,omitempty"`
	Learners             []uint64 `protobuf:"varint,2,rep,packed,name=learners,proto3" json:"learners,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfState) Reset()         { *m = ConfState{} }
func (m *ConfState) String() string { return proto.CompactTextString(m) }
func (*ConfState) ProtoMessage()    {}
func (*ConfState) Descriptor() ([]byte, []int) {
	return fileDescriptor_b042552c306ae59b, []int{5}
}

func (m *ConfState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfState.Unmarshal(m, b)
}
func (m *ConfState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfState.Marshal(b, m, deterministic)
}
func (m *ConfState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfState.Merge(m, src)
}
func (m *ConfState) XXX_Size() int {
	return xxx_messageInfo_ConfState.Size(m)
}
func (m *ConfState) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfState.DiscardUnknown(m)
}

var xxx_messageInfo_ConfState proto.InternalMessageInfo

func (m *ConfState) GetNodes() []uint64 {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *ConfState) GetLearners() []uint64 {
	if m != nil {
		return m.Learners
	}
	return nil
}

type ConfChange struct {
	ID                   uint64         `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Type                 ConfChangeType `protobuf:"varint,2,opt,name=Type,proto3,enum=raftpb.ConfChangeType" json:"Type,omitempty"`
	NodeID               uint64         `protobuf:"varint,3,opt,name=NodeID,proto3" json:"NodeID,omitempty"`
	Context              []byte         `protobuf:"bytes,4,opt,name=Context,proto3" json:"Context,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ConfChange) Reset()         { *m = ConfChange{} }
func (m *ConfChange) String() string { return proto.CompactTextString(m) }
func (*ConfChange) ProtoMessage()    {}
func (*ConfChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_b042552c306ae59b, []int{6}
}

func (m *ConfChange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfChange.Unmarshal(m, b)
}
func (m *ConfChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfChange.Marshal(b, m, deterministic)
}
func (m *ConfChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfChange.Merge(m, src)
}
func (m *ConfChange) XXX_Size() int {
	return xxx_messageInfo_ConfChange.Size(m)
}
func (m *ConfChange) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfChange.DiscardUnknown(m)
}

var xxx_messageInfo_ConfChange proto.InternalMessageInfo

func (m *ConfChange) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *ConfChange) GetType() ConfChangeType {
	if m != nil {
		return m.Type
	}
	return ConfChangeType_ConfChangeAddNode
}

func (m *ConfChange) GetNodeID() uint64 {
	if m != nil {
		return m.NodeID
	}
	return 0
}

func (m *ConfChange) GetContext() []byte {
	if m != nil {
		return m.Context
	}
	return nil
}

func init() {
	proto.RegisterEnum("raftpb.EntryType", EntryType_name, EntryType_value)
	proto.RegisterEnum("raftpb.MessageType", MessageType_name, MessageType_value)
	proto.RegisterEnum("raftpb.ConfChangeType", ConfChangeType_name, ConfChangeType_value)
	proto.RegisterType((*Entry)(nil), "raftpb.Entry")
	proto.RegisterType((*SnapshotMetadata)(nil), "raftpb.SnapshotMetadata")
	proto.RegisterType((*Snapshot)(nil), "raftpb.Snapshot")
	proto.RegisterType((*Message)(nil), "raftpb.Message")
	proto.RegisterType((*HardState)(nil), "raftpb.HardState")
	proto.RegisterType((*ConfState)(nil), "raftpb.ConfState")
	proto.RegisterType((*ConfChange)(nil), "raftpb.ConfChange")
}

func init() { proto.RegisterFile("raft.proto", fileDescriptor_b042552c306ae59b) }

var fileDescriptor_b042552c306ae59b = []byte{
	// 728 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x54, 0xdb, 0x6e, 0xc3, 0x44,
	0x10, 0xad, 0x1d, 0xe7, 0x36, 0x4e, 0xd3, 0xcd, 0xb6, 0x54, 0x2b, 0x84, 0x50, 0x14, 0x09, 0x35,
	0x8a, 0x50, 0x05, 0x85, 0x57, 0x1e, 0x4a, 0x82, 0xd4, 0x88, 0xa6, 0x2a, 0x6e, 0xca, 0x2b, 0xda,
	0xc4, 0x93, 0x0b, 0xc4, 0x5e, 0x6b, 0xbd, 0x29, 0x2d, 0x3f, 0xc8, 0x07, 0xf0, 0x43, 0x68, 0xc7,
	0xeb, 0xc4, 0xe9, 0xdb, 0x9c, 0x39, 0xb3, 0x33, 0x67, 0x2e, 0x36, 0x80, 0x96, 0x2b, 0x73, 0x9b,
	0x69, 0x65, 0x14, 0x6f, 0x58, 0x3b, 0x5b, 0x0c, 0x76, 0x50, 0xff, 0x25, 0x35, 0xfa, 0x83, 0x73,
	0x08, 0xe6, 0xa8, 0x13, 0xe1, 0xf7, 0xbd, 0x61, 0x10, 0x91, 0xcd, 0xaf, 0xa0, 0x3e, 0x4d, 0x63,
	0x7c, 0x17, 0x35, 0x72, 0x16, 0x80, 0x7f, 0x03, 0xc1, 0xfc, 0x23, 0x43, 0xe1, 0xf5, 0xbd, 0x61,
	0xf7, 0xae, 0x77, 0x5b, 0x64, 0xba, 0xa5, 0x34, 0x96, 0x88, 0x88, 0xb6, 0x09, 0x27, 0xd2, 0x48,
	0x11, 0xf4, 0xbd, 0x61, 0x27, 0x22, 0x7b, 0x90, 0x02, 0x7b, 0x49, 0x65, 0x96, 0x6f, 0x94, 0x99,
	0xa1, 0x91, 0xb1, 0x34, 0x92, 0x7f, 0x07, 0xb0, 0x54, 0xe9, 0xea, 0x8f, 0xdc, 0x48, 0x53, 0x24,
	0x0d, 0x8f, 0x49, 0xc7, 0x2a, 0x5d, 0xbd, 0x58, 0x22, 0x6a, 0x2f, 0x4b, 0xd3, 0xca, 0xda, 0x92,
	0xac, 0x42, 0x6b, 0x01, 0x6c, 0x3d, 0x63, 0x1b, 0x28, 0xb4, 0x92, 0x3d, 0x98, 0x43, 0xab, 0xac,
	0x67, 0x79, 0x5b, 0x8f, 0x2a, 0x74, 0x22, 0xb2, 0xf9, 0x8f, 0xd0, 0x4a, 0x9c, 0x0e, 0x4a, 0x16,
	0xde, 0x89, 0xb2, 0xf2, 0x67, 0x9d, 0xd1, 0x21, 0x72, 0xf0, 0xaf, 0x0f, 0xcd, 0x19, 0xe6, 0xb9,
	0x5c, 0x23, 0xbf, 0x81, 0xc0, 0x1c, 0x87, 0x71, 0x59, 0xbe, 0x76, 0x74, 0x31, 0x0e, 0x1b, 0xc0,
	0xbb, 0xe0, 0x1b, 0xe5, 0x14, 0xfb, 0x46, 0x1d, 0xe4, 0x06, 0x47, 0xb9, 0x5c, 0x40, 0x73, 0xa7,
	0xd6, 0xb4, 0x86, 0x3a, 0xb9, 0x4b, 0x78, 0x6c, 0xb9, 0x51, 0x6d, 0xf9, 0x06, 0x9a, 0x98, 0x1a,
	0xbd, 0xc5, 0x5c, 0x34, 0xfb, 0xb5, 0x61, 0x78, 0x77, 0x7e, 0xb2, 0x8c, 0xa8, 0x64, 0xf9, 0x35,
	0x34, 0x96, 0x2a, 0x49, 0xb6, 0x46, 0xb4, 0xe8, 0xbd, 0x43, 0xfc, 0x5b, 0x68, 0xe5, 0xae, 0x4f,
	0xd1, 0xa6, 0xfe, 0xd9, 0xe7, 0xfe, 0xa3, 0x43, 0x84, 0xcd, 0xa2, 0xf1, 0x4f, 0x5c, 0x1a, 0x01,
	0x7d, 0x6f, 0xd8, 0x8a, 0x1c, 0xe2, 0x5f, 0x03, 0x14, 0xd6, 0xc3, 0x36, 0x35, 0x22, 0xa4, 0x0a,
	0x15, 0x8f, 0x6d, 0x6b, 0xa9, 0x52, 0x83, 0xef, 0x46, 0x74, 0x68, 0xf8, 0x25, 0x1c, 0xfc, 0x0a,
	0xed, 0x07, 0xa9, 0xe3, 0x62, 0xad, 0xe5, 0x44, 0xbc, 0xca, 0x44, 0x38, 0x04, 0x6f, 0xca, 0x60,
	0x79, 0x95, 0xd6, 0xae, 0x34, 0x53, 0xab, 0x36, 0x33, 0xf8, 0x09, 0xda, 0xe3, 0xea, 0x8d, 0xa4,
	0x2a, 0xc6, 0x5c, 0x78, 0xfd, 0x9a, 0x1d, 0x18, 0x01, 0xfe, 0x25, 0xb4, 0x76, 0x28, 0x75, 0x8a,
	0x3a, 0x17, 0x3e, 0x11, 0x07, 0x3c, 0xf8, 0x07, 0xc0, 0x3e, 0x1f, 0x6f, 0x64, 0xba, 0xa6, 0x75,
	0x4d, 0x27, 0x4e, 0x8a, 0x3f, 0x9d, 0xf0, 0x91, 0x3b, 0x7a, 0x9f, 0xf6, 0x7c, 0x5d, 0xbd, 0xcf,
	0xe2, 0x45, 0xe5, 0xf2, 0xaf, 0xa1, 0xf1, 0xa4, 0x62, 0x9c, 0x4e, 0x4a, 0x81, 0x05, 0xb2, 0x73,
	0x18, 0xbb, 0x39, 0x14, 0x1f, 0x45, 0x09, 0x47, 0xdf, 0x43, 0xfb, 0xf0, 0xf9, 0xf0, 0x0b, 0x08,
	0x09, 0x3c, 0x29, 0x9d, 0xc8, 0x1d, 0x3b, 0xe3, 0x97, 0x70, 0x41, 0x8e, 0x63, 0x31, 0xe6, 0x8d,
	0xfe, 0xf3, 0x21, 0xac, 0x5c, 0x19, 0x07, 0x68, 0xcc, 0xf2, 0xf5, 0xc3, 0x3e, 0x63, 0x67, 0x3c,
	0x84, 0xe6, 0x2c, 0x5f, 0xff, 0x8c, 0xd2, 0x30, 0xcf, 0x81, 0x67, 0xad, 0x32, 0xe6, 0xbb, 0xa8,
	0xfb, 0x2c, 0x63, 0x35, 0xde, 0x05, 0x28, 0xec, 0x08, 0xf3, 0x8c, 0x05, 0x2e, 0xf0, 0x77, 0x65,
	0x90, 0xd5, 0xad, 0x08, 0x07, 0x88, 0x6d, 0x38, 0xd6, 0x5e, 0x05, 0x6b, 0x72, 0x06, 0x1d, 0x5b,
	0x0c, 0xa5, 0x36, 0x0b, 0x5b, 0xa5, 0xc5, 0xaf, 0x80, 0x55, 0x3d, 0xf4, 0xa8, 0xcd, 0x39, 0x74,
	0x67, 0xf9, 0xfa, 0x35, 0xd5, 0x28, 0x97, 0x1b, 0xb9, 0xd8, 0x21, 0x03, 0xde, 0x83, 0x73, 0x97,
	0xc8, 0x6e, 0x6a, 0x9f, 0xb3, 0xd0, 0x85, 0x8d, 0x37, 0xb8, 0xfc, 0xeb, 0xb7, 0xbd, 0xd2, 0xfb,
	0x84, 0x75, 0xf8, 0x17, 0xd0, 0x9b, 0xe5, 0xeb, 0xb9, 0x96, 0x69, 0xbe, 0x42, 0xfd, 0x88, 0x32,
	0x46, 0xcd, 0xce, 0xdd, 0xeb, 0xf9, 0x36, 0x41, 0xb5, 0x37, 0x4f, 0xea, 0x6f, 0xd6, 0x75, 0x62,
	0x22, 0x94, 0x31, 0xfd, 0x9f, 0xd8, 0x85, 0x13, 0x73, 0xf0, 0x90, 0x18, 0xe6, 0xfa, 0x7d, 0xd6,
	0x48, 0x2d, 0xf6, 0x5c, 0x55, 0x87, 0x29, 0x86, 0x8f, 0x3e, 0xa0, 0x7b, 0xba, 0x52, 0xab, 0xe3,
	0xe8, 0xb9, 0x8f, 0x63, 0xbb, 0x4b, 0x76, 0xc6, 0x05, 0x5c, 0x1d, 0xdd, 0x11, 0x26, 0xea, 0x0d,
	0x89, 0xf1, 0x4e, 0x99, 0xd7, 0x2c, 0x96, 0xa6, 0x60, 0x7c, 0xfe, 0x15, 0x88, 0x93, 0x54, 0x8f,
	0xc5, 0xe9, 0x11, 0x5b, 0x5b, 0x34, 0xe8, 0xc7, 0xfc, 0xc3, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x0e, 0x3b, 0xe6, 0x71, 0xa6, 0x05, 0x00, 0x00,
}
