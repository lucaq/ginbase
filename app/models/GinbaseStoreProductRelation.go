package models

type GinbaseStoreProductRelation struct {
	Uid       int64                `json:"uid"`
	ProductId int64                `json:"productId"`
	Type      string               `json:"type"`
	Category  string               `json:"category"`
	Product   *GinbaseStoreProduct `json:"product" gorm:"foreignKey:ProductId;"`
	BaseModel
}

func (GinbaseStoreProductRelation) TableName() string {
	return "ginbase_store_product_relation"
}

// get all
func GetAllProductRelation(pageNUm int, pageSize int, maps interface{}) (int64, []GinbaseStoreProductRelation) {
	var (
		total int64
		data  []GinbaseStoreProductRelation
	)
	db.Model(&GinbaseStoreProductRelation{}).Where(maps).Count(&total)
	db.Where(maps).Offset(pageNUm).Limit(pageSize).Preload("Product").Order("id desc").Find(&data)

	return total, data
}

func AddStoreProductRelation(m *GinbaseStoreProductRelation) error {
	var err error
	if err = db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByStoreProductRelation(m *GinbaseStoreProductRelation) error {
	var err error
	err = db.Save(m).Error
	if err != nil {
		return err
	}

	return err
}

func DelByStoreProductRelations(ids []int64) error {
	var err error
	err = db.Where("id in (?)", ids).Delete(&GinbaseStoreProductRelation{}).Error
	if err != nil {
		return err
	}

	return err
}
