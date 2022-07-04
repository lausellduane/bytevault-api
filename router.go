package main

import (
    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"
    "net/http"
)

func NewRouter() *mux.Router{
    r := mux.NewRouter()
    // GET Requests
    r.HandleFunc("/api/v1/fragments", getFragmentsHandler).Methods("GET")
    r.HandleFunc("/api/v1/fragments/tags", getFragmentsTagsHandler).Methods("GET")
    // POST Requests
    r.HandleFunc("/api/v1/fragments", postCodeFragmentHandler).Methods("POST")
    // PATCH Requests

    // DELETE Requests
    r.HandleFunc("/api/v1/fragments", deleteCodeFragmentsHandler).Methods("DELETE")
    
    return r
}

func NewHandler(router *mux.Router) http.Handler {
    headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Bearer", "Bearer ", "content-type", "Origin", "Accept"})
    originsOk := handlers.AllowedOrigins([]string{"http://localhost:8000", "http://localhost:3000"})
    methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "DELETE"})
    credentialsOk := handlers.AllowCredentials()

    handler := handlers.CORS(originsOk, headersOk, methodsOk, credentialsOk)(router)
    return handler
}