// +build windows

package systeminfo

import (
	"fmt"
	"path/filepath"
	"syscall"
	"unsafe"
)

type DiskStatus struct {
	Total uint64 `json:"total"`
	Used  uint64 `json:"used"`
	Free  uint64 `json:"free"`
}

// windows 获取磁盘信息
// 不能保证100%获取成功， 获取不到时返回 nil
func DiskInfo(path string) *DiskStatus {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Get DiskInfo Panic:", err)
		}
	}()

	// windows下只需要传入盘符即可
	path = filepath.VolumeName(path)

	h := syscall.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")

	var (
		freeBytesAvailable     uint64 // 当前用户可用容量
		totalNumberOfBytes     uint64 // 总容量
		totalNumberOfFreeBytes uint64 // 磁盘剩余容量
	)

	// 第一个指针为 调用者可用的字节数量， 第二个指针为 磁盘总字节数 第三个指针为 磁盘可用的字节数
	c.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(path))),
		uintptr(unsafe.Pointer(&freeBytesAvailable)),     // 指针1
		uintptr(unsafe.Pointer(&totalNumberOfBytes)),     // 指针2
		uintptr(unsafe.Pointer(&totalNumberOfFreeBytes)), // 指针3
	)

	return &DiskStatus{
		Free:  freeBytesAvailable,
		Total: totalNumberOfBytes,
		Used:  totalNumberOfBytes - freeBytesAvailable,
	}
}
