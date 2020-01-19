package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"ibp/interfaces"
	"log"
	"github.com/joho/godotenv"
	"os"
)

func EnvLoad() {
	if os.Getenv("ENV_MODE") == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	} else {
		// TODO dev環境以外の場合は、.envファイルではなく、HOSTサーバーの環境変数を使用。
	}
}

func main () {
	EnvLoad()
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	interfaces.Routes(e)
	interfaces.Run(e, os.Getenv("PORT"))
}