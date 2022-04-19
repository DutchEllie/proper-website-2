package main

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

const (
	songRepoURL string = "" // URL where the music files are stored. Made a bit like the ApiURL
)

type song struct {
	ITitle string
	IID    string
	IURL   string
}

func newSong() *song {
	return &song{}
}

func (f *song) Title(v string) *song {
	f.ITitle = v
	return f
}

func (f *song) URL(v string) *song {
	f.IURL = v
	return f
}

func (f *song) ID(v string) *song {
	f.IID = v
	return f
}

type musicPlayer struct {
	app.Compo

	songs                 map[string](*song)
	currentlySelectedSong string
}

func newMusicPlayer() *musicPlayer {
	return &musicPlayer{
		songs: make(map[string]*song),
	}
}

// On... handlers

func (f *musicPlayer) OnMount(ctx app.Context) {
	ctx.Handle("switchSong", f.handleSwitchSong)
	// Statically create all the music.
	// I am not making a database for this shit lmao
	// Also do not forget to at least set the currentlySelectedSong to the first one added, or at least make it point somewhere
	f.songs["ievan-polka"] = newSong().
		Title("Ievan Polka - Hatsune Miku").
		URL("https://files.catbox.moe/lh229f.mp3").
		ID("ievan-polka")
	f.songs["tokusya-seizon"] = newSong().
		Title("Tokusya-Seizon Wonder-la-der!! - Amane Kanata").
		URL("https://music-website.s3.nl-ams.scw.cloud/Tokusya-Seizon%20Wonder-la-der%21%21.mp3").
		ID("tokusya-seizon")

}

// Action handlers

// Call with a value called "title" to switch to the right song
func (f *musicPlayer) handleSwitchSong(ctx app.Context, a app.Action) {
	title, ok := a.Value.(string)
	if !ok {
		app.Log("Error calling handleSwitchSong function. Title value was not found")
		return
	}
	v, ok := f.songs[title]
	if !ok {
		app.Log("Error getting song. Song with title does not exist")
		return
	}

	f.currentlySelectedSong = v.IID
	f.Update()
}

func (f *musicPlayer) Render() app.UI {
	// Don't forget to handle the possibility of no songs having been added and the currentlySelectedSong to be empty
	if f.currentlySelectedSong == "" {
		return app.Div().
			Body(
				app.Range(f.songs).Map(func(s string) app.UI {
					return app.Div().
						Style("border", "solid 1px red").
						Style("width", "fit-content").
						Body(
							app.Span().
								Style("text-decoration", "underline").
								Style("width", "fit-content").
								Text(f.songs[s].ITitle).
								OnClick(func(ctx app.Context, e app.Event) {
									ctx.NewActionWithValue("switchSong", f.songs[s].IID)
								}),
						)
				}),
			)
	}
	return app.Div().Body(
		app.P().
			Class("p-h3").
			Text(fmt.Sprintf("Currently playing: %s", f.songs[f.currentlySelectedSong].ITitle)),
		app.Audio().
			Src(f.songs[f.currentlySelectedSong].IURL).
			Controls(true).
			AutoPlay(true),
		// Lots of buttons of songs
		app.Div().
			Body(
				app.Range(f.songs).Map(func(s string) app.UI {
					return app.Div().
						Style("border", "solid 1px red").
						Style("width", "fit-content").
						Body(
							app.Span().
								Style("text-decoration", "underline").
								Style("width", "fit-content").
								Text(f.songs[s].ITitle).
								OnClick(func(ctx app.Context, e app.Event) {
									ctx.NewActionWithValue("switchSong", f.songs[s].IID)
								}),
						)
				}),
			),
	)
}
