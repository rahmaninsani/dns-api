package exception

import (
	"github.com/rahmaninsani/dns-api/helper"
	"net/http"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err any) {
	if notFoundError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err any) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		response := helper.ToResponse(http.StatusNotFound, exception.Error, nil)
		helper.WriteToResponseBody(writer, response)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err any) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	response := helper.ToResponse(http.StatusInternalServerError, "Internal Server Error", nil)
	helper.WriteToResponseBody(writer, response)
}
