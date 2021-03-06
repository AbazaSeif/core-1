// Code generated by protoc-gen-go. DO NOT EDIT.
// source: insonmnia.proto

package sonm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TaskLogsRequest_Type int32

const (
	TaskLogsRequest_STDOUT TaskLogsRequest_Type = 0
	TaskLogsRequest_STDERR TaskLogsRequest_Type = 1
	TaskLogsRequest_BOTH   TaskLogsRequest_Type = 2
)

var TaskLogsRequest_Type_name = map[int32]string{
	0: "STDOUT",
	1: "STDERR",
	2: "BOTH",
}
var TaskLogsRequest_Type_value = map[string]int32{
	"STDOUT": 0,
	"STDERR": 1,
	"BOTH":   2,
}

func (x TaskLogsRequest_Type) String() string {
	return proto.EnumName(TaskLogsRequest_Type_name, int32(x))
}
func (TaskLogsRequest_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor8, []int{10, 0} }

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

type ID struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *ID) Reset()                    { *m = ID{} }
func (m *ID) String() string            { return proto.CompactTextString(m) }
func (*ID) ProtoMessage()               {}
func (*ID) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

func (m *ID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type NumericID struct {
	Id uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *NumericID) Reset()                    { *m = NumericID{} }
func (m *NumericID) String() string            { return proto.CompactTextString(m) }
func (*NumericID) ProtoMessage()               {}
func (*NumericID) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{2} }

func (m *NumericID) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type EthID struct {
	Id *EthAddress `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *EthID) Reset()                    { *m = EthID{} }
func (m *EthID) String() string            { return proto.CompactTextString(m) }
func (*EthID) ProtoMessage()               {}
func (*EthID) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{3} }

func (m *EthID) GetId() *EthAddress {
	if m != nil {
		return m.Id
	}
	return nil
}

type TaskID struct {
	// Id is task ID itself
	Id     string  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	DealID *BigInt `protobuf:"bytes,2,opt,name=dealID" json:"dealID,omitempty"`
}

func (m *TaskID) Reset()                    { *m = TaskID{} }
func (m *TaskID) String() string            { return proto.CompactTextString(m) }
func (*TaskID) ProtoMessage()               {}
func (*TaskID) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{4} }

func (m *TaskID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TaskID) GetDealID() *BigInt {
	if m != nil {
		return m.DealID
	}
	return nil
}

type Count struct {
	Count uint64 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
}

func (m *Count) Reset()                    { *m = Count{} }
func (m *Count) String() string            { return proto.CompactTextString(m) }
func (*Count) ProtoMessage()               {}
func (*Count) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{5} }

func (m *Count) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type CPUUsage struct {
	Total uint64 `protobuf:"varint,1,opt,name=total" json:"total,omitempty"`
}

func (m *CPUUsage) Reset()                    { *m = CPUUsage{} }
func (m *CPUUsage) String() string            { return proto.CompactTextString(m) }
func (*CPUUsage) ProtoMessage()               {}
func (*CPUUsage) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{6} }

func (m *CPUUsage) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

type MemoryUsage struct {
	MaxUsage uint64 `protobuf:"varint,1,opt,name=maxUsage" json:"maxUsage,omitempty"`
}

func (m *MemoryUsage) Reset()                    { *m = MemoryUsage{} }
func (m *MemoryUsage) String() string            { return proto.CompactTextString(m) }
func (*MemoryUsage) ProtoMessage()               {}
func (*MemoryUsage) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{7} }

func (m *MemoryUsage) GetMaxUsage() uint64 {
	if m != nil {
		return m.MaxUsage
	}
	return 0
}

type NetworkUsage struct {
	TxBytes       uint64  `protobuf:"varint,1,opt,name=txBytes" json:"txBytes,omitempty"`
	RxBytes       uint64  `protobuf:"varint,2,opt,name=rxBytes" json:"rxBytes,omitempty"`
	TxPackets     uint64  `protobuf:"varint,3,opt,name=txPackets" json:"txPackets,omitempty"`
	RxPackets     uint64  `protobuf:"varint,4,opt,name=rxPackets" json:"rxPackets,omitempty"`
	TxErrors      uint64  `protobuf:"varint,5,opt,name=txErrors" json:"txErrors,omitempty"`
	RxErrors      uint64  `protobuf:"varint,6,opt,name=rxErrors" json:"rxErrors,omitempty"`
	TxDropped     uint64  `protobuf:"varint,7,opt,name=txDropped" json:"txDropped,omitempty"`
	RxDropped     uint64  `protobuf:"varint,8,opt,name=rxDropped" json:"rxDropped,omitempty"`
	TxBytesRate   float64 `protobuf:"fixed64,9,opt,name=txBytesRate" json:"txBytesRate,omitempty"`
	RxBytesRate   float64 `protobuf:"fixed64,10,opt,name=rxBytesRate" json:"rxBytesRate,omitempty"`
	TxPacketsRate float64 `protobuf:"fixed64,11,opt,name=txPacketsRate" json:"txPacketsRate,omitempty"`
	RxPacketsRate float64 `protobuf:"fixed64,12,opt,name=rxPacketsRate" json:"rxPacketsRate,omitempty"`
	TxErrorsRate  float64 `protobuf:"fixed64,13,opt,name=txErrorsRate" json:"txErrorsRate,omitempty"`
	RxErrorsRate  float64 `protobuf:"fixed64,14,opt,name=rxErrorsRate" json:"rxErrorsRate,omitempty"`
	TxDroppedRate float64 `protobuf:"fixed64,15,opt,name=txDroppedRate" json:"txDroppedRate,omitempty"`
	RxDroppedRate float64 `protobuf:"fixed64,16,opt,name=rxDroppedRate" json:"rxDroppedRate,omitempty"`
}

func (m *NetworkUsage) Reset()                    { *m = NetworkUsage{} }
func (m *NetworkUsage) String() string            { return proto.CompactTextString(m) }
func (*NetworkUsage) ProtoMessage()               {}
func (*NetworkUsage) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{8} }

func (m *NetworkUsage) GetTxBytes() uint64 {
	if m != nil {
		return m.TxBytes
	}
	return 0
}

func (m *NetworkUsage) GetRxBytes() uint64 {
	if m != nil {
		return m.RxBytes
	}
	return 0
}

func (m *NetworkUsage) GetTxPackets() uint64 {
	if m != nil {
		return m.TxPackets
	}
	return 0
}

func (m *NetworkUsage) GetRxPackets() uint64 {
	if m != nil {
		return m.RxPackets
	}
	return 0
}

func (m *NetworkUsage) GetTxErrors() uint64 {
	if m != nil {
		return m.TxErrors
	}
	return 0
}

func (m *NetworkUsage) GetRxErrors() uint64 {
	if m != nil {
		return m.RxErrors
	}
	return 0
}

func (m *NetworkUsage) GetTxDropped() uint64 {
	if m != nil {
		return m.TxDropped
	}
	return 0
}

func (m *NetworkUsage) GetRxDropped() uint64 {
	if m != nil {
		return m.RxDropped
	}
	return 0
}

func (m *NetworkUsage) GetTxBytesRate() float64 {
	if m != nil {
		return m.TxBytesRate
	}
	return 0
}

func (m *NetworkUsage) GetRxBytesRate() float64 {
	if m != nil {
		return m.RxBytesRate
	}
	return 0
}

func (m *NetworkUsage) GetTxPacketsRate() float64 {
	if m != nil {
		return m.TxPacketsRate
	}
	return 0
}

func (m *NetworkUsage) GetRxPacketsRate() float64 {
	if m != nil {
		return m.RxPacketsRate
	}
	return 0
}

func (m *NetworkUsage) GetTxErrorsRate() float64 {
	if m != nil {
		return m.TxErrorsRate
	}
	return 0
}

func (m *NetworkUsage) GetRxErrorsRate() float64 {
	if m != nil {
		return m.RxErrorsRate
	}
	return 0
}

func (m *NetworkUsage) GetTxDroppedRate() float64 {
	if m != nil {
		return m.TxDroppedRate
	}
	return 0
}

func (m *NetworkUsage) GetRxDroppedRate() float64 {
	if m != nil {
		return m.RxDroppedRate
	}
	return 0
}

type ResourceUsage struct {
	Cpu     *CPUUsage                `protobuf:"bytes,1,opt,name=cpu" json:"cpu,omitempty"`
	Memory  *MemoryUsage             `protobuf:"bytes,2,opt,name=memory" json:"memory,omitempty"`
	Network map[string]*NetworkUsage `protobuf:"bytes,3,rep,name=network" json:"network,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ResourceUsage) Reset()                    { *m = ResourceUsage{} }
func (m *ResourceUsage) String() string            { return proto.CompactTextString(m) }
func (*ResourceUsage) ProtoMessage()               {}
func (*ResourceUsage) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{9} }

func (m *ResourceUsage) GetCpu() *CPUUsage {
	if m != nil {
		return m.Cpu
	}
	return nil
}

func (m *ResourceUsage) GetMemory() *MemoryUsage {
	if m != nil {
		return m.Memory
	}
	return nil
}

func (m *ResourceUsage) GetNetwork() map[string]*NetworkUsage {
	if m != nil {
		return m.Network
	}
	return nil
}

type TaskLogsRequest struct {
	Type          TaskLogsRequest_Type `protobuf:"varint,1,opt,name=type,enum=sonm.TaskLogsRequest_Type" json:"type,omitempty"`
	Id            string               `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Since         string               `protobuf:"bytes,3,opt,name=since" json:"since,omitempty"`
	AddTimestamps bool                 `protobuf:"varint,4,opt,name=addTimestamps" json:"addTimestamps,omitempty"`
	Follow        bool                 `protobuf:"varint,5,opt,name=Follow" json:"Follow,omitempty"`
	Tail          string               `protobuf:"bytes,6,opt,name=Tail" json:"Tail,omitempty"`
	Details       bool                 `protobuf:"varint,7,opt,name=Details" json:"Details,omitempty"`
	DealID        *BigInt              `protobuf:"bytes,8,opt,name=dealID" json:"dealID,omitempty"`
}

func (m *TaskLogsRequest) Reset()                    { *m = TaskLogsRequest{} }
func (m *TaskLogsRequest) String() string            { return proto.CompactTextString(m) }
func (*TaskLogsRequest) ProtoMessage()               {}
func (*TaskLogsRequest) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{10} }

func (m *TaskLogsRequest) GetType() TaskLogsRequest_Type {
	if m != nil {
		return m.Type
	}
	return TaskLogsRequest_STDOUT
}

func (m *TaskLogsRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TaskLogsRequest) GetSince() string {
	if m != nil {
		return m.Since
	}
	return ""
}

func (m *TaskLogsRequest) GetAddTimestamps() bool {
	if m != nil {
		return m.AddTimestamps
	}
	return false
}

func (m *TaskLogsRequest) GetFollow() bool {
	if m != nil {
		return m.Follow
	}
	return false
}

func (m *TaskLogsRequest) GetTail() string {
	if m != nil {
		return m.Tail
	}
	return ""
}

func (m *TaskLogsRequest) GetDetails() bool {
	if m != nil {
		return m.Details
	}
	return false
}

func (m *TaskLogsRequest) GetDealID() *BigInt {
	if m != nil {
		return m.DealID
	}
	return nil
}

type TaskLogsChunk struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *TaskLogsChunk) Reset()                    { *m = TaskLogsChunk{} }
func (m *TaskLogsChunk) String() string            { return proto.CompactTextString(m) }
func (*TaskLogsChunk) ProtoMessage()               {}
func (*TaskLogsChunk) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{11} }

func (m *TaskLogsChunk) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Chunk struct {
	Chunk []byte `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (m *Chunk) Reset()                    { *m = Chunk{} }
func (m *Chunk) String() string            { return proto.CompactTextString(m) }
func (*Chunk) ProtoMessage()               {}
func (*Chunk) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{12} }

func (m *Chunk) GetChunk() []byte {
	if m != nil {
		return m.Chunk
	}
	return nil
}

type Progress struct {
	Size int64 `protobuf:"varint,1,opt,name=size" json:"size,omitempty"`
}

func (m *Progress) Reset()                    { *m = Progress{} }
func (m *Progress) String() string            { return proto.CompactTextString(m) }
func (*Progress) ProtoMessage()               {}
func (*Progress) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{13} }

func (m *Progress) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

type Duration struct {
	Nanoseconds int64 `protobuf:"varint,1,opt,name=nanoseconds" json:"nanoseconds,omitempty"`
}

func (m *Duration) Reset()                    { *m = Duration{} }
func (m *Duration) String() string            { return proto.CompactTextString(m) }
func (*Duration) ProtoMessage()               {}
func (*Duration) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{14} }

func (m *Duration) GetNanoseconds() int64 {
	if m != nil {
		return m.Nanoseconds
	}
	return 0
}

type EthAddress struct {
	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *EthAddress) Reset()                    { *m = EthAddress{} }
func (m *EthAddress) String() string            { return proto.CompactTextString(m) }
func (*EthAddress) ProtoMessage()               {}
func (*EthAddress) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{15} }

func (m *EthAddress) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

type DataSize struct {
	Bytes uint64 `protobuf:"varint,1,opt,name=bytes" json:"bytes,omitempty"`
}

func (m *DataSize) Reset()                    { *m = DataSize{} }
func (m *DataSize) String() string            { return proto.CompactTextString(m) }
func (*DataSize) ProtoMessage()               {}
func (*DataSize) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{16} }

func (m *DataSize) GetBytes() uint64 {
	if m != nil {
		return m.Bytes
	}
	return 0
}

type DataSizeRate struct {
	BitsPerSecond uint64 `protobuf:"varint,1,opt,name=bitsPerSecond" json:"bitsPerSecond,omitempty"`
}

func (m *DataSizeRate) Reset()                    { *m = DataSizeRate{} }
func (m *DataSizeRate) String() string            { return proto.CompactTextString(m) }
func (*DataSizeRate) ProtoMessage()               {}
func (*DataSizeRate) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{17} }

func (m *DataSizeRate) GetBitsPerSecond() uint64 {
	if m != nil {
		return m.BitsPerSecond
	}
	return 0
}

type Price struct {
	PerSecond *BigInt `protobuf:"bytes,1,opt,name=perSecond" json:"perSecond,omitempty"`
}

func (m *Price) Reset()                    { *m = Price{} }
func (m *Price) String() string            { return proto.CompactTextString(m) }
func (*Price) ProtoMessage()               {}
func (*Price) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{18} }

func (m *Price) GetPerSecond() *BigInt {
	if m != nil {
		return m.PerSecond
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "sonm.Empty")
	proto.RegisterType((*ID)(nil), "sonm.ID")
	proto.RegisterType((*NumericID)(nil), "sonm.NumericID")
	proto.RegisterType((*EthID)(nil), "sonm.EthID")
	proto.RegisterType((*TaskID)(nil), "sonm.TaskID")
	proto.RegisterType((*Count)(nil), "sonm.Count")
	proto.RegisterType((*CPUUsage)(nil), "sonm.CPUUsage")
	proto.RegisterType((*MemoryUsage)(nil), "sonm.MemoryUsage")
	proto.RegisterType((*NetworkUsage)(nil), "sonm.NetworkUsage")
	proto.RegisterType((*ResourceUsage)(nil), "sonm.ResourceUsage")
	proto.RegisterType((*TaskLogsRequest)(nil), "sonm.TaskLogsRequest")
	proto.RegisterType((*TaskLogsChunk)(nil), "sonm.TaskLogsChunk")
	proto.RegisterType((*Chunk)(nil), "sonm.Chunk")
	proto.RegisterType((*Progress)(nil), "sonm.Progress")
	proto.RegisterType((*Duration)(nil), "sonm.Duration")
	proto.RegisterType((*EthAddress)(nil), "sonm.EthAddress")
	proto.RegisterType((*DataSize)(nil), "sonm.DataSize")
	proto.RegisterType((*DataSizeRate)(nil), "sonm.DataSizeRate")
	proto.RegisterType((*Price)(nil), "sonm.Price")
	proto.RegisterEnum("sonm.TaskLogsRequest_Type", TaskLogsRequest_Type_name, TaskLogsRequest_Type_value)
}

func init() { proto.RegisterFile("insonmnia.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 838 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x95, 0xdd, 0x6e, 0x2b, 0x35,
	0x10, 0xc7, 0xc9, 0xe6, 0xa3, 0x9b, 0x49, 0xda, 0x06, 0xab, 0x42, 0x51, 0xf8, 0x50, 0x64, 0x2a,
	0x94, 0x22, 0x48, 0xa5, 0x1e, 0x2e, 0xd0, 0x41, 0x42, 0xa2, 0x4d, 0x10, 0x95, 0xa0, 0x27, 0x72,
	0xd3, 0x1b, 0xee, 0x9c, 0x5d, 0x2b, 0xb5, 0xb2, 0xbb, 0x5e, 0x6c, 0x2f, 0x27, 0x39, 0x0f, 0xc2,
	0x0b, 0xf0, 0x72, 0x3c, 0x06, 0xf2, 0x57, 0xba, 0x5b, 0x24, 0xee, 0x66, 0xfe, 0xff, 0xdf, 0xc6,
	0x63, 0x6b, 0x66, 0x02, 0xe7, 0xbc, 0x50, 0xa2, 0xc8, 0x0b, 0x4e, 0xe7, 0xa5, 0x14, 0x5a, 0xa0,
	0x8e, 0x49, 0x27, 0xc3, 0x0d, 0xdf, 0xf2, 0x42, 0x3b, 0x6d, 0x82, 0x12, 0x5a, 0xd2, 0x0d, 0xcf,
	0xb8, 0xe6, 0x4c, 0x79, 0xed, 0x5c, 0xf3, 0x9c, 0x29, 0x4d, 0xf3, 0xd2, 0x09, 0xf8, 0x04, 0xba,
	0xcb, 0xbc, 0xd4, 0x07, 0x7c, 0x01, 0xd1, 0xfd, 0x02, 0x9d, 0x41, 0xc4, 0xd3, 0x71, 0x6b, 0xda,
	0x9a, 0xf5, 0x49, 0xc4, 0x53, 0xfc, 0x29, 0xf4, 0x1f, 0xaa, 0x9c, 0x49, 0x9e, 0x34, 0xcc, 0x8e,
	0x35, 0xaf, 0xa0, 0xbb, 0xd4, 0xcf, 0xf7, 0x0b, 0x34, 0x3d, 0x1a, 0x83, 0x9b, 0xd1, 0xdc, 0x94,
	0x32, 0x5f, 0xea, 0xe7, 0x9f, 0xd2, 0x54, 0x32, 0xa5, 0x2c, 0xfa, 0x23, 0xf4, 0xd6, 0x54, 0xed,
	0xfe, 0x7b, 0x02, 0xba, 0x84, 0x5e, 0xca, 0x68, 0x76, 0xbf, 0x18, 0x47, 0xf6, 0xfb, 0xa1, 0xfb,
	0xfe, 0x96, 0x6f, 0xef, 0x0b, 0x4d, 0xbc, 0x87, 0x3f, 0x87, 0xee, 0x9d, 0xa8, 0x0a, 0x8d, 0x2e,
	0xa0, 0x9b, 0x98, 0xc0, 0x97, 0xe1, 0x12, 0x3c, 0x85, 0xf8, 0x6e, 0xf5, 0xf4, 0xa4, 0xe8, 0x96,
	0x19, 0x42, 0x0b, 0x4d, 0xb3, 0x40, 0xd8, 0x04, 0x5f, 0xc1, 0xe0, 0x37, 0x96, 0x0b, 0x79, 0x70,
	0xd0, 0x04, 0xe2, 0x9c, 0xee, 0x6d, 0xec, 0xb9, 0x63, 0x8e, 0xff, 0xea, 0xc0, 0xf0, 0x81, 0xe9,
	0xf7, 0x42, 0xee, 0x1c, 0x3c, 0x86, 0x13, 0xbd, 0xbf, 0x3d, 0x68, 0xa6, 0x3c, 0x1b, 0x52, 0xe3,
	0x48, 0xef, 0x44, 0xce, 0xf1, 0x29, 0xfa, 0x0c, 0xfa, 0x7a, 0xbf, 0xa2, 0xc9, 0x8e, 0x69, 0x35,
	0x6e, 0x5b, 0xef, 0x45, 0x30, 0xae, 0x3c, 0xba, 0x1d, 0xe7, 0x1e, 0x05, 0x53, 0x9c, 0xde, 0x2f,
	0xa5, 0x14, 0x52, 0x8d, 0xbb, 0xae, 0xb8, 0x90, 0x1b, 0x4f, 0x06, 0xaf, 0xe7, 0xbc, 0x90, 0xbb,
	0x33, 0x17, 0x52, 0x94, 0x25, 0x4b, 0xc7, 0x27, 0xe1, 0x4c, 0x2f, 0xb8, 0x33, 0x83, 0x1b, 0x87,
	0x33, 0x83, 0x3b, 0x85, 0x81, 0xbf, 0x14, 0xa1, 0x9a, 0x8d, 0xfb, 0xd3, 0xd6, 0xac, 0x45, 0xea,
	0x92, 0x21, 0x64, 0x8d, 0x00, 0x47, 0xd4, 0x24, 0x74, 0x09, 0xa7, 0xc7, 0x2b, 0x5a, 0x66, 0x60,
	0x99, 0xa6, 0x68, 0x28, 0xd9, 0xa0, 0x86, 0x8e, 0x6a, 0x88, 0x08, 0xc3, 0x30, 0xdc, 0xd9, 0x42,
	0xa7, 0x16, 0x6a, 0x68, 0x86, 0x91, 0x75, 0xe6, 0xcc, 0x31, 0x75, 0xcd, 0xd5, 0xe4, 0x2f, 0x69,
	0xa1, 0xf3, 0x50, 0x53, 0x4d, 0x74, 0x35, 0xd5, 0xa9, 0x51, 0xa8, 0xa9, 0x26, 0xe2, 0x7f, 0x5a,
	0x70, 0x4a, 0x98, 0x12, 0x95, 0x4c, 0x98, 0xeb, 0x8c, 0x29, 0xb4, 0x93, 0xb2, 0xf2, 0x9d, 0x7f,
	0xe6, 0x3a, 0x37, 0x34, 0x22, 0x31, 0x16, 0xba, 0x82, 0x5e, 0x6e, 0xfb, 0xce, 0xb7, 0xf7, 0xc7,
	0x0e, 0xaa, 0xf5, 0x22, 0xf1, 0x00, 0x7a, 0x0b, 0x27, 0x85, 0x6b, 0xbb, 0x71, 0x7b, 0xda, 0x9e,
	0x0d, 0x6e, 0xa6, 0x8e, 0x6d, 0x1c, 0x39, 0xf7, 0x9d, 0xb9, 0x2c, 0xb4, 0x3c, 0x90, 0xf0, 0xc1,
	0xe4, 0xe1, 0xd8, 0xb2, 0xd6, 0x40, 0x23, 0x68, 0xef, 0xd8, 0xc1, 0x8f, 0x99, 0x09, 0xd1, 0x0c,
	0xba, 0x7f, 0xd2, 0xac, 0x62, 0xbe, 0x0e, 0xe4, 0x7e, 0xbb, 0xde, 0xe7, 0xc4, 0x01, 0x6f, 0xa3,
	0xef, 0x5b, 0xf8, 0xef, 0x08, 0xce, 0xcd, 0xc0, 0xfe, 0x2a, 0xb6, 0x8a, 0xb0, 0x3f, 0x2a, 0xa6,
	0x34, 0x9a, 0x43, 0x47, 0x1f, 0x4a, 0x37, 0x2f, 0x67, 0x37, 0x13, 0xf7, 0x03, 0xaf, 0xa0, 0xf9,
	0xfa, 0x50, 0x32, 0x62, 0x39, 0x3f, 0xe9, 0xd1, 0x71, 0xd2, 0x2f, 0xa0, 0xab, 0x78, 0x91, 0x30,
	0x3b, 0x0e, 0x7d, 0xe2, 0x12, 0xf3, 0xf4, 0x34, 0x4d, 0xd7, 0x61, 0x2d, 0xb9, 0x71, 0x88, 0x49,
	0x53, 0x44, 0x9f, 0x40, 0xef, 0x67, 0x91, 0x65, 0xe2, 0xbd, 0x1d, 0x88, 0x98, 0xf8, 0x0c, 0x21,
	0xe8, 0xac, 0x29, 0xcf, 0xec, 0x28, 0xf4, 0x89, 0x8d, 0xcd, 0x50, 0x2e, 0x98, 0xa6, 0x3c, 0x53,
	0x76, 0x08, 0x62, 0x12, 0xd2, 0xda, 0xae, 0x89, 0xff, 0x67, 0xd7, 0xcc, 0xa0, 0x63, 0x6e, 0x81,
	0x00, 0x7a, 0x8f, 0xeb, 0xc5, 0xbb, 0xa7, 0xf5, 0xe8, 0x23, 0x1f, 0x2f, 0x09, 0x19, 0xb5, 0x50,
	0x0c, 0x9d, 0xdb, 0x77, 0xeb, 0x5f, 0x46, 0x11, 0xfe, 0x12, 0x4e, 0xc3, 0xfd, 0xef, 0x9e, 0xab,
	0x62, 0x67, 0xca, 0x49, 0xa9, 0xa6, 0xf6, 0x89, 0x86, 0xc4, 0xc6, 0x76, 0x75, 0x59, 0xd3, 0xac,
	0x2e, 0x13, 0x78, 0xd7, 0x25, 0xf8, 0x0b, 0x88, 0x57, 0x52, 0x6c, 0xcd, 0xa6, 0x34, 0x9f, 0x2b,
	0xfe, 0xc1, 0xbd, 0x70, 0x9b, 0xd8, 0x18, 0x7f, 0x03, 0xf1, 0xa2, 0x92, 0x54, 0x73, 0x51, 0x98,
	0x11, 0x2c, 0x68, 0x21, 0x14, 0x4b, 0x44, 0x91, 0x2a, 0x8f, 0xd5, 0x25, 0xfc, 0x15, 0xc0, 0xcb,
	0xe6, 0x35, 0x2f, 0x41, 0x5d, 0xe8, 0xcf, 0x0c, 0xa9, 0x59, 0x98, 0x0b, 0xaa, 0xe9, 0x23, 0xff,
	0x60, 0x17, 0xe6, 0xa6, 0xb6, 0xdc, 0x5c, 0x82, 0xbf, 0x83, 0x61, 0x20, 0xc2, 0x88, 0x6c, 0xb8,
	0x56, 0x2b, 0x26, 0x1f, 0xed, 0x59, 0x9e, 0x6e, 0x8a, 0xf8, 0x0d, 0x74, 0x57, 0x92, 0x27, 0x0c,
	0x7d, 0x0d, 0xfd, 0xb2, 0x81, 0xbe, 0x7e, 0xed, 0x17, 0xfb, 0xf6, 0xf2, 0x77, 0xbc, 0xe5, 0xfa,
	0xb9, 0xda, 0xcc, 0x13, 0x91, 0x5f, 0x1b, 0xe8, 0x5b, 0x2e, 0xae, 0x13, 0x21, 0xd9, 0xb5, 0xfd,
	0x93, 0xfa, 0xc1, 0x48, 0x9b, 0x9e, 0x8d, 0xdf, 0xfc, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x97,
	0x76, 0xb6, 0xfc, 0x06, 0x00, 0x00,
}
