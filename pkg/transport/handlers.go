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
	s.HandleFunc("/health", health).Methods(http.MethodHead)
	s.HandleFunc("/arithmetic", arithmetic).Methods(http.MethodPost)
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

func arithmetic(w http.ResponseWriter, r *http.Request) {
	//str := r.Body
	//res := 0

	//str, err := ioutil.ReadAll(req.Body)
	//if err != nil {
	//	panic(err)
	//}
	//
	//if str == "2+2"{
	//	fmt.Fprintf(w, "4")
	//}
	//
	//fmt.Fprintf(w, res)
}
