package models

import (
	"time"

	"ginbase/pkg/constant"

	"gorm.io/datatypes"
)

type GinbaseWechatMenu struct {
	Key     string         `json:"key"`
	Result  datatypes.JSON `json:"result"`
	AddTime time.Time      `json:"addTIme" gorm:"autoCreateTime"`
}

func (GinbaseWechatMenu) TableName() string {
	return "ginbase_wechat_menu"
}

// get all
func GetWechatMenu(maps interface{}) GinbaseWechatMenu {
	var (
		data GinbaseWechatMenu
	)

	db.Where(maps).First(&data)

	return data
}

func AddWechatMenu(m *GinbaseWechatMenu) error {
	var err error
	if err = db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByWechatMenu(m *GinbaseWechatMenu) error {
	var err error
	err = db.Model(&GinbaseWechatMenu{}).Where("key", constant.GINBASE_WEICHAT_MENU).Updates(m).Error
	if err != nil {
		return err
	}

	return err
}
