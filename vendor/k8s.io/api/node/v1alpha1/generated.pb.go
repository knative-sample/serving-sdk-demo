/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: k8s.io/kubernetes/vendor/k8s.io/api/node/v1alpha1/generated.proto

/*
	Package v1alpha1 is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/kubernetes/vendor/k8s.io/api/node/v1alpha1/generated.proto

	It has these top-level messages:
		RuntimeClass
		RuntimeClassList
		RuntimeClassSpec
*/
package v1alpha1

import (
	fmt "fmt"

	proto "github.com/gogo/protobuf/proto"

	math "math"

	strings "strings"

	reflect "reflect"

	io "io"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

func (m *RuntimeClass) Reset()                    { *m = RuntimeClass{} }
func (*RuntimeClass) ProtoMessage()               {}
func (*RuntimeClass) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *RuntimeClassList) Reset()                    { *m = RuntimeClassList{} }
func (*RuntimeClassList) ProtoMessage()               {}
func (*RuntimeClassList) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func (m *RuntimeClassSpec) Reset()                    { *m = RuntimeClassSpec{} }
func (*RuntimeClassSpec) ProtoMessage()               {}
func (*RuntimeClassSpec) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{2} }

func init() {
	proto.RegisterType((*RuntimeClass)(nil), "k8s.io.api.node.v1alpha1.RuntimeClass")
	proto.RegisterType((*RuntimeClassList)(nil), "k8s.io.api.node.v1alpha1.RuntimeClassList")
	proto.RegisterType((*RuntimeClassSpec)(nil), "k8s.io.api.node.v1alpha1.RuntimeClassSpec")
}
func (m *RuntimeClass) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RuntimeClass) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ObjectMeta.Size()))
	n1, err := m.ObjectMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Spec.Size()))
	n2, err := m.Spec.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	return i, nil
}

func (m *RuntimeClassList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RuntimeClassList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ListMeta.Size()))
	n3, err := m.ListMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *RuntimeClassSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RuntimeClassSpec) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.RuntimeHandler)))
	i += copy(dAtA[i:], m.RuntimeHandler)
	return i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RuntimeClass) Size() (n int) {
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *RuntimeClassList) Size() (n int) {
	var l int
	_ = l
	l = m.ListMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *RuntimeClassSpec) Size() (n int) {
	var l int
	_ = l
	l = len(m.RuntimeHandler)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *RuntimeClass) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RuntimeClass{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(this.ObjectMeta.String(), "ObjectMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Spec:` + strings.Replace(strings.Replace(this.Spec.String(), "RuntimeClassSpec", "RuntimeClassSpec", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *RuntimeClassList) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RuntimeClassList{`,
		`ListMeta:` + strings.Replace(strings.Replace(this.ListMeta.String(), "ListMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ListMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Items), "RuntimeClass", "RuntimeClass", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *RuntimeClassSpec) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RuntimeClassSpec{`,
		`RuntimeHandler:` + fmt.Sprintf("%v", this.RuntimeHandler) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *RuntimeClass) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RuntimeClass: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RuntimeClass: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *RuntimeClassList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RuntimeClassList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RuntimeClassList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ListMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, RuntimeClass{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *RuntimeClassSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RuntimeClassSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RuntimeClassSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RuntimeHandler", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RuntimeHandler = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
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
				next, err := skipGenerated(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("k8s.io/kubernetes/vendor/k8s.io/api/node/v1alpha1/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x6b, 0xd4, 0x40,
	0x14, 0xc7, 0x33, 0xb5, 0x85, 0x75, 0x5a, 0x4b, 0xc9, 0x41, 0xc2, 0x1e, 0xa6, 0x65, 0x0f, 0x52,
	0x04, 0x67, 0xdc, 0x22, 0xe2, 0x49, 0x30, 0x5e, 0x14, 0x2b, 0x42, 0xbc, 0x89, 0x07, 0x27, 0xc9,
	0x33, 0x19, 0xb3, 0xc9, 0x0c, 0x99, 0x49, 0xc0, 0x9b, 0x1f, 0xc1, 0x2f, 0xa4, 0xe7, 0x3d, 0xf6,
	0xd8, 0x53, 0x71, 0xe3, 0x17, 0x91, 0x99, 0x64, 0xbb, 0xdb, 0x2e, 0xc5, 0xbd, 0xe5, 0xbd, 0xf9,
	0xff, 0x7f, 0xef, 0xfd, 0x5f, 0xf0, 0xab, 0xe2, 0x85, 0xa6, 0x42, 0xb2, 0xa2, 0x89, 0xa1, 0xae,
	0xc0, 0x80, 0x66, 0x2d, 0x54, 0xa9, 0xac, 0xd9, 0xf0, 0xc0, 0x95, 0x60, 0x95, 0x4c, 0x81, 0xb5,
	0x53, 0x3e, 0x53, 0x39, 0x9f, 0xb2, 0x0c, 0x2a, 0xa8, 0xb9, 0x81, 0x94, 0xaa, 0x5a, 0x1a, 0xe9,
	0x07, 0xbd, 0x92, 0x72, 0x25, 0xa8, 0x55, 0xd2, 0xa5, 0x72, 0xfc, 0x24, 0x13, 0x26, 0x6f, 0x62,
	0x9a, 0xc8, 0x92, 0x65, 0x32, 0x93, 0xcc, 0x19, 0xe2, 0xe6, 0xab, 0xab, 0x5c, 0xe1, 0xbe, 0x7a,
	0xd0, 0xf8, 0xd9, 0x6a, 0x64, 0xc9, 0x93, 0x5c, 0x54, 0x50, 0x7f, 0x67, 0xaa, 0xc8, 0x6c, 0x43,
	0xb3, 0x12, 0x0c, 0x67, 0xed, 0xc6, 0xf8, 0x31, 0xbb, 0xcb, 0x55, 0x37, 0x95, 0x11, 0x25, 0x6c,
	0x18, 0x9e, 0xff, 0xcf, 0xa0, 0x93, 0x1c, 0x4a, 0x7e, 0xdb, 0x37, 0xf9, 0x8d, 0xf0, 0x41, 0xd4,
	0x4b, 0x5e, 0xcf, 0xb8, 0xd6, 0xfe, 0x17, 0x3c, 0xb2, 0x4b, 0xa5, 0xdc, 0xf0, 0x00, 0x9d, 0xa0,
	0xd3, 0xfd, 0xb3, 0xa7, 0x74, 0x75, 0x8b, 0x6b, 0x36, 0x55, 0x45, 0x66, 0x1b, 0x9a, 0x5a, 0x35,
	0x6d, 0xa7, 0xf4, 0x43, 0xfc, 0x0d, 0x12, 0xf3, 0x1e, 0x0c, 0x0f, 0xfd, 0xf9, 0xd5, 0xb1, 0xd7,
	0x5d, 0x1d, 0xe3, 0x55, 0x2f, 0xba, 0xa6, 0xfa, 0xe7, 0x78, 0x57, 0x2b, 0x48, 0x82, 0x1d, 0x47,
	0x7f, 0x4c, 0xef, 0xba, 0x34, 0x5d, 0xdf, 0xeb, 0xa3, 0x82, 0x24, 0x3c, 0x18, 0xb8, 0xbb, 0xb6,
	0x8a, 0x1c, 0x65, 0xf2, 0x0b, 0xe1, 0xa3, 0x75, 0xe1, 0xb9, 0xd0, 0xc6, 0xff, 0xbc, 0x11, 0x82,
	0x6e, 0x17, 0xc2, 0xba, 0x5d, 0x84, 0xa3, 0x61, 0xd4, 0x68, 0xd9, 0x59, 0x0b, 0xf0, 0x0e, 0xef,
	0x09, 0x03, 0xa5, 0x0e, 0x76, 0x4e, 0xee, 0x9d, 0xee, 0x9f, 0x3d, 0xda, 0x2e, 0x41, 0xf8, 0x60,
	0x40, 0xee, 0xbd, 0xb5, 0xe6, 0xa8, 0x67, 0x4c, 0xa2, 0x9b, 0xeb, 0xdb, 0x64, 0xfe, 0x4b, 0x7c,
	0x38, 0xfc, 0xb6, 0x37, 0xbc, 0x4a, 0x67, 0x50, 0xbb, 0x10, 0xf7, 0xc3, 0x87, 0x03, 0xe1, 0x30,
	0xba, 0xf1, 0x1a, 0xdd, 0x52, 0x87, 0x74, 0xbe, 0x20, 0xde, 0xc5, 0x82, 0x78, 0x97, 0x0b, 0xe2,
	0xfd, 0xe8, 0x08, 0x9a, 0x77, 0x04, 0x5d, 0x74, 0x04, 0x5d, 0x76, 0x04, 0xfd, 0xe9, 0x08, 0xfa,
	0xf9, 0x97, 0x78, 0x9f, 0x46, 0xcb, 0x35, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x94, 0x34, 0x0e,
	0xef, 0x30, 0x03, 0x00, 0x00,
}
