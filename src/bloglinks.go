package main

import (
	"encoding/json"
	"io"
	"net/http"

	"dutchellie.nl/DutchEllie/proper-website-2/entity"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type blogLinks struct {
	app.Compo

	blogposts []entity.BlogPost
}

func (b *blogLinks) OnMount(ctx app.Context) {
	b.LoadPosts(ctx)
}

func (b *blogLinks) Render() app.UI {
	return newUIBlock().
		Class("left").
		Class("leftbarblock").
		Class("blogpost-bar").
		UI(
			app.P().
				Class("p-h3").
				Style("margin-left", "10px").
				Style("margin-top", "10px").
				Style("text-decoration", "underline").
				Text("Posts"),
			app.Range(b.blogposts).Slice(func(i int) app.UI {
				return app.P().
					Class("blogpost-titles").
					Style("cursor", "pointer").
					Text(b.blogposts[i].Title).
					OnClick(func(ctx app.Context, e app.Event) {
						//e.PreventDefault()

						// Calling the update-blogpost action defined in the blogpage.go file.
						// There it's updated
						ctx.NewAction("update-blogpost", app.T("blogpost", b.blogposts[i].Path))
					})
			}),
		)
}

func (b *blogLinks) LoadPosts(ctx app.Context) {
	// TODO: maybe you can put this in a localbrowser storage?
	url := ApiURL + "/blogpost"
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
			err = json.Unmarshal(jsondata, &b.blogposts)
			if err != nil {
				app.Log(err)
				return
			}
		})
	})
}
