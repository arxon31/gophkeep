package s3

import "github.com/ilyakaznacheev/cleanenv"

type config struct {
	URI      string `env:"S3_URI" env-required`
	User     string `env:"S3_USER" env-required`
	Password string `env:"S3_PASSWORD" env-required`
}

func NewConfig() (*config, error) {
	var cfg *config

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
