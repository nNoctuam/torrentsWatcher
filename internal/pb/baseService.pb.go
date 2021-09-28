// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: baseService.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Torrent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title      string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	PageUrl    string                 `protobuf:"bytes,3,opt,name=page_url,json=pageUrl,proto3" json:"page_url,omitempty"`
	FileUrl    string                 `protobuf:"bytes,4,opt,name=file_url,json=fileUrl,proto3" json:"file_url,omitempty"`
	Forum      string                 `protobuf:"bytes,5,opt,name=forum,proto3" json:"forum,omitempty"`
	Author     string                 `protobuf:"bytes,6,opt,name=author,proto3" json:"author,omitempty"`
	Size       uint64                 `protobuf:"varint,7,opt,name=size,proto3" json:"size,omitempty"`
	Seeders    uint64                 `protobuf:"varint,8,opt,name=seeders,proto3" json:"seeders,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	UploadedAt *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=uploaded_at,json=uploadedAt,proto3" json:"uploaded_at,omitempty"`
}

func (x *Torrent) Reset() {
	*x = Torrent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baseService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Torrent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Torrent) ProtoMessage() {}

func (x *Torrent) ProtoReflect() protoreflect.Message {
	mi := &file_baseService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Torrent.ProtoReflect.Descriptor instead.
func (*Torrent) Descriptor() ([]byte, []int) {
	return file_baseService_proto_rawDescGZIP(), []int{0}
}

func (x *Torrent) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Torrent) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Torrent) GetPageUrl() string {
	if x != nil {
		return x.PageUrl
	}
	return ""
}

func (x *Torrent) GetFileUrl() string {
	if x != nil {
		return x.FileUrl
	}
	return ""
}

func (x *Torrent) GetForum() string {
	if x != nil {
		return x.Forum
	}
	return ""
}

func (x *Torrent) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Torrent) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Torrent) GetSeeders() uint64 {
	if x != nil {
		return x.Seeders
	}
	return 0
}

func (x *Torrent) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Torrent) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Torrent) GetUploadedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UploadedAt
	}
	return nil
}

type PartToRename struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OldName string `protobuf:"bytes,1,opt,name=oldName,proto3" json:"oldName,omitempty"`
	NewName string `protobuf:"bytes,2,opt,name=newName,proto3" json:"newName,omitempty"`
}

func (x *PartToRename) Reset() {
	*x = PartToRename{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baseService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartToRename) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartToRename) ProtoMessage() {}

func (x *PartToRename) ProtoReflect() protoreflect.Message {
	mi := &file_baseService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartToRename.ProtoReflect.Descriptor instead.
func (*PartToRename) Descriptor() ([]byte, []int) {
	return file_baseService_proto_rawDescGZIP(), []int{1}
}

func (x *PartToRename) GetOldName() string {
	if x != nil {
		return x.OldName
	}
	return ""
}

func (x *PartToRename) GetNewName() string {
	if x != nil {
		return x.NewName
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baseService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_baseService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_baseService_proto_rawDescGZIP(), []int{2}
}

type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=Text,proto3" json:"Text,omitempty"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baseService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_baseService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_baseService_proto_rawDescGZIP(), []int{3}
}

func (x *SearchRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type AddTorrentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=Url,proto3" json:"Url,omitempty"`
}

func (x *AddTorrentRequest) Reset() {
	*x = AddTorrentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baseService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTorrentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTorrentRequest) ProtoMessage() {}

func (x *AddTorrentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_baseService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTorrentRequest.ProtoReflect.Descriptor instead.
func (*AddTorrentRequest) Descriptor() ([]byte, []int) {
	return file_baseService_proto_rawDescGZIP(), []int{4}
}

func (x *AddTorrentRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type DeleteTorrentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteTorrentRequest) Reset() {
	*x = DeleteTorrentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baseService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTorrentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTorrentRequest) ProtoMessage() {}

func (x *DeleteTorrentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_baseService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTorrentRequest.ProtoReflect.Descriptor instead.
func (*DeleteTorrentRequest) Descriptor() ([]byte, []int) {
	return file_baseService_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteTorrentRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DownloadTorrentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url    string `protobuf:"bytes,1,opt,name=Url,proto3" json:"Url,omitempty"`
	Folder string `protobuf:"bytes,2,opt,name=Folder,proto3" json:"Folder,omitempty"`
}

func (x *DownloadTorrentRequest) Reset() {
	*x = DownloadTorrentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baseService_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadTorrentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadTorrentRequest) ProtoMessage() {}

func (x *DownloadTorrentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_baseService_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadTorrentRequest.ProtoReflect.Descriptor instead.
func (*DownloadTorrentRequest) Descriptor() ([]byte, []int) {
	return file_baseService_proto_rawDescGZIP(), []int{6}
}

func (x *DownloadTorrentRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *DownloadTorrentRequest) GetFolder() string {
	if x != nil {
		return x.Folder
	}
	return ""
}

type RenameTorrentPartsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Names []*PartToRename `protobuf:"bytes,2,rep,name=names,proto3" json:"names,omitempty"`
}

func (x *RenameTorrentPartsRequest) Reset() {
	*x = RenameTorrentPartsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baseService_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenameTorrentPartsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenameTorrentPartsRequest) ProtoMessage() {}

func (x *RenameTorrentPartsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_baseService_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenameTorrentPartsRequest.ProtoReflect.Descriptor instead.
func (*RenameTorrentPartsRequest) Descriptor() ([]byte, []int) {
	return file_baseService_proto_rawDescGZIP(), []int{7}
}

func (x *RenameTorrentPartsRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RenameTorrentPartsRequest) GetNames() []*PartToRename {
	if x != nil {
		return x.Names
	}
	return nil
}

type TorrentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Torrent *Torrent `protobuf:"bytes,1,opt,name=torrent,proto3" json:"torrent,omitempty"`
}

func (x *TorrentResponse) Reset() {
	*x = TorrentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baseService_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TorrentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TorrentResponse) ProtoMessage() {}

func (x *TorrentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_baseService_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TorrentResponse.ProtoReflect.Descriptor instead.
func (*TorrentResponse) Descriptor() ([]byte, []int) {
	return file_baseService_proto_rawDescGZIP(), []int{8}
}

func (x *TorrentResponse) GetTorrent() *Torrent {
	if x != nil {
		return x.Torrent
	}
	return nil
}

type TorrentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Torrents []*Torrent `protobuf:"bytes,1,rep,name=torrents,proto3" json:"torrents,omitempty"`
}

func (x *TorrentsResponse) Reset() {
	*x = TorrentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baseService_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TorrentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TorrentsResponse) ProtoMessage() {}

func (x *TorrentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_baseService_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TorrentsResponse.ProtoReflect.Descriptor instead.
func (*TorrentsResponse) Descriptor() ([]byte, []int) {
	return file_baseService_proto_rawDescGZIP(), []int{9}
}

func (x *TorrentsResponse) GetTorrents() []*Torrent {
	if x != nil {
		return x.Torrents
	}
	return nil
}

type DownloadFoldersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Folders []string `protobuf:"bytes,1,rep,name=folders,proto3" json:"folders,omitempty"`
}

func (x *DownloadFoldersResponse) Reset() {
	*x = DownloadFoldersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baseService_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadFoldersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadFoldersResponse) ProtoMessage() {}

func (x *DownloadFoldersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_baseService_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadFoldersResponse.ProtoReflect.Descriptor instead.
func (*DownloadFoldersResponse) Descriptor() ([]byte, []int) {
	return file_baseService_proto_rawDescGZIP(), []int{10}
}

func (x *DownloadFoldersResponse) GetFolders() []string {
	if x != nil {
		return x.Folders
	}
	return nil
}

type DownloadTorrentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Hash string `protobuf:"bytes,3,opt,name=Hash,proto3" json:"Hash,omitempty"`
}

func (x *DownloadTorrentResponse) Reset() {
	*x = DownloadTorrentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baseService_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadTorrentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadTorrentResponse) ProtoMessage() {}

func (x *DownloadTorrentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_baseService_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadTorrentResponse.ProtoReflect.Descriptor instead.
func (*DownloadTorrentResponse) Descriptor() ([]byte, []int) {
	return file_baseService_proto_rawDescGZIP(), []int{11}
}

func (x *DownloadTorrentResponse) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *DownloadTorrentResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DownloadTorrentResponse) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

var File_baseService_proto protoreflect.FileDescriptor

var file_baseService_proto_rawDesc = []byte{
	0x0a, 0x11, 0x62, 0x61, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf4,
	0x02, 0x0a, 0x07, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66,
	0x69, 0x6c, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x65, 0x64,
	0x65, 0x72, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x73, 0x65, 0x65, 0x64, 0x65,
	0x72, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x65, 0x64, 0x41, 0x74, 0x22, 0x42, 0x0a, 0x0c, 0x50, 0x61, 0x72, 0x74, 0x54, 0x6f, 0x52,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6e, 0x65, 0x77, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6e, 0x65, 0x77, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x23, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x54, 0x65, 0x78, 0x74, 0x22, 0x25, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x54, 0x6f,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x55, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x72, 0x6c, 0x22, 0x26,
	0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x42, 0x0a, 0x16, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x55, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55,
	0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x22, 0x59, 0x0a, 0x19, 0x52, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x54, 0x6f, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x05,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x3e, 0x0a, 0x0f, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x74, 0x6f, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x74, 0x6f,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x22, 0x41, 0x0a, 0x10, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x74, 0x6f, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x08,
	0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x33, 0x0a, 0x17, 0x44, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x22, 0x51, 0x0a,
	0x17, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x48, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x48, 0x61, 0x73, 0x68,
	0x32, 0x95, 0x04, 0x0a, 0x0b, 0x42, 0x61, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3f, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x45, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65,
	0x64, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x44,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x12, 0x0f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x64,
	0x64, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x6f, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0d,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54,
	0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x58, 0x0a, 0x0f, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x6f, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x12, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x12, 0x52, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x74, 0x73,
	0x12, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x52, 0x65, 0x6e, 0x61,
	0x6d, 0x65, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_baseService_proto_rawDescOnce sync.Once
	file_baseService_proto_rawDescData = file_baseService_proto_rawDesc
)

func file_baseService_proto_rawDescGZIP() []byte {
	file_baseService_proto_rawDescOnce.Do(func() {
		file_baseService_proto_rawDescData = protoimpl.X.CompressGZIP(file_baseService_proto_rawDescData)
	})
	return file_baseService_proto_rawDescData
}

var file_baseService_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_baseService_proto_goTypes = []interface{}{
	(*Torrent)(nil),                   // 0: protobuf.Torrent
	(*PartToRename)(nil),              // 1: protobuf.PartToRename
	(*Empty)(nil),                     // 2: protobuf.Empty
	(*SearchRequest)(nil),             // 3: protobuf.SearchRequest
	(*AddTorrentRequest)(nil),         // 4: protobuf.AddTorrentRequest
	(*DeleteTorrentRequest)(nil),      // 5: protobuf.DeleteTorrentRequest
	(*DownloadTorrentRequest)(nil),    // 6: protobuf.DownloadTorrentRequest
	(*RenameTorrentPartsRequest)(nil), // 7: protobuf.RenameTorrentPartsRequest
	(*TorrentResponse)(nil),           // 8: protobuf.TorrentResponse
	(*TorrentsResponse)(nil),          // 9: protobuf.TorrentsResponse
	(*DownloadFoldersResponse)(nil),   // 10: protobuf.DownloadFoldersResponse
	(*DownloadTorrentResponse)(nil),   // 11: protobuf.DownloadTorrentResponse
	(*timestamppb.Timestamp)(nil),     // 12: google.protobuf.Timestamp
}
var file_baseService_proto_depIdxs = []int32{
	12, // 0: protobuf.Torrent.created_at:type_name -> google.protobuf.Timestamp
	12, // 1: protobuf.Torrent.updated_at:type_name -> google.protobuf.Timestamp
	12, // 2: protobuf.Torrent.uploaded_at:type_name -> google.protobuf.Timestamp
	1,  // 3: protobuf.RenameTorrentPartsRequest.names:type_name -> protobuf.PartToRename
	0,  // 4: protobuf.TorrentResponse.torrent:type_name -> protobuf.Torrent
	0,  // 5: protobuf.TorrentsResponse.torrents:type_name -> protobuf.Torrent
	3,  // 6: protobuf.BaseService.Search:input_type -> protobuf.SearchRequest
	2,  // 7: protobuf.BaseService.GetMonitoredTorrents:input_type -> protobuf.Empty
	2,  // 8: protobuf.BaseService.GetDownloadFolders:input_type -> protobuf.Empty
	4,  // 9: protobuf.BaseService.AddTorrent:input_type -> protobuf.AddTorrentRequest
	5,  // 10: protobuf.BaseService.DeleteTorrent:input_type -> protobuf.DeleteTorrentRequest
	6,  // 11: protobuf.BaseService.DownloadTorrent:input_type -> protobuf.DownloadTorrentRequest
	7,  // 12: protobuf.BaseService.RenameTorrentParts:input_type -> protobuf.RenameTorrentPartsRequest
	9,  // 13: protobuf.BaseService.Search:output_type -> protobuf.TorrentsResponse
	9,  // 14: protobuf.BaseService.GetMonitoredTorrents:output_type -> protobuf.TorrentsResponse
	10, // 15: protobuf.BaseService.GetDownloadFolders:output_type -> protobuf.DownloadFoldersResponse
	8,  // 16: protobuf.BaseService.AddTorrent:output_type -> protobuf.TorrentResponse
	2,  // 17: protobuf.BaseService.DeleteTorrent:output_type -> protobuf.Empty
	11, // 18: protobuf.BaseService.DownloadTorrent:output_type -> protobuf.DownloadTorrentResponse
	2,  // 19: protobuf.BaseService.RenameTorrentParts:output_type -> protobuf.Empty
	13, // [13:20] is the sub-list for method output_type
	6,  // [6:13] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_baseService_proto_init() }
func file_baseService_proto_init() {
	if File_baseService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_baseService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Torrent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_baseService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartToRename); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_baseService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_baseService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_baseService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTorrentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_baseService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTorrentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_baseService_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadTorrentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_baseService_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenameTorrentPartsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_baseService_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TorrentResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_baseService_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TorrentsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_baseService_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadFoldersResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_baseService_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadTorrentResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_baseService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_baseService_proto_goTypes,
		DependencyIndexes: file_baseService_proto_depIdxs,
		MessageInfos:      file_baseService_proto_msgTypes,
	}.Build()
	File_baseService_proto = out.File
	file_baseService_proto_rawDesc = nil
	file_baseService_proto_goTypes = nil
	file_baseService_proto_depIdxs = nil
}
