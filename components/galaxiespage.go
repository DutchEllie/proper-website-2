package components

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type GalaxiesPage struct {
	app.Compo
}

func NewGalaxiesPage() *GalaxiesPage {
	return &GalaxiesPage{}
}

func (f *GalaxiesPage) Render() app.UI {
	return app.Div().Body(
		&header{},
		&navbar{},
		&galaxiesPanel{},
		&bannerPanel{},
	).Class("main")
}

type galaxiesPanel struct {
	app.Compo
}

func (b *galaxiesPanel) Render() app.UI {
	return app.Div().Body(
		app.P().
			Text(`Galaxies`).
			Class("p-h1"),
		app.P().
			Class("content-text").
			Text(`Here you can find some really really really cool pages that I found on the internet.
			Some of these are blogs or even blogposts I found, but the ones on top are special!
			They're the websites of friends of mine! Please visit them, because they worked really hard
			on their websites as well!`),
		app.Div().
			Body(
				app.P().
					Class("p-h2 mt-20 mb-10 bold").
					Text("My friends!"),
				app.Ul().Body(
					app.Li().Body(
						app.Div().Body(
							// TODO: Create a modal popup for each name!!!
							app.A().Href("https://forestofunix.xyz").
								Class("p-h3 m-t5").
								Text("Forest of Unix"),
							app.P().
								Class("m-t5").
								Text(`A website by Sebastiaan. A massive Linux fanboy, runs Gentoo on his
								ThinkPad. Absolutely based.`),
						),
					),
					app.Li().Body(
						app.Div().Body(
							// TODO: Create a modal popup for each name!!!
							app.A().Href("https://nymphali.neocities.org").
								Class("p-h3 m-t5").
								Text("Nymphali"),
							app.P().
								Class("m-t5").
								Text(`The website made by ■■■■■■, whoops Nymphali. They have an awesome
								minimalist website that's just lovely.`),
						),
					),
					app.Li().Body(
						app.Div().Body(
							// TODO: Create a modal popup for each name!!!
							app.A().Href("https://kristypixel.neocities.org").
								Class("p-h3 m-t5").
								Text("Kristy"),
							app.P().
								Class("m-t5").
								Text(`Website made by Kristy. Very cute website, I love it! Keep up the
								awesome work!`),
						),
					),
				),
			),
		app.Div().
			Body(
				app.P().
					Class("p-h2 mt-20 mb-10 bold").
					Text("Neat webspaces"),
				app.P().
					Class("m-t5").
					Style("margin-left", "10px").
					Text(`Just very neat websites I found. Not necessarily by people I know.
					I just thought it would be nice to share them here!`),
				app.Ul().Body(
					app.Li().Body(
						app.Div().Body(
							// TODO: Create a modal popup for each name!!!
							app.A().Href("https://evillious.ylimegirl.com/").
								Class("p-h3 m-t5").
								Text("Evillious Chronicles fan guide"),
							app.P().
								Class("m-t5").
								Text(`A VERY cool website made by Ylimegirl! They wrote a whole
								website dedicated to Evillious Chronicles, which is a super 
								good Japanese light novel and vocaloid series!! Definitely look it up!`),
						),
					),
				),
			),
	).Class("content")
}
