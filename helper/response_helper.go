package helper

import (
	"github.com/rahmaninsani/dns-api/model/domain"
	"github.com/rahmaninsani/dns-api/model/web"
	"net/http"
)

func ToResponse(code int, message string, data any) web.Response {
	return web.Response{
		Code:    code,
		Status:  http.StatusText(code),
		Message: message,
		Data:    data,
	}
}

func ToSecurityCaseResponse(securityCase domain.SecurityCase) web.SecurityCaseResponse {
	return web.SecurityCaseResponse{
		ID:        securityCase.ID,
		Subject:   securityCase.Subject,
		Risk:      securityCase.Risk,
		Analyst:   securityCase.Analyst,
		Timestamp: securityCase.Timestamp,
	}
}

func ToSecurityCaseResponses(securityCases []domain.SecurityCase) []web.SecurityCaseResponse {
	var responses []web.SecurityCaseResponse
	
	for _, securityCase := range securityCases {
		responses = append(responses, ToSecurityCaseResponse(securityCase))
	}
	
	return responses
}
