package json

type WsResp struct {
	SyncId string `json:"syncId"`
	Data   any    `json:"data"`
}

type WsReq struct {
	SyncId  string `json:"syncId"`
	Command string `json:"command"`
	Content any    `json:"content"`
}