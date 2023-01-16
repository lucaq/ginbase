package article_service

import (
	"errors"
	"strings"

	"ginbase/app/models"
	"ginbase/app/models/vo"
	articleEnum "ginbase/pkg/enums/article"
	"ginbase/pkg/global"

	"github.com/silenceper/wechat/v2/officialaccount/material"
)

type Article struct {
	Id   int64
	Name string

	Enabled int

	PageNum  int
	PageSize int

	M *models.GinbaseWechatArticle

	Ids []int64
}

func (d *Article) Get() vo.ResultList {
	var data models.GinbaseWechatArticle
	err := global.GINBASE_DB.Model(&models.GinbaseWechatArticle{}).Where("id = ?", d.Id).First(&data).Error
	if err != nil {
		global.GINBASE_LOG.Error(err)
	}
	return vo.ResultList{Content: data, TotalElements: 0}
}

func (d *Article) GetAll() vo.ResultList {
	maps := make(map[string]interface{})
	if d.Name != "" {
		maps["name"] = d.Name
	}

	total, list := models.GetAllWechatArticle(d.PageNum, d.PageSize, maps)
	return vo.ResultList{Content: list, TotalElements: total}
}

func (d *Article) Pub() error {
	var data models.GinbaseWechatArticle
	err := global.GINBASE_DB.Model(&models.GinbaseWechatArticle{}).Where("id = ?", d.Id).First(&data).Error
	if err != nil {
		global.GINBASE_LOG.Error(err)
	}
	if data.IsPub == articleEnum.IS_PUB_1 {
		return errors.New("已经发布过啦！")
	}
	official := global.GINBASE_OFFICIAL_ACCOUNT
	m := official.GetMaterial()
	ss := strings.Replace(data.Image, global.GINBASE_CONFIG.App.PrefixUrl+"/", global.GINBASE_CONFIG.App.RuntimeRootPath, 1)
	mediaId, url, err := m.AddMaterial(material.MediaTypeThumb, ss)
	if err != nil {
		global.GINBASE_LOG.Error(err)
		return err
	}
	global.GINBASE_LOG.Info(mediaId, url)
	art := &material.Article{
		Title:            data.Title,
		ThumbMediaID:     mediaId,
		ThumbURL:         url,
		Author:           data.Author,
		Digest:           data.Synopsis,
		ShowCoverPic:     1,
		Content:          data.Content,
		ContentSourceURL: "",
	}
	arts := []*material.Article{art}
	id, err := m.AddNews(arts)
	global.GINBASE_LOG.Info(id, err)
	if err != nil {
		global.GINBASE_LOG.Error(err)
		return err
	}

	data.MediaId = id
	data.IsPub = articleEnum.IS_PUB_1

	return models.UpdateByWechatArticle(&data)
}

func (d *Article) Insert() error {
	return models.AddWechatArticle(d.M)
}

func (d *Article) Save() error {
	return models.UpdateByWechatArticle(d.M)
}

func (d *Article) Del() error {
	return models.DelByWechatArticle(d.Ids)
}
