package web

import (
	"github.com/Dominux/clean_architecture_blog/internal/store"
	"github.com/Dominux/clean_architecture_blog/internal/views"
)


type Views struct {
	store *store.Store
	postViews *PostViews
}

func New(store *store.Store) *Views {
	return &Views{store: store}
} 

func (v *Views) Post() *views.PostViews {
	if v.postViews == nil {
		v.postViews = NewPostViews(v)
	}
	return v.postViews
}

