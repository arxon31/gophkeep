package requests

type GetByMetaDTO struct {
	User string
	Meta string
}

func (dto GetByMetaDTO) Validate() error {
	panic("implement me")
}
