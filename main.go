package main

import (
	"log"
	"net/http"
	"context"
	"strings"
	"encoding/json"
	"fmt"

	"github.com/gorilla/mux"
	"strconv"
)
import "api-restful-cake-store/service/query"
import "api-restful-cake-store/service/migration"
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
    // NOTE: See we’re using = to assign the global var
    // instead of := which would assign it only in this function
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
    migration.CreatTable(ctx)

	fmt.Println("API Running...")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", defaults)
	router.HandleFunc("/cakes", CakeService)
	router.HandleFunc("/cakes/{id}", CakeService)
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
	//Dynamic ID Method
	if id != "" && id !="/cakes"  { 
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		var cke models.Cake
		cke.ID, _ = strconv.Atoi(id)
		if r.Method == "DELETE" {

			if err := cakes_query.Delete(ctx, cke); err != nil {
				resp := utils.ResponseAction{Status: "400", Description: "Delete Gagal"}
				utils.ResponseJSON(w, resp, http.StatusOK)
				return 
			}

			resp := utils.ResponseAction{Status: "200", Description: "Delete Succesfully"}
			utils.ResponseJSON(w, resp, http.StatusOK)
			return 
			
		}else if r.Method == "GET"{

			cakes, err := cakes_query.GetDetail(ctx, cke)
			if err != nil {
				fmt.Println(err)
			}
			resp := utils.Response{Status: "200", Data: cakes}
			utils.ResponseJSON(w, resp, http.StatusOK)
			return
			
		}else if r.Method == "PATCH"{

			if err := json.NewDecoder(r.Body).Decode(&cke); err != nil {
				utils.ResponseJSON(w, err, http.StatusBadRequest)
				return
			}

			if err := cakes_query.Update(ctx, cke); err != nil {
				resp := utils.ResponseAction{Status: "400", Description: "Update Gagal "}
				utils.ResponseJSON(w, resp, http.StatusOK)
				return 
			}

			resp := utils.ResponseAction{Status: "200", Description: "Update Succesfully"}
			utils.ResponseJSON(w, resp, http.StatusOK)
			return 
		}
	//Get List
	}else if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		cakes, err := cakes_query.GetAll(ctx)
		if err != nil {
			fmt.Println(err)
		}
		resp := utils.ResponseList{Status: "200", Data: cakes}
		utils.ResponseJSON(w, resp, http.StatusOK)
		return
	//Create / Insert Data
	}else if r.Method == "POST" {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		var cke models.Cake
		if err := json.NewDecoder(r.Body).Decode(&cke); err != nil {
			resp := utils.ResponseAction{Status: "400", Description: "Insert Gagal "}
			utils.ResponseJSON(w, resp, http.StatusOK)
			return 
		}
		if err := cakes_query.Insert(ctx, cke); err != nil {
			resp := utils.ResponseAction{Status: "400", Description: "Insert Gagal "}
			utils.ResponseJSON(w, resp, http.StatusOK)
			return 
		}
		resp := utils.ResponseAction{Status: "200", Description: "Insert Succesfully"}
		utils.ResponseJSON(w, resp, http.StatusOK)
		return 
	}

	http.Error(w, "Tidak di ijinkan", http.StatusNotFound)
	return
}