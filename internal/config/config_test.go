package config

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewConfig(t *testing.T) {

	t.Run("must_return_config", func(t *testing.T) {
		t.Setenv("MONGO_URI", "mongo")
		t.Setenv("MONGO_DB_NAME", "mongo")
		t.Setenv("S3_URI", "s3")
		t.Setenv("S3_USER", "s3")
		t.Setenv("S3_PASSWORD", "s3")
		t.Setenv("CRYPTO_KEY", "key")
		t.Setenv("JWT_KEY", "key")
		cfg, err := NewConfig()
		require.NoError(t, err)
		require.Equal(t, "mongo", cfg.Mongo.URI)
		require.Equal(t, "mongo", cfg.Mongo.DBName)
		require.Equal(t, "s3", cfg.S3.URI)
		require.Equal(t, "s3", cfg.S3.User)
		require.Equal(t, "s3", cfg.S3.Password)
		require.Equal(t, "key", cfg.Secrets.CryptoKey)
		require.Equal(t, "key", cfg.Secrets.JWTKey)
	})

	t.Run("must_return_error", func(t *testing.T) {
		_, err := NewConfig()
		require.Error(t, err)
	})

}
