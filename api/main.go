package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vbmade2000/contacto/config"
	"github.com/vbmade2000/contacto/handlers"
	"github.com/vbmade2000/contacto/contact"
	"net"
	"os"
	"strconv"
)

func main() {
	var conf config.Config
	contactManager := contact.ContactManager{}
	
	if config.ConfigExists() == false {
		fmt.Println("contacto.conf not found. Writing default one...")
		conf = config.GetDefaultConfig()
		configFile, err := os.OpenFile("contacto.conf", os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0644)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		err = config.WriteConfig(conf, configFile)
		configFile.Close()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	} else {
		fmt.Println("Found contacto.conf. Reading configurations...")
		configFile, err := os.Open("contacto.conf")
		conf, err = config.ReadConfig(configFile)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	contact.InitDB(conf)
	handlerManager := handlers.HandlerManager{ConMan: contactManager}

	r := gin.Default()
	r.GET("/contact/:contactID", handlerManager.GetContact)
	r.GET("/contacts", handlerManager.GetContacts)
	r.POST("/contact", handlerManager.CreateContact)
	r.DELETE("/contact/:contactID", handlerManager.DeleteContact)
	r.Run(net.JoinHostPort(conf.Host, strconv.Itoa(conf.ServerPort))) // listen and serve on 0.0.0.0:8080
}
