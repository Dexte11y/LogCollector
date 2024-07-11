package clickhouse

import (
	"database/sql"
	"fmt"
	"log"
	"logcollector/internal/config"
	"time"

	_ "github.com/ClickHouse/clickhouse-go/v2"
)

// ClickHouse представляет собой структуру, которая оборачивает соединение с базой данных ClickHouse
type ClickHouse struct {
	db *sql.DB
}

// NewClickHouse создает новое соединение с базой данных ClickHouse
// cfg - конфигурация подключения к базе данных ClickHouse
func NewClickHouse(cfg *config.ClickHouseConfig) (*ClickHouse, error) {
	connStr := fmt.Sprintf("tcp://%s:%d?username=%s&password=%s&enable_http_compression=1",
		cfg.Host, cfg.Port, cfg.User, cfg.Password)
	db, err := sql.Open("clickhouse", connStr)
	if err != nil {
		log.Printf("internal/storage/clickhouse.go: Failed to start clickhouse: %v", err)
		return nil, err
	}
	return &ClickHouse{db: db}, nil
}

// Close закрывает соединение с базой данных
func (c *ClickHouse) Close() error {
	return c.db.Close()
}

// DB возвращает экземпляр базы данных SQL
func (c *ClickHouse) DB() *sql.DB {
	return c.db
}

// Ping проверяет соединение с базой данных
func (c *ClickHouse) Ping() error {
	return c.db.Ping()
}

// InsertMessage вставляет сообщение в таблицу logs
// date - дата и время сообщения
// key - назваение приложения с которого пришло сообщение
// message - сообщение, которое нужно вставить
func (c *ClickHouse) InsertMessage(date time.Time, key, message string) error {
	_, err := c.db.Exec(`
		INSERT INTO logs 
		(date, app_name, message) 
		VALUES (?, ?, ?)`,
		date, key, message,
	)
	if err != nil {
		log.Printf("internal/storage/clickhouse.go: Failed to insert message into ClickHouse: %v", err)
		return err
	}
	return nil
}
