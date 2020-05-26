/**
 * Project: idorg-ws-resource-recommender
 * Timestamp: 2020-05-26 02:37
 * Author Manuel Bernal Llinares <mbdebian@gmail.com>
 * ---
 */
package health

import (
	"fmt"
	"log"
	"net/http"
)

const (
	apiPrefix = "healthApi"
)

type HandlerContext struct {
	logger *log.Logger
}

// --- API ---
func (handlerContext *HandlerContext) livenessCheck(writer http.ResponseWriter, r *http.Request) {
	// TODO
	writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	writer.WriteHeader(http.StatusOK)
	//writer.Write([]byte("liveness check endpoint"))
	response := NewApiModel(handlerContext.logger).LivenessCheck(r)
	writer.Write([]byte(response))
}

func (handlerContext *HandlerContext) readinessCheck(writer http.ResponseWriter, r *http.Request) {
	// TODO
	writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	writer.WriteHeader(http.StatusOK)
	//writer.Write([]byte("readiness check endpoint"))
	response := NewApiModel(handlerContext.logger).ReadinessCheck(r)
	writer.Write([]byte(response))

}
// END --- API

// --- Routing ---
func (handlerContext *HandlerContext) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc(fmt.Sprintf("/%s/liveness_check", apiPrefix), handlerContext.livenessCheck)
	mux.HandleFunc(fmt.Sprintf("/%s/readiness_check", apiPrefix), handlerContext.readinessCheck)
}

// --- Constructor ---
func NewApiHandler(logger *log.Logger) *HandlerContext {
	return &HandlerContext{
		logger: logger,
	}
}
