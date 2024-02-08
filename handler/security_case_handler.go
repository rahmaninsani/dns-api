package handler

import (
	"github.com/rahmaninsani/dns-api/app"
	"net/http"
)

type SecurityCaseHandler interface {
	Create(writer http.ResponseWriter, request *http.Request, params app.HttpParams)
	Update(writer http.ResponseWriter, request *http.Request, params app.HttpParams)
	Delete(writer http.ResponseWriter, request *http.Request, params app.HttpParams)
	FindAll(writer http.ResponseWriter, request *http.Request, params app.HttpParams)
	FindById(writer http.ResponseWriter, request *http.Request, params app.HttpParams)
}
