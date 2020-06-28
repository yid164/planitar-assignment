package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

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
	w.WriteHeader(201)
	var article Article
	_ = json.NewDecoder(req.Body).Decode(&article)
	article.Name = params["name"]
	articles = append(articles, article)
	json.NewEncoder(w).Encode(article)
}

func main() {

	articles = append(articles, Article{
		Name:    "Ford Mustang",
		Content: "The Ford Mustang is a series of American automobiles manufactured by Ford.",
	})
	articles = append(articles, Article{
		Name:    "Chevrolet Camaro",
		Content: "The Chevrolet Camaro is a mid-size American automobile manufactured by Chevrolet.",
	})

	articles = append(articles, Article{
		Name:    "Dodge Challenger",
		Content: "The Dodge Challenger is the name of three different generations of automobiles (two of those being pony cars) produced by American automobile manufacturer Dodge.",
	})

	articles = append(articles, Article{
		Name:    "Nissan GT-R",
		Content: "The Nissan GT-R is a high-performance sports car and grand tourer produced by Nissan, which was unveiled in 2007. It is the successor to the Skyline GT-R, although no longer part of the Skyline range itself, that name now being used for Nissan's luxury-sport market.",
	})

	articles = append(articles, Article{
		Name:    "Mazda RX-7",
		Content: "The Mazda RX-7 is a front/mid-engine, rear-wheel-drive, rotary engine-powered sports car that was manufactured and marketed by Mazda from 1978 to 2002 across three generations, all of which made use of a compact, lightweight Wankel rotary engine",
	})

	// Get handle function:
	router := mux.NewRouter()
	router.HandleFunc("/articles/", GetArticles).Methods("GET")
	router.HandleFunc("/articles/{name}", GetArticle).Methods("GET")

	// Put handle function
	router.HandleFunc("/articles/{name}", PutArticle).Methods("PUT")

	// 启动 API端口9899
	log.Fatal(http.ListenAndServe(":9090", router))
}
