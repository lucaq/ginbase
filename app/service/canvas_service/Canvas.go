package canvas_service

import (
	"ginbase/app/models"
	"ginbase/app/models/vo"
	"ginbase/pkg/global"
)

type Canvas struct {
	Id       int64
	Terminal int

	Enabled int

	M *models.GinbaseStoreCanvas

	Ids []int64
}

func (d *Canvas) Get() vo.ResultList {
	var data models.GinbaseStoreCanvas
	err := global.GINBASE_DB.Model(&models.GinbaseStoreCanvas{}).Where("terminal = ?", d.Terminal).First(&data).Error
	if err != nil {
		global.GINBASE_LOG.Error(err)
	}
	return vo.ResultList{Content: data, TotalElements: 0}
}

func (d *Canvas) Save() error {
	if d.M.Id == 0 {
		return models.AddCanvas(d.M)
	} else {
		data := &models.GinbaseStoreCanvas{
			Name:     d.M.Name,
			Terminal: d.M.Terminal,
			Json:     d.M.Json,
		}
		return global.GINBASE_DB.Model(&models.GinbaseStoreCanvas{}).Where("id = ?", d.M.Id).Updates(data).Error
	}

}

//func (d *Canvas) Save() error {
//	return models.UpdateByCanvas(d.M)
//}

func (d *Canvas) Del() error {
	return models.DelByCanvas(d.Ids)
}
