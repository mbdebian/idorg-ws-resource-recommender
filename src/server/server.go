/**
 * Project: idorg-ws-resource-recommender
 * Timestamp: 2020-05-26 02:31
 * Author Manuel Bernal Llinares <mbdebian@gmail.com>
 * ---
 */
package server

import (
	"net/http"
	"time"
)

func New(mux *http.ServeMux, serverPort string) *http.Server {
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

