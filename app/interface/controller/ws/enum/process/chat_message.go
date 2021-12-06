package process

type ChatProcess string

const (
	Add    = ChatProcess("add")
	Modify = ChatProcess("modify")
	Delete = ChatProcess("delete")
)
