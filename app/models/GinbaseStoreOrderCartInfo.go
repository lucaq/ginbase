package models

type GinbaseStoreOrderCartInfo struct {
	Id           int64  `gorm:"primary_key" json:"id"`
	Oid          int64  `json:"oid"`
	OrderId      string `json:"order_id"`
	CartId       int64  `json:"cart_id"`
	ProductId    int64  `json:"product_id"`
	CartInfo     string `json:"cart_info"`
	Unique       string `json:"unique"`
	IsAfterSales int8   `json:"is_after_sales"`
}

func (GinbaseStoreOrderCartInfo) TableName() string {
	return "ginbase_store_order_cart_info"
}

func AddStoreOrderCartInfo(m *GinbaseStoreOrderCartInfo) error {
	var err error
	if err = db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByStoreOrderCartInfo(m *GinbaseStoreOrderCartInfo) error {
	var err error
	err = db.Save(m).Error
	if err != nil {
		return err
	}

	return err
}

func DelByStoreOrderCartInfo(ids []int64) error {
	var err error
	err = db.Where("id in (?)", ids).Delete(&GinbaseStoreOrderCartInfo{}).Error
	if err != nil {
		return err
	}

	return err
}
