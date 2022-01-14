package resource

type Resource string

const (
	Chat          = Resource("chat")
	DirectMessage = Resource("direct-message")
	Field         = Resource("field")
)
