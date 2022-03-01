package pages

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHandlers(r *gin.RouterGroup, service Service) {
	res := resource{service}

	r.GET("/index", res.getindex)
}

type resource struct {
	service Service
}

func (r resource) getindex(c *gin.Context) {
	page, err := r.service.Page(c.Request.Context(), "index")
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	page.template.Execute(c.Writer, nil)
}
