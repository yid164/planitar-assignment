package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// structure of wiki
type Article struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

var articles []Article

// Get all articles
func GetArticles(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(articles)
}

// Get article by name
func GetArticle(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range articles {
		if item.Name == params["name"] {
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "text/html; charset=UTF-8")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(404)
	return
}

// Put an article
func PutArticle(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var isIncluded = false
	var count = 0
	for i, item := range articles {
		if item.Name == params["name"] {

			isIncluded = true
			count = i
			fmt.Println(count)
		}
	}

	if isIncluded {
		w.WriteHeader(200)
		var article Article
		_ = json.NewDecoder(req.Body).Decode(&article)
		article.Name = params["name"]
		articles[count] = article
		json.NewEncoder(w).Encode(article)

	} else {
		var article Article
		_ = json.NewDecoder(req.Body).Decode(&article)
		article.Name = params["name"]
		if len(article.Name) > 0 {
			w.WriteHeader(201)
			articles = append(articles, article)
			json.NewEncoder(w).Encode(article)

		} else {
			w.WriteHeader(400)

		}
	}

}

// run
func main() {

	// Get handle function:
	router := mux.NewRouter()
	router.HandleFunc("/articles/", GetArticles).Methods("GET")
	router.HandleFunc("/articles/{name}", GetArticle).Methods("GET")

	// Put handle function
	router.HandleFunc("/articles/{name}", PutArticle).Methods("PUT")

	// start api on 9090
	log.Fatal(http.ListenAndServe(":9090", router))
}
