package core

import (
	"encoding/json"
	"fmt"
)

type MessageType string

const (
	MessageTypeTransaction = "TRANSACTION"
	MessageTypeBlock       = "BLOCK"
)

type Message struct {
	Type    MessageType `json:"type"`
	Payload string      `json:"payload"`
}

func NewMessage(t MessageType, payload string) *Message {
	return &Message{Type: t, Payload: payload}
}

func (m *Message) Serialize() string {
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("Error serializing message: %v\n", err)
		return ""
	}
	return string(data)
}

func DeserializeMessage(data string) (*Message, error) {
	var message Message
	err := json.Unmarshal([]byte(data), &message)
	if err != nil {
		return nil, err
	}
	return &message, nil
}
