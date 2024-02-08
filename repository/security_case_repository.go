package repository

import (
	"github.com/rahmaninsani/dns-api/model/domain"
)

type SecurityCaseRepository interface {
	Save(securityCase domain.SecurityCase) domain.SecurityCase
	Update(securityCase domain.SecurityCase) domain.SecurityCase
	Delete(securityCase domain.SecurityCase)
	FindAll() []domain.SecurityCase
	FindById(securityCaseId string) (domain.SecurityCase, error)
}
