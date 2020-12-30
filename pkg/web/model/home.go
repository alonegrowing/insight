package model

type Poem struct {
	Id        int64    `json:"id"`
	Token     string   `json:"token"`
	Title     string   `json:"title"`
	Author    string   `json:"author"`
	Intro     string   `json:"intro"`
	Imgs      []string `json:"imgs"`
	Content   string   `json:"content"`
	UpdatedAt string   `json:"updated_at"`
	CreatedAt string   `json:"created_at"`
}
