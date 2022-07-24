package v1

import (
	"ddd/pkg/app/rest/dto/res"
	"ddd/pkg/domain/forecast"

	"github.com/gin-gonic/gin"
)

type ForecastController struct {
	svc *forecast.ForecastService
}

// @BasePath

// ListForecasts godoc
// @Summary List forecasts
// @Schemes
// @Description List forecasts
// @Tags role
// @Accept json
// @Produce json
// @Success 200 {object} []res.ForecastDetailed{}
// @Router /api/v1/forecasts/ [get]
func (ctrl *ForecastController) listForecasts(ctx *gin.Context) {
	resbody := ctrl.svc.ListForecasts(ctx)
	ctx.JSON(200, res.MapForecastToDetailedSliceDto(resbody))
}

func (ctrl *ForecastController) RegisterRoutes(grp *gin.RouterGroup) {
	grp.GET("", ctrl.listForecasts)
}

func NewForecastController(
	svc *forecast.ForecastService,
) *ForecastController{
	return &ForecastController{
		svc: svc,
	}
}
