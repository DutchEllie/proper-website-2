package main

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type EmptyPage struct {
	app.Compo
}

func NewEmptyPage() *EmptyPage {
	return &EmptyPage{}
}

func (e *EmptyPage) Render() app.UI {
	return app.Head().Body()
}
