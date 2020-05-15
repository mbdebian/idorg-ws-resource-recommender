/*
 * -----------------------------------------------------------
 * File Created: 2020-05-08T13:57:02.247Z
 * Author: Manuel Bernal Llinares <mbdebian@gmail.com>
 * -----------------------------------------------------------
 */
package main

import (
	"greet"
	"log"
	"net/http"
)

func main() {
	log.Println("Another try with greeting - " + greet.MyGreet)
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("Hello to everybody"))
	})
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Server failed to start")
	}
}
