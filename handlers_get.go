package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo/options"
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
        http.Error(w, fmt.Sprintf("%v", err), 500)
		return
    }

    json.NewEncoder(w).Encode(fragments)
    return
}

func getFragmentsTagsHandler(w http.ResponseWriter, r *http.Request){
    collection := db.Collection("tags")
    opts := options.Find().SetProjection(bson.D{{"_id", 0}})
    cur, err := collection.Find(ctx, bson.D{}, opts)
    
    if err != nil { 
        http.Error(w, fmt.Sprintf("%v", err), 500)
        return 
    }
    defer cur.Close(context.Background())
    var tags []bson.M
    if err = cur.All(ctx, &tags); err != nil {
        http.Error(w, fmt.Sprintf("%v", err), 500)
		return
    }

    json.NewEncoder(w).Encode(tags)
    return
}

func getProgrammingLanguagesHandler(w http.ResponseWriter, r *http.Request){
    collection := db.Collection("programmingLanguages")
    opts := options.Find().SetProjection(bson.D{{"_id", 0}})
    cur, err := collection.Find(ctx, bson.D{}, opts)

    if err != nil { 
        http.Error(w, fmt.Sprintf("%v", err), 500)
        return 
    }
    defer cur.Close(context.Background())
    var languages []bson.M
    if err = cur.All(ctx, &languages); err != nil {
        http.Error(w, fmt.Sprintf("%v", err), 500)
		return
    }

    json.NewEncoder(w).Encode(languages)
    return
}