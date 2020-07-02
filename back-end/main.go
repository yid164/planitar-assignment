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

	a1 := Article{Name: "Ford Mustang", Content: "The Ford Mustang is a series of American automobiles manufactured by Ford."}
	articles = append(articles, a1)

	a2 := Article{Name: "Chevrolet Camaro", Content: "The Chevrolet Camaro is a mid-size American automobile manufactured by Chevrolet, classified as a pony car and some versions also as a muscle car"}
	articles = append(articles, a2)

	a3 := Article{Name: "Toyota 86", Content: "The Toyota 86 is a 2+2 sports car jointly developed by Toyota and Subaru, manufactured at Subaru's Gunma assembly plant along with a badge engineered variant, marketed as the Subaru BRZ."}
	articles = append(articles, a3)

	a4 := Article{Name: "Subaru WRX STI", Content: "Debuting at the Tokyo Auto Show in October 2007, WRX STI versions build further on the standard WRX cars"}
	articles = append(articles, a4)

	a5 := Article{Name: "Mitsubishi Lancer Evolution", Content: "The Mitsubishi Lancer Evolution, commonly referred to as 'Evo', is a sports sedan based on the Lancer that was manufactured by Japanese manufacturer Mitsubishi Motors from 1992 until 2016."}
	articles = append(articles, a5)

	a6 := Article{Name: "Honda Civic Type R", Content: "The Honda Civic Type R (Japanese: 本田・シビックタイプR, Honda Shibikku Taipuāru) is the high performance version of the Honda Civic compact car"}
	articles = append(articles, a6)

	a7 := Article{Name: "Volkswagen Golf R", Content: "Volkswagen Golf R is built as a three- or five-door hatchback. It is powered by a newly developed version of the 1,984 cc (2.0 L; 121.1 cu in) turbocharged EA888 petrol FSI Inline-four engine"}
	articles = append(articles, a7)

	a8 := Article{Name: "Mercedes-Benz A45 AMG", Content: "The A45 AMG is a performance version of the A-Class, fitted with a 2-litre four-cylinder turbo engine rated at 381 PS (280 kW; 376 hp) at 6000 rpm and 475 N⋅m (350 lbf⋅ft) at 2250-5000 rpm, AMG Speedshift DCT 7-speed sports transmission with Momentary M mode, 4MATIC all-wheel drive"}
	articles = append(articles, a8)

	a9 := Article{Name: "Nissan GT-R", Content: "The Nissan GT-R is a high-performance sports car and grand tourer produced by Nissan, which was unveiled in 2007. It is the successor to the Skyline GT-R, although no longer part of the Skyline range itself, that name now being used for Nissan's luxury-sport market."}
	articles = append(articles, a9)

	a0 := Article{Name: "Audi R8", Content: "The Audi R8 is a mid-engine, 2-seater sports car, which uses Audi's trademark quattro permanent all-wheel drive system. It was introduced by the German car manufacturer Audi AG in 2006."}
	articles = append(articles, a0)

	// Get handle function:
	router := mux.NewRouter()
	router.HandleFunc("/articles/", GetArticles).Methods("GET")
	router.HandleFunc("/articles/{name}", GetArticle).Methods("GET")

	// Put handle function
	router.HandleFunc("/articles/{name}", PutArticle).Methods("PUT")

	// start api on 9090
	log.Fatal(http.ListenAndServe(":9090", router))
}
