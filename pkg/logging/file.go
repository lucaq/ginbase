package logging

import (
	"fmt"
	"ginbase/pkg/global"
	"time"
)

// getLogFilePath get the log file save path
func getLogFilePath() string {
	return fmt.Sprintf("%s%s", global.GINBASE_CONFIG.App.RuntimeRootPath, global.GINBASE_CONFIG.App.LogSavePath)
}

// getLogFileName get the save name of the log file
func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		global.GINBASE_CONFIG.App.LogSaveName,
		time.Now().Format(global.GINBASE_CONFIG.App.TimeFormat),
		global.GINBASE_CONFIG.App.LogFileExt,
	)
}
