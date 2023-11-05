package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	host := "http://localhost:8080"
	// use http request to test the api local host 8080
	// try post request to create a todo item
	// url: http://localhost:8080/create/todo
	if res, err := http.Post(host+"/create/todo", "application/json", bytes.NewReader([]byte(`{"aid":"aid","title":"title","completed":false}`))); err != nil {
		panic(err)
	} else {
		// print the response body to string
		buf := new(bytes.Buffer)
		_, err := buf.ReadFrom(res.Body)
		if err != nil {
			return
		}
		fmt.Println(buf.String())
	}
	// try get request to get all todo items
	// url: http://localhost:8080/get/todo?aid=aid
	if res, err := http.Get(host + "/get/todo?aid=b51eb899-ca9d-4b89-9684-50c8c5d12dd7"); err != nil {
		panic(err)
	} else {
		// print the response body to string
		buf := new(bytes.Buffer)
		_, err := buf.ReadFrom(res.Body)
		if err != nil {
			return
		}
		fmt.Println(buf.String())
	}
	// try post request to update a todo item
	// url: http://localhost:8080/update/todo
	if res, err := http.Post(host+"/update/todo", "application/json", bytes.NewReader([]byte(`{"_id":"65290f26d2f37cc8d81d885f","completed":false}`))); err != nil {
		panic(err)
	} else {
		// print the response body to string
		buf := new(bytes.Buffer)
		_, err := buf.ReadFrom(res.Body)
		if err != nil {
			return
		}
		fmt.Println(buf.String())
	}
	// try post request to delete a todo item
	// url: http://localhost:8080/delete/todo
	if res, err := http.Post(host+"/delete/todo", "application/json", bytes.NewReader([]byte(`{"_id":"65290f26d2f37cc8d81d885f"}`))); err != nil {
		panic(err)
	} else {
		// print the response body to string
		buf := new(bytes.Buffer)
		_, err := buf.ReadFrom(res.Body)
		if err != nil {
			return
		}
		fmt.Println(buf.String())
	}

}
