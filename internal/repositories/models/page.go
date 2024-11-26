package models

import (
	"encoding/json"
	"github.com/google/uuid"
)

type Page struct {
	Id        uuid.UUID       `json:"id"`
	OwnerId   uuid.UUID       `json:"owner_id"`
	ProjectId uuid.UUID       `json:"project_id"`
	Title     string          `json:"title"`
	Data      json.RawMessage `json:"data"`
}
