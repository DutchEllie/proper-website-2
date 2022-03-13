package components

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"dutchellie.nl/DutchEllie/proper-website-2/entity"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

const (
	apiurl = "https://api.nicecock.eu/"
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
				url := apiurl + "api/comment"
				if os.Getenv("TESTING") == "true" {
					url = apiurl + "api/testingcomment"
				}

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
