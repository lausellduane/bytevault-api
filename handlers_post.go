package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"go.mongodb.org/mongo-driver/bson"
)

func postCodeFragmentHandler(w http.ResponseWriter, r *http.Request){
	collectionFragments := db.Collection("fragments")
	collectionTags := db.Collection("tags")
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

	// Handle tags
	var newRawTags Tags
	var newTagIndex int64 = 0
	for _, v := range fragment.Tags {
		if (v.ID == 0){
			newRawTags = append(newRawTags, v)
		}
		if (newTagIndex < v.ID){
			newTagIndex = v.ID
		}
	}
	
	var newTags []interface{}
	var counter int64 = int64(len(newRawTags)) + newTagIndex
	for i := newTagIndex; i < counter; i++ {
		var newTag = bson.D{
			{"id", int(i + int64(1))}, 
			{"label", newRawTags[i - newTagIndex].Label}}
		newTags = append(newTags, newTag)
	}

	if(len(newTags) > 0){
		_, err = collectionTags.InsertMany(ctx, newTags)
		if err != nil {
			http.Error(w, fmt.Sprintf("InsertMany() ERROR: %s", err), 500)
			return
		}
	}

	fragmentResult, err := collectionFragments.InsertOne(ctx, bson.D{
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
