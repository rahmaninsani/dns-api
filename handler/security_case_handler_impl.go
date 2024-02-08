package handler

import (
	"github.com/rahmaninsani/dns-api/app"
	"github.com/rahmaninsani/dns-api/helper"
	"github.com/rahmaninsani/dns-api/model/web"
	"github.com/rahmaninsani/dns-api/usecase"
	"net/http"
)

type SecurityCaseHandlerImpl struct {
	SecurityCaseUseCase usecase.SecurityCaseUseCase
}

func NewSecurityCaseHandler(securityCaseUseCase usecase.SecurityCaseUseCase) SecurityCaseHandler {
	return &SecurityCaseHandlerImpl{
		SecurityCaseUseCase: securityCaseUseCase,
	}
}

func (handler *SecurityCaseHandlerImpl) Create(writer http.ResponseWriter, request *http.Request, params app.HttpParams) {
	securityCaseCreateRequest := web.SecurityCaseCreateRequest{}
	helper.ReadFromRequestBody(request, &securityCaseCreateRequest)

	securityCaseResponse := handler.SecurityCaseUseCase.Create(securityCaseCreateRequest)

	response := helper.ToResponse(http.StatusCreated, "Success", securityCaseResponse)
	helper.WriteToResponseBody(writer, response)
}

func (handler *SecurityCaseHandlerImpl) Update(writer http.ResponseWriter, request *http.Request, params app.HttpParams) {
	securityCaseId := params["id"]
	securityCaseUpdateRequest := web.SecurityCaseUpdateRequest{}
	securityCaseUpdateRequest.ID = securityCaseId
	helper.ReadFromRequestBody(request, &securityCaseUpdateRequest)

	securityCaseResponse := handler.SecurityCaseUseCase.Update(securityCaseUpdateRequest)

	response := helper.ToResponse(http.StatusOK, "Success", securityCaseResponse)
	helper.WriteToResponseBody(writer, response)
}

func (handler *SecurityCaseHandlerImpl) Delete(writer http.ResponseWriter, request *http.Request, params app.HttpParams) {
	securityCaseId := params["id"]
	handler.SecurityCaseUseCase.Delete(securityCaseId)

	response := helper.ToResponse(http.StatusOK, "Success", nil)
	helper.WriteToResponseBody(writer, response)
}

func (handler *SecurityCaseHandlerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params app.HttpParams) {
	securityCaseResponses := handler.SecurityCaseUseCase.FindAll()

	response := helper.ToResponse(http.StatusOK, "Success", securityCaseResponses)
	helper.WriteToResponseBody(writer, response)
}

func (handler *SecurityCaseHandlerImpl) FindById(writer http.ResponseWriter, request *http.Request, params app.HttpParams) {
	securityCaseId := params["id"]
	securityCaseResponse := handler.SecurityCaseUseCase.FindById(securityCaseId)

	response := helper.ToResponse(http.StatusOK, "Success", securityCaseResponse)
	helper.WriteToResponseBody(writer, response)
}
