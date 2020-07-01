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
	Name    string `json:"name,omitempty"`
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
	for _, item := range articles {
		if item.Name == params["name"] {
			w.WriteHeader(200)
			var article Article
			_ = json.NewDecoder(req.Body).Decode(&article)
			article.Name = params["name"]
			item = article
			fmt.Println(item)
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	var article Article
	_ = json.NewDecoder(req.Body).Decode(&article)
	article.Name = params["name"]
	if len(article.Name) > 0 {
		w.WriteHeader(201)
		articles = append(articles, article)
		json.NewEncoder(w).Encode(article)
	} else {
		w.WriteHeader(400)
		return
	}
}

// run
func main() {

	articles = append(articles, Article{
		Name:    "Ford Mustang",
		Content: "The Ford Mustang is a series of American automobiles manufactured by Ford.",
	})
	articles = append(articles, Article{
		Name:    "Chevrolet Camaro",
		Content: "The Chevrolet Camaro is a mid-size American automobile manufactured by Chevrolet.",
	})

	// Get handle function:
	router := mux.NewRouter()
	router.HandleFunc("/articles/", GetArticles).Methods("GET")
	router.HandleFunc("/articles/{name}", GetArticle).Methods("GET")

	// Put handle function
	router.HandleFunc("/articles/{name}", PutArticle).Methods("PUT")

	// start api on 9090
	log.Fatal(http.ListenAndServe(":9090", router))
}
