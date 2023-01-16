package front

import (
	"net/http"

	"ginbase/app/service/cate_service"
	"ginbase/pkg/app"
	"ginbase/pkg/constant"

	"github.com/gin-gonic/gin"
)

// category api
type CategoryController struct {
}

// @Title 获取树形数据
// @Description 获取树形数据
// @Success 200 {object} app.Response
// @router /api/v1/category [get]
func (e *CategoryController) GetCateList(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	cateService := cate_service.Cate{Enabled: 1}
	vo := cateService.GetAll()
	appG.Response(http.StatusOK, constant.SUCCESS, vo)

}
