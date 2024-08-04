package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

type encryptor struct {
	nonce       []byte
	blockCipher cipher.AEAD
}

func New(key []byte) (*encryptor, error) {
	aesBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(aesBlock)
	if err != nil {
		return nil, err
	}

	nonce, err := generateRandom(aesGCM.NonceSize())
	if err != nil {
		return nil, err
	}

	return &encryptor{nonce: nonce, blockCipher: aesGCM}, nil
}

func (e *encryptor) Encrypt(src []byte) ([]byte, error) {
	return e.blockCipher.Seal(nil, e.nonce, src, nil), nil
}

func (e *encryptor) Decrypt(src []byte) ([]byte, error) {
	return e.blockCipher.Open(nil, e.nonce, src, nil)
}

func generateRandom(size int) ([]byte, error) {
	b := make([]byte, size)

	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
