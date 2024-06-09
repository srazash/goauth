package Post

import (
	"goauth/internal/User"
	"time"
)

type Post struct {
	pid       int
	title     string
	content   string
	author    *User.User
	published time.Time
	likes     int
	comments  int
}

func (p *Post) GetPid() int {
	return p.pid
}

func (p *Post) SetPid(pid int) {
	p.pid = pid
}

func (p *Post) GetTitle() string {
	return p.title
}

func (p *Post) SetTitle(title string) {
	p.title = title
}

func (p *Post) GetContent() string {
	return p.content
}

func (p *Post) SetContent(content string) {
	p.content = content
}

func (p *Post) GetAuthor() *User.User {
	return p.author
}

func (p *Post) SetAuthor(author *User.User) {
	p.author = author
}

func (p *Post) GetPublished() time.Time {
	return p.published
}

func (p *Post) SetPublished(published time.Time) {
	p.published = published
}

func (p *Post) IncrementLikes() {
	p.likes++
}

func (p *Post) GetLikes() int {
	return p.likes
}

func (p *Post) SetLikes(likes int) {
	p.likes = likes
}

func (p *Post) IncrementComments() {
	p.comments++
}

func (p *Post) GetComments() int {
	return p.comments
}

func (p *Post) SetComments(comments int) {
	p.comments = comments
}
