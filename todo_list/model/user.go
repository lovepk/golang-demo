package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"unique"`
	PasswrodDigest string  // 存储的是密文
}

// 加密
func (user *User) SetPassword(password string) error {
	byte, err := bcrypt.GenerateFromPassword([]byte(password), 12);
	if err != nil {
		return err
	}
	user.PasswrodDigest = string(byte)
	return nil
}

// 解密
func (user *User) CheckPassword(password string) bool  {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswrodDigest), []byte(password))
	return err == nil
}

