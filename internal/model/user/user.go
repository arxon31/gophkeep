package user

type User string

func (u User) Validate() error {
	if u == "" {
		return ErrEmptyUser
	}

	return nil
}
