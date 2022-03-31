package main

import (
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=pustaka password=pustaka dbname=pustaka port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db connection error")
	}

	db.AutoMigrate(&book.Book{})
	//CRUD
	//===========
	//CREATE data
	//===========

	// book := book.Book{}
	// book.Title = "Sebuah Seni bersikap Bodo Amat"
	// book.Price = 120000
	// book.Discount = 15
	// book.Rating = 5
	// book.Description = "Ini buku yang bagus dari Max Manson"

	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("=====================")
	// 	fmt.Println("Error creating book record")
	// 	fmt.Println("=====================")
	// }

	//CRUD
	//===========
	//READ data
	//===========

	// var books []book.Book

	// err = db.Debug().Where("title = ?", "Man Tiger").Find(&books).Error
	// if err != nil {
	// 	fmt.Println("=====================")
	// 	fmt.Println("Error finding book record")
	// 	fmt.Println("=====================")
	// }

	// for _, b := range books {
	// 	fmt.Println("Title:", b.Title)
	// 	fmt.Println("book object %v", b)
	// }

	//CRUD
	//===========
	//UPDATE data
	//===========

	// var book book.Book

	// err = db.Debug().Where("id = ?", 1).First(&book).Error
	// if err != nil {
	// 	fmt.Println("=====================")
	// 	fmt.Println("Error finding book record")
	// 	fmt.Println("=====================")
	// }

	// book.Title = "Man Tiger (Revised Edition)"
	// err = db.Save(&book).Error
	// if err != nil {
	// 	fmt.Println("=====================")
	// 	fmt.Println("Error Updating book record")
	// 	fmt.Println("=====================")
	// }

	//CRUD
	//===========
	//DELETE data
	//===========

	// var book book.Book

	// err = db.Debug().Where("id = ?", 1).First(&book).Error
	// if err != nil {
	// 	fmt.Println("=====================")
	// 	fmt.Println("Error finding book record")
	// 	fmt.Println("=====================")
	// }

	// err = db.Delete(&book).Error
	// if err != nil {
	// 	fmt.Println("=====================")
	// 	fmt.Println("Error Deleting book record")
	// 	fmt.Println("=====================")
	// }

	bookRepository := book.NewRepository(db)

	books, err := bookRepository.FindAll()

	for _, book := range books {
		fmt.Println("Title:", book.Title)
	}

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBookHandler)

	router.Run(":8888")
}
