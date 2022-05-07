package models

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseModel
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Nickname   string `json:"nickname"`
	TotalSpace uint64 `json:"totalSpace"` // 为用户分配的最大存储空间， 若为0则代表不限制
	UsedSpace  uint64 `json:"usedSpace"`  // 已上传的空间占用
}

// Register 用户注册，写入到DB
func (u *User) Register() error {
	u.Password, _ = u.HashAndSalt(u.Password)
	result := u.DB().Create(u)
	return result.Error
}

// HashAndSalt 密码hash并自动加盐
// DefaultCost=10, 大约耗时40-50ms
func (u *User) HashAndSalt(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(hash), err
}

// ComparePasswords 比对用户密码是否正确
// DefaultCost=10, 大约耗时40-50ms
func (u *User) ComparePasswords(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		log.Println(err)
		return false
	}
	return true
}

// GetUserById Get
func (u *User) GetUserById(id uint) error {
	err := u.DB().Where(id).First(u).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *User) UpdateUsedSpace(size uint64) error {
	u.UsedSpace += size
	return u.DB().Model(u).Update("used_space", u.UsedSpace).Error
}

func (u *User) ChangePassword() error {
	u.Password, _ = u.HashAndSalt(u.Password)
	return u.DB().Model(u).Update("password", u.Password).Error
}
