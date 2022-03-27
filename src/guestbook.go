package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"dutchellie.nl/DutchEllie/proper-website-2/entity"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type guestbook struct {
	app.Compo

	comments []entity.Comment

	name    string
	email   string
	website string
	message string

	lastHash    [32]byte
	gbModalOpen bool
	OnSubmit    func(
		ctx app.Context,
		name string,
		email string,
		website string,
		message string,
	) // Handler to implement which calls the api
}

// TODO: The comments are loaded like 2 or 3 times every time the page is loaded...
func (g *guestbook) OnMount(ctx app.Context) {
	ctx.Handle("guestbook-loadcomments", g.onHandleLoadComments)
	ctx.NewAction("guestbook-loadcomments")
}

/*
func (g *guestbook) OnNav(ctx app.Context) {
	g.LoadComments(ctx)
}

func (g *guestbook) OnUpdate(ctx app.Context) {
	g.LoadComments(ctx)
}*/

func (g guestbook) Render() app.UI {
	return app.Div().Body(
		app.Form().
			Class("guestbook-form").
			Body(
				app.Div().
					Class("input-groups").
					Body(
						app.Div().
							Class("fr").
							Body(
								app.Div().
									Class("input-group").
									Class("input-group-name").
									Body(
										app.Label().
											For("name").
											Text("Name:"),
										app.Input().
											Type("text").
											Name("name").
											Class("input").
											Required(true).
											OnChange(g.ValueTo(&g.name)),
									),
								app.Div().
									Class("input-group").
									Class("input-group-email").
									Body(
										app.Label().
											For("email").
											Text("Email: (optional)"),
										app.Input().
											Type("text").
											Name("email").
											Class("input").
											Required(false).
											OnChange(g.ValueTo(&g.email)),
									),
							),
						app.Div().
							Class("input-group").
							Class("input-group-website").
							Body(
								app.Label().
									For("website").
									Text("Website: (optional)"),
								app.Input().
									Type("text").
									Name("website").
									Class("input").
									Required(false).
									OnChange(g.ValueTo(&g.website)),
							),
						app.Div().
							Class("input-group").
							Class("input-group-message").
							Body(
								app.Label().
									For("message").
									Text("Message:"),
								app.Textarea().
									Name("message").
									Class("input").
									Rows(5).
									Cols(30).
									Required(true).
									OnChange(g.ValueTo(&g.message)),
							),
						app.Div().
							Class("submit-field").
							Body(
								app.Input().
									Type("submit").
									Value("Send!"),
							),
					),
			).OnSubmit(func(ctx app.Context, e app.Event) {
			// This was to prevent the page from reloading
			e.PreventDefault()
			if g.name == "" || g.message == "" {
				fmt.Printf("Error: one or more field(s) are empty. For now unhandled\n")
				return
			}
			if len(g.name) > 40 || len(g.message) > 360 {
				fmt.Printf("Error: Your message is too long fucker\n")
				g.gbModalOpen = true
				return
			}
			g.OnSubmit(ctx, g.name, g.email, g.website, g.message)
			ctx.Dispatch(func(ctx app.Context) {
				g.clear()
			})
			ctx.NewAction("guestbook-loadcomments")
			//g.LoadComments(ctx)
		}),
		app.If(
			g.gbModalOpen,
			&guestbookAlertModal{
				OnClose: func() {
					g.gbModalOpen = false
					g.Update()
				},
			},
		),
		app.Div().Body(
			app.Range(g.comments).Slice(func(i int) app.UI {
				return &guestbookComment{
					Comment: g.comments[i],
				}
			},
			),
		),
	)
}

func (g *guestbook) SmartLoadComments(ctx app.Context) {
	var lasthash []byte
	err := ctx.LocalStorage().Get("lasthash", &lasthash)
	if err != nil {
		app.Log(err)
		return
	}

	if lasthash == nil {
		fmt.Printf("Program thinks lasthash is empty\n")
		g.LoadComments(ctx)
		return
	}
	url := ApiURL + "/commenthash"
	ctx.Async(func() {
		res, err := http.Get(url)
		if err != nil {
			app.Log(err)
			return
		}
		defer res.Body.Close()

		hash, err := io.ReadAll(res.Body)
		if err != nil {
			app.Log(err)
			return
		}

		fmt.Printf("hash: %v\n", hash)
		fmt.Printf("lasthash: %v\n", lasthash)

		// If the hash is different, aka there was an update in the comments
		if !bytes.Equal(hash, lasthash) {
			fmt.Printf("Hash calculation is different\n")
			g.LoadComments(ctx)
			return
		}

		// if the hash is not different, then there is no need to load new comments
		jsondata := make([]byte, 0)
		err = ctx.LocalStorage().Get("comments", &jsondata)
		if err != nil {
			app.Log(err)
			return
		}
		fmt.Printf("jsondata: %v\n", jsondata)
		ctx.Dispatch(func(ctx app.Context) {
			err = json.Unmarshal(jsondata, &g.comments)
			if err != nil {
				app.Log(err)
				return
			}
		})
		return
	})
}

func (g *guestbook) LoadComments(ctx app.Context) {
	// TODO: maybe you can put this in a localbrowser storage?
	fmt.Printf("Called LoadComments()\n")
	url := ApiURL + "/comment"
	ctx.Async(func() {
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

		ctx.Dispatch(func(ctx app.Context) {
			err = json.Unmarshal(jsondata, &g.comments)
			if err != nil {
				app.Log(err)
				return
			}
		})

		ctx.LocalStorage().Set("comments", jsondata)
		// Calculating the hash
		fmt.Printf("Calculating the hash from LoadComments\n")
		hash := sha256.Sum256(jsondata)
		fmt.Printf("hash fresh from calculation: %v\n", hash)
		//g.lastHash = hash
		err = ctx.LocalStorage().Set("lasthash", []byte(fmt.Sprintf("%x\n", hash)))
		if err != nil {
			app.Log(err)
			return
		}
	})
}

func (g *guestbook) clear() {
	g.name = ""
	g.message = ""
}

func (g *guestbook) onHandleLoadComments(ctx app.Context, a app.Action) {
	g.SmartLoadComments(ctx)
	ctx.Dispatch(func(ctx app.Context) {
		g.Update()
	})
}

type guestbookAlertModal struct {
	app.Compo

	PreviousAttempts int
	OnClose          func() // For when we close the modal
}

func (g *guestbookAlertModal) Render() app.UI {
	return app.Div().
		Class("gb-modal").
		ID("gbModal").
		OnClick(func(ctx app.Context, e app.Event) {
			g.OnClose()
		}).
		Body(
			app.Div().
				Class("gb-modal-content").
				Body(
					app.Span().Class("close").Text("X").
						OnClick(func(ctx app.Context, e app.Event) {
							//modal := app.Window().GetElementByID("gbModal")
							//modal.Set("style", "none")
							g.OnClose()
						}),
					app.P().Text("Your name must be <= 40 and your message must be <= 360 characters"),
				),
		)
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
