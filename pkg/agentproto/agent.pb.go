// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/agentproto/agent.proto

package agentproto

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

type ReshardRequest struct {
}

func (m *ReshardRequest) Reset()      { *m = ReshardRequest{} }
func (*ReshardRequest) ProtoMessage() {}
func (*ReshardRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_11e9fe65e2a59325, []int{0}
}
func (m *ReshardRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReshardRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReshardRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReshardRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReshardRequest.Merge(m, src)
}
func (m *ReshardRequest) XXX_Size() int {
	return m.Size()
}
func (m *ReshardRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReshardRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReshardRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ReshardRequest)(nil), "agentproto.ReshardRequest")
}

func init() { proto.RegisterFile("pkg/agentproto/agent.proto", fileDescriptor_11e9fe65e2a59325) }

var fileDescriptor_11e9fe65e2a59325 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0xc8, 0x4e, 0xd7,
	0x4f, 0x4c, 0x4f, 0xcd, 0x2b, 0x29, 0x28, 0xca, 0x2f, 0xc9, 0x87, 0x30, 0xf5, 0xc0, 0x6c, 0x21,
	0x2e, 0x84, 0xb8, 0x94, 0x74, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x3e, 0x98, 0x97, 0x54, 0x9a,
	0xa6, 0x9f, 0x9a, 0x5b, 0x50, 0x52, 0x09, 0x51, 0xa8, 0x24, 0xc0, 0xc5, 0x17, 0x94, 0x5a, 0x9c,
	0x91, 0x58, 0x94, 0x12, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x62, 0x14, 0xc0, 0xc5, 0x1f, 0x9c,
	0x5c, 0x94, 0x58, 0x90, 0x99, 0x97, 0x1e, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0x64, 0xcb,
	0xc5, 0x0e, 0x55, 0x24, 0x24, 0xa5, 0x87, 0x30, 0x59, 0x0f, 0x55, 0xa7, 0x94, 0x98, 0x1e, 0xc4,
	0x26, 0x3d, 0x98, 0x4d, 0x7a, 0xae, 0x20, 0x9b, 0x9c, 0x62, 0x2f, 0x3c, 0x94, 0x63, 0xb8, 0xf1,
	0x50, 0x8e, 0xe1, 0xc3, 0x43, 0x39, 0xc6, 0x86, 0x47, 0x72, 0x8c, 0x2b, 0x1e, 0xc9, 0x31, 0x9e,
	0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x2f, 0x1e, 0xc9, 0x31,
	0x7c, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb,
	0x31, 0x44, 0xa9, 0xa7, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xa7, 0x17,
	0x25, 0xa6, 0x25, 0xe6, 0x25, 0x42, 0xbc, 0xa8, 0x8f, 0xea, 0xef, 0x24, 0x36, 0x30, 0x65, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x79, 0x37, 0xd4, 0x10, 0x01, 0x00, 0x00,
}

func (this *ReshardRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ReshardRequest)
	if !ok {
		that2, ok := that.(ReshardRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}
func (this *ReshardRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&agentproto.ReshardRequest{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringAgent(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ScrapingServiceClient is the client API for ScrapingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ScrapingServiceClient interface {
	// Reshard tells the implementing service to reshard all of its running
	// configs.
	Reshard(ctx context.Context, in *ReshardRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type scrapingServiceClient struct {
	cc *grpc.ClientConn
}

func NewScrapingServiceClient(cc *grpc.ClientConn) ScrapingServiceClient {
	return &scrapingServiceClient{cc}
}

func (c *scrapingServiceClient) Reshard(ctx context.Context, in *ReshardRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/agentproto.ScrapingService/Reshard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScrapingServiceServer is the server API for ScrapingService service.
type ScrapingServiceServer interface {
	// Reshard tells the implementing service to reshard all of its running
	// configs.
	Reshard(context.Context, *ReshardRequest) (*empty.Empty, error)
}

// UnimplementedScrapingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedScrapingServiceServer struct {
}

func (*UnimplementedScrapingServiceServer) Reshard(ctx context.Context, req *ReshardRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reshard not implemented")
}

func RegisterScrapingServiceServer(s *grpc.Server, srv ScrapingServiceServer) {
	s.RegisterService(&_ScrapingService_serviceDesc, srv)
}

func _ScrapingService_Reshard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReshardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScrapingServiceServer).Reshard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentproto.ScrapingService/Reshard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScrapingServiceServer).Reshard(ctx, req.(*ReshardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ScrapingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "agentproto.ScrapingService",
	HandlerType: (*ScrapingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Reshard",
			Handler:    _ScrapingService_Reshard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/agentproto/agent.proto",
}

func (m *ReshardRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReshardRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReshardRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintAgent(dAtA []byte, offset int, v uint64) int {
	offset -= sovAgent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ReshardRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovAgent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAgent(x uint64) (n int) {
	return sovAgent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ReshardRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ReshardRequest{`,
		`}`,
	}, "")
	return s
}
func valueToStringAgent(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ReshardRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAgent
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
			return fmt.Errorf("proto: ReshardRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReshardRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipAgent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAgent
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAgent
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
func skipAgent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAgent
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
					return 0, ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAgent
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
				return 0, ErrInvalidLengthAgent
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthAgent
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAgent
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipAgent(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthAgent
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthAgent = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAgent   = fmt.Errorf("proto: integer overflow")
)
