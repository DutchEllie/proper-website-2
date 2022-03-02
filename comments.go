package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"git.home.dutchellie.nl/DutchEllie/proper-website-2/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (a *application) Comment(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var comment entity.Comment
		body, err := io.ReadAll(r.Body)
		if err != nil {
			a.WriteError(w, http.StatusInternalServerError, err.Error())
			return
		}
		err = json.Unmarshal(body, &comment)
		if err != nil {
			a.WriteError(w, http.StatusInternalServerError, err.Error())
			return
		}
		comment.PostDate = time.Now()

		if comment.Name == "" || comment.Message == "" {
			a.WriteError(w, http.StatusBadRequest, "one or more fields empty")
			return
		}

		_, err = a.collection.InsertOne(context.Background(), comment)
		if err != nil {
			a.WriteError(w, http.StatusInternalServerError, err.Error())
			return
		}
		w.WriteHeader(200)
		return
	case "GET":
		comments := make([]entity.Comment, 0)
		filter := bson.D{}
		sort := options.Find()
		sort.SetSort(bson.D{{"time", -1}})
		cur, err := a.collection.Find(context.Background(), filter, sort)
		if err != nil {
			a.WriteError(w, http.StatusInternalServerError, err.Error())
			return
		}
		err = cur.All(context.Background(), &comments)
		if err != nil {
			a.WriteError(w, http.StatusInternalServerError, err.Error())
			return
		}
		jsondata, err := json.Marshal(comments)
		if err != nil {
			a.WriteError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(200)
		w.Write(jsondata)
		return
	}
}

func (a *application) WriteError(w http.ResponseWriter, code int, err string) {
	w.WriteHeader(code)
	w.Write([]byte(err))
}
