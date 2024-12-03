package model

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	))
	if err != nil {
		log.Fatal(err)
	}

	// Create table if not exists
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS videos (
			bv_number VARCHAR(36) PRIMARY KEY,
			file_path VARCHAR(255) NOT NULL,
			info JSON NOT NULL
		)
	`)
	if err != nil {
		log.Fatal(err)
	}
}

func SaveVideo(video *Video) error {
	infoJSON, err := json.Marshal(video.Info)
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO videos (bv_number, file_path, info) VALUES (?, ?, ?)",
		video.BVNumber, video.FilePath, infoJSON)
	return err
}
