package services

import (
	"encoding/json"
	"net/http"
)

type ResponseReject struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func CORSMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "*")
		w.Header().Set("Access-Control-Max-Age", "86400")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, UPDATE, OPTIONS")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			APIErrorResponse(w, "", http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)

	})
}

// Wrapper function that returns an API error response (JSON only)
func APIErrorResponse(w http.ResponseWriter, err string, httpErrorCode int) {
	responseHandl := ResponseReject{}
	responseHandl.Code = httpErrorCode
	responseHandl.Message = err
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpErrorCode)
	json.NewEncoder(w).Encode(responseHandl)
}
