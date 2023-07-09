package utils

type Message struct {
	payload []byte
}

func CreateMessage(payload []byte) Message {
	m := Message{payload: payload}
	return m
}

// For Testing
// func (m *Message) readString() string {
// 	return string(m.payload)
// }
