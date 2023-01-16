package models

type GinbaseStoreCanvas struct {
	Terminal int    `json:"terminal"`
	Json     string `json:"json"`
	Ttype    int    `json:"type" gorm:"column:type"`
	Name     string `json:"name"`
	BaseModel
}

func (GinbaseStoreCanvas) TableName() string {
	return "ginbase_store_canvas"
}

func AddCanvas(m *GinbaseStoreCanvas) error {
	var err error
	if err = db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByCanvas(m *GinbaseStoreCanvas) error {
	var err error
	err = db.Save(m).Error
	if err != nil {
		return err
	}

	return err
}

func DelByCanvas(ids []int64) error {
	var err error
	err = db.Where("id in (?)", ids).Delete(&GinbaseStoreCanvas{}).Error
	if err != nil {
		return err
	}

	return err
}
