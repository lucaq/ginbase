package models

import "time"

type GinbaseUserExtract struct {
	Uid          int64     `json:"uid"`
	RealName     string    `json:"real_name"`
	ExtractType  string    `json:"extract_type"`
	BankCode     string    `json:"bank_code"`
	BankAddress  string    `json:"bank_address"`
	AlipayCode   string    `json:"alipay_code"`
	ExtractPrice float64   `json:"extract_price"`
	Mark         string    `json:"mark"`
	Balance      float64   `json:"balance"`
	FailMsg      string    `json:"fail_msg"`
	FailTime     time.Time `json:"fail_time"`
	Status       int8      `json:"status"`
	Wechat       string    `json:"wechat"`
	BaseModel
}

func (GinbaseUserExtract) TableName() string {
	return "ginbase_user_extract"
}

func AddUserExtract(m *GinbaseUserExtract) error {
	var err error
	if err = db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByUserExtract(m *GinbaseUserExtract) error {
	var err error
	err = db.Save(m).Error
	if err != nil {
		return err
	}

	return err
}

func DelByUserExtract(ids []int64) error {
	var err error
	err = db.Where("id in (?)", ids).Delete(&GinbaseUserExtract{}).Error
	if err != nil {
		return err
	}

	return err
}
