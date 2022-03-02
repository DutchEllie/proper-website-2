package components

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"git.home.dutchellie.nl/DutchEllie/proper-website-2/entity"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Homepage struct {
	app.Compo

	showGuestbook    bool
	guestbookUpdated bool
}

func NewHomepage() *Homepage {
	return &Homepage{showGuestbook: true, guestbookUpdated: false}
}

func (p *Homepage) Render() app.UI {
	gbp := newGuestbookPanel()
	return app.Div().Body(
		&header{},
		app.Div().Body(
			newNavbar(),
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
					url := "/api/comment"
					req, err := http.Post(url, "application/json", bytes.NewBuffer(jsondata))
					if err != nil {
						fmt.Printf("err: %v\n", err)
						return
					}
					if req.StatusCode == 200 {
					}
					defer req.Body.Close()
				},
			},
			app.If(p.showGuestbook, gbp),
		).Class("main"),
	)
}
