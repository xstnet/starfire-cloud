package models

import (
	"github.com/xstnet/starfire-cloud/pkg/helper/d"
	"log"

	"github.com/xstnet/starfire-cloud/configs"
	"github.com/xstnet/starfire-cloud/internal/common"
	"github.com/xstnet/starfire-cloud/pkg/systeminfo"
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
	u.GetScene()
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
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		log.Println(err)
		return false
	}
	return true
}

// Get
func (u *User) GetUserById(id uint) error {
	err := u.DB().Where(id).First(u).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *User) UpdateUsedSpace(size uint64) error {
	u.UsedSpace += size
	return u.DB().Model(u).Update("used_space", u.UsedSpace).Error
}

// 用户信息转化
func (u *User) ToDetail() d.StringMap {
	// 处理总存储空间
	diskInfo := systeminfo.DiskInfo(configs.Upload.UploadRootPath)
	return d.StringMap{
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

func (u *User) ChangePassword() error {
	u.Password, _ = u.HashAndSalt(u.Password)
	return u.DB().Model(u).Update("password", u.Password).Error
}
