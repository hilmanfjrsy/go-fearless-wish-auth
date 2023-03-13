package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"-"`
	Email    string `gorm:"size:255;not null;unique" json:"email"`
}

type UserOrm struct {
	Db *gorm.DB
}

func (d *UserOrm) GetById(id uint) (User, error) {
	user := User{}
	err := d.Db.Model(&User{}).Where("id = ?", id).First(&user).Error
	return user, err
}

func (d *UserOrm) GetByUsername(username string) (User, error) {
	user := User{}
	err := d.Db.Model(&User{}).Where("username = ?", username).First(&user).Error
	return user, err
}

func (d *UserOrm) InsertUser(user User) (uint, error) {
	err := d.Db.Model(&User{}).Create(&user).Error
	return user.ID, err
}

func (d *UserOrm) UpdateUser(user User) error {
	return d.Db.Model(&user).Updates(&user).Error
}
