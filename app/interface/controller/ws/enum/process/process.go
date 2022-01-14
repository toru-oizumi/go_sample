package process

type Process string

const (
	Add    = Process("add")
	Modify = Process("modify")
	Delete = Process("delete")
)
