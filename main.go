package main

import (
	"log"
	"net/http"
	"context"
	"strings"
	"encoding/json"
	"fmt"

	"github.com/gorilla/mux"
)
import "api-restful-cake-store/service/query"
import "api-restful-cake-store/config"
import "api-restful-cake-store/utils"
import "api-restful-cake-store/models"

func main() {

	db, e := config.MySQL()

	if e != nil {
		log.Fatal(e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error()) 
	}

	fmt.Println("API Running...")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", defaults)
	router.HandleFunc("/cakes", CakeService).Methods("GET")
	router.HandleFunc("/cakes", CakeService).Methods("POST")
	router.HandleFunc("/cakes/{id}", CakeService).Methods("GET")
	err := http.ListenAndServe(":7000", router)

	if err != nil {
		log.Fatal(err)
	}
}

func defaults(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API Cake.")
}

// Cakes Service
func CakeService(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/cakes/")
	if id != "" { 
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		cakes, err := cakes_query.GetDetail(ctx, id)
		if err != nil {
			fmt.Println(err)
		}
		utils.ResponseJSON(w, cakes, http.StatusOK)
		return
	}
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		cakes, err := cakes_query.GetAll(ctx)
		if err != nil {
			fmt.Println(err)
		}
		utils.ResponseJSON(w, cakes, http.StatusOK)
		return
	}else if r.Method == "POST" {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		var cke models.Cake
		if err := json.NewDecoder(r.Body).Decode(&cke); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}
		if err := cakes_query.Insert(ctx, cke); err != nil {
			utils.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}
		res := map[string]string{
			"status": "Succesfully",
		}
		utils.ResponseJSON(w, res, http.StatusCreated)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusNotFound)
	return
}