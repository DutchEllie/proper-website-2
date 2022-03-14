package components

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

	showGuestbook bool

	page string
}

func NewHomepage() *Homepage {
	return &Homepage{showGuestbook: true, page: "home"}
}

func (p *Homepage) Render() app.UI {
	gbp := newGuestbookPanel()
	return app.Div().Body(
		&header{},
		&navbar{},
		&homePanel{
			onShowClick: func() {
				p.showGuestbook = !p.showGuestbook
			},
		},
		&bannerPanel{},
		&guestbookForm{
			OnSubmit: func(name, message string) {
				var comment entity.Comment
				comment.Name = name
				comment.Message = message

				jsondata, err := json.Marshal(comment)
				if err != nil {
					fmt.Printf("err: %v\n", err)
					return
				}
				url := ApiURL

				req, err := http.Post(url, "application/json", bytes.NewBuffer(jsondata))
				if err != nil {
					fmt.Printf("err: %v\n", err)
					return
				}
				if req.StatusCode == 200 {
					p.Update()
				}
				defer req.Body.Close()
			},
		},
		//app.If(p.showGuestbook, gbp),
		gbp.Render(),
	).Class("main")
}
