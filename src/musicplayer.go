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
	//	f.songs["ievan-polka"] = newSong().
	//		Title("Ievan Polka - Hatsune Miku").
	//		URL("https://files.catbox.moe/lh229f.mp3").
	//		ID("ievan-polka")
	f.songs["god-ish"] = newSong().
		Title("God-ish (神っぽいな) feat. Hatsune Miku - PinocchioP").
		URL("https://music-website.s3.nl-ams.scw.cloud/%E7%A5%9E%E3%81%A3%E3%81%BD%E3%81%84%E3%81%AA.m4a").
		ID("god-ish")
	f.songs["servant-of-evil"] = newSong().
		Title("Servant of Evil (悪ノ召使) feat. Kagamine Rin - mothy / AkunoP").
		URL("https://music-website.s3.nl-ams.scw.cloud/Servant%20of%20Evil%20with%20English%20Sub%20-%20%E6%82%AA%E3%83%8E%E5%8F%AC%E4%BD%BF%20-%20Kagamine%20Len%20-%20HQ.m4a").
		ID("servant-of-evil")
	f.songs["im-glad-youre-evil-too"] = newSong().
		Title("I'm glad you're evil too (feat. Hatsune Miku) - PinocchioP").
		URL("https://music-website.s3.nl-ams.scw.cloud/%E3%83%94%E3%83%8E%E3%82%AD%E3%82%AA%E3%83%94%E3%83%BC%20-%20%E3%81%8D%E3%81%BF%E3%82%82%E6%82%AA%E3%81%84%E4%BA%BA%E3%81%A7%E3%82%88%E3%81%8B%E3%81%A3%E3%81%9F%20feat.%20%E5%88%9D%E9%9F%B3%E3%83%9F%E3%82%AF%20_%20I%27m%20glad%20you%27re%20evil%20too.m4a").
		ID("im-glad-youre-evil-too")
	f.songs["tokusya-seizon"] = newSong().
		Title("Tokusya-Seizon Wonder-la-der!! - Amane Kanata").
		URL("https://music-website.s3.nl-ams.scw.cloud/Tokusya-Seizon%20Wonder-la-der%21%21.mp3").
		ID("tokusya-seizon")
	f.songs["kegarenaki-barajuuji"] = newSong().
		Title("Kegarenaki Barajuuji - Ariabl'eyeS").
		URL("https://music-website.s3.nl-ams.scw.cloud/kegarenaki-barajuuji.mp3").
		ID("kegarenaki-barajuuji")
	f.songs["error-towa"] = newSong().
		Title("-ERROR (Cover) - Tokoyami Towa").
		URL("https://music-website.s3.nl-ams.scw.cloud/error-towa.mp3").
		ID("error-towa")
	f.songs["diamond-city-lights"] = newSong().
		Title("Diamond City Lights - LazuLight").
		URL("https://music-website.s3.nl-ams.scw.cloud/diamond-city-lights-lazulight.opus").
		ID("diamond-city-lights")
	f.songs["tsunami-finana"] = newSong().
		Title("TSUNAMI - Finana Ryugu").
		URL("https://music-website.s3.nl-ams.scw.cloud/tsunami-finana.opus").
		ID("tsunami-finana")
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
								Class("finger-hover").
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
								Class("finger-hover").
								Text(f.songs[s].ITitle).
								OnClick(func(ctx app.Context, e app.Event) {
									ctx.NewActionWithValue("switchSong", f.songs[s].IID)
								}),
						)
				}),
			),
	)
}
