package upload

import (
	"fmt"
	"ginbase/pkg/file"
	"ginbase/pkg/global"
	"ginbase/pkg/logging"
	"ginbase/pkg/util"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

// GetImageFullUrl get the full access path
func GetImageFullUrl(name string) string {
	return global.GINBASE_CONFIG.App.PrefixUrl + "/" + GetImagePath() + name
}

// GetImageName get image name
func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)

	return fileName + ext
}

// GetImagePath get save path
func GetImagePath() string {
	return global.GINBASE_CONFIG.App.ImageSavePath
}

// GetImageFullPath get full save path
func GetImageFullPath() string {
	return global.GINBASE_CONFIG.App.RuntimeRootPath + GetImagePath()
}

// CheckImageExt check image file ext
func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range global.GINBASE_CONFIG.App.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

// CheckImageSize check image size
func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		logging.Warn(err)
		return false
	}

	return size <= global.GINBASE_CONFIG.App.ImageMaxSize
}

// CheckImage check if the file exists
func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
