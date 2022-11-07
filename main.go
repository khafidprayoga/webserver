package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
)

func main() {
	serve := echo.New()
	serve.HideBanner = true

	serve.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "./public",
		Index:  "index.html",
		Browse: os.Getenv("APP_BROWSE") == "1",
		HTML5:  true,
	}))

	serve.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
			`"latency_human":"${latency_human}"` +
			`,"bytes_in":${bytes_in},"bytes_out":${bytes_out}}` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
	}))

	Port := os.Getenv("APP_PORT")

	if Port == "" {
		//	Running on default port
		log.Fatal(serve.Start(":8080"))
	} else {
		socket := fmt.Sprintf(":%v", Port)
		log.Fatal(serve.Start(socket))
	}
}
