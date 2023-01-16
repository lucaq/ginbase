package models

import (
	"strings"

	dto2 "ginbase/app/service/product_service/dto"
)

type GinbaseStoreProductAttr struct {
	Id         int64  `json:"id"`
	ProductId  int64  `json:"productId" valid:"Required;"`
	AttrName   string `json:"attrName" valid:"Required;"`
	AttrValues string `json:"attrValues" valid:"Required;"`
}

func (GinbaseStoreProductAttr) TableName() string {
	return "ginbase_store_product_attr"
}

func AddProductAttr(items []dto2.FormatDetail, productId int64) error {
	var err error
	tx := db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var attrGroup []GinbaseStoreProductAttr
	for _, val := range items {
		detailStr := strings.Join(val.Detail, ",")
		var storeProductAttr = GinbaseStoreProductAttr{
			ProductId:  productId,
			AttrName:   val.Value,
			AttrValues: detailStr,
		}
		attrGroup = append(attrGroup, storeProductAttr)
	}

	err = tx.Create(&attrGroup).Error
	if err != nil {
		return err
	}

	return err
}

func DelByProductttr(productId int64) (err error) {
	err = db.Where("product_id = ?", productId).Delete(GinbaseStoreProductAttr{}).Error
	return err
}
