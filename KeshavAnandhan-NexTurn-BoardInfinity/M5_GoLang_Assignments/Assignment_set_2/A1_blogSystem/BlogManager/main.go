package main

import (
	"blogmanager/controller"
	"blogmanager/db"
	"blogmanager/middleware"
	"blogmanager/repository"
	"blogmanager/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	err := db.InitDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	blogRepo := repository.NewBlogRepo(db.GetDatabase())
	blogService := service.NewBlogSvc(blogRepo)
	blogHandler := controller.NewBlogHandler(blogService)
	r := gin.Default()
	r.Use(middleware.LogRequest())
	r.Use(middleware.AuthMiddleware(db.GetDatabase()))
	r.GET("/blogs", blogHandler.GetAllBlogs)
	r.GET("/blog/:id", blogHandler.GetBlogByID)
	r.POST("/blog", blogHandler.CreateNewBlog)
	r.PUT("/blog/:id", blogHandler.UpdateBlog)
	r.DELETE("/blog/:id", blogHandler.DeleteBlog)
	r.Run(":8080")
}
