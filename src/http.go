package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func get(ctx app.Context, path string) ([]byte, error) {
	url := path
	if !strings.HasPrefix(url, "http") {
		u := ctx.Page().URL()
		u.Path = path
		url = u.String()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("Error at getting html page\n")
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Which means either client or server error
	if res.StatusCode >= 400 {
		return nil, err
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}
