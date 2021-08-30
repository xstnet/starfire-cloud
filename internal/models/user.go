package models

import (
	"log"

	"github.com/xstnet/starfire-cloud/internal/common"
	"github.com/xstnet/starfire-cloud/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseModel
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Nickname   string `json:"nickname"`
	TotalSpace uint64 `json:"total_space"` // 为用户分配的最大存储空间， 若为0则代表不限制
	UsedSpace  uint64 `json:"used_space"`  // 已上传的空间占用
}

// 用户注册，写入到DB
func (u *User) Register() error {
	u.Password, _ = u.HashAndSalt(u.Password)
	result := u.DB().Create(u)
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
	diskInfo := utils.DiskInfo("E:")
	if diskInfo == nil {
		diskInfo = &utils.DiskStatus{}
	}

	return map[string]interface{}{
		"id":            u.ID,
		"username":      u.Username,
		"email":         u.Email,
		"nickname":      u.Nickname,
		"total_space":   u.TotalSpace,
		"used_space":    u.UsedSpace,
		"disk_info":     diskInfo,
		"register_time": common.FormatTimestamp(int64(u.CreatedAt)),
	}
}
