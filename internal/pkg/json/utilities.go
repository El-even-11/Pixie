package json

type WsRespData struct {
	SyncId string `json:"syncId"`
	Data   any    `json:"data"`
}

type WsReqData struct {
	SyncId  string `json:"syncId"`
	Command string `json:"command"`
	Content any    `json:"content"`
}

type TypeProbe struct {
	Type string `json:"type"`
}
