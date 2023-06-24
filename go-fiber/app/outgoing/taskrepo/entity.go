package taskrepo

import "github.com/google/uuid"

type TaskEntity struct {
	TaskId      uuid.UUID
	Title       string
	Description *string
}
