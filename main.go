package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

var serverStartTime time.Time

func main() {
	serverStartTime = time.Now()
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		payload := map[string]string{
			"status": "ok",
			"uptime": formatTime(time.Since(serverStartTime)),
		}
		respondWithJSON(w, http.StatusOK, payload)
	})

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		statusCode = http.StatusInternalServerError
		payload = map[string]string{"error": err.Error()}
		response, _ = json.Marshal(payload)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func formatTime(d time.Duration) string {
	days := d / (24 * time.Hour)
	d %= 24 * time.Hour
	hours := d / time.Hour
	d %= time.Hour
	minutes := d / time.Minute
	d %= time.Minute
	seconds := d / time.Second

	return fmt.Sprintf("%dd %02dh %02dm %02ds", days, hours, minutes, seconds)
}
