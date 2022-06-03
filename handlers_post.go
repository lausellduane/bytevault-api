package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"go.mongodb.org/mongo-driver/bson"
)

func postCodeFragmentHandler(w http.ResponseWriter, r *http.Request){
	collection := db.Collection("fragments")
	// Parse body
	bodyData, err := ioutil.ReadAll(r.Body)
	if err != nil{
		http.Error(w, fmt.Sprintf("%v", err), 500)
		return
	}
	var fragment CodeFragment
	if err := json.Unmarshal(bodyData, &fragment); err != nil {
		http.Error(w, fmt.Sprintf("%v", err), 500)
		return
    }

	fragmentResult, err := collection.InsertOne(ctx, bson.D{
		{"title", fragment.Title},
		{"description", fragment.Description},
		{"notes", fragment.Notes},
		{"value", fragment.Value},
		{"tags", fragment.Tags},
		{"language", fragment.Language},
	})

	if err != nil {
		http.Error(w, fmt.Sprintf("InsertOne() ERROR: %s", err), 500)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fragmentResult)
}
