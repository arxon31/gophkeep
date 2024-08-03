package requests

type SaveBankCredentialsDTO struct {
	User       string
	Meta       string
	CardNumber string
	Owner      string
	CVV        int64
}

func (dto SaveBankCredentialsDTO) Validate() error {
	panic("implement me")
}
