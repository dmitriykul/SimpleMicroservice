package transport

import (
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Router() http.Handler {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/health", health).Methods(http.MethodGet)
	s.HandleFunc("/arithmetic", arithmetic).Methods(http.MethodGet)
	return logMiddleWare(r)
}

func logMiddleWare(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"method":     r.Method,
			"url":        r.URL,
			"remoteAddr": r.RemoteAddr,
			"userAgent":  r.UserAgent(),
		}).Info("got a new request")
		h.ServeHTTP(w, r)
	})
}

func health(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "200")
}

func arithmetic(w http.ResponseWriter, _ *http.Request){

}
