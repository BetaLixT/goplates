package controllers

import (
	"ddd-01/pkg/app/rest/dto/req"
	"ddd-01/pkg/domain/role"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
  svc role.RoleService
}

func (ctrl *RoleController) CreateRole(ctx *gin.Context) {
  sp := ctx.MustGet("tx-context").(role.IServiceProvider)
  rbody := req.CreateRole{}
  if err := ctx.Bind(&rbody); err != nil {
    ctx.Error(err)
    return
  }
  ctrl.svc.CreateRole(
    sp,
    rbody.Id,
    rbody.Title,
    rbody.Description,
  )
}
