package app

import (
	"ginbase/app/models/dto"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetParams(c *gin.Context) dto.BasePage {
	var (
		page   int
		size   int
		blurry string
	)

	page = com.StrTo(c.DefaultQuery("page", "1")).MustInt()
	size = com.StrTo(c.DefaultQuery("size", "1")).MustInt()
	blurry = c.DefaultQuery("blurry", "")

	return dto.BasePage{Page: page, Size: size, Blurry: blurry}
}
