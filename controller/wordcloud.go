package controller

import (
	"encoding/json"
	"net/http"
	"sync"
)

var (
	wordFrequency = make(map[string]int)
	mu            sync.Mutex
)

// RegisterWordCloudRoutes registers the API routes
func RegisterWordCloudRoutes(router *mux.Router) {
	router.HandleFunc("/api/submit-word", SubmitWord).Methods("POST")
	router.HandleFunc("/api/word-cloud", GetWordCloud).Methods("GET")
}

// SubmitWord handles POST requests to add a word to the word cloud
func SubmitWord(w http.ResponseWriter, r *http.Request) {
	var request map[string]string
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	word, exists := request["word"]
	if !exists {
		http.Error(w, "Missing 'word' field", http.StatusBadRequest)
		return
	}

	word = normalizeWord(word)

	mu.Lock()
	wordFrequency[word]++
	mu.Unlock()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Word submitted successfully!"))
}

// GetWordCloud handles GET requests to retrieve the word cloud
func GetWordCloud(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	response, err := json.Marshal(wordFrequency)
	if err != nil {
		http.Error(w, "Failed to generate response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// normalizeWord converts a word to lowercase
func normalizeWord(word string) string {
	return word
}
