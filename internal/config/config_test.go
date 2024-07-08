package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	// Создаем временный файл конфигурации
	tempFile, err := os.CreateTemp("", "config.yaml")
	assert.NoError(t, err)
	defer os.Remove(tempFile.Name())

	// Записываем пример конфигурации в файл
	_, err = tempFile.Write([]byte(`
kafka:
  brokers: ["kafka:9092"]
  topic: "test_topic"
  group: "test_group"
postgres:
  host: "localhost"
  port: 5432
  user: "user"
  password: "password"
  dbname: "dbname"
api:
  port: 8080
`))
	assert.NoError(t, err)
	tempFile.Close()

	// Загружаем конфигурацию
	cfg, err := LoadConfig(tempFile.Name())
	assert.NoError(t, err)

	// Проверяем загруженные данные
	assert.Equal(t, "kafka:9092", cfg.Kafka.Brokers[0])
	assert.Equal(t, "test_topic", cfg.Kafka.Topic)
	assert.Equal(t, "test_group", cfg.Kafka.Group)
	assert.Equal(t, "localhost", cfg.Postgres.Host)
	assert.Equal(t, 5432, cfg.Postgres.Port)
	assert.Equal(t, "user", cfg.Postgres.User)
	assert.Equal(t, "password", cfg.Postgres.Password)
	assert.Equal(t, "dbname", cfg.Postgres.DBName)
	assert.Equal(t, 8080, cfg.API.Port)
}