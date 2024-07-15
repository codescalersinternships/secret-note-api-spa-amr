package model

import "time"

// Note struct to model note data
type Note struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	URL         string    `json:"url"`
	Content     string    `json:"content"`
	PublishDate time.Time `json:"publish_date"`
	ExpireAfter int       `json:"expire_after"`
	Views       int       `json:"views"`
	MaxViews    int       `json:"max_views"`
}

// NoteIsExpired method to check if a note is expired based on max views and expiration time
func (n *Note) NoteIsExpired() bool {
	expirationTime := n.PublishDate.Add(time.Duration(n.ExpireAfter) * time.Hour)
	if n.Views >= n.MaxViews || time.Now().After(expirationTime) {
		return true
	}
	return false
}

