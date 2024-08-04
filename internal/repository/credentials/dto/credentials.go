package dto

type Credentials struct {
	User              string
	Meta              string `bson:"meta"`
	EncryptedUserName []byte `bson:"username_enc"`
	EncryptedPassword []byte `bson:"password_enc"`
}

type GetCredentials struct {
	User string
	Meta string `bson:"meta"`
}
