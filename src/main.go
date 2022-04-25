package main

import (
	"os"
    "fmt"
    "net/http"
	"strconv"
	"encoding/json"
)

const MAX_CLIPS_LIMIT int = 1000

func getIntParam(req *http.Request, key string, defaultValue int) (int, error) {
	param := req.URL.Query().Get(key)

	if param != "" {
		val, err := strconv.Atoi(param)

		if err != nil {
			return 0, err
		}

		return val, nil
	}

	return defaultValue, nil
}

func getClips(w http.ResponseWriter, req *http.Request) {

	limit, err := getIntParam(req, "limit", 200)

	if err != nil || limit < 0 {
		w.WriteHeader(400)
		return
	}

	page, err := getIntParam(req, "page", 1)

	if err != nil || limit < 0 {
		w.WriteHeader(400)
		return
	}

	fromTimestamp, err := getIntParam(req, "fromTimestamp", 0)

	if err != nil || limit < 0 {
		w.WriteHeader(400)
		return
	}

	if limit > MAX_CLIPS_LIMIT {
		limit = MAX_CLIPS_LIMIT
	}

	result := getVideosClips(page, limit, fromTimestamp)
	jsonString, err := json.Marshal(result)
	
	if err != nil {
		w.WriteHeader(500)
		return
	}

    fmt.Fprintf(w, string(jsonString))
}

func main() {

    http.HandleFunc("/getClips", getClips)

	port, ok := os.LookupEnv("PORT")

    if !ok {
        port = "8090"
    }

	fmt.Printf("Starting server at: localhost:%s\n", port)

    http.ListenAndServe(":" + port, nil)
}