package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" 
)

var Db *sql.DB

func InitDB(filepath string) error {
    var err error
    Db, err = sql.Open("sqlite3", filepath)
    if err != nil {
        return err
    }

    if err = Db.Ping(); err != nil {
        return err
    }

    return nil
}
