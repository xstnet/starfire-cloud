package crypto

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Md5FilePath(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Errorf("打开文件失败，filename=%v, err=%v", path, err)
		return "", err
	}
	defer file.Close()
	md5h := md5.New()
	io.Copy(md5h, file)

	return hex.EncodeToString(md5h.Sum(nil)), nil
}

func Md5File(file io.Reader) string {
	md5h := md5.New()
	io.Copy(md5h, file)

	return hex.EncodeToString(md5h.Sum(nil))
}
