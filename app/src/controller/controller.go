package controller

import (
	"net/http"

	"github.com/biswash-grg/test-webservice/app/src/service"
	"github.com/gin-gonic/gin"
)

/*
RegisterController Controller for the web app
*/
func RegisterController() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world from GIN! A web framework for GO!")
	})

	r.GET("/students", func(c *gin.Context) {
		ok := service.IsAuthenticated(c.Request)
		if !ok {
			c.String(http.StatusUnauthorized, "401 - You are not authorized to access this resource!")
			return
		}
		c.JSON(http.StatusOK, service.GetAllStudents())
	})

	r.Run(":3000")

}
