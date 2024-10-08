// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tabi/captains/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the module's genesis state
type GenesisState struct {
	// params stores all the parameters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// base_state stores the current epoch state.
	BaseState BaseState `protobuf:"bytes,2,opt,name=base_state,json=baseState,proto3" json:"base_state"`
	// divisions stores all divisions base info.
	Divisions []Division `protobuf:"bytes,3,rep,name=divisions,proto3" json:"divisions"`
	// nodes stores all nodes base info.
	Nodes []Node `protobuf:"bytes,4,rep,name=nodes,proto3" json:"nodes"`
	// epoches_emission
	EpochesEmission []EpochEmission `protobuf:"bytes,5,rep,name=epoches_emission,json=epochesEmission,proto3" json:"epoches_emission"`
	// nodes_claimed_emission
	NodesClaimedEmission []NodeClaimedEmission `protobuf:"bytes,6,rep,name=nodes_claimed_emission,json=nodesClaimedEmission,proto3" json:"nodes_claimed_emission"`
	// nodes_cumulative_emission
	NodesCumulativeEmission []NodeCumulativeEmission `protobuf:"bytes,7,rep,name=nodes_cumulative_emission,json=nodesCumulativeEmission,proto3" json:"nodes_cumulative_emission"`
	// globals_pledge
	GlobalsPledge []GlobalPledge `protobuf:"bytes,8,rep,name=globals_pledge,json=globalsPledge,proto3" json:"globals_pledge"`
	// owners_pledge
	OwnersPledge []OwnerPledge `protobuf:"bytes,9,rep,name=owners_pledge,json=ownersPledge,proto3" json:"owners_pledge"`
	// owners_claimable_computing_power computing powers infos.
	OwnersClaimableComputingPower []ClaimableComputingPower `protobuf:"bytes,10,rep,name=owners_claimable_computing_power,json=ownersClaimableComputingPower,proto3" json:"owners_claimable_computing_power"`
	// globals_computing_power
	GlobalsComputingPower []GlobalComputingPower `protobuf:"bytes,11,rep,name=globals_computing_power,json=globalsComputingPower,proto3" json:"globals_computing_power"`
	// nodes_computing_power
	NodesComputingPower []NodesComputingPower `protobuf:"bytes,12,rep,name=nodes_computing_power,json=nodesComputingPower,proto3" json:"nodes_computing_power"`
	// batches
	Batches []BatchBase `protobuf:"bytes,13,rep,name=batches,proto3" json:"batches"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b875c06e10d2c08, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetBaseState() BaseState {
	if m != nil {
		return m.BaseState
	}
	return BaseState{}
}

func (m *GenesisState) GetDivisions() []Division {
	if m != nil {
		return m.Divisions
	}
	return nil
}

func (m *GenesisState) GetNodes() []Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *GenesisState) GetEpochesEmission() []EpochEmission {
	if m != nil {
		return m.EpochesEmission
	}
	return nil
}

func (m *GenesisState) GetNodesClaimedEmission() []NodeClaimedEmission {
	if m != nil {
		return m.NodesClaimedEmission
	}
	return nil
}

func (m *GenesisState) GetNodesCumulativeEmission() []NodeCumulativeEmission {
	if m != nil {
		return m.NodesCumulativeEmission
	}
	return nil
}

func (m *GenesisState) GetGlobalsPledge() []GlobalPledge {
	if m != nil {
		return m.GlobalsPledge
	}
	return nil
}

func (m *GenesisState) GetOwnersPledge() []OwnerPledge {
	if m != nil {
		return m.OwnersPledge
	}
	return nil
}

func (m *GenesisState) GetOwnersClaimableComputingPower() []ClaimableComputingPower {
	if m != nil {
		return m.OwnersClaimableComputingPower
	}
	return nil
}

func (m *GenesisState) GetGlobalsComputingPower() []GlobalComputingPower {
	if m != nil {
		return m.GlobalsComputingPower
	}
	return nil
}

func (m *GenesisState) GetNodesComputingPower() []NodesComputingPower {
	if m != nil {
		return m.NodesComputingPower
	}
	return nil
}

func (m *GenesisState) GetBatches() []BatchBase {
	if m != nil {
		return m.Batches
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "tabi.captains.v1.GenesisState")
}

func init() { proto.RegisterFile("tabi/captains/v1/genesis.proto", fileDescriptor_6b875c06e10d2c08) }

var fileDescriptor_6b875c06e10d2c08 = []byte{
	// 541 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0xcf, 0x6e, 0xd3, 0x30,
	0x1c, 0xc7, 0x5b, 0xb6, 0x75, 0xab, 0xdb, 0xc2, 0x64, 0xf6, 0x27, 0x14, 0x35, 0xad, 0x90, 0x40,
	0xdd, 0x25, 0xd1, 0x8a, 0xc4, 0x05, 0x09, 0xa1, 0x96, 0x69, 0x48, 0x48, 0x50, 0xc1, 0x8d, 0x4b,
	0xe4, 0x24, 0x56, 0x6a, 0x94, 0xc4, 0x51, 0xec, 0x76, 0xe3, 0x2d, 0x78, 0x16, 0x9e, 0x62, 0xc7,
	0x1d, 0x39, 0x21, 0xd4, 0xbe, 0x08, 0xf2, 0x2f, 0x4e, 0x2b, 0xe2, 0x96, 0x5b, 0xe4, 0xdf, 0xf7,
	0xfb, 0xf9, 0x44, 0x4e, 0x6c, 0x64, 0x4b, 0xe2, 0x33, 0x37, 0x20, 0x99, 0x24, 0x2c, 0x15, 0xee,
	0xe2, 0xd2, 0x8d, 0x68, 0x4a, 0x05, 0x13, 0x4e, 0x96, 0x73, 0xc9, 0xf1, 0xb1, 0x9a, 0x3b, 0xe5,
	0xdc, 0x59, 0x5c, 0x76, 0x4f, 0x22, 0x1e, 0x71, 0x18, 0xba, 0xea, 0xa9, 0xc8, 0x75, 0x7b, 0x06,
	0x27, 0xa7, 0x19, 0xcf, 0xa5, 0x1e, 0xf7, 0x8d, 0xf1, 0x1a, 0x09, 0x81, 0x67, 0x3f, 0x8f, 0x50,
	0xfb, 0xba, 0x30, 0x7f, 0x91, 0x44, 0x52, 0xfc, 0x0a, 0x35, 0x32, 0x92, 0x93, 0x44, 0x58, 0xf5,
	0x41, 0x7d, 0xd8, 0x1a, 0x59, 0x4e, 0xf5, 0x4d, 0x9c, 0x29, 0xcc, 0xc7, 0xfb, 0x77, 0xbf, 0xfb,
	0xb5, 0xcf, 0x3a, 0x8d, 0xdf, 0x22, 0xe4, 0x13, 0x41, 0x3d, 0xa1, 0x28, 0xd6, 0x03, 0xe8, 0x3e,
	0x35, 0xbb, 0x63, 0x22, 0x28, 0x88, 0x74, 0xbd, 0xe9, 0x97, 0x0b, 0xf8, 0x0d, 0x6a, 0x86, 0x6c,
	0xc1, 0x04, 0xe3, 0xa9, 0xb0, 0xf6, 0x06, 0x7b, 0xc3, 0xd6, 0xa8, 0x6b, 0x02, 0xde, 0xe9, 0x48,
	0xd9, 0x5f, 0x57, 0xf0, 0x08, 0x1d, 0xa4, 0x3c, 0xa4, 0xc2, 0xda, 0x87, 0xee, 0x99, 0xd9, 0xfd,
	0xc8, 0xc3, 0xd2, 0x5b, 0x44, 0xf1, 0x14, 0x1d, 0xd3, 0x8c, 0x07, 0x33, 0x2a, 0x3c, 0x9a, 0x30,
	0xa1, 0x40, 0xd6, 0x01, 0xd4, 0xfb, 0x66, 0xfd, 0x4a, 0x25, 0xaf, 0x74, 0x4c, 0x73, 0x1e, 0xe9,
	0x7a, 0xb9, 0x8c, 0x09, 0x3a, 0x03, 0xb4, 0x17, 0xc4, 0x84, 0x25, 0x34, 0xdc, 0x70, 0x1b, 0xc0,
	0x7d, 0xbe, 0xfd, 0xb5, 0x26, 0x45, 0xba, 0x42, 0x3f, 0x01, 0x54, 0x65, 0x86, 0xbf, 0xa1, 0x27,
	0x5a, 0x31, 0x4f, 0xe6, 0x31, 0x91, 0x6c, 0x41, 0x37, 0x96, 0x43, 0xb0, 0x0c, 0x77, 0x58, 0xd6,
	0x85, 0x8a, 0xe8, 0xbc, 0x10, 0x19, 0x63, 0xfc, 0x01, 0x3d, 0x8c, 0x62, 0xee, 0x93, 0x58, 0x78,
	0x59, 0x4c, 0xc3, 0x88, 0x5a, 0x47, 0x20, 0xb0, 0x4d, 0xc1, 0x35, 0xe4, 0xa6, 0x90, 0xd2, 0xd8,
	0x8e, 0xee, 0x16, 0x8b, 0xf8, 0x3d, 0xea, 0xf0, 0x9b, 0x94, 0xe6, 0x6b, 0x56, 0x13, 0x58, 0x3d,
	0x93, 0xf5, 0x49, 0xc5, 0xfe, 0x41, 0xb5, 0x8b, 0xa6, 0x26, 0xdd, 0xa2, 0x81, 0x26, 0xc1, 0x36,
	0x13, 0x3f, 0xa6, 0x5e, 0xc0, 0x93, 0x6c, 0x2e, 0x59, 0x1a, 0x79, 0x19, 0xbf, 0xa1, 0xb9, 0x85,
	0x00, 0x7e, 0x61, 0xc2, 0x27, 0x65, 0x65, 0x52, 0x36, 0xa6, 0xaa, 0xa0, 0x45, 0xbd, 0x02, 0xbc,
	0x23, 0x84, 0x43, 0x74, 0x5e, 0x6e, 0x48, 0x55, 0xd8, 0x02, 0xe1, 0x8b, 0x5d, 0x3b, 0xb3, 0xd5,
	0x76, 0xaa, 0x61, 0x15, 0x8b, 0x87, 0x4e, 0xf5, 0x27, 0xae, 0x38, 0xda, 0xff, 0xfb, 0x89, 0xc4,
	0x56, 0xc5, 0xe3, 0xd4, 0x1c, 0xe1, 0xd7, 0xe8, 0xd0, 0x27, 0x52, 0xfd, 0xb9, 0x56, 0x07, 0x90,
	0x5b, 0xcf, 0xaa, 0x0c, 0x66, 0xea, 0xc0, 0x6a, 0x50, 0xd9, 0x18, 0x4f, 0xee, 0x96, 0x76, 0xfd,
	0x7e, 0x69, 0xd7, 0xff, 0x2c, 0xed, 0xfa, 0x8f, 0x95, 0x5d, 0xbb, 0x5f, 0xd9, 0xb5, 0x5f, 0x2b,
	0xbb, 0xf6, 0xf5, 0x22, 0x62, 0x72, 0x36, 0xf7, 0x9d, 0x80, 0x27, 0xae, 0xe2, 0xc5, 0xc4, 0x17,
	0xf0, 0xe0, 0xde, 0x6e, 0x6e, 0x21, 0xf9, 0x3d, 0xa3, 0xc2, 0x6f, 0xc0, 0x05, 0xf4, 0xf2, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x55, 0xaa, 0x80, 0xdb, 0x0a, 0x05, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Batches) > 0 {
		for iNdEx := len(m.Batches) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Batches[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x6a
		}
	}
	if len(m.NodesComputingPower) > 0 {
		for iNdEx := len(m.NodesComputingPower) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NodesComputingPower[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x62
		}
	}
	if len(m.GlobalsComputingPower) > 0 {
		for iNdEx := len(m.GlobalsComputingPower) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GlobalsComputingPower[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x5a
		}
	}
	if len(m.OwnersClaimableComputingPower) > 0 {
		for iNdEx := len(m.OwnersClaimableComputingPower) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OwnersClaimableComputingPower[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.OwnersPledge) > 0 {
		for iNdEx := len(m.OwnersPledge) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OwnersPledge[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.GlobalsPledge) > 0 {
		for iNdEx := len(m.GlobalsPledge) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GlobalsPledge[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.NodesCumulativeEmission) > 0 {
		for iNdEx := len(m.NodesCumulativeEmission) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NodesCumulativeEmission[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.NodesClaimedEmission) > 0 {
		for iNdEx := len(m.NodesClaimedEmission) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NodesClaimedEmission[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.EpochesEmission) > 0 {
		for iNdEx := len(m.EpochesEmission) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.EpochesEmission[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Nodes) > 0 {
		for iNdEx := len(m.Nodes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Nodes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Divisions) > 0 {
		for iNdEx := len(m.Divisions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Divisions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size, err := m.BaseState.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.BaseState.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Divisions) > 0 {
		for _, e := range m.Divisions {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Nodes) > 0 {
		for _, e := range m.Nodes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.EpochesEmission) > 0 {
		for _, e := range m.EpochesEmission {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.NodesClaimedEmission) > 0 {
		for _, e := range m.NodesClaimedEmission {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.NodesCumulativeEmission) > 0 {
		for _, e := range m.NodesCumulativeEmission {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.GlobalsPledge) > 0 {
		for _, e := range m.GlobalsPledge {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.OwnersPledge) > 0 {
		for _, e := range m.OwnersPledge {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.OwnersClaimableComputingPower) > 0 {
		for _, e := range m.OwnersClaimableComputingPower {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.GlobalsComputingPower) > 0 {
		for _, e := range m.GlobalsComputingPower {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.NodesComputingPower) > 0 {
		for _, e := range m.NodesComputingPower {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Batches) > 0 {
		for _, e := range m.Batches {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BaseState.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Divisions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Divisions = append(m.Divisions, Division{})
			if err := m.Divisions[len(m.Divisions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nodes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nodes = append(m.Nodes, Node{})
			if err := m.Nodes[len(m.Nodes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochesEmission", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EpochesEmission = append(m.EpochesEmission, EpochEmission{})
			if err := m.EpochesEmission[len(m.EpochesEmission)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodesClaimedEmission", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodesClaimedEmission = append(m.NodesClaimedEmission, NodeClaimedEmission{})
			if err := m.NodesClaimedEmission[len(m.NodesClaimedEmission)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodesCumulativeEmission", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodesCumulativeEmission = append(m.NodesCumulativeEmission, NodeCumulativeEmission{})
			if err := m.NodesCumulativeEmission[len(m.NodesCumulativeEmission)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GlobalsPledge", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GlobalsPledge = append(m.GlobalsPledge, GlobalPledge{})
			if err := m.GlobalsPledge[len(m.GlobalsPledge)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnersPledge", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OwnersPledge = append(m.OwnersPledge, OwnerPledge{})
			if err := m.OwnersPledge[len(m.OwnersPledge)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnersClaimableComputingPower", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OwnersClaimableComputingPower = append(m.OwnersClaimableComputingPower, ClaimableComputingPower{})
			if err := m.OwnersClaimableComputingPower[len(m.OwnersClaimableComputingPower)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GlobalsComputingPower", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GlobalsComputingPower = append(m.GlobalsComputingPower, GlobalComputingPower{})
			if err := m.GlobalsComputingPower[len(m.GlobalsComputingPower)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodesComputingPower", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodesComputingPower = append(m.NodesComputingPower, NodesComputingPower{})
			if err := m.NodesComputingPower[len(m.NodesComputingPower)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Batches", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Batches = append(m.Batches, BatchBase{})
			if err := m.Batches[len(m.Batches)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
