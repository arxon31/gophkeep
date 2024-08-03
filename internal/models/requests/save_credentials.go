package requests

type SaveCredentialsDTO struct {
	User     string
	Meta     string
	UserName string
	Password string
}

func (dto SaveCredentialsDTO) Validate() error {
	panic("implement me")
}
