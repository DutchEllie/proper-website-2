package components

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"git.home.dutchellie.nl/DutchEllie/proper-website-2/entity"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

/*
What is this supposed to do:
- It should call on the API to give it a certain number of comments, in the range x to y (this has to be implemented in the api)
- When it has called that, it should display those
- Dynamic links are there to navigate the user along the pages
- Comments are shown dynamically
- This panel can be shown or hidden (maybe)

AND VERY IMPORTANT!
- If a user submits a new comment, automatically put it on the page, no reloading

*/
type guestbookPanel struct {
	app.Compo

	comments []entity.Comment
}

func newGuestbookPanel() *guestbookPanel {
	g := &guestbookPanel{}
	g.LoadComments()
	return g
}

func (g *guestbookPanel) Render() app.UI {
	return app.Div().Body(
		app.Range(g.comments).Slice(func(i int) app.UI {
			return &guestbookComment{
				Comment: g.comments[i],
			}
		}),
	).Class("content gbp")
}

func (g *guestbookPanel) LoadComments() {
	// TODO: maybe you can put this in a localbrowser storage?
	url := apiurl + "api/comment"
	res, err := http.Get(url)
	if err != nil {
		app.Log(err)
		return
	}
	defer res.Body.Close()
	jsondata, err := io.ReadAll(res.Body)
	if err != nil {
		app.Log(err)
		return
	}

	err = json.Unmarshal(jsondata, &g.comments)
	if err != nil {
		app.Log(err)
		return
	}
}

type guestbookComment struct {
	app.Compo

	Comment entity.Comment
	time    string
}

func (c *guestbookComment) Render() app.UI {
	c.time = c.Comment.PostDate.Format(time.RFC1123)
	return app.Div().Body(
		app.Div().Class().Body(
			app.P().Text(c.Comment.Name).Class("name"),
			app.P().Text(c.time).Class("date"),
		).Class("comment-header"),
		app.Div().Class("comment-message").Body(
			app.P().Text(c.Comment.Message),
		),
	).Class("comment")
}