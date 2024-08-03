package responses

type GetBankCredentialsResponseDTO struct {
	CardNumber string `bson:"card_number"`
	Owner      string `bson:"owner"`
	CVV        int64  `bson:"cvv"`
}
