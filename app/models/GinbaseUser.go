package models

import (
	"time"

	"gorm.io/datatypes"
)

type GinbaseUser struct {
	Username       string         `json:"username"`
	Password       string         `json:"password"`
	RealName       string         `json:"real_name"`
	Birthday       int            `json:"birthday"`
	CardId         string         `json:"card_id"`
	Mark           string         `json:"mark"`
	Nickname       string         `json:"nickname"`
	Avatar         string         `json:"avatar"`
	Phone          string         `json:"phone"`
	AddIp          string         `json:"add_ip"`
	LastIp         string         `json:"last_ip"`
	NowMoney       float64        `json:"nowMoney"`
	BrokeragePrice float64        `json:"brokeragePrice"`
	Integral       int            `json:"integral"`
	SignNum        int            `json:"sign_num"`
	Status         int8           `json:"status"`
	Level          int8           `json:"level"`
	SpreadUid      string         `json:"spreadUid"`
	SpreadTime     time.Time      `json:"spread_time"`
	UserType       string         `json:"userType"`
	PayCount       int            `json:"payCount"`
	SpreadCount    int            `json:"spread_count"`
	Address        string         `json:"address"`
	LoginType      string         `json:"login_type"`
	WxProfile      datatypes.JSON `json:"wx_profile"`
	BaseModel
}

func (GinbaseUser) TableName() string {
	return "ginbase_user"
}

// get all
func GetAllWechatUser(pageNUm int, pageSize int, maps interface{}) (int64, []GinbaseUser) {
	var (
		total int64
		data  []GinbaseUser
	)

	db.Model(&GinbaseUser{}).Where(maps).Count(&total)
	db.Where(maps).Offset(pageNUm).Limit(pageSize).Order("id desc").Find(&data)

	return total, data
}

func AddWechatUser(m *GinbaseUser) error {
	var err error
	if err = db.Select("username", "nickname", "password", "real_name", "avatar", "add_ip", "last_ip", "user_type", "wx_profile").Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByWechatUsere(id int64, m *GinbaseUser) error {
	var err error
	err = db.Model(&GinbaseUser{}).Where("id = ?", id).Updates(m).Error
	if err != nil {
		return err
	}

	return err
}

func DelByWechatUser(ids []int64) error {
	var err error
	err = db.Where("id in (?)", ids).Delete(&GinbaseUser{}).Error
	if err != nil {
		return err
	}

	return err
}
