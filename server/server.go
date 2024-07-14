package server

import (
	"encoding/json"
	"fmt"
	"io"
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
		panic("Unable to create DB Connection")
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	if !checkType(r, http.MethodGet) {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "up"})
}

func getUrl(w http.ResponseWriter, r *http.Request) {
	if !checkType(r, http.MethodGet) {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	key := r.URL.Path[1:]
	url, err := db.GetURL(key)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}

func putUrl(w http.ResponseWriter, r *http.Request) {
	if !checkType(r, http.MethodPut) {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var requestData map[string]string
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	url, exists := requestData["url"]
	if !exists || url == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	shortURL := nextUrl()

	err = db.PutURL(shortURL, url)
	if err != nil {
		http.Error(w, "Unable to save URL", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"short_url": fmt.Sprintf("m.wxs.us/%s", shortURL),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func nextUrl() string {
	index, err := db.GetLatestIndex()
	if err != nil {
		fmt.Println(fmt.Errorf("%w", err))
		os.Exit(1)
	}
	newIndex := base68.IncrementBase68String(index)
	db.UpdateLatestIndex(newIndex)
	return newIndex
}
func StartServer() {
	http.HandleFunc("/healthCheck", healthCheck)
	http.HandleFunc("/", getUrl)
	http.HandleFunc("/shorten", putUrl)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(fmt.Errorf("%w", err))
		os.Exit(1)
	}
}

func checkType(r *http.Request, allowedMethods ...string) bool {
	allowed := map[string]bool{}
	for _, method := range allowedMethods {
		allowed[method] = true
	}
	return allowed[r.Method]
}
