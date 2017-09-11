// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: redis_meta.proto

/*
	Package gxredis is a generated protocol buffer package.

	It is generated from these files:
		redis_meta.proto

	It has these top-level messages:
		IPAddr
		Slave
		Instance
*/
package gxredis

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strconv "strconv"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type RedisRole int32

const (
	RR_Default  RedisRole = 0
	RR_Master   RedisRole = 1
	RR_Slave    RedisRole = 2
	RR_Sentinel RedisRole = 3
)

var RedisRole_name = map[int32]string{
	0: "RR_Default",
	1: "RR_Master",
	2: "RR_Slave",
	3: "RR_Sentinel",
}
var RedisRole_value = map[string]int32{
	"RR_Default":  0,
	"RR_Master":   1,
	"RR_Slave":    2,
	"RR_Sentinel": 3,
}

func (RedisRole) EnumDescriptor() ([]byte, []int) { return fileDescriptorRedisMeta, []int{0} }

func getRedisRole(role string) RedisRole {
	switch role {
	case "master":
		return RR_Master
	case "slave":
		return RR_Slave
	case "sentinel":
		return RR_Sentinel
	}

	return RR_Default
}

// TCPAddr represents the address of a TCP end point.
type IPAddr struct {
	IP   string `protobuf:"bytes,1,opt,name=IP,proto3" json:"IP,omitempty"`
	Port uint32 `protobuf:"varint,2,opt,name=Port,proto3" json:"Port,omitempty"`
}

func (m *IPAddr) Reset()                    { *m = IPAddr{} }
func (*IPAddr) ProtoMessage()               {}
func (*IPAddr) Descriptor() ([]byte, []int) { return fileDescriptorRedisMeta, []int{0} }

// Slave represents a Redis slave instance which is known by Sentinel.
type Slave struct {
	Addr  *IPAddr `protobuf:"bytes,1,opt,name=Addr" json:"Addr,omitempty"`
	Flags string  `protobuf:"bytes,2,opt,name=Flags,proto3" json:"Flags,omitempty"`
}

func (m *Slave) Reset()                    { *m = Slave{} }
func (*Slave) ProtoMessage()               {}
func (*Slave) Descriptor() ([]byte, []int) { return fileDescriptorRedisMeta, []int{1} }

// Addr returns an address of slave.
func (s *Slave) Address() string {
	return s.String()
}

// Available returns if slave is in working state at moment based on information in slave flags.
func (s *Slave) Available() bool {
	return !strings.Contains(s.Flags, "disconnected") && !strings.Contains(s.Flags, "s_down")
}

type Instance struct {
	Name   string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Master *IPAddr  `protobuf:"bytes,2,opt,name=Master" json:"Master,omitempty"`
	Slaves []*Slave `protobuf:"bytes,3,rep,name=Slaves" json:"Slaves,omitempty"`
}

func (m *Instance) Reset()                    { *m = Instance{} }
func (*Instance) ProtoMessage()               {}
func (*Instance) Descriptor() ([]byte, []int) { return fileDescriptorRedisMeta, []int{2} }

func init() {
	proto.RegisterType((*IPAddr)(nil), "gxredis.IPAddr")
	proto.RegisterType((*Slave)(nil), "gxredis.Slave")
	proto.RegisterType((*Instance)(nil), "gxredis.Instance")
	proto.RegisterEnum("gxredis.RedisRole", RedisRole_name, RedisRole_value)
}
func (x RedisRole) String() string {
	s, ok := RedisRole_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *IPAddr) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*IPAddr)
	if !ok {
		that2, ok := that.(IPAddr)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *IPAddr")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *IPAddr but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *IPAddr but is not nil && this == nil")
	}
	if this.IP != that1.IP {
		return fmt.Errorf("IP this(%v) Not Equal that(%v)", this.IP, that1.IP)
	}
	if this.Port != that1.Port {
		return fmt.Errorf("Port this(%v) Not Equal that(%v)", this.Port, that1.Port)
	}
	return nil
}
func (this *IPAddr) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*IPAddr)
	if !ok {
		that2, ok := that.(IPAddr)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.IP != that1.IP {
		return false
	}
	if this.Port != that1.Port {
		return false
	}
	return true
}
func (this *Slave) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Slave)
	if !ok {
		that2, ok := that.(Slave)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *Slave")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *Slave but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *Slave but is not nil && this == nil")
	}
	if !this.Addr.Equal(that1.Addr) {
		return fmt.Errorf("Addr this(%v) Not Equal that(%v)", this.Addr, that1.Addr)
	}
	if this.Flags != that1.Flags {
		return fmt.Errorf("Flags this(%v) Not Equal that(%v)", this.Flags, that1.Flags)
	}
	return nil
}
func (this *Slave) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Slave)
	if !ok {
		that2, ok := that.(Slave)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Addr.Equal(that1.Addr) {
		return false
	}
	if this.Flags != that1.Flags {
		return false
	}
	return true
}
func (this *Instance) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Instance)
	if !ok {
		that2, ok := that.(Instance)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *Instance")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *Instance but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *Instance but is not nil && this == nil")
	}
	if this.Name != that1.Name {
		return fmt.Errorf("Name this(%v) Not Equal that(%v)", this.Name, that1.Name)
	}
	if !this.Master.Equal(that1.Master) {
		return fmt.Errorf("Master this(%v) Not Equal that(%v)", this.Master, that1.Master)
	}
	if len(this.Slaves) != len(that1.Slaves) {
		return fmt.Errorf("Slaves this(%v) Not Equal that(%v)", len(this.Slaves), len(that1.Slaves))
	}
	for i := range this.Slaves {
		if !this.Slaves[i].Equal(that1.Slaves[i]) {
			return fmt.Errorf("Slaves this[%v](%v) Not Equal that[%v](%v)", i, this.Slaves[i], i, that1.Slaves[i])
		}
	}
	return nil
}
func (this *Instance) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Instance)
	if !ok {
		that2, ok := that.(Instance)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if !this.Master.Equal(that1.Master) {
		return false
	}
	if len(this.Slaves) != len(that1.Slaves) {
		return false
	}
	for i := range this.Slaves {
		if !this.Slaves[i].Equal(that1.Slaves[i]) {
			return false
		}
	}
	return true
}
func (this *IPAddr) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&gxredis.IPAddr{")
	s = append(s, "IP: "+fmt.Sprintf("%#v", this.IP)+",\n")
	s = append(s, "Port: "+fmt.Sprintf("%#v", this.Port)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Slave) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&gxredis.Slave{")
	if this.Addr != nil {
		s = append(s, "Addr: "+fmt.Sprintf("%#v", this.Addr)+",\n")
	}
	s = append(s, "Flags: "+fmt.Sprintf("%#v", this.Flags)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Instance) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&gxredis.Instance{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	if this.Master != nil {
		s = append(s, "Master: "+fmt.Sprintf("%#v", this.Master)+",\n")
	}
	if this.Slaves != nil {
		s = append(s, "Slaves: "+fmt.Sprintf("%#v", this.Slaves)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringRedisMeta(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *IPAddr) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IPAddr) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.IP) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRedisMeta(dAtA, i, uint64(len(m.IP)))
		i += copy(dAtA[i:], m.IP)
	}
	if m.Port != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintRedisMeta(dAtA, i, uint64(m.Port))
	}
	return i, nil
}

func (m *Slave) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Slave) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Addr != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRedisMeta(dAtA, i, uint64(m.Addr.Size()))
		n1, err := m.Addr.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Flags) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRedisMeta(dAtA, i, uint64(len(m.Flags)))
		i += copy(dAtA[i:], m.Flags)
	}
	return i, nil
}

func (m *Instance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Instance) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRedisMeta(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Master != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRedisMeta(dAtA, i, uint64(m.Master.Size()))
		n2, err := m.Master.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.Slaves) > 0 {
		for _, msg := range m.Slaves {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintRedisMeta(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeFixed64RedisMeta(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32RedisMeta(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintRedisMeta(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *IPAddr) Size() (n int) {
	var l int
	_ = l
	l = len(m.IP)
	if l > 0 {
		n += 1 + l + sovRedisMeta(uint64(l))
	}
	if m.Port != 0 {
		n += 1 + sovRedisMeta(uint64(m.Port))
	}
	return n
}

func (m *Slave) Size() (n int) {
	var l int
	_ = l
	if m.Addr != nil {
		l = m.Addr.Size()
		n += 1 + l + sovRedisMeta(uint64(l))
	}
	l = len(m.Flags)
	if l > 0 {
		n += 1 + l + sovRedisMeta(uint64(l))
	}
	return n
}

func (m *Instance) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRedisMeta(uint64(l))
	}
	if m.Master != nil {
		l = m.Master.Size()
		n += 1 + l + sovRedisMeta(uint64(l))
	}
	if len(m.Slaves) > 0 {
		for _, e := range m.Slaves {
			l = e.Size()
			n += 1 + l + sovRedisMeta(uint64(l))
		}
	}
	return n
}

func sovRedisMeta(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRedisMeta(x uint64) (n int) {
	return sovRedisMeta(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *IPAddr) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&IPAddr{`,
		`IP:` + fmt.Sprintf("%v", this.IP) + `,`,
		`Port:` + fmt.Sprintf("%v", this.Port) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Slave) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Slave{`,
		`Addr:` + strings.Replace(fmt.Sprintf("%v", this.Addr), "IPAddr", "IPAddr", 1) + `,`,
		`Flags:` + fmt.Sprintf("%v", this.Flags) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Instance) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Instance{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Master:` + strings.Replace(fmt.Sprintf("%v", this.Master), "IPAddr", "IPAddr", 1) + `,`,
		`Slaves:` + strings.Replace(fmt.Sprintf("%v", this.Slaves), "Slave", "Slave", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringRedisMeta(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *IPAddr) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRedisMeta
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
			return fmt.Errorf("proto: IPAddr: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IPAddr: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IP", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRedisMeta
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
				return ErrInvalidLengthRedisMeta
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IP = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			m.Port = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRedisMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Port |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRedisMeta(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRedisMeta
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
func (m *Slave) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRedisMeta
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
			return fmt.Errorf("proto: Slave: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Slave: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addr", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRedisMeta
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
				return ErrInvalidLengthRedisMeta
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Addr == nil {
				m.Addr = &IPAddr{}
			}
			if err := m.Addr.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Flags", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRedisMeta
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
				return ErrInvalidLengthRedisMeta
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Flags = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRedisMeta(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRedisMeta
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
func (m *Instance) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRedisMeta
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
			return fmt.Errorf("proto: Instance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Instance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRedisMeta
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
				return ErrInvalidLengthRedisMeta
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Master", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRedisMeta
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
				return ErrInvalidLengthRedisMeta
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Master == nil {
				m.Master = &IPAddr{}
			}
			if err := m.Master.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Slaves", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRedisMeta
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
				return ErrInvalidLengthRedisMeta
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Slaves = append(m.Slaves, &Slave{})
			if err := m.Slaves[len(m.Slaves)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRedisMeta(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRedisMeta
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
func skipRedisMeta(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRedisMeta
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
					return 0, ErrIntOverflowRedisMeta
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
					return 0, ErrIntOverflowRedisMeta
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
				return 0, ErrInvalidLengthRedisMeta
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRedisMeta
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
				next, err := skipRedisMeta(dAtA[start:])
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
	ErrInvalidLengthRedisMeta = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRedisMeta   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("redis_meta.proto", fileDescriptorRedisMeta) }

var fileDescriptorRedisMeta = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0xcd, 0x4e, 0xea, 0x40,
	0x18, 0xed, 0x14, 0xe8, 0xa5, 0x1f, 0x17, 0x68, 0x26, 0x77, 0x41, 0xee, 0x62, 0x42, 0xb8, 0xc9,
	0x95, 0x18, 0x2d, 0x09, 0xfa, 0x02, 0x12, 0x63, 0xd2, 0x85, 0xa6, 0x19, 0x1f, 0xa0, 0x19, 0x60,
	0xa8, 0x24, 0xa5, 0x63, 0xda, 0xc1, 0xb8, 0xf4, 0x11, 0x7c, 0x0c, 0x1f, 0x85, 0x25, 0x4b, 0x97,
	0x74, 0xdc, 0xb8, 0xe4, 0x11, 0x4c, 0xbf, 0x36, 0xae, 0xdc, 0x9d, 0x9f, 0xef, 0x9c, 0x33, 0x19,
	0xf0, 0x32, 0xb9, 0x5c, 0xe7, 0xd1, 0x46, 0x6a, 0xe1, 0x3f, 0x66, 0x4a, 0x2b, 0xfa, 0x2b, 0x7e,
	0x46, 0xed, 0xef, 0x79, 0xbc, 0xd6, 0x0f, 0xdb, 0xb9, 0xbf, 0x50, 0x9b, 0x49, 0xac, 0x62, 0x35,
	0x41, 0x7f, 0xbe, 0x5d, 0x21, 0x43, 0x82, 0xa8, 0xca, 0x8d, 0xce, 0xc0, 0x09, 0xc2, 0xab, 0xe5,
	0x32, 0xa3, 0x3d, 0xb0, 0x83, 0x70, 0x40, 0x86, 0x64, 0xec, 0x72, 0x3b, 0x08, 0x29, 0x85, 0x66,
	0xa8, 0x32, 0x3d, 0xb0, 0x87, 0x64, 0xdc, 0xe5, 0x88, 0x47, 0x33, 0x68, 0xdd, 0x27, 0xe2, 0x49,
	0xd2, 0x7f, 0xd0, 0x2c, 0x43, 0x78, 0xde, 0x99, 0xf6, 0xfd, 0x7a, 0xdd, 0xaf, 0xba, 0x38, 0x9a,
	0xf4, 0x0f, 0xb4, 0x6e, 0x12, 0x11, 0xe7, 0x58, 0xe1, 0xf2, 0x8a, 0x8c, 0x14, 0xb4, 0x83, 0x34,
	0xd7, 0x22, 0x5d, 0xc8, 0x72, 0xe3, 0x4e, 0x6c, 0x64, 0xbd, 0x8a, 0x98, 0x9e, 0x80, 0x73, 0x2b,
	0x72, 0x2d, 0x33, 0x8c, 0xfd, 0x50, 0x5e, 0xdb, 0xf4, 0x3f, 0x38, 0xf8, 0x98, 0x7c, 0xd0, 0x18,
	0x36, 0xc6, 0x9d, 0x69, 0xef, 0xfb, 0x10, 0x65, 0x5e, 0xbb, 0xa7, 0x01, 0xb8, 0xbc, 0x94, 0xb9,
	0x4a, 0x24, 0xed, 0x01, 0x70, 0x1e, 0x5d, 0xcb, 0x95, 0xd8, 0x26, 0xda, 0xb3, 0x68, 0x17, 0x5c,
	0xce, 0xa3, 0xaa, 0xd1, 0x23, 0xf4, 0x37, 0xb4, 0x39, 0x8f, 0x30, 0xe8, 0xd9, 0xb4, 0x0f, 0x9d,
	0x92, 0xc9, 0x54, 0xaf, 0x53, 0x99, 0x78, 0x8d, 0xd9, 0xe5, 0xae, 0x60, 0xd6, 0xbe, 0x60, 0xd6,
	0x7b, 0xc1, 0xac, 0x43, 0xc1, 0xc8, 0xb1, 0x60, 0xe4, 0xc5, 0x30, 0xf2, 0x66, 0x18, 0xd9, 0x19,
	0x46, 0xf6, 0x86, 0x91, 0x83, 0x61, 0xe4, 0xd3, 0x30, 0xeb, 0x68, 0x18, 0x79, 0xfd, 0x60, 0xd6,
	0xdc, 0xc1, 0xaf, 0xbe, 0xf8, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x22, 0x54, 0xf5, 0xb6, 0x01,
	0x00, 0x00,
}
