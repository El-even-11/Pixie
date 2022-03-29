package json

func BuildMessage(t []string, objects []string) []Message {
	if len(t) != len(objects) {
		panic("arrays length mismatch")
	}

	messages := make([]Message, 0)
	for i := 0; i < len(t); i++ {
		switch t[i] {
		case "Plain":
			messages = append(messages,
				Message{
					Type: t[i],
					Text: objects[i],
				},
			)
		case "Image":
			messages = append(messages,
				Message{
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
