package errors

type Error int

const (
	BadRequest   Error = 4001
	TaskNotFound Error = 4004
	Panic        Error = 9009
)
