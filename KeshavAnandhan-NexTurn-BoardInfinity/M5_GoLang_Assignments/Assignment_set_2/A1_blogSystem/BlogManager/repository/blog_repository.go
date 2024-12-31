package repository

import (
	"blogmanager/model"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type BlogRepo struct {
	DB *sql.DB
}

func NewBlogRepo(db *sql.DB) *BlogRepo {
	return &BlogRepo{DB: db}
}

func (repo *BlogRepo) Create(blog *model.Blog) (*model.Blog, error) {
	stmt, err := repo.DB.Prepare("INSERT INTO blog_posts (title, content, author, created_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(blog.Title, blog.Content, blog.Author, time.Now().String())
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	blog.ID = int(id)
	return blog, nil
}

func (repo *BlogRepo) Get(id int) (*model.Blog, error) {
	row := repo.DB.QueryRow("SELECT id, title, content, author, created_at FROM blog_posts WHERE id = ?", id)
	blog := &model.Blog{}
	err := row.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.CreatedAt)
	if err != nil {
		return nil, err
	}
	return blog, nil
}

func (repo *BlogRepo) GetAll() ([]model.Blog, error) {
	rows, err := repo.DB.Query("SELECT id, title, content, author, created_at FROM blog_posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var blogs []model.Blog
	for rows.Next() {
		var blog model.Blog
		err := rows.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		blogs = append(blogs, blog)
	}
	return blogs, nil
}

func (repo *BlogRepo) Update(blog *model.Blog) (*model.Blog, error) {
	stmt, err := repo.DB.Prepare("UPDATE blog_posts SET title = ?, content = ?, author = ?, created_at = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(blog.Title, blog.Content, blog.Author, time.Now().String(), blog.ID)
	if err != nil {
		return nil, err
	}
	return blog, nil
}

func (repo *BlogRepo) Delete(id int) error {
	stmt, err := repo.DB.Prepare("DELETE FROM blog_posts WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
