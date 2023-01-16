package wechat

import (
	"ginbase/pkg/global"

	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
)

func InitWechat() {
	wc := wechat.NewWechat()
	//这里本地内存保存access_token，也可选择redis，memcache或者自定cache
	redisOpts := &cache.RedisOpts{
		Host:        global.GINBASE_CONFIG.Redis.Host,
		Password:    global.GINBASE_CONFIG.Redis.Password,
		Database:    0,
		MaxActive:   global.GINBASE_CONFIG.Redis.MaxActive,
		MaxIdle:     global.GINBASE_CONFIG.Redis.MaxIdle,
		IdleTimeout: 200,
	}
	redisCache := cache.NewRedis(redisOpts)
	wc.SetCache(redisCache)
	cfg := &offConfig.Config{
		AppID:          global.GINBASE_CONFIG.Wechat.AppID,
		AppSecret:      global.GINBASE_CONFIG.Wechat.AppSecret,
		Token:          global.GINBASE_CONFIG.Wechat.Token,
		EncodingAESKey: global.GINBASE_CONFIG.Wechat.EncodingAESKey,
	}

	officialAccount := wc.GetOfficialAccount(cfg)

	global.GINBASE_OFFICIAL_ACCOUNT = officialAccount
}
