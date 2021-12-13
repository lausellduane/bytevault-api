package main

// import (
// 	"fmt"
// 	"encoding/json"
// 	// "errors"
// 	"net/http"
// 	// "strconv"
// 	// "strings"
// 	// "context"
// 	"io/ioutil"
// 	"go.mongodb.org/mongo-driver/bson"
// )

// func postCodeFragmentHandler(w http.ResponseWriter, r *http.Request){
// 	fmt.Println("postCodeFragmentHandler")
// 	fmt.Println("r.Body: ", r.Body)
// 	bodyData, err := ioutil.ReadAll(r.Body)
// 	fmt.Println("bodyData: ", string(bodyData))
// 	if err != nil{
// 		fmt.Sprintf("%v", err)
// 	}
// 	// var body map[string]interface{}
// 	var codefragment CodeFragment
// 	// if err := json.Unmarshal(bodyData, &body); err != nil {
//     //     fmt.Sprintf("json unmarshal err: %v", err)
//     // }
// 	if err := json.Unmarshal(bodyData, &codefragment); err != nil {
// 		fmt.Println("json unmarshal err: ", err)
// 		http.Error(w, http.StatusText(500), 500)
// 		return
//     }
//     fmt.Println("codefragment: ", codefragment)
// 	fmt.Println("codefragment.Title: ", codefragment.Title)
// 	// fmt.Println("codefragment.Title string: ", string(codefragment.Title))
// 	// fmt.Println("body: ", body)

// 	codeFragmentResult, err := init().Database("byteVaultDB").database.Collection("codefragments").InsertOne(ctx, bson.M{
// 		"_id": codefragment.ID, "title": codefragment.Title, "description": codefragment.Description,
// 		"notes": codefragment.Notes, "value": codefragment.Value, "tags": codefragment.Tags,
// 		"language": codefragment.Language})
// 	if err != nil{
// 		fmt.Println("codeFragmentResult err: ", err)
// 		http.Error(w, http.StatusText(500), 500)
// 		return
// 	}
// 	fmt.Println("codeFragmentResult: ", codeFragmentResult)
// }
