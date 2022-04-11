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

func StartingService(endpoints map[string]Endpoint) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()
	var methods []string
	for url, endpoint := range endpoints {
		r.HandleFunc(url, endpoint.Execution).Methods(endpoint.Method)
		if endpoint.Mdw != nil {
			http.Handle(url, endpoint.Mdw(r))
		}
		if !commonutils.ArrayExists(methods, func(v string) bool { return v == endpoint.Method }) {
			methods = append(methods, endpoint.Method)
		}
	}
	port, path, headers, origins := os.Getenv("SERVICE.PORT"), os.Getenv("CONTEXT.PATH"), os.Getenv("ALLOWED.HEADERS"), os.Getenv("ORIGINS")
	r.PathPrefix(path)
	log.Fatal(
		http.ListenAndServe(":"+port,
			handlers.CORS(
				handlers.ExposedHeaders(strings.Split(headers, ",")),
				handlers.AllowedMethods(methods),
				handlers.AllowedOrigins(strings.Split(origins, ",")))(r)))
	log.Println("Deployed on port " + port + " and with context path " + path)
}
