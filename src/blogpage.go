package main

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type BlogPage struct {
	app.Compo
}

// TODO: write the backend API for this
// TODO: Find a proper way of rendering blog posts in markdown
// Backend ideas: In the DB, create an entry for each post and a link to where the html file is located!
// 	That way, I don't have to parse and render markdown!!

// Layout, the leftbar contains the blogpost links and the mainbar contains the post itself!
// Function: After pressing a link for a blog post, that blog post ID gets put in the state instead of the URL

func NewBlogPage() *BlogPage {
	return &BlogPage{}
}

func (b *BlogPage) Render() app.UI {
	return newPage().
		Title("Blog").
		LeftBar(
			newHTMLBlock().
				Class("left leftbarblock"),
		)
}
