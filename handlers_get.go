package main

import (
    // "fmt"
    "net/http"
    "encoding/json"
    "go.mongodb.org/mongo-driver/bson"
    "context"
    "log"
)

func getFragmentsHandler(w http.ResponseWriter, r *http.Request){
    collection := db.Collection("fragments")
    cur, err := collection.Find(ctx, bson.D{})
    if err != nil { log.Fatal(err) }
    defer cur.Close(context.Background())
    var fragments []bson.M
    if err = cur.All(ctx, &fragments); err != nil {
        log.Fatal(err)
    }

    json.NewEncoder(w).Encode(fragments)
    return
}