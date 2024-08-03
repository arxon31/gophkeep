package model

type Credentials struct {
	User         string
	Meta         string `bson:"meta"`
	UserNameHash []byte `bson:"username_hash"`
	UserNameSalt []byte `bson:"username_salt"`
	PasswordHash []byte `bson:"password_hash"`
	PasswordSalt []byte `bson:"password_salt"`
}

type GetCredentials struct {
	User string
	Meta string `bson:"meta"`
}
