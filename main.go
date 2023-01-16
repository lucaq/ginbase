package main

import (
	"fmt"
	"ginbase/app/listen"
	"ginbase/app/models"
	"ginbase/pkg/base"
	"ginbase/pkg/global"
	"ginbase/pkg/jwt"
	"ginbase/pkg/logging"
	"ginbase/pkg/redis"
	"ginbase/pkg/wechat"
	"ginbase/routers"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	global.GINBASE_VP = base.Viper()
	global.GINBASE_LOG = base.SetupLogger()
	models.Setup()
	logging.Setup()
	redis.Setup()
	jwt.Setup()
	listen.Setup()
	wechat.InitWechat()
}

// @title ginbase  API
// @version 1.0
// @description ginbase后台系统
// @termsOfService ...
// @license.name apache2
func main() {
	gin.SetMode(global.GINBASE_CONFIG.Server.RunMode)

	routersInit := routers.InitRouter()
	endPoint := fmt.Sprintf(":%d", global.GINBASE_CONFIG.Server.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: maxHeaderBytes,
	}

	global.GINBASE_LOG.Info("[info] start http server listening %s", endPoint)
	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()

}
