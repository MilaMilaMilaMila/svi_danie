package models

import "github.com/google/uuid"

type Img struct {
	Id   uuid.UUID
	Data []byte
}
