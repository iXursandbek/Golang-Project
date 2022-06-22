package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	Id       int
	Name     string
	TaskName string
}

type UserStatus struct {
	UserId   int
	StatusId int
}

func main() {
	dsn := "host=localhost user=postgres password=x3127106 dbname=fruitdb port=5432 sslmode=disable TimeZone=Asia/Tashkent"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}

	user := User{Name: "Khurshid", TaskName: "Bir ish"}
	us := UserStatus{UserId: 2, StatusId: 3}
	// user := User{Name: "Palonkas", TaskName: "web"}
	// Continuous session mode
	// tx := db.Session(&gorm.Session{SkipDefaultTransaction: true})
	// tx.First(&user, 1)
	// tx.Create(user)
	// tx.Find(&user)
	// tx.Model(&user).Update("Age", 18)
	// fmt.Println(user)
	db.Transaction(func(tx *gorm.DB) error {
		if err != nil {
			err := tx.Create(&user).Error
			if err != nil {
				return err
			}
			err = tx.Create(&us).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}
