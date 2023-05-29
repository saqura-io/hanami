package clickhouse

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/saqura-io/hanami/config"
)

var (
	db   *sql.DB
	once sync.Once
)

type ClickHouse struct {
	Conn *sql.DB
}

func Connect() *ClickHouse {
	once.Do(func() {
		var err error
		db, err = sql.Open(
			"clickhouse",
			fmt.Sprintf(
				"clickhouse://@%s:%s?username=%s&password=%s&debug=true",
				config.Get("CLICKHOUSE_HOST", "127.0.0.1"),
				config.Get("CLICKHOUSE_PORT", "9000"),
				config.Get("CLICKHOUSE_USERNAME", "root"),
				config.Get("CLICKHOUSE_PASSWORD", "toor"),
			),
		)

		if err != nil {
			log.Fatal(err)
		}

		if err = db.Ping(); err != nil {
			log.Fatal(err)
		}

		clickhouseConnMaxLifetime, err := time.ParseDuration(config.Get("CLICKHOUSE_CONNECTION_MAX_LIFETIME", "0"))
		if err != nil {
			log.Fatal(err)
		}
		db.SetConnMaxLifetime(clickhouseConnMaxLifetime)

		clickhouseOpenConns, err := strconv.Atoi(config.Get("CLICKHOUSE_MAX_OPEN_CONNECTIONS", "100"))
		if err != nil {
			log.Fatal(err)
		}
		db.SetMaxOpenConns(clickhouseOpenConns)

		clickhouseMaxIdleConns, err := strconv.Atoi(config.Get("CLICKHOUSE_MAX_IDLE_CONNECTIONS", "50"))
		if err != nil {
			log.Fatal(err)
		}
		db.SetMaxIdleConns(clickhouseMaxIdleConns)
	})

	return &ClickHouse{
		Conn: db,
	}
}

func (c *ClickHouse) GetRowsByQuery(query string) ([]map[string]interface{}, error) {
	rows, err := c.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	var results []map[string]interface{}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}

		result := make(map[string]interface{})
		for i, col := range values {
			if col != nil {
				result[columns[i]] = string(col)
			}
		}
		results = append(results, result)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (c *ClickHouse) ExecuteQueryRaw(query string) (sql.Result, error) {
	return c.Conn.Exec(query)
}
