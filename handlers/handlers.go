package handlers

import (
	"github.com/gin-gonic/gin"
)

func GetContact(c *gin.Context) {
	c.String(200, "This is handler for retrieving a single contact")
}

func GetContacts(c *gin.Context) {
	c.String(200, "This is handler for retrieving all contacts")
}

func CreateContact(c *gin.Context) {
	c.String(200, "This is handler for creating new contact")
}

func DeleteContact(c *gin.Context) {
	c.String(200, "This is handler for deleting contact")
}
