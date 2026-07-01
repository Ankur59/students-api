package sqlite

import (
	"database/sql"

	"github.com/ankur59/students-api/internal/config"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	Db *sql.DB
}

// func Run(cfg *config.Config) (*Sqlite, error) {
// 	db, err := sql.Open("sqlite3", cfg.StoragePath)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// db.Exec(`CREATE TABLE IF NOT EXISTS students(
// 	// 	id INT PRIMARY KEY AUTOINCREMENT,
// 	// 	name TEXT NOT NULL,
// 	// 	age INTEGER,
// 	// 	email TEXT
// 	// )`)
// }
