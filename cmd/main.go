package main

import (
	"SimpleMicroservice/pkg/transport"
	"net/http"
	"os"
)

type expression struct {
	exp string
}

const serverUrl = ":8000"

func main() {
	//log.SetFormatter(&log.JSONFormatter{})
	//log.SetOutput(os.Stdout)
	//log.WithFields(log.Fields{"url": serverUrl}).Info("starting the server")
	//r := transport.Router()
	//fmt.Println(http.ListenAndServe(serverUrl, r))

	port := os.Getenv("PORT")

	r := transport.Router()
	http.ListenAndServe(":"+port, r)

}
