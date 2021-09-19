
package main

import (
	"github.com/DaffaBarin/simple-books-mvc/config"
	"github.com/DaffaBarin/simple-books-mvc/routes"
)

func main() {
	config.InitDB()
	echoApp := routes.NewRoutes()
	echoApp.Start(":8080")
}