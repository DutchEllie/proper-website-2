package templatelib

import (
	"html/template"
	"path/filepath"
)

type TemplateLibrary struct {
	Templates map[string]*template.Template
}

func NewTemplateLibrary(dir string) (TemplateLibrary, error) {
	t := make(map[string]*template.Template)

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return TemplateLibrary{}, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.ParseFiles(page)
		if err != nil {
			return TemplateLibrary{}, err
		}

		// Use the ParseGlob method to add any 'layout' templates to the
		// template set (in our case, it's just the 'base' layout at the
		// moment).
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return TemplateLibrary{}, err
		}

		// Use the ParseGlob method to add any 'partial' templates to the
		// template set (in our case, it's just the 'footer' partial at the
		// moment).
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return TemplateLibrary{}, err
		}
		t[name] = ts
	}

	return TemplateLibrary{Templates: t}, nil
}
