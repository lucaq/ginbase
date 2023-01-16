package vo

import (
	"ginbase/app/models"
)

type ProductDetail struct {
	ProductAttr  []ProductAttr                                  `json:"productAttr"`
	ProductValue map[string]models.GinbaseStoreProductAttrValue `json:"productValue"`
	Reply        models.GinbaseStoreProductReply                `json:"reply"`
	ReplyChance  string                                         `json:"replyChance"`
	ReplyCount   string                                         `json:"replyCount"`
	StoreInfo    Product                                        `json:"storeInfo"`
	Uid          int64                                          `json:"uid"`
	TempName     string                                         `json:"tempName"`
}
