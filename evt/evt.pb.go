// Code generated by protoc-gen-go. DO NOT EDIT.
// source: evt.proto

package evt

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

type Evt struct {
	// Types that are valid to be assigned to Event:
	//	*Evt_PrepareToStart
	//	*Evt_GameStarted
	//	*Evt_PlayerDisconnected
	//	*Evt_PlayerReconnected
	//	*Evt_DoMoveReq
	//	*Evt_DoMoveRes
	//	*Evt_PlayerMove
	//	*Evt_EndGame
	//	*Evt_PlayerLeft
	//	*Evt_ErrorMessage
	Event                isEvt_Event `protobuf_oneof:"event"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Evt) Reset()         { *m = Evt{} }
func (m *Evt) String() string { return proto.CompactTextString(m) }
func (*Evt) ProtoMessage()    {}
func (*Evt) Descriptor() ([]byte, []int) {
	return fileDescriptor_f42bf8acbb5981a6, []int{0}
}

func (m *Evt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Evt.Unmarshal(m, b)
}
func (m *Evt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Evt.Marshal(b, m, deterministic)
}
func (m *Evt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Evt.Merge(m, src)
}
func (m *Evt) XXX_Size() int {
	return xxx_messageInfo_Evt.Size(m)
}
func (m *Evt) XXX_DiscardUnknown() {
	xxx_messageInfo_Evt.DiscardUnknown(m)
}

var xxx_messageInfo_Evt proto.InternalMessageInfo

type isEvt_Event interface {
	isEvt_Event()
}

type Evt_PrepareToStart struct {
	PrepareToStart *PrepareToStart `protobuf:"bytes,1,opt,name=prepare_to_start,json=prepareToStart,proto3,oneof"`
}

type Evt_GameStarted struct {
	GameStarted *GameStarted `protobuf:"bytes,2,opt,name=game_started,json=gameStarted,proto3,oneof"`
}

type Evt_PlayerDisconnected struct {
	PlayerDisconnected *PlayerDisconnected `protobuf:"bytes,3,opt,name=player_disconnected,json=playerDisconnected,proto3,oneof"`
}

type Evt_PlayerReconnected struct {
	PlayerReconnected *PlayerReconnected `protobuf:"bytes,4,opt,name=player_reconnected,json=playerReconnected,proto3,oneof"`
}

type Evt_DoMoveReq struct {
	DoMoveReq *DoMoveReq `protobuf:"bytes,5,opt,name=do_move_req,json=doMoveReq,proto3,oneof"`
}

type Evt_DoMoveRes struct {
	DoMoveRes *DoMoveRes `protobuf:"bytes,6,opt,name=do_move_res,json=doMoveRes,proto3,oneof"`
}

type Evt_PlayerMove struct {
	PlayerMove *PlayerMove `protobuf:"bytes,7,opt,name=player_move,json=playerMove,proto3,oneof"`
}

type Evt_EndGame struct {
	EndGame *EndGame `protobuf:"bytes,8,opt,name=end_game,json=endGame,proto3,oneof"`
}

type Evt_PlayerLeft struct {
	PlayerLeft *PlayerLeft `protobuf:"bytes,9,opt,name=player_left,json=playerLeft,proto3,oneof"`
}

type Evt_ErrorMessage struct {
	ErrorMessage *ErrorMessage `protobuf:"bytes,10,opt,name=error_message,json=errorMessage,proto3,oneof"`
}

func (*Evt_PrepareToStart) isEvt_Event() {}

func (*Evt_GameStarted) isEvt_Event() {}

func (*Evt_PlayerDisconnected) isEvt_Event() {}

func (*Evt_PlayerReconnected) isEvt_Event() {}

func (*Evt_DoMoveReq) isEvt_Event() {}

func (*Evt_DoMoveRes) isEvt_Event() {}

func (*Evt_PlayerMove) isEvt_Event() {}

func (*Evt_EndGame) isEvt_Event() {}

func (*Evt_PlayerLeft) isEvt_Event() {}

func (*Evt_ErrorMessage) isEvt_Event() {}

func (m *Evt) GetEvent() isEvt_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *Evt) GetPrepareToStart() *PrepareToStart {
	if x, ok := m.GetEvent().(*Evt_PrepareToStart); ok {
		return x.PrepareToStart
	}
	return nil
}

func (m *Evt) GetGameStarted() *GameStarted {
	if x, ok := m.GetEvent().(*Evt_GameStarted); ok {
		return x.GameStarted
	}
	return nil
}

func (m *Evt) GetPlayerDisconnected() *PlayerDisconnected {
	if x, ok := m.GetEvent().(*Evt_PlayerDisconnected); ok {
		return x.PlayerDisconnected
	}
	return nil
}

func (m *Evt) GetPlayerReconnected() *PlayerReconnected {
	if x, ok := m.GetEvent().(*Evt_PlayerReconnected); ok {
		return x.PlayerReconnected
	}
	return nil
}

func (m *Evt) GetDoMoveReq() *DoMoveReq {
	if x, ok := m.GetEvent().(*Evt_DoMoveReq); ok {
		return x.DoMoveReq
	}
	return nil
}

func (m *Evt) GetDoMoveRes() *DoMoveRes {
	if x, ok := m.GetEvent().(*Evt_DoMoveRes); ok {
		return x.DoMoveRes
	}
	return nil
}

func (m *Evt) GetPlayerMove() *PlayerMove {
	if x, ok := m.GetEvent().(*Evt_PlayerMove); ok {
		return x.PlayerMove
	}
	return nil
}

func (m *Evt) GetEndGame() *EndGame {
	if x, ok := m.GetEvent().(*Evt_EndGame); ok {
		return x.EndGame
	}
	return nil
}

func (m *Evt) GetPlayerLeft() *PlayerLeft {
	if x, ok := m.GetEvent().(*Evt_PlayerLeft); ok {
		return x.PlayerLeft
	}
	return nil
}

func (m *Evt) GetErrorMessage() *ErrorMessage {
	if x, ok := m.GetEvent().(*Evt_ErrorMessage); ok {
		return x.ErrorMessage
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Evt) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Evt_PrepareToStart)(nil),
		(*Evt_GameStarted)(nil),
		(*Evt_PlayerDisconnected)(nil),
		(*Evt_PlayerReconnected)(nil),
		(*Evt_DoMoveReq)(nil),
		(*Evt_DoMoveRes)(nil),
		(*Evt_PlayerMove)(nil),
		(*Evt_EndGame)(nil),
		(*Evt_PlayerLeft)(nil),
		(*Evt_ErrorMessage)(nil),
	}
}

type PrepareToStart struct {
	Time                 int64    `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrepareToStart) Reset()         { *m = PrepareToStart{} }
func (m *PrepareToStart) String() string { return proto.CompactTextString(m) }
func (*PrepareToStart) ProtoMessage()    {}
func (*PrepareToStart) Descriptor() ([]byte, []int) {
	return fileDescriptor_f42bf8acbb5981a6, []int{1}
}

func (m *PrepareToStart) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrepareToStart.Unmarshal(m, b)
}
func (m *PrepareToStart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrepareToStart.Marshal(b, m, deterministic)
}
func (m *PrepareToStart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrepareToStart.Merge(m, src)
}
func (m *PrepareToStart) XXX_Size() int {
	return xxx_messageInfo_PrepareToStart.Size(m)
}
func (m *PrepareToStart) XXX_DiscardUnknown() {
	xxx_messageInfo_PrepareToStart.DiscardUnknown(m)
}

var xxx_messageInfo_PrepareToStart proto.InternalMessageInfo

func (m *PrepareToStart) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

type GameStarted struct {
	First                int64    `protobuf:"varint,1,opt,name=first,proto3" json:"first,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameStarted) Reset()         { *m = GameStarted{} }
func (m *GameStarted) String() string { return proto.CompactTextString(m) }
func (*GameStarted) ProtoMessage()    {}
func (*GameStarted) Descriptor() ([]byte, []int) {
	return fileDescriptor_f42bf8acbb5981a6, []int{2}
}

func (m *GameStarted) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameStarted.Unmarshal(m, b)
}
func (m *GameStarted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameStarted.Marshal(b, m, deterministic)
}
func (m *GameStarted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameStarted.Merge(m, src)
}
func (m *GameStarted) XXX_Size() int {
	return xxx_messageInfo_GameStarted.Size(m)
}
func (m *GameStarted) XXX_DiscardUnknown() {
	xxx_messageInfo_GameStarted.DiscardUnknown(m)
}

var xxx_messageInfo_GameStarted proto.InternalMessageInfo

func (m *GameStarted) GetFirst() int64 {
	if m != nil {
		return m.First
	}
	return 0
}

type PlayerDisconnected struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlayerDisconnected) Reset()         { *m = PlayerDisconnected{} }
func (m *PlayerDisconnected) String() string { return proto.CompactTextString(m) }
func (*PlayerDisconnected) ProtoMessage()    {}
func (*PlayerDisconnected) Descriptor() ([]byte, []int) {
	return fileDescriptor_f42bf8acbb5981a6, []int{3}
}

func (m *PlayerDisconnected) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerDisconnected.Unmarshal(m, b)
}
func (m *PlayerDisconnected) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerDisconnected.Marshal(b, m, deterministic)
}
func (m *PlayerDisconnected) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerDisconnected.Merge(m, src)
}
func (m *PlayerDisconnected) XXX_Size() int {
	return xxx_messageInfo_PlayerDisconnected.Size(m)
}
func (m *PlayerDisconnected) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerDisconnected.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerDisconnected proto.InternalMessageInfo

func (m *PlayerDisconnected) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type PlayerReconnected struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlayerReconnected) Reset()         { *m = PlayerReconnected{} }
func (m *PlayerReconnected) String() string { return proto.CompactTextString(m) }
func (*PlayerReconnected) ProtoMessage()    {}
func (*PlayerReconnected) Descriptor() ([]byte, []int) {
	return fileDescriptor_f42bf8acbb5981a6, []int{4}
}

func (m *PlayerReconnected) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerReconnected.Unmarshal(m, b)
}
func (m *PlayerReconnected) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerReconnected.Marshal(b, m, deterministic)
}
func (m *PlayerReconnected) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerReconnected.Merge(m, src)
}
func (m *PlayerReconnected) XXX_Size() int {
	return xxx_messageInfo_PlayerReconnected.Size(m)
}
func (m *PlayerReconnected) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerReconnected.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerReconnected proto.InternalMessageInfo

func (m *PlayerReconnected) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DoMoveReq struct {
	Move                 int64    `protobuf:"varint,1,opt,name=move,proto3" json:"move,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoMoveReq) Reset()         { *m = DoMoveReq{} }
func (m *DoMoveReq) String() string { return proto.CompactTextString(m) }
func (*DoMoveReq) ProtoMessage()    {}
func (*DoMoveReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f42bf8acbb5981a6, []int{5}
}

func (m *DoMoveReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoMoveReq.Unmarshal(m, b)
}
func (m *DoMoveReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoMoveReq.Marshal(b, m, deterministic)
}
func (m *DoMoveReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoMoveReq.Merge(m, src)
}
func (m *DoMoveReq) XXX_Size() int {
	return xxx_messageInfo_DoMoveReq.Size(m)
}
func (m *DoMoveReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DoMoveReq.DiscardUnknown(m)
}

var xxx_messageInfo_DoMoveReq proto.InternalMessageInfo

func (m *DoMoveReq) GetMove() int64 {
	if m != nil {
		return m.Move
	}
	return 0
}

type DoMoveRes struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Move                 int64    `protobuf:"varint,2,opt,name=move,proto3" json:"move,omitempty"`
	Accept               bool     `protobuf:"varint,3,opt,name=accept,proto3" json:"accept,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoMoveRes) Reset()         { *m = DoMoveRes{} }
func (m *DoMoveRes) String() string { return proto.CompactTextString(m) }
func (*DoMoveRes) ProtoMessage()    {}
func (*DoMoveRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_f42bf8acbb5981a6, []int{6}
}

func (m *DoMoveRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoMoveRes.Unmarshal(m, b)
}
func (m *DoMoveRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoMoveRes.Marshal(b, m, deterministic)
}
func (m *DoMoveRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoMoveRes.Merge(m, src)
}
func (m *DoMoveRes) XXX_Size() int {
	return xxx_messageInfo_DoMoveRes.Size(m)
}
func (m *DoMoveRes) XXX_DiscardUnknown() {
	xxx_messageInfo_DoMoveRes.DiscardUnknown(m)
}

var xxx_messageInfo_DoMoveRes proto.InternalMessageInfo

func (m *DoMoveRes) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DoMoveRes) GetMove() int64 {
	if m != nil {
		return m.Move
	}
	return 0
}

func (m *DoMoveRes) GetAccept() bool {
	if m != nil {
		return m.Accept
	}
	return false
}

type PlayerMove struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Move                 int64    `protobuf:"varint,2,opt,name=move,proto3" json:"move,omitempty"`
	Next                 int64    `protobuf:"varint,3,opt,name=next,proto3" json:"next,omitempty"`
	TimeLeft             int64    `protobuf:"varint,4,opt,name=time_left,json=timeLeft,proto3" json:"time_left,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlayerMove) Reset()         { *m = PlayerMove{} }
func (m *PlayerMove) String() string { return proto.CompactTextString(m) }
func (*PlayerMove) ProtoMessage()    {}
func (*PlayerMove) Descriptor() ([]byte, []int) {
	return fileDescriptor_f42bf8acbb5981a6, []int{7}
}

func (m *PlayerMove) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerMove.Unmarshal(m, b)
}
func (m *PlayerMove) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerMove.Marshal(b, m, deterministic)
}
func (m *PlayerMove) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerMove.Merge(m, src)
}
func (m *PlayerMove) XXX_Size() int {
	return xxx_messageInfo_PlayerMove.Size(m)
}
func (m *PlayerMove) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerMove.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerMove proto.InternalMessageInfo

func (m *PlayerMove) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PlayerMove) GetMove() int64 {
	if m != nil {
		return m.Move
	}
	return 0
}

func (m *PlayerMove) GetNext() int64 {
	if m != nil {
		return m.Next
	}
	return 0
}

func (m *PlayerMove) GetTimeLeft() int64 {
	if m != nil {
		return m.TimeLeft
	}
	return 0
}

type EndGame struct {
	Winner               int64    `protobuf:"varint,1,opt,name=winner,proto3" json:"winner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EndGame) Reset()         { *m = EndGame{} }
func (m *EndGame) String() string { return proto.CompactTextString(m) }
func (*EndGame) ProtoMessage()    {}
func (*EndGame) Descriptor() ([]byte, []int) {
	return fileDescriptor_f42bf8acbb5981a6, []int{8}
}

func (m *EndGame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndGame.Unmarshal(m, b)
}
func (m *EndGame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndGame.Marshal(b, m, deterministic)
}
func (m *EndGame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndGame.Merge(m, src)
}
func (m *EndGame) XXX_Size() int {
	return xxx_messageInfo_EndGame.Size(m)
}
func (m *EndGame) XXX_DiscardUnknown() {
	xxx_messageInfo_EndGame.DiscardUnknown(m)
}

var xxx_messageInfo_EndGame proto.InternalMessageInfo

func (m *EndGame) GetWinner() int64 {
	if m != nil {
		return m.Winner
	}
	return 0
}

type PlayerLeft struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlayerLeft) Reset()         { *m = PlayerLeft{} }
func (m *PlayerLeft) String() string { return proto.CompactTextString(m) }
func (*PlayerLeft) ProtoMessage()    {}
func (*PlayerLeft) Descriptor() ([]byte, []int) {
	return fileDescriptor_f42bf8acbb5981a6, []int{9}
}

func (m *PlayerLeft) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerLeft.Unmarshal(m, b)
}
func (m *PlayerLeft) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerLeft.Marshal(b, m, deterministic)
}
func (m *PlayerLeft) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerLeft.Merge(m, src)
}
func (m *PlayerLeft) XXX_Size() int {
	return xxx_messageInfo_PlayerLeft.Size(m)
}
func (m *PlayerLeft) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerLeft.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerLeft proto.InternalMessageInfo

func (m *PlayerLeft) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ErrorMessage struct {
	Err                  string   `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorMessage) Reset()         { *m = ErrorMessage{} }
func (m *ErrorMessage) String() string { return proto.CompactTextString(m) }
func (*ErrorMessage) ProtoMessage()    {}
func (*ErrorMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_f42bf8acbb5981a6, []int{10}
}

func (m *ErrorMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorMessage.Unmarshal(m, b)
}
func (m *ErrorMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorMessage.Marshal(b, m, deterministic)
}
func (m *ErrorMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorMessage.Merge(m, src)
}
func (m *ErrorMessage) XXX_Size() int {
	return xxx_messageInfo_ErrorMessage.Size(m)
}
func (m *ErrorMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorMessage proto.InternalMessageInfo

func (m *ErrorMessage) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*Evt)(nil), "evt.Evt")
	proto.RegisterType((*PrepareToStart)(nil), "evt.PrepareToStart")
	proto.RegisterType((*GameStarted)(nil), "evt.GameStarted")
	proto.RegisterType((*PlayerDisconnected)(nil), "evt.PlayerDisconnected")
	proto.RegisterType((*PlayerReconnected)(nil), "evt.PlayerReconnected")
	proto.RegisterType((*DoMoveReq)(nil), "evt.DoMoveReq")
	proto.RegisterType((*DoMoveRes)(nil), "evt.DoMoveRes")
	proto.RegisterType((*PlayerMove)(nil), "evt.PlayerMove")
	proto.RegisterType((*EndGame)(nil), "evt.EndGame")
	proto.RegisterType((*PlayerLeft)(nil), "evt.PlayerLeft")
	proto.RegisterType((*ErrorMessage)(nil), "evt.ErrorMessage")
}

func init() { proto.RegisterFile("evt.proto", fileDescriptor_f42bf8acbb5981a6) }

var fileDescriptor_f42bf8acbb5981a6 = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xdf, 0x6b, 0xd3, 0x50,
	0x14, 0xc7, 0xb3, 0x26, 0x6b, 0x9b, 0x93, 0x5a, 0xdb, 0x33, 0x99, 0x01, 0x05, 0x67, 0xb6, 0x07,
	0x7d, 0x19, 0x32, 0x11, 0x7c, 0x13, 0x64, 0xa3, 0x41, 0x1c, 0xc8, 0x9d, 0xef, 0x21, 0x36, 0xa7,
	0x25, 0xb0, 0xe4, 0x66, 0x37, 0x97, 0xa8, 0xff, 0xa3, 0x7f, 0x94, 0xdc, 0x73, 0xb3, 0x26, 0x6d,
	0x7c, 0xf0, 0xed, 0xfc, 0xfa, 0x7c, 0x73, 0x38, 0xdf, 0x4b, 0xc0, 0xa7, 0x46, 0x5f, 0x56, 0x4a,
	0x6a, 0x89, 0x2e, 0x35, 0x3a, 0xfa, 0xe3, 0x81, 0x7b, 0xd3, 0x68, 0xfc, 0x04, 0x8b, 0x4a, 0x51,
	0x95, 0x2a, 0x4a, 0xb4, 0x4c, 0x6a, 0x9d, 0x2a, 0x1d, 0x1e, 0x9d, 0x1d, 0xbd, 0x09, 0xae, 0x4e,
	0x2e, 0x0d, 0xf2, 0xcd, 0x36, 0xbf, 0xcb, 0x3b, 0xd3, 0x8a, 0x1d, 0x31, 0xaf, 0xf6, 0x2a, 0xf8,
	0x01, 0x66, 0xdb, 0xb4, 0x20, 0x8b, 0x52, 0x16, 0x8e, 0x18, 0x5e, 0x30, 0xbc, 0x4a, 0x0b, 0xba,
	0xb3, 0xf5, 0xd8, 0x11, 0xc1, 0xb6, 0x4b, 0xf1, 0x0b, 0x9c, 0x54, 0xf7, 0xe9, 0x6f, 0x52, 0x49,
	0x96, 0xd7, 0x6b, 0x59, 0x96, 0xb4, 0x36, 0xb4, 0xcb, 0xf4, 0x73, 0xfb, 0x69, 0xee, 0x5f, 0xf7,
	0xda, 0xb1, 0x23, 0xb0, 0x1a, 0x54, 0x71, 0x05, 0x6d, 0x35, 0x51, 0xd4, 0x49, 0x79, 0x2c, 0x75,
	0xda, 0x93, 0x12, 0xd4, 0x57, 0x5a, 0x56, 0x87, 0x45, 0x7c, 0x07, 0x41, 0x26, 0x93, 0x42, 0x36,
	0x94, 0x28, 0x7a, 0x08, 0x8f, 0x59, 0x61, 0xce, 0x0a, 0xd7, 0xf2, 0x56, 0x36, 0x24, 0xe8, 0x21,
	0x76, 0x84, 0x9f, 0x3d, 0x26, 0xfb, 0x44, 0x1d, 0x8e, 0xff, 0x41, 0xd4, 0x7d, 0xa2, 0xc6, 0x2b,
	0x08, 0xda, 0x65, 0x0d, 0x15, 0x4e, 0x98, 0x78, 0xda, 0xdb, 0xd2, 0x0c, 0xc6, 0x8e, 0x80, 0x6a,
	0x97, 0xe1, 0x5b, 0x98, 0x52, 0x99, 0x25, 0xe6, 0x7e, 0xe1, 0x94, 0x81, 0x19, 0x03, 0x37, 0x65,
	0x66, 0x4e, 0x1c, 0x3b, 0x62, 0x42, 0x36, 0xec, 0xc9, 0xdf, 0xd3, 0x46, 0x87, 0xfe, 0x40, 0xfe,
	0x2b, 0x6d, 0x74, 0x27, 0x6f, 0x32, 0xfc, 0x08, 0x4f, 0x48, 0x29, 0xa9, 0x92, 0x82, 0xea, 0x3a,
	0xdd, 0x52, 0x08, 0x4c, 0x2d, 0xed, 0x37, 0x4c, 0xe7, 0xd6, 0x36, 0x62, 0x47, 0xcc, 0xa8, 0x97,
	0x7f, 0x9e, 0xc0, 0x31, 0x35, 0x54, 0xea, 0xe8, 0x02, 0xe6, 0xfb, 0x2f, 0x05, 0x11, 0x3c, 0x9d,
	0x17, 0xc4, 0x8f, 0xc9, 0x15, 0x1c, 0x47, 0xe7, 0x10, 0xf4, 0x9e, 0x04, 0x3e, 0x83, 0xe3, 0x4d,
	0xae, 0x6a, 0xdd, 0xce, 0xd8, 0x24, 0xba, 0x00, 0x1c, 0x3a, 0x8f, 0x73, 0x18, 0xe5, 0x59, 0x3b,
	0x38, 0xca, 0xb3, 0xe8, 0x1c, 0x96, 0x03, 0x53, 0x07, 0x43, 0xaf, 0xc0, 0xdf, 0xf9, 0x66, 0x16,
	0xe2, 0x8b, 0xb7, 0x0b, 0x99, 0x38, 0x5a, 0x75, 0x03, 0xf5, 0x21, 0xbd, 0x03, 0x46, 0x1d, 0x80,
	0xa7, 0x30, 0x4e, 0xd7, 0x6b, 0xaa, 0x34, 0xbf, 0xd4, 0xa9, 0x68, 0xb3, 0x28, 0x05, 0xe8, 0xdc,
	0xfb, 0x2f, 0x25, 0x04, 0xaf, 0xa4, 0x5f, 0x56, 0xc7, 0x15, 0x1c, 0xe3, 0x0b, 0xf0, 0xcd, 0x9d,
	0xac, 0x75, 0x1e, 0x37, 0xa6, 0xa6, 0x60, 0x5c, 0x8a, 0x5e, 0xc3, 0xa4, 0xf5, 0xdb, 0x6c, 0xf1,
	0x33, 0x2f, 0x4b, 0x52, 0xed, 0x37, 0xda, 0x2c, 0x7a, 0xf9, 0xb8, 0x05, 0xdb, 0x7a, 0x78, 0x8d,
	0x33, 0x98, 0xf5, 0xcd, 0xc4, 0x05, 0xb8, 0xa4, 0xac, 0x84, 0x2f, 0x4c, 0xf8, 0x63, 0xcc, 0x3f,
	0x88, 0xf7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x94, 0x4d, 0xbd, 0x2d, 0x04, 0x00, 0x00,
}
