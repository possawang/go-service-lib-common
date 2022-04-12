package errorutils

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/possawang/go-service-lib-common/domain"
)

func Handle4xx(w http.ResponseWriter, msg string, code int, status string) {
	log.Printf("Invalidate %s\n", msg)
	response := domain.BaseResponse[any]{
		Data:       nil,
		StatusCode: strconv.Itoa(code),
		Msg:        msg,
	}
	handleError(w, response, status)
}

func Handle5xx(w http.ResponseWriter, e error, code int, status string) {
	log.Printf("Internal error %s\n", e)
	response := domain.BaseResponse[any]{
		Data:       nil,
		StatusCode: strconv.Itoa(code),
		Msg:        e.Error(),
	}
	handleError(w, response, status)
}

func handleError(w http.ResponseWriter, response domain.BaseResponse[any], status string) {
	json, _ := json.Marshal(response)
	w.Header().Set("status", status)
	w.Write(json)
}
