package route

import (
	"github.com/rahmaninsani/dns-api/app"
	"github.com/rahmaninsani/dns-api/handler"
)

func NewSecurityCaseRouter(router *app.Router, securityCaseHandler handler.SecurityCaseHandler) {
	basePath := "/api/security-cases"

	router.POST(basePath, securityCaseHandler.Create)
	router.PUT(basePath+"/:id", securityCaseHandler.Update)
	router.DELETE(basePath+"/:id", securityCaseHandler.Delete)
	router.GET(basePath, securityCaseHandler.FindAll)
	router.GET(basePath+"/:id", securityCaseHandler.FindById)
}
