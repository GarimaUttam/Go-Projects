// this file is for connecting to the database

package config

import(
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/gorm"
    "gorm.io/driver/mysql"
)

var (
	db * gorm.DB
)

// func Connect(){
// 	d, err :=  gorm.Open("mysql", "root:password/simplerest?charset=utf8parseTime=True&loc=Local")
// 	if err != nil{
// 		panic(err)
// 	}
// 	db = d
// }

func Connect() {
    dsn := "root:password@tcp(127.0.0.1:3306)/simplerest?charset=utf8mb4&parseTime=True&loc=Local"
    d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database: " + err.Error())
    }
    db = d
}


func GetDB() *gorm.DB{
	return db
}