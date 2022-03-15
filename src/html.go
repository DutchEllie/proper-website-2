package main

import (
	"errors"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

const (
	getHTML = "/html/get"
)

func handleGetHTML(ctx app.Context, a app.Action) {
	path := a.Tags.Get("path")
	if path == "" {
		app.Log(errors.New("getting html failed"))
		return
	}

	state := htmlState(path)

	var ht htmlContent
	ctx.GetState(state, &ht)
	switch ht.Status {
	case loading, loaded:
		return
	}

	ht.Status = loading
	ht.Error = nil
	ctx.SetState(state, ht)

	res, err := get(ctx, path)
	if err != nil {
		ht.Status = loadingErr
		ht.Error = err
		ctx.SetState(state, ht)
		return
	}

	ht.Status = loaded
	ht.Data = string(res)
	ctx.SetState(state, ht)
}

func htmlState(src string) string {
	return src
}

type htmlContent struct {
	Status status
	Error  error
	Data   string
}

type status int

const (
	neverLoaded status = iota
	loading
	loadingErr
	loaded
)
