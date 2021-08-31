package common

import (
	"fmt"
	"reflect"
	"regexp"
	"time"

	"github.com/xstnet/starfire-cloud/internal/errors"
)

func FormatFileSize(fileSize uint64) (size string) {
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

func FormatTimestamp(timestamp int64) string {
	return time.Unix(int64(timestamp), 0).Format("2006-01-02 15:04:05")
}

func CheckDirname(dirname string) error {
	if dirname == "" {
		return errors.New("文件夹名称不能为空")
	}
	matched, err := regexp.MatchString(`^[^/\\\\:\\*\\?\\<\\>\\|\"]{1,255}$`, dirname)
	if err != nil || !matched {
		return errors.New(`文件夹名称不能包含\/:*?"<>|`)
	}
	return nil
}

func CheckFilename(filename string) error {
	if filename == "" {
		return errors.New("文件名称不能为空")
	}
	matched, err := regexp.MatchString(`^[^/\\\\:\\*\\?\\<\\>\\|\"]{1,255}$`, filename)
	if err != nil || !matched {
		return errors.New(`文件名称不能包含\/:*?"<>|`)
	}
	return nil
}

// 处理列表分页
func ProcessPageByList(page, pageSize, defaultPageSize int) (limit int, offset int) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = defaultPageSize
	}

	// 限制最大值
	if pageSize > 10000 {
		pageSize = 10000
	}

	offset = (page - 1) * pageSize
	limit = pageSize

	return
}

func DumpVal(val interface{}) {
	fmt.Println("[dump-val], type:", reflect.TypeOf(val))
	fmt.Println("[dump-val], value:", val)
}
