package models

import "bytes"

type FileDTO struct {
	User string
	Meta string
	Name string
	Data bytes.Buffer
}

func (dto FileDTO) Validate() error {
	panic("implement me")
}
