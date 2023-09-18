package main

import (
	"flinkcdc-createData/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strconv"
	"time"
)

func main() {

	dsn := "etl:@98Kur69pvGf@tcp(10.122.37.15:9030)/test2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 开启 SQL 日志
		//Logger: logger.Default.LogMode(logger.Silent), // 设置为 logger.Silent 可以关闭日志
	})
	if err != nil {
		panic("无法连接数据库")
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	defer sqlDB.Close()

	//db.AutoMigrate(&model.User{})
	users := make([]model.User, 0)
	idx := 1
	fmt.Println("start time: ", time.Now().Format("2006-01-02 15:04:05"))
	for j := 1; j <= 400; j++ {
		for i := 1; i <= 25000; i++ {
			users = append(users, model.User{
				Id:   idx,
				Name: "user" + strconv.Itoa(idx),
			})
			idx++
		}
		if res := db.Create(&users); res.Error != nil {
			panic(res.Error)
		}
		users = users[:0]
	}
	fmt.Println("end time: ", time.Now().Format("2006-01-02 15:04:05"))
}
