package mysql

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
)

func SetUpDB(db *sql.DB) error {
	b, err := os.ReadFile("./pkg/models/mysql/setup.sql")
	if err != nil {
		return err
	}
	reqs := strings.Split(string(b), ";")
	for _, req := range reqs {
		_, err := db.Exec(fmt.Sprintf("%s;", req))
		if err != nil {
			return err
		}
	}
	return nil
}

func CleanUpDB(db *sql.DB) error {
	b, err := os.ReadFile("./pkg/models/mysql/cleanup.sql")
	if err != nil {
		return err
	}
	reqs := strings.Split(string(b), ";")
	for _, req := range reqs {
		_, err := db.Exec(fmt.Sprintf("%s;", req))
		if err != nil {
			return err
		}
	}
	return nil
}
