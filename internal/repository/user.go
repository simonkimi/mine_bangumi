package repository

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"default:''"`
	Password string `gorm:"default:''"`
}

type UserExtraAuth struct {
	gorm.Model
	UserID     uint64
	User       *User  `gorm:"foreignKey:UserID"`
	Type       string `gorm:"type:enum('totp', 'passkey', 'recovery_code')"`
	Identifier string `gorm:"type:varchar(255)"`
	Secret     string `gorm:"type:varchar(255)"`
}

func AuthUser(username string, password string) {

}
