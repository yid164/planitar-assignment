# Planitar Home Assignment For the Interview

## Part I

Build a web server implementing the API described below to run a wiki. In this case, a wiki is
  basically a key-value storage. You can store any content under a certain name. Later you
  can retrieve the latest stored content for a particular name.

#### Solution: 

I implemented the GET, PUT API on the port 9090

#### Testing: 

I tested the code on the Postman, it returns all the code

#### How to run it

`go run main.go`

##### How to test it
```go test ```


## Part II

Build a single page application for managing a Markdown-based wiki. The app should use the web server written in Part I as a store for the articles.

#### Solution: 

1. Vue CLI 
2. Vue Router + Vuetify
3. Axios

#### How to run it
``npm install``
```npm run serve```
