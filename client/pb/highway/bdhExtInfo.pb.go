// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/highway/bdhExtInfo.proto

package highway

import (
	proto "github.com/RomiChan/protobuf/proto"
)

type CommFileExtReq struct {
	ActionType proto.Option[uint32] `protobuf:"varint,1,opt"`
	Uuid       []byte               `protobuf:"bytes,2,opt"`
}

type CommFileExtRsp struct {
	Retcode     proto.Option[int32] `protobuf:"varint,1,opt"`
	DownloadUrl []byte              `protobuf:"bytes,2,opt"`
}

type PicInfo struct {
	Idx    proto.Option[uint32] `protobuf:"varint,1,opt"`
	Size   proto.Option[uint32] `protobuf:"varint,2,opt"`
	BinMd5 []byte               `protobuf:"bytes,3,opt"`
	Type   proto.Option[uint32] `protobuf:"varint,4,opt"`
}

type QQVoiceExtReq struct {
	Qid     []byte               `protobuf:"bytes,1,opt"`
	Fmt     proto.Option[uint32] `protobuf:"varint,2,opt"`
	Rate    proto.Option[uint32] `protobuf:"varint,3,opt"`
	Bits    proto.Option[uint32] `protobuf:"varint,4,opt"`
	Channel proto.Option[uint32] `protobuf:"varint,5,opt"`
	Pinyin  proto.Option[uint32] `protobuf:"varint,6,opt"`
}

type QQVoiceExtRsp struct {
	Qid     []byte              `protobuf:"bytes,1,opt"`
	Retcode proto.Option[int32] `protobuf:"varint,2,opt"`
	Result  []*QQVoiceResult    `protobuf:"bytes,3,rep"`
}

type QQVoiceResult struct {
	Text   []byte               `protobuf:"bytes,1,opt"`
	Pinyin []byte               `protobuf:"bytes,2,opt"`
	Source proto.Option[uint32] `protobuf:"varint,3,opt"`
}

type ShortVideoReqExtInfo struct {
	Cmd                  proto.Option[uint32]   `protobuf:"varint,1,opt"`
	SessionId            proto.Option[uint64]   `protobuf:"varint,2,opt"`
	Thumbinfo            *PicInfo               `protobuf:"bytes,3,opt"`
	Videoinfo            *VideoInfo             `protobuf:"bytes,4,opt"`
	ShortvideoSureReq    *ShortVideoSureReqInfo `protobuf:"bytes,5,opt"`
	IsMergeCmdBeforeData proto.Option[bool]     `protobuf:"varint,6,opt"`
	_                    [0]func()
}

type ShortVideoRspExtInfo struct {
	Cmd               proto.Option[uint32]   `protobuf:"varint,1,opt"`
	SessionId         proto.Option[uint64]   `protobuf:"varint,2,opt"`
	Retcode           proto.Option[int32]    `protobuf:"varint,3,opt"`
	Errinfo           []byte                 `protobuf:"bytes,4,opt"`
	Thumbinfo         *PicInfo               `protobuf:"bytes,5,opt"`
	Videoinfo         *VideoInfo             `protobuf:"bytes,6,opt"`
	ShortvideoSureRsp *ShortVideoSureRspInfo `protobuf:"bytes,7,opt"`
	RetryFlag         proto.Option[uint32]   `protobuf:"varint,8,opt"`
}

type ShortVideoSureReqInfo struct {
	Fromuin         proto.Option[uint64] `protobuf:"varint,1,opt"`
	ChatType        proto.Option[uint32] `protobuf:"varint,2,opt"`
	Touin           proto.Option[uint64] `protobuf:"varint,3,opt"`
	GroupCode       proto.Option[uint64] `protobuf:"varint,4,opt"`
	ClientType      proto.Option[uint32] `protobuf:"varint,5,opt"`
	Thumbinfo       *PicInfo             `protobuf:"bytes,6,opt"`
	MergeVideoinfo  []*VideoInfo         `protobuf:"bytes,7,rep"`
	DropVideoinfo   []*VideoInfo         `protobuf:"bytes,8,rep"`
	BusinessType    proto.Option[uint32] `protobuf:"varint,9,opt"`
	SubBusinessType proto.Option[uint32] `protobuf:"varint,10,opt"`
}

type ShortVideoSureRspInfo struct {
	Fileid    []byte               `protobuf:"bytes,1,opt"`
	Ukey      []byte               `protobuf:"bytes,2,opt"`
	Videoinfo *VideoInfo           `protobuf:"bytes,3,opt"`
	MergeCost proto.Option[uint32] `protobuf:"varint,4,opt"`
}

type StoryVideoExtReq struct {
	_ [0]func()
}

type StoryVideoExtRsp struct {
	Retcode proto.Option[int32] `protobuf:"varint,1,opt"`
	Msg     []byte              `protobuf:"bytes,2,opt"`
	CdnUrl  []byte              `protobuf:"bytes,3,opt"`
	FileKey []byte              `protobuf:"bytes,4,opt"`
	FileId  []byte              `protobuf:"bytes,5,opt"`
}

type UploadPicExtInfo struct {
	FileResid        []byte `protobuf:"bytes,1,opt"`
	DownloadUrl      []byte `protobuf:"bytes,2,opt"`
	ThumbDownloadUrl []byte `protobuf:"bytes,3,opt"`
}

type VideoInfo struct {
	Idx       proto.Option[uint32] `protobuf:"varint,1,opt"`
	Size      proto.Option[uint32] `protobuf:"varint,2,opt"`
	BinMd5    []byte               `protobuf:"bytes,3,opt"`
	Format    proto.Option[uint32] `protobuf:"varint,4,opt"`
	ResLen    proto.Option[uint32] `protobuf:"varint,5,opt"`
	ResWidth  proto.Option[uint32] `protobuf:"varint,6,opt"`
	Time      proto.Option[uint32] `protobuf:"varint,7,opt"`
	Starttime proto.Option[uint64] `protobuf:"varint,8,opt"`
	IsAudio   proto.Option[uint32] `protobuf:"varint,9,opt"`
}
