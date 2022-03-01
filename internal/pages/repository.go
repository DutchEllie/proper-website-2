package pages

import (
	"context"
	"html/template"

	"git.home.dutchellie.nl/DutchEllie/proper-website-2/internal/templatelib"
)

type Repository interface {
	Page(ctx context.Context, name string) (*template.Template, error)
}

func NewRepository(lib *templatelib.TemplateLibrary) Repository {
	return repository{lib}
}

type repository struct {
	library *templatelib.TemplateLibrary
}

func (r repository) Page(ctx context.Context, name string) (*template.Template, error) {
	return r.library.Templates[name], nil
}
