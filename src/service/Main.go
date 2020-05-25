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
	"os"
	"time"
)

var (
	configServerPort = getEnvDefault("IDORG_RESOURCE_RECOMMENDER_CONFIG_SERVER_PORT", "8080")

)

// Env helper
func getEnvDefault(key, defValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defValue
}

func main() {
	log.Println("Another try with greeting - " + greet.MyGreet)
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("Hello to everybody"))
	})

	srv := newServer(mux, configServerPort)

	// Start the server (HTTP)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed to start - '%s'", err)
	}
}

func newServer(mux *http.ServeMux, serverPort string) *http.Server {
	// Custom Server
	srv := &http.Server{
		Addr:         ":" + serverPort,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	return srv
}
