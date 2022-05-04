package fileUtil

import (
	"fmt"
	"regexp"

	"github.com/xstnet/starfire-cloud/internal/errors"
)

// 检查文件名
func CheckName(filename string) error {
	if filename == "" {
		return errors.New("文件名称不能为空")
	}
	matched, err := regexp.MatchString(`^[^/\\\\:\\*\\?\\<\\>\\|\"]{1,255}$`, filename)
	if err != nil || !matched {
		return errors.New(`文件名称不能包含\/:*?"<>|`)
	}
	return nil
}

// 格式化文件大小
func FormatSize(fileSize uint64) (sizeLabel string) {
	if fileSize < 1024 {
		return fmt.Sprintf("%.2fB", float64(fileSize)/float64(1))
	} else if fileSize < (1024 * 1024) {
		return fmt.Sprintf("%.2fKB", float64(fileSize)/float64(1024))
	} else if fileSize < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fMB", float64(fileSize)/float64(1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fGB", float64(fileSize)/float64(1024*1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fTB", float64(fileSize)/float64(1024*1024*1024*1024))
	} else { //if fileSize < (1024 * 1024 * 1024 * 1024 * 1024 * 1024)
		return fmt.Sprintf("%.2fEB", float64(fileSize)/float64(1024*1024*1024*1024*1024))
	}
}
