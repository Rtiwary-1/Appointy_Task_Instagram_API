package main

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"net/http"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	//"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"

	"github.com/gorilla/mux"
)

// This struct refers to the users
type Users struct {
	ID       primitive.ObjectID `json:"Id,omitempty"`
	Name     string             `json:"name,omitempty"`
	EmailID  string             `json:"email,omitempty"`
	Password string             `json:"password,omitempty"`
}

// Declaration of database interface connection
var client, _ = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

var collection_users = client.Database("thepolyglotdeveloper").Collection("Users")
var ctx_users, _ = context.WithTimeout(context.Background(), 10*time.Second)

// POST request for user
func CreateUserEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var user Users
	json.NewDecoder(request.Body).Decode(&user)

	result, _ := collection_users.InsertOne(ctx_users, user)
	json.NewEncoder(response).Encode(result)
}

// Get request for user
func GetUserEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var user Users
	err := collection_users.FindOne(ctx_users, Users{ID: id}).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(user)
}

func PasswordHash(password string) [32]uint8 {
	return sha256.Sum256([]byte(password))
}


