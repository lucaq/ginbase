package models

type GinbaseShippingTemplatesRegion struct {
	Id            int64   `gorm:"primary_key" json:"id"`
	ProvinceId    int     `json:"province_id"`
	TempId        int64   `json:"temp_id"`
	CityId        int     `json:"city_id"`
	First         float64 `json:"first"`
	FirstPrice    float64 `json:"first_price"`
	Continues     float64 `json:"continues"`
	ContinuePrice float64 `json:"continue_price"`
	Type          int8    `json:"type"`
	Uniqid        string  `json:"uniqid"`
}

func (GinbaseShippingTemplatesRegion) TableName() string {
	return "ginbase_shipping_templates_region"
}

func AddShippingTemplatesRegion(m *GinbaseShippingTemplatesRegion) error {
	var err error
	if err = db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByShippingTemplatesRegion(m *GinbaseShippingTemplatesRegion) error {
	var err error
	err = db.Save(m).Error
	if err != nil {
		return err
	}

	return err
}

func DelByShippingTemplatesRegion(ids []int64) error {
	var err error
	err = db.Where("id in (?)", ids).Delete(&GinbaseShippingTemplatesRegion{}).Error
	if err != nil {
		return err
	}

	return err
}
