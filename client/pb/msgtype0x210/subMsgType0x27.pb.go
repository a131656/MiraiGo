// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/msgtype0x210/subMsgType0x27.proto

package msgtype0x210

import (
	proto "github.com/RomiChan/protobuf/proto"
)

type AddGroup struct {
	Groupid   proto.Option[uint32] `protobuf:"varint,1,opt"`
	Sortid    proto.Option[uint32] `protobuf:"varint,2,opt"`
	Groupname []byte               `protobuf:"bytes,3,opt"`
}

type AppointmentNotify struct {
	FromUin         proto.Option[uint64] `protobuf:"varint,1,opt"`
	AppointId       proto.Option[string] `protobuf:"bytes,2,opt"`
	Notifytype      proto.Option[uint32] `protobuf:"varint,3,opt"`
	TipsContent     proto.Option[string] `protobuf:"bytes,4,opt"`
	UnreadCount     proto.Option[uint32] `protobuf:"varint,5,opt"`
	JoinWording     proto.Option[string] `protobuf:"bytes,6,opt"`
	ViewWording     proto.Option[string] `protobuf:"bytes,7,opt"`
	Sig             []byte               `protobuf:"bytes,8,opt"`
	EventInfo       []byte               `protobuf:"bytes,9,opt"`
	NearbyEventInfo []byte               `protobuf:"bytes,10,opt"`
	FeedEventInfo   []byte               `protobuf:"bytes,11,opt"`
}

type BinaryMsg struct {
	OpType  proto.Option[uint32] `protobuf:"varint,1,opt"`
	OpValue []byte               `protobuf:"bytes,2,opt"`
}

type ChatMatchInfo struct {
	Sig              []byte               `protobuf:"bytes,1,opt"`
	Uin              proto.Option[uint64] `protobuf:"varint,2,opt"`
	MatchUin         proto.Option[uint64] `protobuf:"varint,3,opt"`
	TipsWording      []byte               `protobuf:"bytes,4,opt"`
	LeftChatTime     proto.Option[uint32] `protobuf:"varint,5,opt"`
	TimeStamp        proto.Option[uint64] `protobuf:"varint,6,opt"`
	MatchExpiredTime proto.Option[uint32] `protobuf:"varint,7,opt"`
	C2CExpiredTime   proto.Option[uint32] `protobuf:"varint,8,opt"`
	MatchCount       proto.Option[uint32] `protobuf:"varint,9,opt"`
	Nick             []byte               `protobuf:"bytes,10,opt"`
}

type ConfMsgRoamFlag struct {
	Confid    proto.Option[uint64] `protobuf:"varint,1,opt"`
	Flag      proto.Option[uint32] `protobuf:"varint,2,opt"`
	Timestamp proto.Option[uint64] `protobuf:"varint,3,opt"`
	_         [0]func()
}

type DaRenNotify struct {
	Uin             proto.Option[uint64] `protobuf:"varint,1,opt"`
	LoginDays       proto.Option[uint32] `protobuf:"varint,2,opt"`
	Days            proto.Option[uint32] `protobuf:"varint,3,opt"`
	IsYestodayLogin proto.Option[uint32] `protobuf:"varint,4,opt"`
	IsTodayLogin    proto.Option[uint32] `protobuf:"varint,5,opt"`
	_               [0]func()
}

type DelFriend struct {
	Uins []uint64 `protobuf:"varint,1,rep"`
}

type DelGroup struct {
	Groupid proto.Option[uint32] `protobuf:"varint,1,opt"`
	_       [0]func()
}

type FanpaiziNotify struct {
	FromUin     proto.Option[uint64] `protobuf:"varint,1,opt"`
	FromNick    proto.Option[string] `protobuf:"bytes,2,opt"`
	TipsContent []byte               `protobuf:"bytes,3,opt"`
	Sig         []byte               `protobuf:"bytes,4,opt"`
}

type ForwardBody struct {
	// optional uint32 notifyType = 1;
	// optional uint32 opType = 2;
	// optional AddGroup addGroup = 3;
	// optional DelGroup delGroup = 4;
	// optional ModGroupName modGroupName = 5;
	// optional ModGroupSort modGroupSort = 6;
	// optional ModFriendGroup modFriendGroup = 7;
	// optional ModProfile modProfile = 8;
	// optional ModFriendRemark modFriendRemark = 9;
	// optional ModLongNick modLongNick = 10;
	// optional ModCustomFace modCustomFace = 11;
	ModGroupProfile *ModGroupProfile `protobuf:"bytes,12,opt"`
	// optional ModGroupMemberProfile modGroupMemberProfile = 13;
	DelFriend *DelFriend `protobuf:"bytes,14,opt"`
	_         [0]func()
}

type FrdCustomOnlineStatusChange struct {
	Uin proto.Option[uint64] `protobuf:"varint,1,opt"`
	_   [0]func()
}

type FriendGroup struct {
	Fuin       proto.Option[uint64] `protobuf:"varint,1,opt"`
	OldGroupId []uint32             `protobuf:"varint,2,rep"`
	NewGroupId []uint32             `protobuf:"varint,3,rep"`
}

type FriendRemark struct {
	Type      proto.Option[uint32] `protobuf:"varint,1,opt"`
	Fuin      proto.Option[uint64] `protobuf:"varint,2,opt"`
	RmkName   []byte               `protobuf:"bytes,3,opt"`
	GroupCode proto.Option[uint64] `protobuf:"varint,4,opt"`
}

type GPS struct {
	Lat  proto.Option[int32] `protobuf:"varint,1,opt"`
	Lon  proto.Option[int32] `protobuf:"varint,2,opt"`
	Alt  proto.Option[int32] `protobuf:"varint,3,opt"`
	Type proto.Option[int32] `protobuf:"varint,4,opt"`
	_    [0]func()
}

type GroupMemberProfileInfo struct {
	Field proto.Option[uint32] `protobuf:"varint,1,opt"`
	Value []byte               `protobuf:"bytes,2,opt"`
}

type GroupProfileInfo struct {
	Field proto.Option[uint32] `protobuf:"varint,1,opt"`
	Value []byte               `protobuf:"bytes,2,opt"`
}

type GroupSort struct {
	Groupid proto.Option[uint32] `protobuf:"varint,1,opt"`
	Sortid  proto.Option[uint32] `protobuf:"varint,2,opt"`
	_       [0]func()
}

type GrpMsgRoamFlag struct {
	Groupcode proto.Option[uint64] `protobuf:"varint,1,opt"`
	Flag      proto.Option[uint32] `protobuf:"varint,2,opt"`
	Timestamp proto.Option[uint64] `protobuf:"varint,3,opt"`
	_         [0]func()
}

type HotFriendNotify struct {
	DstUin         proto.Option[uint64] `protobuf:"varint,1,opt"`
	PraiseHotLevel proto.Option[uint32] `protobuf:"varint,2,opt"`
	ChatHotLevel   proto.Option[uint32] `protobuf:"varint,3,opt"`
	PraiseHotDays  proto.Option[uint32] `protobuf:"varint,4,opt"`
	ChatHotDays    proto.Option[uint32] `protobuf:"varint,5,opt"`
	CloseLevel     proto.Option[uint32] `protobuf:"varint,6,opt"`
	CloseDays      proto.Option[uint32] `protobuf:"varint,7,opt"`
	PraiseFlag     proto.Option[uint32] `protobuf:"varint,8,opt"`
	ChatFlag       proto.Option[uint32] `protobuf:"varint,9,opt"`
	CloseFlag      proto.Option[uint32] `protobuf:"varint,10,opt"`
	NotifyTime     proto.Option[uint64] `protobuf:"varint,11,opt"`
	LastPraiseTime proto.Option[uint64] `protobuf:"varint,12,opt"`
	LastChatTime   proto.Option[uint64] `protobuf:"varint,13,opt"`
	QzoneHotLevel  proto.Option[uint32] `protobuf:"varint,14,opt"`
	QzoneHotDays   proto.Option[uint32] `protobuf:"varint,15,opt"`
	QzoneFlag      proto.Option[uint32] `protobuf:"varint,16,opt"`
	LastQzoneTime  proto.Option[uint64] `protobuf:"varint,17,opt"`
	_              [0]func()
}

type MQQCampusNotify struct {
	FromUin proto.Option[uint64] `protobuf:"varint,1,opt"`
	Wording proto.Option[string] `protobuf:"bytes,2,opt"`
	Target  proto.Option[string] `protobuf:"bytes,3,opt"`
	Type    proto.Option[uint32] `protobuf:"varint,4,opt"`
	Source  proto.Option[string] `protobuf:"bytes,5,opt"`
	_       [0]func()
}

type ModConfProfile struct {
	Uin          proto.Option[uint64] `protobuf:"varint,1,opt"`
	ConfUin      proto.Option[uint32] `protobuf:"varint,2,opt"`
	ProfileInfos []*ProfileInfo       `protobuf:"bytes,3,rep"`
}

type ModCustomFace struct {
	Type      proto.Option[uint32] `protobuf:"varint,1,opt"`
	Uin       proto.Option[uint64] `protobuf:"varint,2,opt"`
	GroupCode proto.Option[uint64] `protobuf:"varint,3,opt"`
	CmdUin    proto.Option[uint64] `protobuf:"varint,4,opt"`
	_         [0]func()
}

type ModFrdRoamPriv struct {
	RoamPriv []*OneRoamPriv `protobuf:"bytes,1,rep"`
}

type ModFriendGroup struct {
	FrdGroup []*FriendGroup `protobuf:"bytes,1,rep"`
}

type ModFriendRemark struct {
	FrdRmk []*FriendRemark `protobuf:"bytes,1,rep"`
}

type ModGroupMemberProfile struct {
	GroupUin                proto.Option[uint64]      `protobuf:"varint,1,opt"`
	Uin                     proto.Option[uint64]      `protobuf:"varint,2,opt"`
	GroupMemberProfileInfos []*GroupMemberProfileInfo `protobuf:"bytes,3,rep"`
	GroupCode               proto.Option[uint64]      `protobuf:"varint,4,opt"`
}

type ModGroupName struct {
	Groupid   proto.Option[uint32] `protobuf:"varint,1,opt"`
	Groupname []byte               `protobuf:"bytes,2,opt"`
}

type ModGroupProfile struct {
	GroupUin          proto.Option[uint64] `protobuf:"varint,1,opt"`
	GroupProfileInfos []*GroupProfileInfo  `protobuf:"bytes,2,rep"`
	GroupCode         proto.Option[uint64] `protobuf:"varint,3,opt"`
	CmdUin            proto.Option[uint64] `protobuf:"varint,4,opt"`
}

type ModGroupSort struct {
	Groupsort []*GroupSort `protobuf:"bytes,1,rep"`
}

type ModLongNick struct {
	Uin   proto.Option[uint64] `protobuf:"varint,1,opt"`
	Value []byte               `protobuf:"bytes,2,opt"`
}

type ModProfile struct {
	Uin          proto.Option[uint64] `protobuf:"varint,1,opt"`
	ProfileInfos []*ProfileInfo       `protobuf:"bytes,2,rep"`
}

type ModSnsGeneralInfo struct {
	SnsGeneralInfos []*SnsUpateBuffer `protobuf:"bytes,1,rep"`
}

type SubMsg0X27Body struct {
	ModInfos []*ForwardBody `protobuf:"bytes,1,rep"`
}

type NewComeinUser struct {
	Uin    proto.Option[uint64] `protobuf:"varint,1,opt"`
	IsFrd  proto.Option[uint32] `protobuf:"varint,2,opt"`
	Remark []byte               `protobuf:"bytes,3,opt"`
	Nick   []byte               `protobuf:"bytes,4,opt"`
}

type NewComeinUserNotify struct {
	MsgType       proto.Option[uint32] `protobuf:"varint,1,opt"`
	OngNotify     proto.Option[bool]   `protobuf:"varint,2,opt"`
	PushTime      proto.Option[uint32] `protobuf:"varint,3,opt"`
	NewComeinUser *NewComeinUser       `protobuf:"bytes,4,opt"`
	NewGroup      *NewGroup            `protobuf:"bytes,5,opt"`
	NewGroupUser  *NewGroupUser        `protobuf:"bytes,6,opt"`
	_             [0]func()
}

type NewGroup struct {
	GroupCode proto.Option[uint64] `protobuf:"varint,1,opt"`
	GroupName []byte               `protobuf:"bytes,2,opt"`
	OwnerUin  proto.Option[uint64] `protobuf:"varint,3,opt"`
	OwnerNick []byte               `protobuf:"bytes,4,opt"`
	Distance  []byte               `protobuf:"bytes,5,opt"`
}

type NewGroupUser struct {
	Uin      proto.Option[uint64] `protobuf:"varint,1,opt"`
	Sex      proto.Option[int32]  `protobuf:"varint,2,opt"`
	Age      proto.Option[int32]  `protobuf:"varint,3,opt"`
	Nick     proto.Option[string] `protobuf:"bytes,4,opt"`
	Distance []byte               `protobuf:"bytes,5,opt"`
}

type OneRoamPriv struct {
	Fuin      proto.Option[uint64] `protobuf:"varint,1,opt"`
	PrivTag   proto.Option[uint32] `protobuf:"varint,2,opt"`
	PrivValue proto.Option[uint32] `protobuf:"varint,3,opt"`
	_         [0]func()
}

type PraiseRankNotify struct {
	IsChampion proto.Option[uint32] `protobuf:"varint,11,opt"`
	RankNum    proto.Option[uint32] `protobuf:"varint,12,opt"`
	Msg        proto.Option[string] `protobuf:"bytes,13,opt"`
	_          [0]func()
}

type ProfileInfo struct {
	Field proto.Option[uint32] `protobuf:"varint,1,opt"`
	Value []byte               `protobuf:"bytes,2,opt"`
}

type PushReportDev struct {
	MsgType      proto.Option[uint32] `protobuf:"varint,1,opt"`
	Cookie       []byte               `protobuf:"bytes,4,opt"`
	ReportMaxNum proto.Option[uint32] `protobuf:"varint,5,opt"`
	Sn           []byte               `protobuf:"bytes,6,opt"`
}

type PushSearchDev struct {
	MsgType  proto.Option[uint32] `protobuf:"varint,1,opt"`
	GpsInfo  *GPS                 `protobuf:"bytes,2,opt"`
	DevTime  proto.Option[uint32] `protobuf:"varint,3,opt"`
	PushTime proto.Option[uint32] `protobuf:"varint,4,opt"`
	Din      proto.Option[uint64] `protobuf:"varint,5,opt"`
	Data     proto.Option[string] `protobuf:"bytes,6,opt"`
	_        [0]func()
}

type QQPayPush struct {
	Uin   proto.Option[uint64] `protobuf:"varint,1,opt"`
	PayOk proto.Option[bool]   `protobuf:"varint,2,opt"`
	_     [0]func()
}

type SnsUpateBuffer struct {
	Uin           proto.Option[uint64] `protobuf:"varint,1,opt"`
	Code          proto.Option[uint64] `protobuf:"varint,2,opt"`
	Result        proto.Option[uint32] `protobuf:"varint,3,opt"`
	SnsUpdateItem []*SnsUpdateItem     `protobuf:"bytes,400,rep"`
	Idlist        []uint32             `protobuf:"varint,401,rep"`
}

type SnsUpdateFlag struct {
	UpdateSnsFlag []*SnsUpdateOneFlag `protobuf:"bytes,1,rep"`
}

type SnsUpdateItem struct {
	UpdateSnsType proto.Option[uint32] `protobuf:"varint,1,opt"`
	Value         []byte               `protobuf:"bytes,2,opt"`
}

type SnsUpdateOneFlag struct {
	XUin proto.Option[uint64] `protobuf:"varint,1,opt"`
	Id   proto.Option[uint64] `protobuf:"varint,2,opt"`
	Flag proto.Option[uint32] `protobuf:"varint,3,opt"`
	_    [0]func()
}
