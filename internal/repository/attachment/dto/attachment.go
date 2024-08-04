package dto

import "bytes"

type GetAttachment struct {
	User string
	Meta string `bson:"meta"`
}

type Attachment struct {
	User    string
	Meta    string `bson:"meta"`
	Name    string `bson:"attachment_name"`
	Content *bytes.Buffer
}
