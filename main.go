package main

import (
	"github.com/laironacosta/ms-gin-go/routes"
)

func main() {
	r := routes.NewRouter()
	r.Run() // listen and serve on 0.0.0.0:8080
}
