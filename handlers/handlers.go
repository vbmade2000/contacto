package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/vbmade2000/contacto/contact"
	"github.com/vbmade2000/contacto/config"
)

type HandlerManager struct {
	Conf config.Config
	ConMan contact.ContactManager
}

func (hMan *HandlerManager) GetContact(c *gin.Context) {
	// TODO: Get contactID from url and pass it to function
	// Retrieve a contact with supplied ID
	cntact, err:= hMan.ConMan.GetContact(2)
	if err!=nil {
		c.String(500, err.Error())
	} else {
		// Convert interface type to Contact struct type
		cntact = cntact.(contact.Contact)
		// Convert retrieved contact to json format to output
		if jsonContact, err := json.Marshal(cntact); err!=nil {
			c.String(500, string(err.Error()))
		} else {
			c.String(200, string(jsonContact))
		}
	}
}

func (hMan *HandlerManager) GetContacts(c *gin.Context) {
	
	c.String(200, "This is handler for retrieving all contacts")
}

func (hMan *HandlerManager) CreateContact(c *gin.Context) {
	c.String(200, "This is handler for creating new contact")
}

func (hMan *HandlerManager) DeleteContact(c *gin.Context) {
	c.String(200, "This is handler for deleting contact")
}

func (hMan *HandlerManager) UpdateContact(c *gin.Context) {
	c.String(200, "This is handler for updating contact")
}
