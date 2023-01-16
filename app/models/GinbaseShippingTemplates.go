package models

type GinbaseShippingTemplates struct {
	Name        string `json:"name"`
	Type        int8   `json:"type"`
	RegionInfo  string `json:"region_info"`
	Appoint     int8   `json:"appoint"`
	AppointInfo string `json:"appoint_info"`
	Sort        int8   `json:"sort"`
	BaseModel
}

func (GinbaseShippingTemplates) TableName() string {
	return "ginbase_shipping_templates"
}

func AddShippingTemplates(m *GinbaseShippingTemplates) error {
	var err error
	if err = db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByShippingTemplates(m *GinbaseShippingTemplates) error {
	var err error
	err = db.Save(m).Error
	if err != nil {
		return err
	}

	return err
}

func DelByShippingTemplatess(ids []int64) error {
	var err error
	err = db.Where("id in (?)", ids).Delete(&GinbaseShippingTemplates{}).Error
	if err != nil {
		return err
	}

	return err
}
