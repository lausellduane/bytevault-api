package main

import (
    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"
    "net/http"
)

func NewRouter() *mux.Router{
    r := mux.NewRouter()
    r.HandleFunc("/api/v1/fragments", getFragmentsHandler).Methods("GET")
    // r.HandleFunc("/api/v1/codefragment", postCodeFragmentHandler).Methods("POST")
    return r
}

func NewHandler(router *mux.Router) http.Handler {
    headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Bearer", "Bearer ", "content-type", "Origin", "Accept"})
    originsOk := handlers.AllowedOrigins([]string{"http://localhost:8000", "http://localhost:3000"})
    methodsOk := handlers.AllowedMethods([]string{"GET"})
    credentialsOk := handlers.AllowCredentials()

    handler := handlers.CORS(originsOk, headersOk, methodsOk, credentialsOk)(router)
    return handler
}