package dto

import "ginbase/app/models/dto"

type UserQuery struct {
	dto.BasePage
	Sort    string
	Blurry  string
	Enabled bool
}
