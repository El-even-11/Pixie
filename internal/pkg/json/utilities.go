package json

type MessageChainType int

const (
	FriendMessage MessageChainType = 1
	GroupMessage  MessageChainType = 2
)

type MessageType int

const (
	Source MessageType = 1
	At     MessageType = 2
	AtAll  MessageType = 3
	Face   MessageType = 4
	Plain  MessageType = 5
	Image  MessageType = 6
)

type EventType int

const (
	FriendInputStatusChangedEvent   EventType = 1
	GroupRecallEvent                EventType = 2
	FriendRecallEvent               EventType = 3
	NudgeEvent                      EventType = 4
	MemberCardChangeEvent           EventType = 5
	MemberHonorChangeEvent          EventType = 6
	NewFriendRequestEvent           EventType = 7
	BotInvitedJoinGroupRequestEvent EventType = 8
)

type WsRespData struct {
	SyncId string `json:"syncId"`
	Data   any    `json:"data"`
}

type WsReqData struct {
	SyncId  string `json:"syncId"`
	Command string `json:"command"`
	Content []byte `json:"content"`
}

type TypeProbe struct {
	Type string `json:"type"`
}
