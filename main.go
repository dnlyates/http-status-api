package main

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

func createResponseHandler(statusCode int) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(statusCode, http.StatusText(statusCode))
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	statusCodes := [5]int{http.StatusOK, http.StatusBadRequest, http.StatusNotFound, http.StatusUnprocessableEntity, http.StatusInternalServerError}
	for _, s := range statusCodes {
		r.GET(fmt.Sprintf("/%d", s), createResponseHandler(s))
		r.PUT(fmt.Sprintf("/%d", s), createResponseHandler(s))
		r.PATCH(fmt.Sprintf("/%d", s), createResponseHandler(s))
		r.POST(fmt.Sprintf("/%d", s), createResponseHandler(s))
		r.DELETE(fmt.Sprintf("/%d", s), createResponseHandler(s))
	}

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
