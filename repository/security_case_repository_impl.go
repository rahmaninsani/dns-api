package repository

import (
	"github.com/rahmaninsani/dns-api/app"
	"github.com/rahmaninsani/dns-api/helper"
	"github.com/rahmaninsani/dns-api/model/domain"
)

type SecurityCaseRepositoryImpl struct {
	DB *app.DB
}

func NewSecurityCaseRepository(DB *app.DB) SecurityCaseRepository {
	return &SecurityCaseRepositoryImpl{
		DB: DB,
	}
}

func (repository *SecurityCaseRepositoryImpl) Save(securityCase domain.SecurityCase) domain.SecurityCase {
	err := repository.DB.Create(securityCase)
	helper.PanicIfError(err)

	return securityCase
}

func (repository *SecurityCaseRepositoryImpl) Update(securityCase domain.SecurityCase) domain.SecurityCase {
	err := repository.DB.Update(securityCase)
	helper.PanicIfError(err)

	updatedSecurityCase, err := repository.DB.FindOne(securityCase.ID)
	helper.PanicIfError(err)

	return updatedSecurityCase
}

func (repository *SecurityCaseRepositoryImpl) Delete(securityCase domain.SecurityCase) {
	err := repository.DB.Delete(securityCase.ID)
	helper.PanicIfError(err)
}

func (repository *SecurityCaseRepositoryImpl) FindAll() []domain.SecurityCase {
	securityCases := repository.DB.Find()

	return securityCases
}

func (repository *SecurityCaseRepositoryImpl) FindById(securityCaseId string) (domain.SecurityCase, error) {
	updatedSecurityCase, err := repository.DB.FindOne(securityCaseId)
	if err != nil {
		return domain.SecurityCase{}, err
	}

	if updatedSecurityCase.ID == "" {
		return domain.SecurityCase{}, nil
	}

	return updatedSecurityCase, nil
}
