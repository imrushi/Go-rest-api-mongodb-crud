package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/imrushi/restapi/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Collection {
	config, err := util.LoadConfig("./")
	if err != nil {
		log.Fatal("Cannot load Config:", err)
	}
	//set client opetions
	clientOptions := options.Client().ApplyURI(config.MONGO_URI)

	//connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Connected to MongoDB!")

	collection := client.Database(config.MONGO_DATABASE).Collection(config.MONGO_COLLECTION)

	return collection
}

//ErrorResponse : This is error model
type ErrorResponse struct {
	StatusCode  int    `json:"status"`
	ErrorMesage string `json:"message"`
}

// GetError : This is helper function to prepare error model.
// If you want to export your function. You must to start upper case function name. Otherwise you won't see your function when you import that on other class.
func GetError(err error, w http.ResponseWriter) {
	log.Fatal(err.Error())
	var response = ErrorResponse{
		ErrorMesage: err.Error(),
		StatusCode:  http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}
