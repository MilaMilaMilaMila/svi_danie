package models

import (
	"encoding/json"
	"github.com/google/uuid"
)

type Page struct {
	Id        uuid.UUID
	OwnerId   uuid.UUID
	ProjectId uuid.UUID
	Title     string
	Data      json.RawMessage
}
