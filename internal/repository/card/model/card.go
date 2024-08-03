package model

type GetCard struct {
	User string
	Meta string `bson:"meta"`
}

type Card struct {
	User       string
	Meta       string `bson:"meta"`
	Owner      string `bson:"owner"`
	NumberHash []byte `bson:"number_hash"`
	NumberSalt []byte `bson:"number_salt"`
	CVVHash    []byte `bson:"cvv_hash"`
	CVVSalt    []byte `bson:"cvv_salt"`
}
