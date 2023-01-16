package models

import "gorm.io/gorm"

type GinbaseUserBill struct {
	Uid      int64   `json:"uid"`
	LinkId   string  `json:"link_id"`
	Pm       int8    `json:"pm"`
	Title    string  `json:"title"`
	Category string  `json:"category"`
	Type     string  `json:"type"`
	Number   float64 `json:"number"`
	Balance  float64 `json:"balance"`
	Mark     string  `json:"mark"`
	Status   int8    `json:"status"`
	BaseModel
}

func (GinbaseUserBill) TableName() string {
	return "ginbase_user_bill"
}

//增加支出流水
func Expend(tx *gorm.DB, uid int64, title, category, typestr, mark, linkId string, number, balance float64) error {
	data := &GinbaseUserBill{
		Uid:      uid,
		Title:    title,
		Category: category,
		Type:     typestr,
		Number:   number,
		Balance:  balance,
		Mark:     mark,
		Pm:       0,
		LinkId:   linkId,
	}
	return tx.Model(&GinbaseUserBill{}).Create(data).Error
}

//增加收入流水
func Income(tx *gorm.DB, uid int64, title, category, typestr, mark, linkId string, number, balance float64) error {
	data := &GinbaseUserBill{
		Uid:      uid,
		Title:    title,
		Category: category,
		Type:     typestr,
		Number:   number,
		Balance:  balance,
		Mark:     mark,
		Pm:       1,
		LinkId:   linkId,
	}
	return tx.Model(&GinbaseUserBill{}).Create(data).Error
}

func AddUserBill(m *GinbaseUserBill) error {
	var err error
	if err = db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByUserBill(m *GinbaseUserBill) error {
	var err error
	err = db.Save(m).Error
	if err != nil {
		return err
	}

	return err
}

func DelByUserBill(ids []int64) error {
	var err error
	err = db.Where("id in (?)", ids).Delete(&GinbaseUserBill{}).Error
	if err != nil {
		return err
	}

	return err
}
