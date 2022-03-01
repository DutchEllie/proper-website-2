package main

import (
	"fmt"
	"os"

	"git.home.dutchellie.nl/DutchEllie/proper-website-2/internal/pages"
	"git.home.dutchellie.nl/DutchEllie/proper-website-2/internal/templatelib"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("SRV_PORT")
	if port == "" {
		msg := `Website writen in Go!

Please include all the proper environment variables:
- "SRV_PORT" <number of the exposing port>
		`
		fmt.Println(msg)
		return
	}
	address := ":" + port

	lib, err := templatelib.NewTemplateLibrary("../../website/templates")
	if err != nil {
		fmt.Println(err)
		return
	}

	router, _ := setupRouter(lib)

	router.Run(address)
}

func setupRouter(lib templatelib.TemplateLibrary) (*gin.Engine, error) {
	router := gin.Default()

	pages.RegisterHandlers(router.Group("/"), pages.NewService(pages.NewRepository(&lib)))
	return router, nil
}
