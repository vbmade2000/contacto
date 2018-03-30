package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vbmade2000/contacto/config"
	"github.com/vbmade2000/contacto/handlers"
	"net"
	"os"
	"strconv"
)

func main() {
	var conf config.Config
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

	r := gin.Default()
	r.GET("/contact/:contactID", handlers.GetContact)
	r.GET("/contacts", handlers.GetContacts)
	r.POST("/contact", handlers.CreateContact)
	r.DELETE("/contact/:contactID", handlers.DeleteContact)
	r.Run(net.JoinHostPort(conf.Host, strconv.Itoa(conf.ServerPort))) // listen and serve on 0.0.0.0:8080
}
