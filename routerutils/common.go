package routerutils

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/possawang/go-service-lib-common/commonutils"
)

type Endpoint struct {
	Execution func(w http.ResponseWriter, r *http.Request)
	Method    string
	Mdw       func(r http.Handler) http.Handler
}

func StartingService(endpoints map[string]Endpoint, mdw func(h http.Handler) http.Handler) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	path := os.Getenv("CONTEXT.PATH")
	r := mux.NewRouter().PathPrefix(path).Subrouter()
	r.Use(mdw)
	var methods []string
	for url, endpoint := range endpoints {
		sr := r.PathPrefix(url).Subrouter()
		if endpoint.Mdw != nil {
			sr.Use(endpoint.Mdw)
		}
		sr.HandleFunc(url, endpoint.Execution).Methods(endpoint.Method)
		if !commonutils.ArrayExists(methods, func(v string) bool { return v == endpoint.Method }) {
			methods = append(methods, endpoint.Method)
		}
	}
	port, headers, origins := os.Getenv("SERVICE.PORT"), os.Getenv("ALLOWED.HEADERS"), os.Getenv("ORIGINS")
	log.Fatal(
		http.ListenAndServe(":"+port,
			handlers.CORS(
				handlers.ExposedHeaders(strings.Split(headers, ",")),
				handlers.AllowedMethods(methods),
				handlers.AllowedOrigins(strings.Split(origins, ",")))(r)))
	log.Println("Deployed on port " + port + " and with context path " + path)
}
