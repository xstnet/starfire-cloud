package models

import (
	"log"

	"github.com/xstnet/starfire-cloud/internal/common"
	"github.com/xstnet/starfire-cloud/internal/db"
	"github.com/xstnet/starfire-cloud/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseField
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Nickname   string `json:"nickname"`
	TotalSpace uint64
	UsedSpace  uint64
}

func (u *User) Register() error {
	u.Password, _ = u.HashAndSalt(u.Password)
	result := db.DB.Create(u)
	return result.Error
}

// 密码hash并自动加盐
// DefaultCost=10, 大约耗时40-50ms
func (u *User) HashAndSalt(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(hash), err
}

// 比对用户密码是否正确
// DefaultCost=10, 大约耗时40-50ms
func (u *User) ComparePasswords(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// 用户信息转化
func (u *User) ToDetail() map[string]interface{} {
	// 处理总存储空间
	totalSpace := u.TotalSpace
	if totalSpace == 0 {
		if diskInfo := utils.DiskInfo("E:"); diskInfo != nil {
			// 没有对该用户限制存储容量，使用整个磁盘的空间做为存储容量
			totalSpace = diskInfo.Total
		}
	}

	// todo: 应该把剩余空间传到前端
	return map[string]interface{}{
		"id":               u.ID,
		"username":         u.Username,
		"email":            u.Email,
		"nickname":         u.Nickname,
		"totalSpace":       totalSpace,
		"totalSpaceString": common.FormatFileSize(totalSpace),
		"usedSpace":        u.UsedSpace,
		"usedSpaceString":  common.FormatFileSize(u.UsedSpace),
	}
}
