package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product2 struct {
	Name  string
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&Product2{})

	// Create
	db.Create(&Product2{Code: "A42", Price: 100, Name: "a book"})

	// Read
	var product Product2
	db.First(&product, 2)                 // 根据整形主键查找
	db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	db.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	db.Model(&product).Updates(Product2{Price: 200, Code: "F42"}) // 仅更新非零值字段
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "Z42"})

	// Delete - 删除 product
	db.Delete(&product, 1)
}
