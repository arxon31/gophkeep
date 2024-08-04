package dto

type GetCard struct {
	User string
	Meta string `bson:"meta"`
}

type Card struct {
	User             string
	Meta             string `bson:"meta"`
	Owner            string `bson:"owner"`
	EncpryptedNumber []byte `bson:"number_enc"`
	EncrpytedCVV     []byte `bson:"cvv_enc"`
}
