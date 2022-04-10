package json

func BuildWsReq(target int64, command string, t, objects []string) WsReq {
	return WsReq{
		SyncId:  "0",
		Command: command,
		Content: Message{
			Target:       target,
			MessageChain: BuildMessage(t, objects),
		},
	}
}

func BuildMessage(t, objects []string) []MessageItem {
	if len(t) != len(objects) {
		panic("arrays length mismatch")
	}

	messages := make([]MessageItem, 0)
	for i := 0; i < len(t); i++ {
		switch t[i] {
		case "Plain":
			messages = append(messages,
				MessageItem{
					Type: t[i],
					Text: objects[i],
				},
			)
		case "Image":
			messages = append(messages,
				MessageItem{
					Type: t[i],
					URL:  objects[i],
				},
			)
		default:
			panic("unknown message type")
		}
	}

	return messages
}
