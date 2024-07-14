package model

import "time"

type Note struct {
	ID          uint
	URL         string
	Content     string
	PublishDate time.Time
	ExpireAfter time.Time
	Views       int
	MaxViews    int
}

func (n Note) NoteIsExpired() bool {
	if n.MaxViews >= n.Views || n.ExpireAfter.After(n.PublishDate) {
		return false
	}
	return true
}
