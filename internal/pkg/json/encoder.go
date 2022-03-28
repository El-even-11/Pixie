package json

import (
	"encoding/json"
)

func Encode(req WsReq) ([]byte, error) {
	return json.Marshal(req)
}
