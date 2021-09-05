// +build !windows

package systeminfo

import (
	"fmt"
	"syscall"
)

type DiskStatus struct {
	Total uint64 `json:"total"`
	Used  uint64 `json:"used"`
	Free  uint64 `json:"free"`
}

// Linux 获取磁盘信息
func DiskInfo(path string) (diskInfo *DiskStatus) {
	defer func() {
		if err := recover(); err != nil {
			// panic时也要返回初值
			diskInfo = &DiskStatus{}
			fmt.Println("Get DiskInfo Panic:", err)
		}
	}()

	fs := syscall.Statfs_t{}

	err := syscall.Statfs(path, &fs)

	if err != nil {
		// todo: log
		return &DiskStatus{}
	}

	diskInfo = &DiskStatus{
		Total: fs.Blocks * uint64(fs.Bsize),
		Free:  fs.Bfree * uint64(fs.Bsize),
		Used:  0,
	}
	diskInfo.Used = diskInfo.Total - diskInfo.Free
	return
}
