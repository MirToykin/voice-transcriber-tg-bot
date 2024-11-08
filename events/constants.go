package events

type EventType int
type ChatType string

const (
	Unknown EventType = iota
	TextMessage
	VoiceMessage
	GroupTextMessage
)

const GroupChat ChatType = "group"
