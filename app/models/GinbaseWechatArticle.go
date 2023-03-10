package models

import "github.com/astaxie/beego/validation"

type GinbaseWechatArticle struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Image     string `json:"image"`
	Synopsis  string `json:"synopsis"`
	Content   string `json:"content"`
	Visit     int    `json:"visit"`
	Sort      int    `json:"sort"`
	Url       string `json:"url"`
	Status    int    `json:"status"`
	ProductId int    `json:"product_id"`
	MediaId   string `json:"media_id"`
	IsPub     int    `json:"is_pub"`
	BaseModel
}

func (GinbaseWechatArticle) TableName() string {
	return "ginbase_wechat_article"
}

func (a *GinbaseWechatArticle) Valid(v *validation.Validation) {
	if a.Title == "" {
		v.SetError("title", "标题不能为空")
	}
	if a.Author == "" {
		v.SetError("author", "作者不能为空")
	}
}

// get all
func GetAllWechatArticle(pageNUm int, pageSize int, maps interface{}) (int64, []GinbaseWechatArticle) {
	var (
		total int64
		data  []GinbaseWechatArticle
	)

	db.Model(&GinbaseWechatArticle{}).Where(maps).Count(&total)
	db.Where(maps).Offset(pageNUm).Limit(pageSize).Order("id desc").Find(&data)

	return total, data
}

func AddWechatArticle(m *GinbaseWechatArticle) error {
	var err error
	if err = db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByWechatArticle(m *GinbaseWechatArticle) error {
	var err error
	err = db.Save(m).Error
	if err != nil {
		return err
	}

	return err
}

func DelByWechatArticle(ids []int64) error {
	var err error
	err = db.Where("id in (?)", ids).Delete(&GinbaseWechatArticle{}).Error
	if err != nil {
		return err
	}

	return err
}
