/*
 * -----------------------------------------------------------
 * File Created: 2020-05-08T13:57:02.247Z
 * Author: Manuel Bernal Llinares <mbdebian@gmail.com>
 * -----------------------------------------------------------
 */
package main

import (
	"log"
	"net/http"
	"os"
	"server"
	"api/health"
)

var (
	configServerPort = getEnvDefault("IDORG_RESOURCE_RECOMMENDER_CONFIG_SERVER_PORT", "8083")
	configLinkCheckerHost = getEnvDefault("IDORG_RESOURCE_RECOMMENDER_CONFIG_BACKEND_SERVICE_LINK_CHECKER_HOST", "localhost")
	configLinkCheckerPort = getEnvDefault("IDORG_RESOURCE_RECOMMENDER_CONFIG_BACKEND_SERVICE_LINK_CHECKER_PORT", "8084")
)

// Env helper
func getEnvDefault(key, defValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defValue
}

func main() {
	// Create service logger
	logger := log.New(os.Stdout, "[IDORG_RESOURCE_RECOMMENDER] ", log.LstdFlags | log.Lshortfile)

	// Bootstrap
	logger.Println("[BOOTSTRAP] identifiers.org Resource Recommender Service")
	mux := http.NewServeMux()
	// APIs SETUP
	health.NewApiHandler(logger).SetupRoutes(mux)
	logger.Println("[BOOTSTRAP] APIs registration completed")
	// Start the Server
	logger.Printf("[BOOTSTRAP] Service LISTENING on port %s\n", configServerPort)
	srv := server.New(mux, configServerPort)

	// Start the server (HTTP)
	err := srv.ListenAndServe()
	if err != nil {
		logger.Fatalf("Server failed to start - '%s'", err)
	}
}
