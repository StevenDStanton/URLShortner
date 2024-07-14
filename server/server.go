package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/StevenDStanton/URLShortner/base68"
	turso "github.com/StevenDStanton/URLShortner/database"
)

var (
	db *turso.DBConnection
)

func init() {
	var err error
	db, err = turso.NewDBConnection()
	if err != nil {
		panic("No DB")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func NextURL(w http.ResponseWriter, r *http.Request) {
	index, err := db.GetLatestIndex()
	if err != nil {
		fmt.Println(fmt.Errorf("%w", err))
		os.Exit(1)
	}
	newIndex := base68.IncrementBase68String(index)
	db.UpdateLatestIndex(newIndex)
	w.WriteHeader(http.StatusOK)
}
func StartServer() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/AddURL", NextURL)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(fmt.Errorf("%w", err))
		os.Exit(1)
	}
}
