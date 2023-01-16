package dto

import "ginbase/app/models/dto"

type DictQuery struct {
	dto.BasePage
	Blurry string
}
