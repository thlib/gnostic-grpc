// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bookstore.proto

package bookstore

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/protoc-gen-go/descriptor"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Book struct {
	Author               string   `protobuf:"bytes,1,opt,name=author,proto3" json:"author,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Book) Reset()         { *m = Book{} }
func (m *Book) String() string { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()    {}
func (*Book) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f82f486e563a88c, []int{0}
}

func (m *Book) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Book.Unmarshal(m, b)
}
func (m *Book) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Book.Marshal(b, m, deterministic)
}
func (m *Book) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Book.Merge(m, src)
}
func (m *Book) XXX_Size() int {
	return xxx_messageInfo_Book.Size(m)
}
func (m *Book) XXX_DiscardUnknown() {
	xxx_messageInfo_Book.DiscardUnknown(m)
}

var xxx_messageInfo_Book proto.InternalMessageInfo

func (m *Book) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *Book) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Book) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type ListBooksResponse struct {
	Books                []*Book  `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBooksResponse) Reset()         { *m = ListBooksResponse{} }
func (m *ListBooksResponse) String() string { return proto.CompactTextString(m) }
func (*ListBooksResponse) ProtoMessage()    {}
func (*ListBooksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f82f486e563a88c, []int{1}
}

func (m *ListBooksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBooksResponse.Unmarshal(m, b)
}
func (m *ListBooksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBooksResponse.Marshal(b, m, deterministic)
}
func (m *ListBooksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBooksResponse.Merge(m, src)
}
func (m *ListBooksResponse) XXX_Size() int {
	return xxx_messageInfo_ListBooksResponse.Size(m)
}
func (m *ListBooksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBooksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListBooksResponse proto.InternalMessageInfo

func (m *ListBooksResponse) GetBooks() []*Book {
	if m != nil {
		return m.Books
	}
	return nil
}

type ListShelvesResponse struct {
	Shelves              []*Shelf `protobuf:"bytes,1,rep,name=shelves,proto3" json:"shelves,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListShelvesResponse) Reset()         { *m = ListShelvesResponse{} }
func (m *ListShelvesResponse) String() string { return proto.CompactTextString(m) }
func (*ListShelvesResponse) ProtoMessage()    {}
func (*ListShelvesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f82f486e563a88c, []int{2}
}

func (m *ListShelvesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListShelvesResponse.Unmarshal(m, b)
}
func (m *ListShelvesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListShelvesResponse.Marshal(b, m, deterministic)
}
func (m *ListShelvesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListShelvesResponse.Merge(m, src)
}
func (m *ListShelvesResponse) XXX_Size() int {
	return xxx_messageInfo_ListShelvesResponse.Size(m)
}
func (m *ListShelvesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListShelvesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListShelvesResponse proto.InternalMessageInfo

func (m *ListShelvesResponse) GetShelves() []*Shelf {
	if m != nil {
		return m.Shelves
	}
	return nil
}

type Shelf struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Theme                string   `protobuf:"bytes,2,opt,name=theme,proto3" json:"theme,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Shelf) Reset()         { *m = Shelf{} }
func (m *Shelf) String() string { return proto.CompactTextString(m) }
func (*Shelf) ProtoMessage()    {}
func (*Shelf) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f82f486e563a88c, []int{3}
}

func (m *Shelf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Shelf.Unmarshal(m, b)
}
func (m *Shelf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Shelf.Marshal(b, m, deterministic)
}
func (m *Shelf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Shelf.Merge(m, src)
}
func (m *Shelf) XXX_Size() int {
	return xxx_messageInfo_Shelf.Size(m)
}
func (m *Shelf) XXX_DiscardUnknown() {
	xxx_messageInfo_Shelf.DiscardUnknown(m)
}

var xxx_messageInfo_Shelf proto.InternalMessageInfo

func (m *Shelf) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Shelf) GetTheme() string {
	if m != nil {
		return m.Theme
	}
	return ""
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f82f486e563a88c, []int{4}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type CreateShelfParameters struct {
	Shelf                *Shelf   `protobuf:"bytes,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateShelfParameters) Reset()         { *m = CreateShelfParameters{} }
func (m *CreateShelfParameters) String() string { return proto.CompactTextString(m) }
func (*CreateShelfParameters) ProtoMessage()    {}
func (*CreateShelfParameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f82f486e563a88c, []int{5}
}

func (m *CreateShelfParameters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateShelfParameters.Unmarshal(m, b)
}
func (m *CreateShelfParameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateShelfParameters.Marshal(b, m, deterministic)
}
func (m *CreateShelfParameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateShelfParameters.Merge(m, src)
}
func (m *CreateShelfParameters) XXX_Size() int {
	return xxx_messageInfo_CreateShelfParameters.Size(m)
}
func (m *CreateShelfParameters) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateShelfParameters.DiscardUnknown(m)
}

var xxx_messageInfo_CreateShelfParameters proto.InternalMessageInfo

func (m *CreateShelfParameters) GetShelf() *Shelf {
	if m != nil {
		return m.Shelf
	}
	return nil
}

type GetShelfParameters struct {
	Shelf                int32    `protobuf:"varint,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetShelfParameters) Reset()         { *m = GetShelfParameters{} }
func (m *GetShelfParameters) String() string { return proto.CompactTextString(m) }
func (*GetShelfParameters) ProtoMessage()    {}
func (*GetShelfParameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f82f486e563a88c, []int{6}
}

func (m *GetShelfParameters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetShelfParameters.Unmarshal(m, b)
}
func (m *GetShelfParameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetShelfParameters.Marshal(b, m, deterministic)
}
func (m *GetShelfParameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShelfParameters.Merge(m, src)
}
func (m *GetShelfParameters) XXX_Size() int {
	return xxx_messageInfo_GetShelfParameters.Size(m)
}
func (m *GetShelfParameters) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShelfParameters.DiscardUnknown(m)
}

var xxx_messageInfo_GetShelfParameters proto.InternalMessageInfo

func (m *GetShelfParameters) GetShelf() int32 {
	if m != nil {
		return m.Shelf
	}
	return 0
}

type DeleteShelfParameters struct {
	Shelf                int32    `protobuf:"varint,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteShelfParameters) Reset()         { *m = DeleteShelfParameters{} }
func (m *DeleteShelfParameters) String() string { return proto.CompactTextString(m) }
func (*DeleteShelfParameters) ProtoMessage()    {}
func (*DeleteShelfParameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f82f486e563a88c, []int{7}
}

func (m *DeleteShelfParameters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteShelfParameters.Unmarshal(m, b)
}
func (m *DeleteShelfParameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteShelfParameters.Marshal(b, m, deterministic)
}
func (m *DeleteShelfParameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteShelfParameters.Merge(m, src)
}
func (m *DeleteShelfParameters) XXX_Size() int {
	return xxx_messageInfo_DeleteShelfParameters.Size(m)
}
func (m *DeleteShelfParameters) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteShelfParameters.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteShelfParameters proto.InternalMessageInfo

func (m *DeleteShelfParameters) GetShelf() int32 {
	if m != nil {
		return m.Shelf
	}
	return 0
}

type ListBooksParameters struct {
	Shelf                int32    `protobuf:"varint,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBooksParameters) Reset()         { *m = ListBooksParameters{} }
func (m *ListBooksParameters) String() string { return proto.CompactTextString(m) }
func (*ListBooksParameters) ProtoMessage()    {}
func (*ListBooksParameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f82f486e563a88c, []int{8}
}

func (m *ListBooksParameters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBooksParameters.Unmarshal(m, b)
}
func (m *ListBooksParameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBooksParameters.Marshal(b, m, deterministic)
}
func (m *ListBooksParameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBooksParameters.Merge(m, src)
}
func (m *ListBooksParameters) XXX_Size() int {
	return xxx_messageInfo_ListBooksParameters.Size(m)
}
func (m *ListBooksParameters) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBooksParameters.DiscardUnknown(m)
}

var xxx_messageInfo_ListBooksParameters proto.InternalMessageInfo

func (m *ListBooksParameters) GetShelf() int32 {
	if m != nil {
		return m.Shelf
	}
	return 0
}

type CreateBookParameters struct {
	Shelf                int32    `protobuf:"varint,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
	Book                 *Book    `protobuf:"bytes,2,opt,name=book,proto3" json:"book,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBookParameters) Reset()         { *m = CreateBookParameters{} }
func (m *CreateBookParameters) String() string { return proto.CompactTextString(m) }
func (*CreateBookParameters) ProtoMessage()    {}
func (*CreateBookParameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f82f486e563a88c, []int{9}
}

func (m *CreateBookParameters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBookParameters.Unmarshal(m, b)
}
func (m *CreateBookParameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBookParameters.Marshal(b, m, deterministic)
}
func (m *CreateBookParameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBookParameters.Merge(m, src)
}
func (m *CreateBookParameters) XXX_Size() int {
	return xxx_messageInfo_CreateBookParameters.Size(m)
}
func (m *CreateBookParameters) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBookParameters.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBookParameters proto.InternalMessageInfo

func (m *CreateBookParameters) GetShelf() int32 {
	if m != nil {
		return m.Shelf
	}
	return 0
}

func (m *CreateBookParameters) GetBook() *Book {
	if m != nil {
		return m.Book
	}
	return nil
}

type GetBookParameters struct {
	Shelf                int32    `protobuf:"varint,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
	Book                 int32    `protobuf:"varint,2,opt,name=book,proto3" json:"book,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBookParameters) Reset()         { *m = GetBookParameters{} }
func (m *GetBookParameters) String() string { return proto.CompactTextString(m) }
func (*GetBookParameters) ProtoMessage()    {}
func (*GetBookParameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f82f486e563a88c, []int{10}
}

func (m *GetBookParameters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBookParameters.Unmarshal(m, b)
}
func (m *GetBookParameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBookParameters.Marshal(b, m, deterministic)
}
func (m *GetBookParameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBookParameters.Merge(m, src)
}
func (m *GetBookParameters) XXX_Size() int {
	return xxx_messageInfo_GetBookParameters.Size(m)
}
func (m *GetBookParameters) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBookParameters.DiscardUnknown(m)
}

var xxx_messageInfo_GetBookParameters proto.InternalMessageInfo

func (m *GetBookParameters) GetShelf() int32 {
	if m != nil {
		return m.Shelf
	}
	return 0
}

func (m *GetBookParameters) GetBook() int32 {
	if m != nil {
		return m.Book
	}
	return 0
}

type DeleteBookParameters struct {
	Shelf                int32    `protobuf:"varint,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
	Book                 int32    `protobuf:"varint,2,opt,name=book,proto3" json:"book,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBookParameters) Reset()         { *m = DeleteBookParameters{} }
func (m *DeleteBookParameters) String() string { return proto.CompactTextString(m) }
func (*DeleteBookParameters) ProtoMessage()    {}
func (*DeleteBookParameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f82f486e563a88c, []int{11}
}

func (m *DeleteBookParameters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBookParameters.Unmarshal(m, b)
}
func (m *DeleteBookParameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBookParameters.Marshal(b, m, deterministic)
}
func (m *DeleteBookParameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBookParameters.Merge(m, src)
}
func (m *DeleteBookParameters) XXX_Size() int {
	return xxx_messageInfo_DeleteBookParameters.Size(m)
}
func (m *DeleteBookParameters) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBookParameters.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBookParameters proto.InternalMessageInfo

func (m *DeleteBookParameters) GetShelf() int32 {
	if m != nil {
		return m.Shelf
	}
	return 0
}

func (m *DeleteBookParameters) GetBook() int32 {
	if m != nil {
		return m.Book
	}
	return 0
}

func init() {
	proto.RegisterType((*Book)(nil), "bookstore.Book")
	proto.RegisterType((*ListBooksResponse)(nil), "bookstore.ListBooksResponse")
	proto.RegisterType((*ListShelvesResponse)(nil), "bookstore.ListShelvesResponse")
	proto.RegisterType((*Shelf)(nil), "bookstore.Shelf")
	proto.RegisterType((*Error)(nil), "bookstore.Error")
	proto.RegisterType((*CreateShelfParameters)(nil), "bookstore.CreateShelfParameters")
	proto.RegisterType((*GetShelfParameters)(nil), "bookstore.GetShelfParameters")
	proto.RegisterType((*DeleteShelfParameters)(nil), "bookstore.DeleteShelfParameters")
	proto.RegisterType((*ListBooksParameters)(nil), "bookstore.ListBooksParameters")
	proto.RegisterType((*CreateBookParameters)(nil), "bookstore.CreateBookParameters")
	proto.RegisterType((*GetBookParameters)(nil), "bookstore.GetBookParameters")
	proto.RegisterType((*DeleteBookParameters)(nil), "bookstore.DeleteBookParameters")
}

func init() { proto.RegisterFile("bookstore.proto", fileDescriptor_6f82f486e563a88c) }

var fileDescriptor_6f82f486e563a88c = []byte{
	// 605 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xdd, 0x6e, 0x12, 0x51,
	0x10, 0x0e, 0xa5, 0x0b, 0x65, 0x88, 0x29, 0x8c, 0x14, 0x37, 0x2b, 0xa5, 0xe4, 0x68, 0x4d, 0x83,
	0x91, 0x8d, 0x35, 0xde, 0x90, 0x18, 0xb5, 0xda, 0xd4, 0x0b, 0x2f, 0x2c, 0xc6, 0x98, 0x78, 0x63,
	0x96, 0x72, 0x80, 0x4d, 0x81, 0x43, 0xf6, 0x9c, 0x9a, 0x98, 0xa6, 0x37, 0xbe, 0x82, 0x8f, 0xe6,
	0x2b, 0x78, 0xeb, 0x3b, 0x98, 0x9d, 0xb3, 0x7f, 0xec, 0x2e, 0xc5, 0x78, 0x05, 0x67, 0x66, 0xbe,
	0x6f, 0xbe, 0xd9, 0xf9, 0x06, 0x76, 0x87, 0x42, 0x5c, 0x4a, 0x25, 0x3c, 0xde, 0x5b, 0x7a, 0x42,
	0x09, 0xac, 0x44, 0x01, 0xab, 0x35, 0x11, 0x62, 0x32, 0xe3, 0xb6, 0xb3, 0x74, 0x6d, 0x67, 0xb1,
	0x10, 0xca, 0x51, 0xae, 0x58, 0x48, 0x5d, 0x68, 0xdd, 0x0f, 0xb2, 0xf4, 0x1a, 0x5e, 0x8d, 0x6d,
	0x3e, 0x5f, 0xaa, 0xef, 0x41, 0xb2, 0x93, 0x4e, 0x8e, 0xb8, 0xbc, 0xf0, 0xdc, 0xa5, 0x12, 0x9e,
	0xae, 0x60, 0xef, 0x60, 0xfb, 0x44, 0x88, 0x4b, 0x6c, 0x42, 0xc9, 0xb9, 0x52, 0x53, 0xe1, 0x99,
	0x85, 0x4e, 0xe1, 0xa8, 0x32, 0x08, 0x5e, 0x88, 0xb0, 0xbd, 0x70, 0xe6, 0xdc, 0xdc, 0xa2, 0x28,
	0xfd, 0xc7, 0x06, 0x18, 0xca, 0x55, 0x33, 0x6e, 0x16, 0x29, 0xa8, 0x1f, 0xac, 0x0f, 0xf5, 0xf7,
	0xae, 0x54, 0x3e, 0x9b, 0x1c, 0x70, 0xb9, 0x14, 0x0b, 0xc9, 0xf1, 0x10, 0x0c, 0x1a, 0xc4, 0x2c,
	0x74, 0x8a, 0x47, 0xd5, 0xe3, 0xdd, 0x5e, 0x3c, 0xa7, 0x5f, 0x38, 0xd0, 0x59, 0xf6, 0x1a, 0xee,
	0xfa, 0xd8, 0x8f, 0x53, 0x3e, 0xfb, 0xc6, 0x63, 0x74, 0x17, 0xca, 0x52, 0x87, 0x02, 0x7c, 0x2d,
	0x81, 0xf7, 0x8b, 0xc7, 0x83, 0xb0, 0x80, 0x3d, 0x05, 0x83, 0x22, 0x91, 0xe2, 0x42, 0x4a, 0xf1,
	0x94, 0x47, 0x63, 0xe8, 0x07, 0x7b, 0x0e, 0xc6, 0xa9, 0xe7, 0xe9, 0x21, 0x2f, 0xc4, 0x48, 0x43,
	0x8c, 0x01, 0xfd, 0x47, 0x13, 0xca, 0x73, 0x2e, 0xa5, 0x33, 0x09, 0x41, 0xe1, 0x93, 0xbd, 0x84,
	0xbd, 0x37, 0x1e, 0x77, 0x14, 0xa7, 0x7e, 0x1f, 0x1c, 0xcf, 0x99, 0x73, 0xc5, 0x3d, 0x89, 0x8f,
	0xc0, 0xf0, 0xd5, 0x8c, 0x89, 0x27, 0x4f, 0xac, 0x4e, 0xb3, 0x2e, 0xe0, 0x19, 0x57, 0x69, 0x74,
	0x23, 0x89, 0x2e, 0x86, 0xb5, 0x4f, 0x60, 0xef, 0x2d, 0x9f, 0xf1, 0x6c, 0xb3, 0xfc, 0xf2, 0xc7,
	0xfa, 0x43, 0xd2, 0x12, 0x36, 0x16, 0x9f, 0x43, 0x43, 0x0f, 0xe2, 0x97, 0x6f, 0xaa, 0xc6, 0x07,
	0xb0, 0xed, 0xcf, 0x43, 0x5f, 0x23, 0x67, 0x93, 0x94, 0x64, 0x2f, 0xa0, 0x7e, 0xc6, 0xd5, 0x3f,
	0xf1, 0x61, 0x82, 0xaf, 0x18, 0xc0, 0x5f, 0x41, 0x43, 0x4f, 0xfb, 0xbf, 0x0c, 0xc7, 0x7f, 0x4a,
	0x50, 0x39, 0x09, 0x95, 0xe1, 0x67, 0xa8, 0x26, 0x7c, 0x85, 0xcd, 0x9e, 0xbe, 0x87, 0x5e, 0x78,
	0x0f, 0xbd, 0x53, 0xff, 0x58, 0xac, 0x76, 0x62, 0x98, 0x1c, 0x1f, 0xb2, 0xda, 0x8f, 0x5f, 0xbf,
	0x7f, 0x6e, 0x01, 0xee, 0xd8, 0x81, 0xdb, 0xf0, 0x0b, 0x54, 0x13, 0x1e, 0xc0, 0x4e, 0x82, 0x20,
	0xd7, 0x1b, 0x56, 0xc6, 0x0c, 0xec, 0x1e, 0x91, 0xd6, 0x59, 0x44, 0xda, 0x0f, 0xc6, 0x3a, 0x87,
	0x3b, 0xf1, 0xca, 0x6f, 0x93, 0xbd, 0x26, 0x1e, 0xca, 0xed, 0xc6, 0x72, 0x3f, 0xc1, 0x4e, 0xe8,
	0x38, 0xdc, 0x4f, 0x28, 0xc9, 0xda, 0x30, 0x47, 0xa8, 0x49, 0x74, 0x88, 0xb5, 0x90, 0xce, 0xbe,
	0x26, 0xa1, 0x37, 0xe8, 0x40, 0x35, 0x61, 0xce, 0x95, 0xaf, 0x90, 0x6b, 0xda, 0xb5, 0x8a, 0x83,
	0x16, 0xdd, 0x6c, 0x0b, 0x17, 0x2a, 0x91, 0xa1, 0x31, 0xbd, 0xa7, 0x94, 0xcd, 0xad, 0x56, 0x5e,
	0x3e, 0xda, 0x62, 0x9b, 0x9a, 0x98, 0xd8, 0x4c, 0x37, 0xb1, 0x09, 0x85, 0x23, 0x80, 0xf8, 0x1c,
	0xf0, 0x20, 0xb3, 0xd2, 0x55, 0x4f, 0x5a, 0xe9, 0x0b, 0x60, 0x0f, 0x89, 0xbf, 0xcd, 0xd6, 0xf0,
	0xf7, 0xc9, 0xa0, 0xf8, 0x15, 0xca, 0xc1, 0x85, 0x60, 0x6b, 0x75, 0x13, 0x9b, 0xf8, 0x0f, 0x89,
	0xff, 0x00, 0xf7, 0xf3, 0xf9, 0xed, 0x6b, 0xff, 0xe7, 0x06, 0x67, 0x00, 0xf1, 0x0d, 0xad, 0x8c,
	0x91, 0x77, 0x5a, 0x6b, 0x57, 0x12, 0x74, 0xeb, 0xde, 0xde, 0x6d, 0x58, 0x22, 0xd8, 0xb3, 0xbf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x75, 0xd6, 0x7d, 0xc1, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BookstoreClient is the client API for Bookstore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BookstoreClient interface {
	ListShelves(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ListShelvesResponse, error)
	CreateShelf(ctx context.Context, in *CreateShelfParameters, opts ...grpc.CallOption) (*Shelf, error)
	DeleteShelves(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	GetShelf(ctx context.Context, in *GetShelfParameters, opts ...grpc.CallOption) (*Shelf, error)
	DeleteShelf(ctx context.Context, in *DeleteShelfParameters, opts ...grpc.CallOption) (*empty.Empty, error)
	ListBooks(ctx context.Context, in *ListBooksParameters, opts ...grpc.CallOption) (*ListBooksResponse, error)
	CreateBook(ctx context.Context, in *CreateBookParameters, opts ...grpc.CallOption) (*Book, error)
	GetBook(ctx context.Context, in *GetBookParameters, opts ...grpc.CallOption) (*Book, error)
	DeleteBook(ctx context.Context, in *DeleteBookParameters, opts ...grpc.CallOption) (*empty.Empty, error)
}

type bookstoreClient struct {
	cc *grpc.ClientConn
}

func NewBookstoreClient(cc *grpc.ClientConn) BookstoreClient {
	return &bookstoreClient{cc}
}

func (c *bookstoreClient) ListShelves(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ListShelvesResponse, error) {
	out := new(ListShelvesResponse)
	err := c.cc.Invoke(ctx, "/bookstore.Bookstore/ListShelves", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) CreateShelf(ctx context.Context, in *CreateShelfParameters, opts ...grpc.CallOption) (*Shelf, error) {
	out := new(Shelf)
	err := c.cc.Invoke(ctx, "/bookstore.Bookstore/CreateShelf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) DeleteShelves(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/bookstore.Bookstore/DeleteShelves", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) GetShelf(ctx context.Context, in *GetShelfParameters, opts ...grpc.CallOption) (*Shelf, error) {
	out := new(Shelf)
	err := c.cc.Invoke(ctx, "/bookstore.Bookstore/GetShelf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) DeleteShelf(ctx context.Context, in *DeleteShelfParameters, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/bookstore.Bookstore/DeleteShelf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) ListBooks(ctx context.Context, in *ListBooksParameters, opts ...grpc.CallOption) (*ListBooksResponse, error) {
	out := new(ListBooksResponse)
	err := c.cc.Invoke(ctx, "/bookstore.Bookstore/ListBooks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) CreateBook(ctx context.Context, in *CreateBookParameters, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/bookstore.Bookstore/CreateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) GetBook(ctx context.Context, in *GetBookParameters, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/bookstore.Bookstore/GetBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) DeleteBook(ctx context.Context, in *DeleteBookParameters, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/bookstore.Bookstore/DeleteBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookstoreServer is the server API for Bookstore service.
type BookstoreServer interface {
	ListShelves(context.Context, *empty.Empty) (*ListShelvesResponse, error)
	CreateShelf(context.Context, *CreateShelfParameters) (*Shelf, error)
	DeleteShelves(context.Context, *empty.Empty) (*empty.Empty, error)
	GetShelf(context.Context, *GetShelfParameters) (*Shelf, error)
	DeleteShelf(context.Context, *DeleteShelfParameters) (*empty.Empty, error)
	ListBooks(context.Context, *ListBooksParameters) (*ListBooksResponse, error)
	CreateBook(context.Context, *CreateBookParameters) (*Book, error)
	GetBook(context.Context, *GetBookParameters) (*Book, error)
	DeleteBook(context.Context, *DeleteBookParameters) (*empty.Empty, error)
}

// UnimplementedBookstoreServer can be embedded to have forward compatible implementations.
type UnimplementedBookstoreServer struct {
}

func (*UnimplementedBookstoreServer) ListShelves(ctx context.Context, req *empty.Empty) (*ListShelvesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListShelves not implemented")
}
func (*UnimplementedBookstoreServer) CreateShelf(ctx context.Context, req *CreateShelfParameters) (*Shelf, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShelf not implemented")
}
func (*UnimplementedBookstoreServer) DeleteShelves(ctx context.Context, req *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteShelves not implemented")
}
func (*UnimplementedBookstoreServer) GetShelf(ctx context.Context, req *GetShelfParameters) (*Shelf, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShelf not implemented")
}
func (*UnimplementedBookstoreServer) DeleteShelf(ctx context.Context, req *DeleteShelfParameters) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteShelf not implemented")
}
func (*UnimplementedBookstoreServer) ListBooks(ctx context.Context, req *ListBooksParameters) (*ListBooksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBooks not implemented")
}
func (*UnimplementedBookstoreServer) CreateBook(ctx context.Context, req *CreateBookParameters) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBook not implemented")
}
func (*UnimplementedBookstoreServer) GetBook(ctx context.Context, req *GetBookParameters) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}
func (*UnimplementedBookstoreServer) DeleteBook(ctx context.Context, req *DeleteBookParameters) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBook not implemented")
}

func RegisterBookstoreServer(s *grpc.Server, srv BookstoreServer) {
	s.RegisterService(&_Bookstore_serviceDesc, srv)
}

func _Bookstore_ListShelves_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).ListShelves(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookstore.Bookstore/ListShelves",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).ListShelves(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_CreateShelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShelfParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).CreateShelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookstore.Bookstore/CreateShelf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).CreateShelf(ctx, req.(*CreateShelfParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_DeleteShelves_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).DeleteShelves(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookstore.Bookstore/DeleteShelves",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).DeleteShelves(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_GetShelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShelfParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).GetShelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookstore.Bookstore/GetShelf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).GetShelf(ctx, req.(*GetShelfParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_DeleteShelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteShelfParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).DeleteShelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookstore.Bookstore/DeleteShelf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).DeleteShelf(ctx, req.(*DeleteShelfParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_ListBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBooksParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).ListBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookstore.Bookstore/ListBooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).ListBooks(ctx, req.(*ListBooksParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_CreateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).CreateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookstore.Bookstore/CreateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).CreateBook(ctx, req.(*CreateBookParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookstore.Bookstore/GetBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).GetBook(ctx, req.(*GetBookParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_DeleteBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBookParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).DeleteBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookstore.Bookstore/DeleteBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).DeleteBook(ctx, req.(*DeleteBookParameters))
	}
	return interceptor(ctx, in, info, handler)
}

var _Bookstore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bookstore.Bookstore",
	HandlerType: (*BookstoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListShelves",
			Handler:    _Bookstore_ListShelves_Handler,
		},
		{
			MethodName: "CreateShelf",
			Handler:    _Bookstore_CreateShelf_Handler,
		},
		{
			MethodName: "DeleteShelves",
			Handler:    _Bookstore_DeleteShelves_Handler,
		},
		{
			MethodName: "GetShelf",
			Handler:    _Bookstore_GetShelf_Handler,
		},
		{
			MethodName: "DeleteShelf",
			Handler:    _Bookstore_DeleteShelf_Handler,
		},
		{
			MethodName: "ListBooks",
			Handler:    _Bookstore_ListBooks_Handler,
		},
		{
			MethodName: "CreateBook",
			Handler:    _Bookstore_CreateBook_Handler,
		},
		{
			MethodName: "GetBook",
			Handler:    _Bookstore_GetBook_Handler,
		},
		{
			MethodName: "DeleteBook",
			Handler:    _Bookstore_DeleteBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bookstore.proto",
}
