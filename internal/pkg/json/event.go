package json

type Event struct {
	// common
	Type string `json:"type"`

	// FriendInputStatusChangedEvent
	Friend    friend `json:"friend,omitempty"`
	Inputting bool   `json:"inputting,omitempty"`

	// GroupRecallEvent & FriendRecallEvent
	Group     group `json:"group,omitempty"`
	AuthorID  int64 `json:"authorId,omitempty"`
	MessageID int   `json:"messageId,omitempty"`

	// NudgeEvent
	FromID  int64   `json:"fromId,omitempty"`
	Subject subject `json:"subject,omitempty"`
	Suffix  string  `json:"suffix,omitempty"`
	Target  int64   `json:"target,omitempty"`

	// MemberCardChangeEvent
	Origin  string `json:"origin,omitempty"`
	Current string `json:"current,omitempty"`

	// MemberHonorChangeEvent
	Honor string `json:"honor,omitempty"`

	// MemberCardChangeEvent & MemberHonorChangeEvent
	Member member `json:"member,omitempty"`

	// NudgeEvent & MemberHonorChangeEvent
	Action string `json:"action,omitempty"`
}

type friend struct {
	ID       int64  `json:"id"`
	Nickname string `json:"nickname"`
}

type group struct {
	ID int64 `json:"id"`
}

type subject struct {
	ID   int64  `json:"id"`
	Kind string `json:"kind"`
}

type member struct {
	ID    int64 `json:"id"`
	Group group `json:"group"`
}
