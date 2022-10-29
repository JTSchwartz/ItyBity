package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Err struct {
	Msg string `bson:"err" json:"err"`
}

var (
	db *mongo.Collection
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, ignore if running in production")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		os.Getenv("MONGODB_URI"),
	))
	if err != nil {
		log.Fatal("Database connection: ", err)
	}

	db = client.Database("ityBity").Collection("shortUrls")
	router := mux.NewRouter()

	router.Path("/robots.txt").Handler(http.FileServer(http.Dir("./client/dist")))
	router.HandleFunc("/{slug}", reroute).Methods("GET")
	router.HandleFunc("/change", change).Methods("POST")
	router.HandleFunc("/create", create).Methods("POST")
	router.HandleFunc("/remove", remove).Methods("POST")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./client/dist")))

	serverAddress := ":8080"
	log.Println("starting server at", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, router))
}

func change(w http.ResponseWriter, r *http.Request) {
	var api Api

	err := json.NewDecoder(r.Body).Decode(&api)
	if err != nil {
		log.Println(err.Error())
		JSONResponse(w, Err{Msg: "Error occurred while removing slug"}, http.StatusBadRequest)
		return
	}

	exists := slugExists(api.Slug)
	if !exists {
		log.Println("Slug not found")
		JSONResponse(w, Err{Msg: "Slug not found"}, http.StatusBadRequest)
		return
	}

	matches, err := secretMatchesSlug(api.Secret, api.Slug)
	if err != nil {
		log.Println(err.Error())
		JSONResponse(w, Err{Msg: "Error occurred while matching secret to slug"}, http.StatusBadRequest)
		return
	}

	if !matches {
		JSONResponse(w, Err{Msg: "Secret does not match slug"}, http.StatusBadRequest)
		return
	}

	err = updateDestination(api.Full, api.Slug)
	if err != nil {
		log.Println(err.Error())
		JSONResponse(w, Err{Msg: "Error occurred attempting to change destination"}, http.StatusBadRequest)
		return
	}

	JSONResponse(w, api, http.StatusOK)
}

func create(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering creation")
	var api Api

	err := json.NewDecoder(r.Body).Decode(&api)
	if err != nil {
		JSONResponse(w, Err{Msg: "Failed to create short link"}, http.StatusBadRequest)
		return
	}

	if api.Slug != "" {
		exists := slugExists(api.Slug)

		if exists {
			log.Println("Slug already exists")
			JSONResponse(w, Err{Msg: "Slug already exists"}, http.StatusBadRequest)
			return
		}
	} else {
		api.createSlug()

		exists := slugExists(api.Slug)
		for exists {
			api.createSlug()
		}
	}

	api.createSecret()

	err = createShortUrl(api)
	if err != nil {
		log.Println(err.Error())
		JSONResponse(w, Err{Msg: "Error occurred while creating short url"}, http.StatusBadRequest)
		return
	}

	JSONResponse(w, api, http.StatusOK)
}

func remove(w http.ResponseWriter, r *http.Request) {
	var api Api

	err := json.NewDecoder(r.Body).Decode(&api)
	if err != nil {
		log.Println(err.Error())
		JSONResponse(w, Err{Msg: "Error occurred while removing slug"}, http.StatusBadRequest)
		return
	}

	exists := slugExists(api.Slug)
	if !exists {
		log.Println("Slug not found")
		JSONResponse(w, Err{Msg: "Slug not found"}, http.StatusBadRequest)
		return
	}

	matches, err := secretMatchesSlug(api.Secret, api.Slug)
	if err != nil {
		log.Println(err.Error())
		JSONResponse(w, Err{Msg: "Error occurred while matching secret to slug"}, http.StatusBadRequest)
		return
	}

	if !matches {
		log.Println("Secret and slug does not exist")
		JSONResponse(w, Err{Msg: "Secret does not match slug"}, http.StatusBadRequest)
		return
	}

	err = removeSlug(api.Slug)
	if err != nil {
		log.Println(err.Error())
		JSONResponse(w, Err{Msg: "Error occurred while removing slug"}, http.StatusBadRequest)
		return
	}

	JSONResponse(w, api, http.StatusOK)
}

func reroute(w http.ResponseWriter, r *http.Request) {
	slug := mux.Vars(r)["slug"]
	full, err := getFullUrl(slug)

	if err != nil {
		log.Println("Redirect cannot be attempted")
		http.Redirect(w, r, "https://itybity.xyz", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, full, http.StatusPermanentRedirect)
}

// Database
func createShortUrl(new Api) error {
	_, err := db.InsertOne(context.Background(), new)
	return err
}

func getFullUrl(slug string) (string, error) {
	result := Api{}
	filter := bson.D{{"slug", slug}}
	err := db.FindOne(context.Background(), filter).Decode(&result)
	return result.Full, err
}

func secretMatchesSlug(secret string, slug string) (bool, error) {
	result := Api{}
	filter := bson.D{{"slug", slug}}
	err := db.FindOne(context.Background(), filter).Decode(&result)

	return result.Secret == secret, err
}

func removeSlug(slug string) error {
	filter := bson.D{{"slug", slug}}
	_, err := db.DeleteOne(context.Background(), filter)

	return err
}

func slugExists(slug string) bool {
	log.Println(slug)
	filter := bson.D{{"slug", slug}}
	result := db.FindOne(context.Background(), filter)

	return result.Err() != mongo.ErrNoDocuments
}

func updateDestination(dest string, slug string) error {
	filter := bson.D{{"slug", slug}}
	update := bson.D{{"$set", bson.D{{"full", dest}}}}
	_, err := db.UpdateOne(context.Background(), filter, update)

	return err
}

// JSON

func JSONResponse(w http.ResponseWriter, d interface{}, c int) {
	dj, err := json.Marshal(d)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(c)
	_, _ = fmt.Fprintf(w, "%s", dj)
}
