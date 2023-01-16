package vo

import (
	"ginbase/app/models"
	"ginbase/app/service/cart_service/vo"
	dto2 "ginbase/app/service/order_service/dto"
	vo2 "ginbase/app/service/wechat_user_service/vo"
)

type ConfirmOrder struct {
	AddressInfo       models.GinbaseUserAddress `json:"addressInfo"`
	BargainId         int64                     `json:"bargainId"`
	CartInfo          []vo.Cart                 `json:"cartInfo"`
	CombinationId     int64                     `json:"combinationId"`
	Deduction         bool                      `json:"deduction"`
	EnableIntegral    bool                      `json:"enableIntegral"`
	SeckillId         int64                     `json:"seckillId"`
	EnableIntegralNum int                       `json:"enableIntegralNum"`
	IntegralRadio     int                       `json:"integralRadio"`
	OrderKey          string                    `json:"orderKey"`
	StoreSelfMention  int                       `json:"storeSelfMention"`
	//UsableCoupon string `json:"usableCoupon"`
	//SystemStore string `json:"systemStore"`
	UserInfo   vo2.User        `json:"userInfo"`
	PriceGroup dto2.PriceGroup `json:"priceGroup"`
}
