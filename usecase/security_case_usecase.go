package usecase

import (
	"github.com/rahmaninsani/dns-api/model/web"
)

type SecurityCaseUseCase interface {
	Create(request web.SecurityCaseCreateRequest) web.SecurityCaseResponse
	Update(request web.SecurityCaseUpdateRequest) web.SecurityCaseResponse
	Delete(securityCaseId string)
	FindAll() []web.SecurityCaseResponse
	FindById(securityCaseId string) web.SecurityCaseResponse
}
