package models

import "github.com/google/uuid"

type Project struct {
	Id      uuid.UUID
	OwnerId uuid.UUID
	Title   string
}
