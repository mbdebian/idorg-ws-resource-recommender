/**
 * Project: idorg-ws-resource-recommender
 * Timestamp: 2020-05-26 03:18
 * Author: Manuel Bernal Llinares <mbdebian@gmail.com>
 * ---
 */
package health

import (
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
)

// Context
type HandlerModelContext struct {
	logger *log.Logger
}

var (
	sessionId string
)

// --- INIT ---
func init() {
	if newId, error := uuid.NewRandom(); error != nil {
		sessionId = "error_no_session_id"
	} else {
		sessionId = fmt.Sprintf("%v", newId)
	}
}

// Constructor
func NewApiModel(logger *log.Logger) *HandlerModelContext {
	return &HandlerModelContext{logger: logger}
}

// --- API ---
func (handlerModelContext *HandlerModelContext) LivenessCheck(r *http.Request) string {
	// TODO
	return "TODO"
}

func (handlerModelContext *HandlerModelContext) ReadinessCheck(r *http.Request) string {
	// TODO
	return "TODO"
}
