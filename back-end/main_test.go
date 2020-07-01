package main

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

// Test Store an article in code 201
func TestPutCreated(t *testing.T) {

	url := "http://localhost:9090/articles/test0"
	method := "PUT"

	payload := strings.NewReader("")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Test Store an article (Created) with non-empty name and empty content")
	res, err := client.Do(req)
	if status := res.StatusCode; status != http.StatusCreated {
		fmt.Println("Error")
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

}

// Test Store an article in code 201
func TestPutCreated1(t *testing.T) {
	url := "http://localhost:9090/articles/test1"
	method := "PUT"

	payload := strings.NewReader("{\n    \"content\": \"this is test 1\"\n}")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Test Store an article (Created) with non-empty name and non-empty content")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if status := res.StatusCode; status != http.StatusCreated {
		fmt.Println("Error")
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}

// Test Store an article in code 405 error
func TestPutCreated2(t *testing.T) {
	url := "http://localhost:9090/articles/"
	method := "PUT"

	payload := strings.NewReader("")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Test Store an article (Created) with empty name")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if status := res.StatusCode; status != http.StatusMethodNotAllowed {
		fmt.Println("Error")
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}

// Test Store an article in code 200
func TestPutUpdate(t *testing.T) {
	url := "http://localhost:9090/articles/test0"
	method := "PUT"

	payload := strings.NewReader("{\n    \"content\": \"This is test 0\"\n}")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Test Store an article (Updated), update the content")

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if status := res.StatusCode; status != http.StatusOK {
		fmt.Println("Error")
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}

// Test List stored articles in code 200
func TestGetArticles(t *testing.T) {

	url := "http://localhost:9090/articles/"
	method := "GET"

	payload := strings.NewReader("")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Test Get All Articles")
	res, err := client.Do(req)
	if status := res.StatusCode; status != http.StatusOK {
		fmt.Println("Error")
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

}

// Test read an article in code 200
func TestGetArticle1(t *testing.T) {

	url := "http://localhost:9090/articles/test0"
	method := "GET"

	payload := strings.NewReader("")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Test Get valid Article 0")
	res, err := client.Do(req)
	if status := res.StatusCode; status != http.StatusOK {
		fmt.Println("Error")
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

}

// Test read an article in code 200
func TestGetArticle2(t *testing.T) {
	url := "http://localhost:9090/articles/test1"
	method := "GET"

	payload := strings.NewReader("")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Test Get valid Article 1")
	res, err := client.Do(req)
	if status := res.StatusCode; status != http.StatusOK {
		fmt.Println("Error")
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}

// Test read an article in code 404
func TestGetArticle3(t *testing.T) {
	url := "http://localhost:9090/articles/test"
	method := "GET"

	payload := strings.NewReader("")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Test Get unvalid Article")
	res, err := client.Do(req)
	if status := res.StatusCode; status != http.StatusNotFound {
		fmt.Println("Error")
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}
