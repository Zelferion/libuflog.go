// Copyright (c) 2026 Jane Doe
// Licensed under the MIT License. See LICENSE file in the project root.
package libuflog

import (
	"regexp"
	"time"

	"github.com/zelferion/libuflog.go/formatting"
)

type Message struct {
	rawMessage       string
	formattedMessage string
	time             time.Time
	messageType      string
	typeStyle        []formatting.Ansi
}

var ansiEscape = regexp.MustCompile(`\033\[[0-9;]*m`)

func NewMessage(msg string) Message {
	return Message{
		rawMessage:       ansiEscape.ReplaceAllString(msg, ""),
		formattedMessage: msg,
	}
}

func (m *Message) SetTime(t time.Time) {
	m.time = t
}

func (m *Message) SetType(t string) {
	m.messageType = t
}

func (m *Message) SetRawMessage(str string) {
	m.rawMessage = str
}

func (m *Message) SetFormattedMessage(str string) {
	m.formattedMessage = str
}

func (m *Message) SetTypeStyle(args ...formatting.Ansi) {
	m.typeStyle = args
}

func (m *Message) GetRawMessage() string {
	return m.rawMessage
}

func (m *Message) GetFormattedMessage() string {
	return m.formattedMessage
}

func (m *Message) GetMessageType() string {
	return m.messageType
}

func (m *Message) GetTypeStyle() []formatting.Ansi {
	return m.typeStyle
}
