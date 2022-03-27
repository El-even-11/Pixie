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
	BotOnlineEvent                  EventType = 1
	BotOfflineEventActive           EventType = 2
	BotOfflineEventForce            EventType = 3
	BotOfflineEventDropped          EventType = 4
	FriendInputStatusChangedEvent   EventType = 5
	FriendNickChangedEvent          EventType = 6
	BotJoinGroupEvent               EventType = 7
	GroupRecallEvent                EventType = 8
	FriendRecallEvent               EventType = 9
	NudgeEvent                      EventType = 10
	GroupNameChangeEvent            EventType = 11
	MemberCardChangeEvent           EventType = 12
	MemberHonorChangeEvent          EventType = 13
	NewFriendRequestEvent           EventType = 14
	BotInvitedJoinGroupRequestEvent EventType = 15
	Poke                            EventType = 16
	Dice                            EventType = 17
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
