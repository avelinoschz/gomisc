package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct {
	DB Database
}

func main() {

	mux := http.NewServeMux()

	h := &Handler{
		DB: &MockDatabase{},
	}

	mux.HandleFunc("/", LoggerMiddleware(h.HomeHandler))
	mux.HandleFunc("/post", LoggerMiddleware(h.PostHandler))

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Endpoint called:", r.URL.Path)
		next.ServeHTTP(w, r)
	}
}

func (h *Handler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	name := params.Get("name")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(map[string]string{"message": "Hello, " + name})
	if err != nil {
		log.Println("err:", err.Error())
	}
}

func VerifyingMessage(message string) bool {
	return message == "Hello"
}

type PostPayload struct {
	Message string `json:"message"`
}

func (h *Handler) PostHandler(w http.ResponseWriter, r *http.Request) {
	var payload PostPayload
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.DB.Save("key", payload.Message)
	if err != nil {
		http.Error(w, "Failed to save message to database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type Database interface {
	Get(key string) (string, error)
	Save(key, value string) error
}

type MockDatabase struct {
	Data map[string]string
}

func (m *MockDatabase) Get(key string) (string, error) {
	if value, ok := m.Data[key]; ok {
		return value, nil
	}
	return "", nil
}

func (m *MockDatabase) Save(key, value string) error {
	m.Data[key] = value
	return nil
}
