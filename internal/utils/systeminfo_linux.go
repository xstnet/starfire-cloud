// +build !windows

package utils

import (
	"fmt"
	"runtime"
	"syscall"
	"unsafe"
)

type DiskStatus struct {
	Total uint64 `json:"total"`
	Used  uint64 `json:"used"`
	Free  uint64 `json:"free"`
}

// Linux 获取磁盘信息
func DiskInfo(path string) *DiskStatus {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Get DiskInfo Panic:", err)
		}
	}()

	fs := syscall.Statfs_t{}

	err := syscall.Statfs(path, &fs)

	if err != nil {
		// todo: log
		return &DiskStatus{}
	}

	t := &DiskStatus{
		Total: fs.Blocks * uint64(fs.Bsize),
		Free:  fs.Bfree * uint64(fs.Bsize),
		Used:  0,
	}
	t.Used = t.Total - t.Free
	return t
}
