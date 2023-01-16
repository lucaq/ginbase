package models

import (
	"time"

	"gorm.io/gorm"
)

type GinbaseStoreOrderStatus struct {
	Id            int64     `gorm:"primary_key" json:"id"`
	Oid           int64     `json:"oid"`
	ChangeType    string    `json:"change_type"`
	ChangeMessage string    `json:"change_message"`
	ChangeTime    time.Time `json:"change_time" gorm:"autoCreateTime"`
}

func (GinbaseStoreOrderStatus) TableName() string {
	return "ginbase_store_order_status"
}

func AddStoreOrderStatus(tx *gorm.DB, oid int64, change, msg string) error {
	data := &GinbaseStoreOrderStatus{
		Oid:           oid,
		ChangeType:    change,
		ChangeMessage: msg,
	}
	return tx.Model(&GinbaseStoreOrderStatus{}).Create(data).Error
}

func UpdateByStoreOrderStatus(m *GinbaseStoreOrderStatus) error {
	var err error
	err = db.Save(m).Error
	if err != nil {
		return err
	}

	return err
}

func DelByStoreOrderStatus(ids []int64) error {
	var err error
	err = db.Where("id in (?)", ids).Delete(&GinbaseStoreOrderStatus{}).Error
	if err != nil {
		return err
	}

	return err
}
