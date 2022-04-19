// Idea for this page:
// - Make a navbar on the top for different genres and switch the pages content when clicked

package main

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type MusicPage struct {
	app.Compo
}

func NewMusicPage() *MusicPage {
	return &MusicPage{}
}

func (f *MusicPage) Render() app.UI {
	return newPage().
		Title("Music!").
		LeftBar(
			newHTMLBlock().
				Class("left").
				Class("leftbarblock").
				Src("/web/blocks/snippets/bannerpanel.html"),
		).
		Main(
			// Genre navbar above this
			newUIBlock().
				Class("right").
				Class("contentblock").
				UI(
					app.Div().
						Body(
							app.P().
								Class("m-t5").
								Text(`I am quite picky with my music most of the time. I rarely enjoy an entire album of an artist and most artists for me have only a couple amazing songs.
						My tastes in music are almost exclusively Japanese songs. Vocaloid is how I began and nowadays I listen to all sorts of Japanese music.
						Here are some of the songs, artists and albums I like the most.`),
							app.P().
								Class("p-h3").
								Style("color", "red").
								Text("Warning! Player feature still in beta. Stuff can break and design is most certainly not final at all!"),
							app.P().
								Text("Just click one of the songs to play it."),
							app.P().
								Class("p-h2").
								Text("Songs"),
							newMusicPlayer(),
						),
				),
		)
}
