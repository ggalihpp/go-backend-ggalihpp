package main

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/joho/godotenv"
)

// H -> Will be dynamic struct
type H map[string]interface{}

func main() {
	envFile := os.Getenv("ENV_FILE")
	if envFile == "" {
		envFile = ".env"
	}

	err := godotenv.Load(envFile)
	if err != nil {
		panic(err)
	}

	e := echo.New()

	// Activation CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.PATCH},
	}))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "IP=${remote_ip}, method=${method}, uri=${uri}, status=${status}\n",
	}))

	//e.Use(middleware.BasicAuth(CheckBasicAuth))

	setupHandlers(e)

	if err = e.Start(os.Getenv("PORT")); err != nil {
		panic(err)
	}

}

func checkBasicAuth(username, password string, c echo.Context) (bool, error) {
	if username == os.Getenv("AUTHUSERNAME") && password == os.Getenv("AUTHPASSWORD") {
		return true, nil
	}
	return false, nil
}
