package responses

type GetCredentialsResponseDTO struct {
	UserName string `bson:"user_name"`
	Password string `bson:"password"`
}
