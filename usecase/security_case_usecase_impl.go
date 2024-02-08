package usecase

import (
	"fmt"
	"github.com/rahmaninsani/dns-api/exception"
	"github.com/rahmaninsani/dns-api/helper"
	"github.com/rahmaninsani/dns-api/model/domain"
	"github.com/rahmaninsani/dns-api/model/web"
	"github.com/rahmaninsani/dns-api/repository"
	"time"
)

type SecurityCaseUseCaseImpl struct {
	SecurityCaseRepository repository.SecurityCaseRepository
}

func NewSecurityCaseUseCase(securityCaseRepository repository.SecurityCaseRepository) SecurityCaseUseCase {
	return &SecurityCaseUseCaseImpl{
		SecurityCaseRepository: securityCaseRepository,
	}
}

func (useCase *SecurityCaseUseCaseImpl) Create(request web.SecurityCaseCreateRequest) web.SecurityCaseResponse {
	securityCase := domain.SecurityCase{
		ID:        helper.GenerateUUID(),
		Subject:   request.Subject,
		Risk:      request.Risk,
		Analyst:   request.Analyst,
		Timestamp: time.Now(),
	}

	newSecurityCase := useCase.SecurityCaseRepository.Save(securityCase)

	return helper.ToSecurityCaseResponse(newSecurityCase)
}

func (useCase *SecurityCaseUseCaseImpl) Update(request web.SecurityCaseUpdateRequest) web.SecurityCaseResponse {
	securityCase, err := useCase.SecurityCaseRepository.FindById(request.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	fmt.Println(request.ID)
	if securityCase.ID == "" {
		panic(exception.NewNotFoundError("Security Case with ID '" + request.ID + "' not found"))
	}

	securityCase.Subject = request.Subject
	securityCase.Risk = request.Risk
	securityCase.Analyst = request.Analyst
	securityCase.Timestamp = time.Now()

	updatedSecurityCase := useCase.SecurityCaseRepository.Update(securityCase)

	return helper.ToSecurityCaseResponse(updatedSecurityCase)

}

func (useCase *SecurityCaseUseCaseImpl) Delete(securityCaseId string) {
	securityCase, err := useCase.SecurityCaseRepository.FindById(securityCaseId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	if securityCase.ID == "" {
		panic(exception.NewNotFoundError("Security Case with ID '" + securityCaseId + "' not found"))
	}

	useCase.SecurityCaseRepository.Delete(securityCase)
}

func (useCase *SecurityCaseUseCaseImpl) FindAll() []web.SecurityCaseResponse {
	securityCases := useCase.SecurityCaseRepository.FindAll()

	return helper.ToSecurityCaseResponses(securityCases)
}

func (useCase *SecurityCaseUseCaseImpl) FindById(securityCaseId string) web.SecurityCaseResponse {
	securityCase, err := useCase.SecurityCaseRepository.FindById(securityCaseId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	if securityCase.ID == "" {
		panic(exception.NewNotFoundError("Security Case with ID '" + securityCaseId + "' not found"))
	}

	return helper.ToSecurityCaseResponse(securityCase)
}
