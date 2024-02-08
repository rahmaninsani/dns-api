package test

import (
	"encoding/json"
	"github.com/rahmaninsani/dns-api/app"
	"github.com/rahmaninsani/dns-api/exception"
	"github.com/rahmaninsani/dns-api/handler"
	"github.com/rahmaninsani/dns-api/helper"
	"github.com/rahmaninsani/dns-api/model/domain"
	"github.com/rahmaninsani/dns-api/repository"
	"github.com/rahmaninsani/dns-api/route"
	"github.com/rahmaninsani/dns-api/usecase"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const baseURL = "http://localhost:8080/api/security-cases"

func setupTestDB() *app.DB {
	db := app.NewDB("../data_test.json")
	return db
}

func setupRouter(db *app.DB) http.Handler {
	db.TruncateDB()

	router := app.NewRouter()
	router.PanicHandler = exception.ErrorHandler

	securityCaseRepository := repository.NewSecurityCaseRepository(db)
	securityCaseUseCase := usecase.NewSecurityCaseUseCase(securityCaseRepository)
	securityCaseHandler := handler.NewSecurityCaseHandler(securityCaseUseCase)
	route.NewSecurityCaseRouter(router, securityCaseHandler)

	return router
}

func TestCreateSecurityCaseSuccess(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)

	securityCase := domain.SecurityCase{
		Subject: "Subject Test",
		Risk:    "Risk Test",
		Analyst: "Analyst Test",
	}

	securityCaseJSON, err := json.Marshal(securityCase)
	helper.PanicIfError(err)

	requestBody := strings.NewReader(string(securityCaseJSON))
	request := httptest.NewRequest(http.MethodPost, baseURL, requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, err := io.ReadAll(response.Body)
	helper.PanicIfError(err)

	var responseBody map[string]any
	err = json.Unmarshal(body, &responseBody)
	helper.PanicIfError(err)

	if _, ok := responseBody["code"].(float64); !ok {
		t.Errorf("Expected code %v, got %v", "number", responseBody["code"])
	}

	if responseBody["code"].(float64) != http.StatusCreated {
		t.Errorf("Expected code %d, got %v", http.StatusCreated, responseBody["code"])
	}

	if responseBody["status"] != http.StatusText(http.StatusCreated) {
		t.Errorf("Expected status %s, got %v", http.StatusText(http.StatusCreated), responseBody["status"])
	}

	if responseBody["message"] != "Success" {
		t.Errorf("Expected message %s, got %v", "Success", responseBody["message"])
	}

	securityCaseResponse := responseBody["data"].(map[string]any)

	if securityCaseResponse["id"] == "" {
		t.Errorf("Expected id not empty, got %v", securityCaseResponse["id"])
	}

	if securityCaseResponse["subject"] != securityCase.Subject {
		t.Errorf("Expected subject %s, got %v", securityCase.Subject, securityCaseResponse["subject"])
	}

	if securityCaseResponse["risk"] != securityCase.Risk {
		t.Errorf("Expected risk %s, got %v", securityCase.Risk, securityCaseResponse["risk"])
	}

	if securityCaseResponse["analyst"] != securityCase.Analyst {
		t.Errorf("Expected analyst %s, got %v", securityCase.Analyst, securityCaseResponse["analyst"])
	}

	if securityCaseResponse["timestamp"] == "" {
		t.Errorf("Expected timestamp not empty, got %v", securityCaseResponse["timestamp"])
	}

	db.TruncateDB()
}
