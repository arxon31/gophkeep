package docs

import "github.com/ilyakaznacheev/cleanenv"

type config struct {
	URI    string `env:"MONGO_URI" env-required`
	DBname string `env:"MONGO_DB_NAME" env-required`
}

func NewConfig() (*config, error) {
	var cfg *config
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
