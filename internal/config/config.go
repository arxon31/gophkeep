package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Mongo   Mongo
	S3      S3
	Secrets Secrets
}

type Mongo struct {
	URI    string `env:"MONGO_URI" env-required:"true"`
	DBName string `env:"MONGO_DB_NAME" env-required:"true"`
}

type S3 struct {
	URI      string `env:"S3_URI" env-required:"true"`
	User     string `env:"S3_USER" env-required:"true"`
	Password string `env:"S3_PASSWORD" env-required:"true"`
}

type Secrets struct {
	CryptoKey string `env:"CRYPTO_KEY" env-required:"true"`
	JWTKey    string `env:"JWT_KEY" env-required:"true"`
}

func NewConfig() (*Config, error) {
	var cfg Config

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
