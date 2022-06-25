package api

import (
	"log"
	"net/http"
	"net/http/cookiejar"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"golang.org/x/net/publicsuffix"
)

type ApiApplication interface {
	NewRouter() *mux.Router
	Visit(w http.ResponseWriter, r *http.Request)
}

func NewApiApp() (ApiApplication, error) {
	cookiejar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return nil, err
	}
	return &apiapp{
		cj: cookiejar,
	}, nil
}

type apiapp struct {
	router mux.Router
	cj     *cookiejar.Jar
}

func (a *apiapp) NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/visit", a.Visit)

	return router
}

// Called when someone visits any page of the website
// Calls for the spyware cookie and sets it if it doesn't yet exist
func (a *apiapp) Visit(w http.ResponseWriter, r *http.Request) {
	log.Printf("Visit called\n")

	c, err := r.Cookie("spyware")
	if err != nil && err != http.ErrNoCookie {
		log.Printf("Error: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if err == http.ErrNoCookie {
		// Create cookie and send it
		log.Printf("No cookie sent by client, sending cookie to them!\n")
		c = &http.Cookie{Name: "spyware", Value: uuid.NewString(), MaxAge: 0}
		http.SetCookie(w, c)
	}

	w.WriteHeader(200)
	w.Write([]byte("yeet"))

	log.Printf("Someone visited: %s\n", c.Value)

}
