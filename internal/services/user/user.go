package user

import (
	"github.com/xstnet/starfire-cloud/configs"
	"github.com/xstnet/starfire-cloud/internal/errors"
	"github.com/xstnet/starfire-cloud/internal/models"
	"github.com/xstnet/starfire-cloud/pkg/helper/fileHelper"
	"github.com/xstnet/starfire-cloud/pkg/systeminfo"
)

// CheckRemainSpace 检查剩余的存储空间是否足够
func CheckRemainSpace(user *models.User, size uint64) error {
	// 判断用户的剩余空间是否足够
	if user.TotalSpace > 0 && user.TotalSpace-user.UsedSpace < size {
		return errors.New("可用上传空间不足，当前剩余: " + fileHelper.FormatSize(user.TotalSpace-user.UsedSpace))
	}
	// 判断磁盘余量
	diskInfo := systeminfo.DiskInfo(configs.Upload.UploadRootPath)
	if diskInfo.Total == 0 {
		return errors.New("获取磁盘信息失败")
	}
	if diskInfo.Free < size {
		return errors.New("磁盘剩余空间不足，当前余量: " + fileHelper.FormatSize(diskInfo.Free))
	}
	return nil
}
