package json

import (
	"encoding/json"
)

func Encode(req WsReqData) ([]byte, error) {
	return json.Marshal(req)
}
