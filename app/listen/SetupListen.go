package listen

import (
	"fmt"
	"ginbase/pkg/global"
)

func Setup() {
	var sub PSubscriber
	fmt.Printf(global.GINBASE_CONFIG.Redis.Host)
	conn := PConnect(global.GINBASE_CONFIG.Redis.Host, global.GINBASE_CONFIG.Redis.Password)
	sub.ReceiveKeySpace(conn)
	sub.Psubscribe()
}
