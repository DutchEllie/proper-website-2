package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"dutchellie.nl/DutchEllie/proper-website-2/entity"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

var (
	ApiURL string
)

type Homepage struct {
	app.Compo
}

func NewHomepage() *Homepage {
	return &Homepage{}
}

func (p *Homepage) Render() app.UI {
	return newPage().
		Title("Homepage").
		LeftBar(
			newHTMLBlock().
				Class("left").
				Class("leftbarblock").
				Src("/web/blocks/snippets/bannerpanel.html"),
		).
		Main(
			newHTMLBlock().
				Class("right").
				Class("contentblock").
				Src("/web/blocks/pages/intro.html"),
			newUIBlock().
				Class("right").
				Class("contentblock").
				UI(
					&guestbook{
						OnSubmit: func(ctx app.Context, name, email, website, message, uuid string) {
							var comment entity.Comment
							comment.Name = name
							comment.Email = email
							comment.Website = website
							comment.Message = message
							comment.UUID = uuid

							jsondata, err := json.Marshal(comment)
							if err != nil {
								fmt.Printf("err: %v\n", err)
								return
							}
							url := ApiURL + "/comment"

							// This is not Async'ed, because otherwise you run into a race
							// condition where you reload the comments before the server had time
							// to process the request!
							{
								req, err := http.Post(url, "application/json", bytes.NewBuffer(jsondata))
								if err != nil {
									fmt.Printf("err: %v\n", err)
									return
								}
								if req.StatusCode == 200 {
									p.Update()
								}
								defer req.Body.Close()
							}
						},
					},
				),
		)
}
