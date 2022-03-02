package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type AboutPage struct {
	app.Compo
}

func NewAboutPage() *AboutPage {
	return &AboutPage{}
}

func (a *AboutPage) Render() app.UI {
	return app.Div().Body(
		&header{},
		app.Div().Body(
			newNavbar(),
			&aboutPanel{},
		).Class("main"),
	)
}

type aboutPanel struct {
	app.Compo

	aboutText string
}

func (a *aboutPanel) Render() app.UI {
	return app.Div().Body(
		app.Img().Src("/web/static/images/rin-1.gif").Styles(map[string]string{"width": "100px", "position": "absolute", "top": "10px", "right": "10px"}),
		app.Raw(`<p>I am a 21 year old computer science student, living and studying in The Netherlands. I like Docker, Kubernetes and Golang!
<br>
I made this website because I was inspired again by the amazing Neocities pages that I discovered because of my friends.
They also have their own pages (you can find them on the friends tab, do check them out!) and I just had to get a good website of my own!
<br>
I am not that great at web development, especially design, but I love trying it regardless!
<br><br>
To say a bit more about me personally, I love all things computers. From servers to embedded devices! I love the cloud and all that it brings
(except for big megacorps, but alright) and it's my goal to work for a big cloud company!
<br>
Aside from career path ambitions,　ボーカロイドはすきです！ I love vocaloid and other Japanese music and culture!!
I also like Vtubers, especially from Hololive and it's my goal to one day finally understand them in their native language!
<br><br>
There is a lot more to say in words, but who cares about those! Have a look around my creative digital oasis and see what crazy stuff you can find!</p>`),
	).Class("content")
}
