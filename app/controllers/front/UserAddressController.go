package front

import (
	"net/http"

	"ginbase/app/params"
	"ginbase/app/service/address_service"
	"ginbase/pkg/app"
	"ginbase/pkg/constant"
	"ginbase/pkg/jwt"
	"ginbase/pkg/util"

	"github.com/gin-gonic/gin"
)

// Address api
type UserAddressController struct {
}

// @Title 设置默认地址
// @Description 设置默认地址
// @Success 200 {object} app.Response
// @router /api/v1/address/del [post]
func (e *UserAddressController) Del(c *gin.Context) {
	var (
		param params.IdParam
		appG  = app.Gin{C: c}
	)
	paramErr := app.BindAndValidate(c, &param)
	if paramErr != nil {
		appG.Response(http.StatusBadRequest, paramErr.Error(), nil)
		return
	}

	uid, _ := jwt.GetAppUserId(c)
	addressService := address_service.Address{
		Id:  param.Id,
		Uid: uid,
	}
	err := addressService.DelAddress()
	if err != nil {
		appG.Response(http.StatusInternalServerError, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, constant.SUCCESS, "ok")

}

// @Title 设置默认地址
// @Description 设置默认地址
// @Success 200 {object} app.Response
// @router /api/v1/address/default/set [post]
func (e *UserAddressController) SetDefault(c *gin.Context) {
	var (
		param params.IdParam
		appG  = app.Gin{C: c}
	)
	paramErr := app.BindAndValidate(c, &param)
	if paramErr != nil {
		appG.Response(http.StatusBadRequest, paramErr.Error(), nil)
		return
	}

	uid, _ := jwt.GetAppUserId(c)
	addressService := address_service.Address{
		Id:  param.Id,
		Uid: uid,
	}
	err := addressService.SetDefault()
	if err != nil {
		appG.Response(http.StatusInternalServerError, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, constant.SUCCESS, "ok")

}

// @Title 获取列表数据
// @Description 获取列表数据
// @Success 200 {object} app.Response
// @router /api/v1/address [get]
func (e *UserAddressController) GetList(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	uid, _ := jwt.GetAppUserId(c)
	addressService := address_service.Address{
		Enabled:  1,
		PageNum:  util.GetFrontPage(c),
		PageSize: util.GetFrontLimit(c),
		Uid:      uid,
	}

	vo, total, page := addressService.GetList()

	appG.ResponsePage(http.StatusOK, constant.SUCCESS, vo, total, page)

}

// @Title 添加or更新地址
// @Description 添加or更新地址
// @Success 200 {object} app.Response
// @router /api/v1/address/edit [post]
func (e *UserAddressController) SaveAddress(c *gin.Context) {
	var (
		param params.AddressParan
		appG  = app.Gin{C: c}
	)
	paramErr := app.BindAndValidate(c, &param)
	if paramErr != nil {
		appG.Response(http.StatusBadRequest, paramErr.Error(), nil)
		return
	}

	uid, _ := jwt.GetAppUserId(c)
	addressService := address_service.Address{
		Param: &param,
		Uid:   uid,
	}
	id, err := addressService.AddOrUpdate()
	if err != nil {
		appG.Response(http.StatusInternalServerError, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, constant.SUCCESS, gin.H{"id": id})

}

// @Title 获取树形数据
// @Description 获取树形数据
// @Success 200 {object} app.Response
// @router /api/v1/city_list [get]
func (e *UserAddressController) GetCityList(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	addressService := address_service.Address{Enabled: 1}
	vo := addressService.GetCitys()
	appG.Response(http.StatusOK, constant.SUCCESS, vo)

}
