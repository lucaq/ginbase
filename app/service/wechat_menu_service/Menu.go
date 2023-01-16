package wechat_menu_service

import (
	"encoding/json"

	"ginbase/app/models"
	"ginbase/app/models/vo"
	menuDto "ginbase/app/service/wechat_menu_service/dto"
	"ginbase/pkg/constant"
	"ginbase/pkg/global"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

type Menu struct {
	Id  int64
	Key string

	Dto menuDto.WechatMenu

	M *models.GinbaseWechatMenu
}

func (d *Menu) GetAll() vo.ResultList {
	maps := make(map[string]interface{})
	maps["key"] = constant.GINBASE_WEICHAT_MENU

	data := models.GetWechatMenu(maps)
	return vo.ResultList{Content: data, TotalElements: 0}
}

func (d *Menu) Insert() error {
	button := gin.H{
		"button": d.Dto.Buttons,
	}
	jsonstr, _ := json.Marshal(button)
	str := string(jsonstr)
	global.GINBASE_LOG.Info(str)
	official := global.GINBASE_OFFICIAL_ACCOUNT
	m := official.GetMenu()
	err := m.SetMenuByJSON(str)
	if err != nil {
		global.GINBASE_LOG.Error(err)
	}

	result, _ := json.Marshal(d.Dto.Buttons)
	model := models.GinbaseWechatMenu{
		Key:    constant.GINBASE_WEICHAT_MENU,
		Result: datatypes.JSON(result),
	}
	return models.UpdateByWechatMenu(&model)
}
