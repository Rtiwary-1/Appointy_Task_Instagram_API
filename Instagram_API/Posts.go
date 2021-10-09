package main

import(
	"context"
	"net/http"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
    //"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
	"github.com/gorilla/mux"
)

type Posts struct {
	ID primitive.ObjectID `json:"Id,omitempty"`
	Caption string `json:"caption"`
	Image_Url string `json:"image_url,omitempty"`
	Timestamp primitive.Timestamp `json:"posted_time,omitempty"`
	Post_User primitive.ObjectID  `json:"post_user,omitempty"`
}


var collection_posts = client.Database("thepolyglotdeveloper").Collection("Posts")
var ctx_posts, _ = context.WithTimeout(context.Background(), 10*time.Second)



func CreatePostEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var post Posts
	json.NewDecoder(request.Body).Decode(&post)
	
	result, _ := collection_posts.InsertOne(ctx_posts, post)
	json.NewEncoder(response).Encode(result)
}



func GetPostEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["Id"])
	var post Posts
	err := collection_posts.FindOne(ctx_posts, Posts{ID: id}).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(post)
}


func GetAllPostEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["post_user"])
	var all_post_by_user []Posts
	cursor, err := collection_posts.Find(ctx_posts, bson.M{"Post_User":id})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx_posts)
	for cursor.Next(ctx_posts) {
		var post Posts
		cursor.Decode(&post)
		all_post_by_user = append(all_post_by_user, post)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(all_post_by_user)
}