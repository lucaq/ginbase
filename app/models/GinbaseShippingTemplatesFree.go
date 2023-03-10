package models

type GinbaseShippingTemplatesFree struct {
	Id         int64   `gorm:"primary_key" json:"id"`
	ProvinceId int     `json:"province_id"`
	TempId     int64   `json:"temp_id"`
	CityId     int     `json:"city_id"`
	Number     float64 `json:"number"`
	Price      float64 `json:"price"`
	Type       int8    `json:"type"`
	Uniqid     string  `json:"uniqid"`
}

func (GinbaseShippingTemplatesFree) TableName() string {
	return "ginbase_shipping_templates_free"
}

func AddShippingTemplatesFree(m *GinbaseShippingTemplatesFree) error {
	var err error
	if err = db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByShippingTemplatesFree(m *GinbaseShippingTemplatesFree) error {
	var err error
	err = db.Save(m).Error
	if err != nil {
		return err
	}

	return err
}

func DelByShippingTemplatesFree(ids []int64) error {
	var err error
	err = db.Where("id in (?)", ids).Delete(&GinbaseShippingTemplatesFree{}).Error
	if err != nil {
		return err
	}

	return err
}
