package cart_service

import (
	"errors"
	"strconv"
	"strings"

	"ginbase/app/models"
	"ginbase/app/models/vo"
	"ginbase/app/params"
	cartVo "ginbase/app/service/cart_service/vo"
	"ginbase/app/service/product_service"
	productVo "ginbase/app/service/product_service/vo"
	cartEnum "ginbase/pkg/enums/cart"
	orderEnum "ginbase/pkg/enums/order"
	"ginbase/pkg/global"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type Cart struct {
	Id   int64
	Name string

	Enabled int

	PageNum  int
	PageSize int

	M *models.GinbaseStoreCart

	Ids []int64

	Uid     int64
	Param   *params.CartParam
	CartIds string
	Status  int

	NumParam *params.CartNumParam
	IdsParam *params.CartIdsParam
}

//update cart num
func (d *Cart) ChangeCartNum() error {
	var (
		cart models.GinbaseStoreCart
		err  error
	)
	err = global.GINBASE_DB.Model(&models.GinbaseStoreCart{}).
		Where("uid = ?", d.Uid).
		Where("id = ?", d.NumParam.Id).
		First(&cart).Error
	if err != nil {
		global.GINBASE_LOG.Error(err)
		return errors.New("当前购物车不存在")
	}
	err = CheckStock(cart.ProductId, d.NumParam.Number, cart.ProductAttrUnique)
	if err != nil {
		return errors.New("改产品库存不足" + strconv.Itoa(d.NumParam.Number))
	}
	if d.NumParam.Number == cart.CartNum {
		return nil
	}

	err = global.GINBASE_DB.Model(&models.GinbaseStoreCart{}).
		Where("id = ?", cart.Id).
		Update("cart_num", d.NumParam.Number).Error
	if err != nil {
		global.GINBASE_LOG.Error(err)
		return errors.New("修改失败")
	}
	return nil
}

//get cart list
func (d *Cart) GetCartList() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["uid"] = d.Uid
	maps["is_pay"] = orderEnum.PAY_STATUS_0
	if d.Status == 0 {
		maps["is_new"] = cartEnum.NEW_O
	}
	if d.CartIds != "" {
		maps["id"] = strings.Split(d.CartIds, ",")
	}
	//if len(d.Ids) > 0 {
	//	maps["id"] = d.Ids
	//}
	var (
		carts   []models.GinbaseStoreCart
		valid   = make([]cartVo.Cart, 0)
		invalid = make([]cartVo.Cart, 0)
		error   error
	)
	error = global.GINBASE_DB.Model(&models.GinbaseStoreCart{}).
		Where(maps).Find(&carts).Error
	if error != nil {
		global.GINBASE_LOG.Error(error)
		return gin.H{
			"valid":   valid,
			"invalid": invalid,
		}
	}

	for _, cart := range carts {
		var (
			storeProduct models.GinbaseStoreProduct
			productVo    productVo.Product
			cartVo       cartVo.Cart
		)
		err := global.GINBASE_DB.Model(&models.GinbaseStoreProduct{}).
			Where("id = ?", cart.ProductId).
			First(&storeProduct).Error
		copier.Copy(&productVo, storeProduct)
		copier.Copy(&cartVo, cart)
		if err != nil {
			global.GINBASE_LOG.Error(err)
			global.GINBASE_DB.Where("id = ?", cart.Id).Delete(&models.GinbaseStoreCart{})
			continue
		}
		ee := CheckStock(cart.ProductId, cart.CartNum, cart.ProductAttrUnique)
		if storeProduct.IsShow == 0 || ee != nil {
			cartVo.ProductInfo = productVo
			invalid = append(invalid, cartVo)
			continue
		}
		//获取有效购物车
		var productAttrValue models.GinbaseStoreProductAttrValue
		global.GINBASE_DB.Model(&models.GinbaseStoreProductAttrValue{}).
			Where("`unique` = ?", cart.ProductAttrUnique).First(&productAttrValue)
		productVo.AttrInfo = productAttrValue
		cartVo.ProductInfo = productVo
		cartVo.TruePrice = productAttrValue.Price
		cartVo.VipTruePrice = productAttrValue.Price
		cartVo.CostPrice = productAttrValue.Cost
		cartVo.TrueStock = productAttrValue.Stock
		valid = append(valid, cartVo)
	}

	return gin.H{
		"valid":   valid,
		"invalid": invalid,
	}
}

//get num
func (d *Cart) GetUserCartNum() int {
	var num int
	err := global.GINBASE_DB.
		Raw("select sum(cart_num) from ginbase_store_cart "+
			"where is_pay=0 and is_del = 0 and is_new = 0 and uid = ?", d.Uid).
		Scan(&num).Error
	if err != nil {
		global.GINBASE_LOG.Error(err)
		return 0
	}
	return num
}

//add collect
func (d *Cart) AddCart() (int64, error) {
	err := CheckStock(d.Param.ProductId, d.Param.CartNum, d.Param.UniqueId)
	if err != nil {
		return 0, err
	}
	var cart models.GinbaseStoreCart
	whereCart := &models.GinbaseStoreCart{
		Uid:               d.Uid,
		IsPay:             orderEnum.PAY_STATUS_0,
		ProductAttrUnique: d.Param.UniqueId,
		ProductId:         d.Param.ProductId,
		IsNew:             d.Param.IsNew,
		BargainId:         d.Param.BargainId,
		CombinationId:     d.Param.CombinationId,
		SeckillId:         d.Param.SeckillId,
	}
	err = global.GINBASE_DB.Model(&models.GinbaseStoreCart{}).
		Where(whereCart).
		Where("is_new = ?", d.Param.IsNew).
		Order("id desc").First(&cart).Error
	if err != nil {
		whereCart.CartNum = d.Param.CartNum
		err = global.GINBASE_DB.Model(&models.GinbaseStoreCart{}).Create(whereCart).Error
		if err != nil {
			return 0, errors.New("加入购物车失败")
		}

		return whereCart.Id, nil
	} else {
		if cartEnum.NEW_O == d.Param.IsNew {
			cart.CartNum = cart.CartNum + d.Param.CartNum
		}
		err = global.GINBASE_DB.Model(&models.GinbaseStoreCart{}).
			Where("id = ?", cart.Id).Save(cart).Error
		if err != nil {
			return 0, errors.New("加入购物车失败")
		}
	}

	return cart.Id, nil
}

func CheckStock(productId int64, cartNum int, unique string) error {
	var (
		storeProduct models.GinbaseStoreProduct
		err          error
	)
	err = global.GINBASE_DB.Model(&models.GinbaseStoreProduct{}).
		Where("id = ?", productId).
		Where("is_show", 1).
		First(&storeProduct).Error
	if err != nil {
		global.GINBASE_LOG.Error(err)
		return errors.New("该商品已下架或者删除")
	}
	productService := product_service.Product{
		Id:     productId,
		Unique: unique,
	}

	stock := productService.GetStock()
	if stock < cartNum {
		return errors.New(storeProduct.StoreName + "库存不足" + strconv.Itoa(cartNum))
	}
	return nil
}

func (d *Cart) GetAll() vo.ResultList {
	maps := make(map[string]interface{})
	if d.Name != "" {
		maps["name"] = d.Name
	}

	total, list := models.GetAllStoreCart(d.PageNum, d.PageSize, maps)
	return vo.ResultList{Content: list, TotalElements: total}
}

func (d *Cart) Insert() error {
	return models.AddStoreCart(d.M)
}

func (d *Cart) Save() error {
	return models.UpdateByStoreCart(d.M)
}

func (d *Cart) Del() error {
	return models.DelByStoreCart(d.IdsParam.Ids)
}
