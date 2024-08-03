package requests

type SaveFileS3URLDTO struct {
	User string
	Meta string
	URL  string
}

func (dto SaveFileS3URLDTO) Validate() error {
	panic("implement me")
}
