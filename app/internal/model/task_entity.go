package model

import "github.com/google/uuid"

type TaskEntity struct {
	Uuid        uuid.UUID
	Title       string
	Description *string
}
