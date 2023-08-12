package main

import (
	"energyApp_site/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	//login/reg part

	reg := r.Group("/reg")
	{
		reg.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "reg.html", gin.H{
				"title": "REG",
			})
		})

		reg.POST("/", func(c *gin.Context) {
			username := c.PostForm("username")
			email := c.PostForm("email")
			password := c.PostForm("password")
			addUser := controller.AddUser(username, password, email)

			if addUser == 0 {
				c.JSON(http.StatusOK, gin.H{
					"username": username,
					"password": password,
					"email":    email,
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Error",
				})
			}
		})
	}

	login := r.Group("/login")
	{
		login.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"title": "LOGIN",
			})
		})

		login.POST("/", func(c *gin.Context) {
			username := c.PostForm("username")
			password := c.PostForm("password")

			checkUser := controller.CheckUser(username, password)

			if checkUser == 0 {
				c.JSON(http.StatusOK, gin.H{
					"username": username,
					"password": password,
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"message": "ERROR",
				})
			}
		})
	}
	//end login/reg part

	r.Run(":45600")
}
