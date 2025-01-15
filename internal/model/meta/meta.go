package meta

type Meta string

func (m Meta) Validate() error {
	if m == "" {
		return ErrEmptyMeta
	}
	return nil
}
