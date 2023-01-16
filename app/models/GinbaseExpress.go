package models

type GinbaseExpress struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Sort   int    `json:"sort"`
	IsShow int8   `json:"is_show"`
	BaseModel
}

func (GinbaseExpress) TableName() string {
	return "ginbase_express"
}

// get all
func GetAllExpress(pageNUm int, pageSize int, maps interface{}) (int64, []GinbaseExpress) {
	var (
		total int64
		lists []GinbaseExpress
	)

	db.Model(&GinbaseExpress{}).Where(maps).Count(&total)
	db.Where(maps).Offset(pageNUm).Limit(pageSize).Find(&lists)

	return total, lists
}

func AddExpress(m *GinbaseExpress) error {
	var err error
	if err = db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByExpress(m *GinbaseExpress) error {
	var err error
	err = db.Updates(m).Error
	if err != nil {
		return err
	}

	return err
}

func DelByExpress(ids []int64) error {
	var err error
	err = db.Where("id in (?)", ids).Delete(&GinbaseExpress{}).Error
	if err != nil {
		return err
	}

	return err
}
