package main

import (
    "net/http"
    //"fmt"
)

func main() {
    router := NewRouter()
    handler := NewHandler(router)

    http.ListenAndServe(":8000", handler)
}