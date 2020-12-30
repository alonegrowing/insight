package dao

import "context"

type Topic struct {
	ID      int64  `gorm:"id" json:"id"`
	Name    string `gorm:"name" json:"name"`
	Img     string `gorm:"img" json:"img"`
	Content string `gorm:"content" json:"content"`
	Created int64  `gorm:"created_at" json:"created"`
}

type Poem struct {
	Id        int64  `gorm:"id" json:"id"`
	Token     string `gorm:"token" json:"token"`
	Title     string `gorm:"title" json:"title"`
	Author    string `gorm:"author" json:"author"`
	Imgs      string `gorm:"imgs" json:"imgs"`
	Intro     string `gorm:"intro" json:"intro"`
	Content   string `gorm:"content" json:"content"`
	Status    string `gorm:"status" json:"status"`
	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
	CreatedAt string `gorm:"created_at" json:"created_at"`
}

type PoemDao interface {
	GetPoemById(ctx context.Context, id int64) (poem *Poem)
	GetPoemList(ctx context.Context) (poems []*Poem)
}
