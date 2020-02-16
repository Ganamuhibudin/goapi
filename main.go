package main

import (
	"flag"
	"fmt"
	"github.com/ganamuhibudin/goapi/config"
	"github.com/ganamuhibudin/goapi/utils"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func main() {
	var host, port string
	flag.StringVar(&host, "host", os.Getenv("HOST"), "host of the service")
	flag.StringVar(&port, "port", os.Getenv("PORT"), "port of the service")
	flag.Parse()

	db := utils.GetDBConnection()
	defer db.Close()

	routes := &config.Routes{DB: db}

	//info version service
	fmt.Printf("Service: %s\nVersion: %s\nParams:\n-host: host of the service\n-port: port of the service\nFramework:\n", os.Getenv("APP_NAME"), os.Getenv("APP_VER"))

	if rval := recover(); rval != nil {
		fmt.Printf("Rval: %+v\n", rval)
	}
	// Setup routes and run application
	routes.Setup(host, port)
}
