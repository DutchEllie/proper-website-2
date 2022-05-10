# proper-website-2

[![Build Status](https://drone.dutchellie.nl/api/badges/DutchEllie/proper-website-2/status.svg)](https://drone.dutchellie.nl/DutchEllie/proper-website-2)

A truly proper website this time™

**TODO**:
- Change domain to quenten.nl and staging.quenten.nl
- Dynamically make domains for other branches
- Make a generic page component, so that you don't have to add the standard panels every time (such as navbar, leftbar, header, etc)

App notes:
- leftbar (under nav) is default always the same. 
  This is in a standardized "page" component.
  it can be changed by specifying it when creating the "page" component, or maybe in a prerender/onnav (or both).
  Structure will be:
    Generic page component IS IMPLEMENTED BY "specific page" WHERE OnNav/OnPreRender specify its details WHICH IS CREATED with its own "newXPage()" function. See https://github.com/maxence-charriere/go-app/blob/master/docs/src/actions-page.go for example.
- "main content" is an array of app.UI elements.
  This can also be further abstracted using "block" components (the ones with the cool color)
- THOSE BLOCK COMPONENTS' CONTENT can be set manually on creation, OR by passing HTML/MarkDown (not decided yet)

This website will be done with this:
- Backend written in Golang
- Templating with Go
- Database connection for comments and such
- Markdown article writing

It will be done in the following order:
- First get up a working simple proof of concept
- Then make the comments better
- Database better
- Layout
- Markdown article writing
- Layout