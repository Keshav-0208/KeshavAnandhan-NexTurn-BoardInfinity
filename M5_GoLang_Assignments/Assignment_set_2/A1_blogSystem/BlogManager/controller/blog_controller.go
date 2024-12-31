package controller

import (
	"blogmanager/model"
	"blogmanager/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BlogHandler struct {
	BlogService *service.BlogService
}

func NewBlogHandler(blogService *service.BlogService) *BlogHandler {
	return &BlogHandler{BlogService: blogService}
}

func (handler *BlogHandler) CreateNewBlog(c *gin.Context) {
	fmt.Println("Creating a new blog post")

	var blog model.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		fmt.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	fmt.Printf("Parsed blog details: %+v\n", blog)

	createdBlog, err := handler.BlogService.CreateBlog(&blog)
	if err != nil {
		fmt.Printf("Error creating blog: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error creating blog"})
		return
	}

	c.JSON(http.StatusOK, createdBlog)
}

func (handler *BlogHandler) GetBlogByID(c *gin.Context) {
	id := c.Param("id")
	blogID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	blog, err := handler.BlogService.GetBlog(blogID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}

	c.JSON(http.StatusOK, blog)
}

func (handler *BlogHandler) GetAllBlogs(c *gin.Context) {
	blogs, err := handler.BlogService.GetAllBlogs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve blogs"})
		return
	}

	c.JSON(http.StatusOK, blogs)
}

func (handler *BlogHandler) UpdateBlog(c *gin.Context) {
	id := c.Param("id")
	blogID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	var blog model.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	blog.ID = blogID
	updatedBlog, err := handler.BlogService.UpdateBlog(&blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating blog"})
		return
	}

	c.JSON(http.StatusOK, updatedBlog)
}

func (handler *BlogHandler) DeleteBlog(c *gin.Context) {
	id := c.Param("id")
	blogID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	err = handler.BlogService.DeleteBlog(blogID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete blog"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog deleted successfully"})
}
