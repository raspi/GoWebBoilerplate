package main

import (
	webfw "github.com/gin-gonic/gin"
	"./application"
)

func main() {

	app := webfw.New()
	application.NewApplication(app)

	app.Run(":8080")
}
