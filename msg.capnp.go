package factom

// AUTO GENERATED - DO NOT EDIT

import (
	C "github.com/glycerine/go-capnproto"
	"unsafe"
)

type CommitHeader C.Struct

func NewCommitHeader(s *C.Segment) CommitHeader      { return CommitHeader(s.NewStruct(24, 0)) }
func NewRootCommitHeader(s *C.Segment) CommitHeader  { return CommitHeader(s.NewRootStruct(24, 0)) }
func AutoNewCommitHeader(s *C.Segment) CommitHeader  { return CommitHeader(s.NewStructAR(24, 0)) }
func ReadRootCommitHeader(s *C.Segment) CommitHeader { return CommitHeader(s.Root(0).ToStruct()) }
func (s CommitHeader) Version() int8                 { return int8(C.Struct(s).Get8(0)) }
func (s CommitHeader) SetVersion(v int8)             { C.Struct(s).Set8(0, uint8(v)) }
func (s CommitHeader) TimeNounce() uint64            { return C.Struct(s).Get64(8) }
func (s CommitHeader) SetTimeNounce(v uint64)        { C.Struct(s).Set64(8, v) }
func (s CommitHeader) PayAmount() uint64             { return C.Struct(s).Get64(16) }
func (s CommitHeader) SetPayAmount(v uint64)         { C.Struct(s).Set64(16, v) }

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s CommitHeader) MarshalJSON() (bs []byte, err error) { return }

type CommitHeader_List C.PointerList

func NewCommitHeaderList(s *C.Segment, sz int) CommitHeader_List {
	return CommitHeader_List(s.NewCompositeList(24, 0, sz))
}
func (s CommitHeader_List) Len() int { return C.PointerList(s).Len() }
func (s CommitHeader_List) At(i int) CommitHeader {
	return CommitHeader(C.PointerList(s).At(i).ToStruct())
}
func (s CommitHeader_List) ToArray() []CommitHeader {
	return *(*[]CommitHeader)(unsafe.Pointer(C.PointerList(s).ToArray()))
}
func (s CommitHeader_List) Set(i int, item CommitHeader) { C.PointerList(s).Set(i, C.Object(item)) }

type CommitEntryMsg C.Struct

func NewCommitEntryMsg(s *C.Segment) CommitEntryMsg      { return CommitEntryMsg(s.NewStruct(0, 2)) }
func NewRootCommitEntryMsg(s *C.Segment) CommitEntryMsg  { return CommitEntryMsg(s.NewRootStruct(0, 2)) }
func AutoNewCommitEntryMsg(s *C.Segment) CommitEntryMsg  { return CommitEntryMsg(s.NewStructAR(0, 2)) }
func ReadRootCommitEntryMsg(s *C.Segment) CommitEntryMsg { return CommitEntryMsg(s.Root(0).ToStruct()) }
func (s CommitEntryMsg) Head() CommitHeader              { return CommitHeader(C.Struct(s).GetObject(0).ToStruct()) }
func (s CommitEntryMsg) SetHead(v CommitHeader)          { C.Struct(s).SetObject(0, C.Object(v)) }
func (s CommitEntryMsg) HashEntry() []byte               { return C.Struct(s).GetObject(1).ToData() }
func (s CommitEntryMsg) SetHashEntry(v []byte)           { C.Struct(s).SetObject(1, s.Segment.NewData(v)) }

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s CommitEntryMsg) MarshalJSON() (bs []byte, err error) { return }

type CommitEntryMsg_List C.PointerList

func NewCommitEntryMsgList(s *C.Segment, sz int) CommitEntryMsg_List {
	return CommitEntryMsg_List(s.NewCompositeList(0, 2, sz))
}
func (s CommitEntryMsg_List) Len() int { return C.PointerList(s).Len() }
func (s CommitEntryMsg_List) At(i int) CommitEntryMsg {
	return CommitEntryMsg(C.PointerList(s).At(i).ToStruct())
}
func (s CommitEntryMsg_List) ToArray() []CommitEntryMsg {
	return *(*[]CommitEntryMsg)(unsafe.Pointer(C.PointerList(s).ToArray()))
}
func (s CommitEntryMsg_List) Set(i int, item CommitEntryMsg) { C.PointerList(s).Set(i, C.Object(item)) }

type CommitChainMsg C.Struct

func NewCommitChainMsg(s *C.Segment) CommitChainMsg      { return CommitChainMsg(s.NewStruct(0, 4)) }
func NewRootCommitChainMsg(s *C.Segment) CommitChainMsg  { return CommitChainMsg(s.NewRootStruct(0, 4)) }
func AutoNewCommitChainMsg(s *C.Segment) CommitChainMsg  { return CommitChainMsg(s.NewStructAR(0, 4)) }
func ReadRootCommitChainMsg(s *C.Segment) CommitChainMsg { return CommitChainMsg(s.Root(0).ToStruct()) }
func (s CommitChainMsg) Head() CommitHeader              { return CommitHeader(C.Struct(s).GetObject(0).ToStruct()) }
func (s CommitChainMsg) SetHead(v CommitHeader)          { C.Struct(s).SetObject(0, C.Object(v)) }
func (s CommitChainMsg) HashChain() []byte               { return C.Struct(s).GetObject(1).ToData() }
func (s CommitChainMsg) SetHashChain(v []byte)           { C.Struct(s).SetObject(1, s.Segment.NewData(v)) }
func (s CommitChainMsg) HashChainEntry() []byte          { return C.Struct(s).GetObject(2).ToData() }
func (s CommitChainMsg) SetHashChainEntry(v []byte)      { C.Struct(s).SetObject(2, s.Segment.NewData(v)) }
func (s CommitChainMsg) HashEntry() []byte               { return C.Struct(s).GetObject(3).ToData() }
func (s CommitChainMsg) SetHashEntry(v []byte)           { C.Struct(s).SetObject(3, s.Segment.NewData(v)) }

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s CommitChainMsg) MarshalJSON() (bs []byte, err error) { return }

type CommitChainMsg_List C.PointerList

func NewCommitChainMsgList(s *C.Segment, sz int) CommitChainMsg_List {
	return CommitChainMsg_List(s.NewCompositeList(0, 4, sz))
}
func (s CommitChainMsg_List) Len() int { return C.PointerList(s).Len() }
func (s CommitChainMsg_List) At(i int) CommitChainMsg {
	return CommitChainMsg(C.PointerList(s).At(i).ToStruct())
}
func (s CommitChainMsg_List) ToArray() []CommitChainMsg {
	return *(*[]CommitChainMsg)(unsafe.Pointer(C.PointerList(s).ToArray()))
}
func (s CommitChainMsg_List) Set(i int, item CommitChainMsg) { C.PointerList(s).Set(i, C.Object(item)) }

// commit msg
type CommitMsg C.Struct
type CommitMsg_Which uint16

const (
	COMMITMSG_ENTRY CommitMsg_Which = 0
	COMMITMSG_CHAIN CommitMsg_Which = 1
)

func NewCommitMsg(s *C.Segment) CommitMsg      { return CommitMsg(s.NewStruct(8, 1)) }
func NewRootCommitMsg(s *C.Segment) CommitMsg  { return CommitMsg(s.NewRootStruct(8, 1)) }
func AutoNewCommitMsg(s *C.Segment) CommitMsg  { return CommitMsg(s.NewStructAR(8, 1)) }
func ReadRootCommitMsg(s *C.Segment) CommitMsg { return CommitMsg(s.Root(0).ToStruct()) }
func (s CommitMsg) Which() CommitMsg_Which     { return CommitMsg_Which(C.Struct(s).Get16(0)) }
func (s CommitMsg) Entry() CommitEntryMsg      { return CommitEntryMsg(C.Struct(s).GetObject(0).ToStruct()) }
func (s CommitMsg) SetEntry(v CommitEntryMsg) {
	C.Struct(s).Set16(0, 0)
	C.Struct(s).SetObject(0, C.Object(v))
}
func (s CommitMsg) Chain() CommitChainMsg { return CommitChainMsg(C.Struct(s).GetObject(0).ToStruct()) }
func (s CommitMsg) SetChain(v CommitChainMsg) {
	C.Struct(s).Set16(0, 1)
	C.Struct(s).SetObject(0, C.Object(v))
}

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s CommitMsg) MarshalJSON() (bs []byte, err error) { return }

type CommitMsg_List C.PointerList

func NewCommitMsgList(s *C.Segment, sz int) CommitMsg_List {
	return CommitMsg_List(s.NewCompositeList(8, 1, sz))
}
func (s CommitMsg_List) Len() int           { return C.PointerList(s).Len() }
func (s CommitMsg_List) At(i int) CommitMsg { return CommitMsg(C.PointerList(s).At(i).ToStruct()) }
func (s CommitMsg_List) ToArray() []CommitMsg {
	return *(*[]CommitMsg)(unsafe.Pointer(C.PointerList(s).ToArray()))
}
func (s CommitMsg_List) Set(i int, item CommitMsg) { C.PointerList(s).Set(i, C.Object(item)) }

// signed msg
type SignedMsg C.Struct
type SignedMsgData SignedMsg
type SignedMsgData_Which uint16

const (
	SIGNEDMSGDATA_COMMIT SignedMsgData_Which = 0
	SIGNEDMSGDATA_OTHER  SignedMsgData_Which = 1
)

func NewSignedMsg(s *C.Segment) SignedMsg          { return SignedMsg(s.NewStruct(8, 3)) }
func NewRootSignedMsg(s *C.Segment) SignedMsg      { return SignedMsg(s.NewRootStruct(8, 3)) }
func AutoNewSignedMsg(s *C.Segment) SignedMsg      { return SignedMsg(s.NewStructAR(8, 3)) }
func ReadRootSignedMsg(s *C.Segment) SignedMsg     { return SignedMsg(s.Root(0).ToStruct()) }
func (s SignedMsg) Data() SignedMsgData            { return SignedMsgData(s) }
func (s SignedMsgData) Which() SignedMsgData_Which { return SignedMsgData_Which(C.Struct(s).Get16(0)) }
func (s SignedMsgData) Commit() CommitMsg          { return CommitMsg(C.Struct(s).GetObject(0).ToStruct()) }
func (s SignedMsgData) SetCommit(v CommitMsg) {
	C.Struct(s).Set16(0, 0)
	C.Struct(s).SetObject(0, C.Object(v))
}
func (s SignedMsgData) SetOther()          { C.Struct(s).Set16(0, 1) }
func (s SignedMsg) SignedData() []byte     { return C.Struct(s).GetObject(1).ToData() }
func (s SignedMsg) SetSignedData(v []byte) { C.Struct(s).SetObject(1, s.Segment.NewData(v)) }
func (s SignedMsg) PublicKey() []byte      { return C.Struct(s).GetObject(2).ToData() }
func (s SignedMsg) SetPublicKey(v []byte)  { C.Struct(s).SetObject(2, s.Segment.NewData(v)) }

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s SignedMsg) MarshalJSON() (bs []byte, err error) { return }

type SignedMsg_List C.PointerList

func NewSignedMsgList(s *C.Segment, sz int) SignedMsg_List {
	return SignedMsg_List(s.NewCompositeList(8, 3, sz))
}
func (s SignedMsg_List) Len() int           { return C.PointerList(s).Len() }
func (s SignedMsg_List) At(i int) SignedMsg { return SignedMsg(C.PointerList(s).At(i).ToStruct()) }
func (s SignedMsg_List) ToArray() []SignedMsg {
	return *(*[]SignedMsg)(unsafe.Pointer(C.PointerList(s).ToArray()))
}
func (s SignedMsg_List) Set(i int, item SignedMsg) { C.PointerList(s).Set(i, C.Object(item)) }
