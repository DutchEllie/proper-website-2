package pages

import (
	"context"
	"html/template"
)

type Service interface {
	Page(ctx context.Context, name string) (Page, error)
}

type Page struct {
	template *template.Template
}

type service struct {
	repository Repository
	//logger     log.Logger
}

func NewService(r Repository) Service {
	return service{r}
}

func (s service) Page(ctx context.Context, name string) (Page, error) {
	page, err := s.repository.Page(ctx, name)
	if err != nil {
		return Page{}, err
	}
	return Page{page}, nil
}
