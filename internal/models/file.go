package models

import "bytes"

type FileDTO struct {
	User string
	Meta string
	Name string
	Data bytes.Buffer
}
