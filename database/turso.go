package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type DBConnection struct {
	DB *sql.DB
}

var (
	tursoURL   string
	tursoToken string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Unable to load .env file")
	}

	tursoURL = os.Getenv("TURSO_DATABASE_URL")
	tursoToken = os.Getenv("TURSO_AUTH_TOKEN")
}

func NewDBConnection() (*DBConnection, error) {

	url := fmt.Sprintf("libsql://%s?authToken=%s", tursoURL, tursoToken)
	db, err := sql.Open("libsql", url)

	if err != nil {
		fmt.Println(fmt.Errorf("%w", err))
		return nil, err
	}

	if err := db.Ping(); err != nil {
		fmt.Println(fmt.Errorf("%w", err))
		return nil, err
	}

	return &DBConnection{
		DB: db,
	}, nil
}

func (dbc *DBConnection) GetURL(indexKey string) (string, error) {
	var url string
	query := "SELECT url FROM url_map WHERE index_key = ?"
	err := dbc.DB.QueryRow(query, indexKey).Scan(&url)
	return url, err
}

func (dbc *DBConnection) PutURL(indexKey string, url string) error {
	query := `INSERT INTO url_map (index_key, url) 
				VALUES(?, ?)
				ON CONFLICT(index_key)
				DO UPDATE SET url = excluded.url`
	_, err := dbc.DB.Exec(query, indexKey, url)

	return err
}

func (dbc *DBConnection) GetLatestIndex() (string, error) {
	indexKey := "-1"
	query := "SELECT value from latest_record WHERE id = 0"
	err := dbc.DB.QueryRow(query).Scan(&indexKey)
	return indexKey, err
}

func (dbc *DBConnection) UpdateLatestIndex(value string) error {
	query := "UPDATE latest_record SET value = ? WHERE id = 0"
	_, err := dbc.DB.Exec(query, value)
	return err
}

func (dbc DBConnection) Close() error {
	return dbc.DB.Close()
}
