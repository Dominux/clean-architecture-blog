package web

import (
	"github.com/Dominux/clean_architecture_blog/internal/services"
	"github.com/Dominux/clean_architecture_blog/internal/store"
	"github.com/Dominux/clean_architecture_blog/internal/views"
)


type Views struct {
	store store.Store
	postViews *PostViews
}

func New(s store.Store) *Views {
	return &Views{store: s}
} 

func (v *Views) Post() views.PostViews {
	if v.postViews == nil {
		v.postViews = &PostViews{
			service: services.NewPostService(v.store),
		}
	}
	return v.postViews
}

