package models

import DB "graphql/database"

type User struct {
	ID   uint   `gorm:"primeryKey" json:"id,string"`
	Name string `gorm:"type:string" json:"name,string"`
}

func init() {
	err := DB.Con().AutoMigrate(&User{})
	if err != nil {
		panic(err.Error())
	}
}
