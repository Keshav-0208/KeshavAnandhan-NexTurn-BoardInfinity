package service

import (
	"blogmanager/model"
	"blogmanager/repository"
)

type BlogSvc struct {
	BlogRepo *repository.BlogRepo
}

func NewBlogSvc(blogRepo *repository.BlogRepo) *BlogSvc {
	return &BlogSvc{BlogRepo: blogRepo}
}

func (svc *BlogSvc) CreateBlog(blog *model.Blog) (*model.Blog, error) {
	return svc.BlogRepo.Create(blog)
}

func (svc *BlogSvc) GetBlog(id int) (*model.Blog, error) {
	return svc.BlogRepo.Get(id)
}

func (svc *BlogSvc) GetAllBlogs() ([]model.Blog, error) {
	return svc.BlogRepo.GetAll()
}

func (svc *BlogSvc) UpdateBlog(blog *model.Blog) (*model.Blog, error) {
	return svc.BlogRepo.Update(blog)
}

func (svc *BlogSvc) DeleteBlog(id int) error {
	return svc.BlogRepo.Delete(id)
}
