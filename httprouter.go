package main

import (
    "fmt"
    "encoding/json"
    "net/http"
    "github.com/julienschmidt/httprouter"
)

type Request struct {
     name string
}

type Response struct {
     message string
}

func Hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
    fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
}

func HelloPost(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
    
    decoder := json.NewDecoder(req.Body)
    var reqst Request
    err := decoder.Decode(&reqst)
	if err != nil {
		panic("Error in decoding the JSON")
	}
    response := Response{message: "Hello, " + reqst.name + " !"}
    json.NewEncoder(rw).Encode(response)
    
}

func main() {
    mux := httprouter.New()
    
    mux.GET("/Hello/:name", Hello)
    mux.POST("/Hello/", HelloPost)
    
    server := http.Server{
            Addr: "0.0.0.0:8080",
            Handler: mux,
    }
    server.ListenAndServe()
}
