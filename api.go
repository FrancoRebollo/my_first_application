package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiResponse struct {
	Status string
}

type ApiError struct {
	Error string
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/sign-up", makeHTTPHandleFunc(s.handleUser))
	router.HandleFunc("/sign-in", makeHTTPHandleFunc(s.handleUser))

	log.Println("Users microservice is running on PORT: ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)

}

func (s *APIServer) handleUser(w http.ResponseWriter, r *http.Request) error {

	if r.Method == "GET" {
		if r.URL.Path == "/sign-up" {
			return s.handleLoginUser(w, r)
		}
	}

	if r.Method == "POST" {
		if r.URL.Path == "/sign-up" {
			return s.handleSingUpUser(w, r)
		}
	}

	return nil
}

func (s *APIServer) handleSingUpUser(w http.ResponseWriter, r *http.Request) error {
	if err := s.createUser(w, r); err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, ApiResponse{Status: "Signed up"})
}

func (s *APIServer) handleLoginUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}
