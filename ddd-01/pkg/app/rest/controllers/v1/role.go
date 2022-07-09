package v1

import (
	"ddd-01/pkg/app/rest/dto/req"
	"ddd-01/pkg/app/rest/dto/res"
	"ddd-01/pkg/domain/role"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
  svc *role.RoleService
}

// @BasePath 

// CreateRole godoc
// @Summary create role 
// @Schemes
// @Description create role 
// @Tags role
// @Accept json
// @Param role body req.CreateRole true "The input todo struct"
// @Produce json
// @Success 201 {object} res.Role{}
// @Router /api/v1/roles/ [post]
func (ctrl *RoleController) createRole(ctx *gin.Context) {
  sp := ctx.MustGet("tx-context").(role.IServiceProvider)
  rbody := req.CreateRole{}
  if err := ctx.Bind(&rbody); err != nil {
    ctx.Error(err)
    return
  }
  role, err := ctrl.svc.CreateRole(
    sp,
    rbody.Id,
    rbody.Title,
    rbody.Description,
  )
  if err != nil {
    ctx.Error(err)
    return
  }
  ctx.JSON(201, res.MapRoleToDto(role))
}

func (ctrl *RoleController) RegisterRoutes(grp *gin.RouterGroup) {
  grp.POST("", ctrl.createRole)
}

func NewRoleController(
  svc *role.RoleService,
) *RoleController {
  return &RoleController{
    svc: svc,
  }
}
