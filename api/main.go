package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vbmade2000/contacto/handlers"
)

func main() {
	r := gin.Default()
	r.GET("/contact/:contactID", handlers.GetContact)
	r.GET("/contacts", handlers.GetContacts)
	r.POST("/contact", handlers.CreateContact)
	r.DELETE("/contact/:contactID", handlers.DeleteContact)
	r.Run() // listen and serve on 0.0.0.0:8080
}
