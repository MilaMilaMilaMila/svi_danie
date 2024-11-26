package models

import "github.com/google/uuid"

type Project struct {
	Id      uuid.UUID `json:"id"`
	OwnerId uuid.UUID `json:"owner_id"`
	Title   string    `json:"title"`
	Pages   []*Page   `json:"pages"`
}
