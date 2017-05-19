package functions

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// HTTPAPI for function discovery
type HTTPAPI struct {
	Functions *Functions
}

// RegisterRoutes register HTTP API routes
func (h HTTPAPI) RegisterRoutes(router *httprouter.Router) {
	router.GET("/v0/api/function/:name", h.getFunction)
	router.POST("/v0/api/function", h.registerFunction)
}

func (h HTTPAPI) getFunction(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fn, err := h.Functions.GetFunction(params.ByName("name"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		encoder := json.NewEncoder(w)
		encoder.Encode(fn)
	}
}

func (h HTTPAPI) registerFunction(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fn := new(Function)
	dec := json.NewDecoder(r.Body)
	dec.Decode(fn)

	output, err := h.Functions.RegisterFunction(fn)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		encoder := json.NewEncoder(w)
		encoder.Encode(output)
	}
}