package main

import (
	"fmt"
	"encoding/json"
	// "errors"
	"net/http"
	// "strconv"
	// "strings"
	// "context"
	"io/ioutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func deleteCodeFragmentsHandler(w http.ResponseWriter, r *http.Request) {
	collection := db.Collection("fragments")
	// Parse body
	bodyData, err := ioutil.ReadAll(r.Body)
	if err != nil{
		http.Error(w, fmt.Sprintf("%v", err), 500)
		return
	}
	var body map[string]interface{}
	json.Unmarshal(bodyData, &body)
	str := fmt.Sprint(body["id"])

	idPrimitive, err := primitive.ObjectIDFromHex(str)
	if err != nil {
		http.Error(w, fmt.Sprintf("primitive.ObjectIDFromHex ERROR: %s", err), 500)
		return
	} 

	// Call the DeleteOne() method by passing BSON
	res, err := collection.DeleteOne(ctx, bson.M{"_id": idPrimitive})
	
	if err != nil {
		http.Error(w, fmt.Sprintf("DeleteOne() ERROR: %s", err), 500)
		return
	} else {
		if res.DeletedCount == 0 {
			http.Error(w, "not found", 404)
			return
		} 
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(idPrimitive)
	return
}

