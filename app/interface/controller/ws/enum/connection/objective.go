package connection

type Objective string

const (
	Chat          = Objective("chat")
	DirectMessage = Objective("direct-message")
	Field         = Objective("field")
)
