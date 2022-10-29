package user_controller

import (
	"hackthon/common/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddEnterpriseByCnpj(ctx *gin.Context) {
	addEnterpriseError := service.CreateUser(ctx)

	if addEnterpriseError != nil {
		util.HttpBadRequestMessage(ctx, addEnterpriseError)
		return
	}

	ctx.JSON(http.StatusCreated, "Enterprise Added successfuly")
}
