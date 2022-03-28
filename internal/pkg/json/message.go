package json

type MessageChain struct {
	// common
	Messages []Message `json:"messageChain"`
	
	// receive
	Type   string `json:"type,omitempty"`
	Sender Sender `json:"sender,omitempty"`

	// send
	Target int64 `json:"target,omitempty"`
}

type Sender struct {
	// common
	ID int64 `json:"id"`

	// friend sender
	Nickname string `json:"nickname,omitempty"`

	// group sender
	MemberName         string    `json:"memberName,omitempty"`
	LastSpeakTimeStamp int64     `json:"lastSpeakTimestamp,omitempty"`
	Group              GroupInfo `json:"group,omitempty"`
}

type GroupInfo struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Message struct {
	// common
	Type string `json:"type"`

	// source
	ID   int   `json:"id,omitempty"`
	Time int64 `json:"time,omitempty"`

	// at
	Target int64 `json:"target,omitempty"`

	// face
	FaceID int `json:"faceId,omitempty"`

	// plain
	Text string `json:"text,omitempty"`

	// image
	URL string `json:"url,omitempty"`
}
