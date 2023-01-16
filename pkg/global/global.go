package global

import (
	"ginbase/conf"

	"github.com/silenceper/wechat/v2/officialaccount"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GINBASE_DB               *gorm.DB
	GINBASE_VP               *viper.Viper
	GINBASE_LOG              *zap.SugaredLogger
	GINBASE_CONFIG           conf.Config
	GINBASE_OFFICIAL_ACCOUNT *officialaccount.OfficialAccount
)
